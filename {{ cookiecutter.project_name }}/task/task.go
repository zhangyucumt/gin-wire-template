package task

import (
	"context"
	"github.com/google/wire"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"time"
)

type Task struct {
	log              *logger.Logger
}

func NewTask(log *logger.Logger) *Task {
	return &Task{log: log}
}

func (t *Task) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)

	//eg.Go(func() error {
	//	ticker := time.NewTicker(10 * time.Second)
	//	defer ticker.Stop()
	//	return t.run(ctx, ticker, t.checkClusterTask.Run, "检查数据库集群和 Prometheus 集群同步状态")
	//})
	// eg.Go(func() error {
	// 	return t.runOnce(ctx, t.checkClusterTask.Run, "检查数据库集群和 Prometheus 集群同步状态")
	// })

	// eg.Go(func() error {
	// 	ticker := time.NewTicker(15 * time.Second)
	// 	defer ticker.Stop()
	// 	return t.run(ctx, ticker, t.syncSwitch.Run, "检查交换机和 Prometheus 集群同步状态")
	// })

	// eg.Go(func() error {
	// 	ticker := time.NewTicker(15 * time.Second)
	// 	defer ticker.Stop()
	// 	return t.run(ctx, ticker, t.sendAlertsTask.Run, "发送告警短信和邮件")
	// })

	err := eg.Wait()
	if err != nil {
		t.log.Errorf("task eg error: %v", err)
	}
	return err
}

func (t *Task) run(ctx context.Context, ticker *time.Ticker, f func(context.Context) error, name string) error {
	t.log.Infof("注册任务: %v", name)
	for {
		select {
		case <-ticker.C:
			err := f(ctx)
			if err != nil {
				t.log.Errorf("task %v running error: %v", name, err)
			}
		case <-ctx.Done():
			ticker.Stop()
			t.log.Infof("task %v stopped", name)
			return nil
		}
	}
}

func (t *Task) runOnce(ctx context.Context, f func(context.Context) error, name string) error {
	t.log.Infof("运行一次性任务: %v", name)
	err := f(ctx)
	if err != nil {
		t.log.Infof("运行一次性任务失败: %v", name)
		return err
	}
	t.log.Infof("运行一次性任务成功: %v", name)
	return nil
}

var ProviderSet = wire.NewSet(NewTask)
