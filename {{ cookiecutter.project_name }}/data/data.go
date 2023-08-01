package data

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"{{ cookiecutter.project_name }}/data/ent"
	"{{ cookiecutter.project_name }}/data/ent/migrate"
	"{{ cookiecutter.project_name }}/setting"
)

// Data .
type Data struct {
	db     *ent.Client
	redis  *redis.Client
	logger *logrus.Logger
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserData, NewPermissionData, NewGroupData, NewOperateLogData)

func NewData(log *logrus.Logger) (*Data, func(), error) {
	client, err := ent.Open(setting.DatabaseSetting.Type, setting.DatabaseSetting.Source)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:            setting.RedisSetting.Host,
		Password:        setting.RedisSetting.Password,
		DB:              setting.RedisSetting.Db,
		Username:        setting.RedisSetting.Username,
		MaxIdleConns:    setting.RedisSetting.MaxIdle,
		ConnMaxIdleTime: setting.RedisSetting.IdleTimeout,
		//PoolSize:        setting.RedisSetting.MaxActive,
	})
	sc := rdb.Ping(context.TODO())
	if sc.Err() != nil {
		return nil, nil, sc.Err()
	}

	d := &Data{
		db:     client,
		redis:  rdb,
		logger: log,
	}
	return d, func() {
		log.Info("关闭数据库和 redis 链接")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.redis.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func (d *Data) parserPage(page, pageSize int) (int, int) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}
	return page, pageSize
}

func (d *Data) rollback(tx *ent.Tx, err error, r any) error {
	if err != nil {
		d.logger.Debugf("rollback transaction because err: %v", err)
		err2 := tx.Rollback()
		if err2 != nil {
			d.logger.Errorf("rollback transaction fail: %v", err2)
		}
		return err2
	}
	if r != nil {
		d.logger.Debugf("rollback transaction because panic: %v", r)
		err2 := tx.Rollback()
		if err2 != nil {
			d.logger.Errorf("rollback transaction fail: %v", err2)
		}
		panic(r)
	}
	if err == nil {
		err2 := tx.Commit()
		if err2 != nil {
			d.logger.Errorf("commit transaction fail: %v", err2)
		}
		return err2
	}
	return nil
}
