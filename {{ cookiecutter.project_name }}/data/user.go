package data

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sync"
	"time"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/data/ent"
	"{{ cookiecutter.project_name }}/data/ent/predicate"
	u "{{ cookiecutter.project_name }}/data/ent/user"
	"{{ cookiecutter.project_name }}/pkg/e"
)

type userData struct {
	log  *logger.Logger
	data *Data
	lock sync.RWMutex
}

func NewUserData(log *logger.Logger, data *Data) biz.UserRepo {
	return &userData{log: log, data: data}
}

func (s *userData) Login(ctx context.Context, username, password string) (*biz.User, error) {
	user, err := s.data.db.User.Query().WithGroups().WithPermissions().Where(u.Username(username)).First(ctx)
	if err != nil {
		return nil, e.RaiseApiException(http.StatusBadRequest, e.AuthError)
	}
	if !s.ValidatePassword(user.Password, password) {
		return nil, e.RaiseApiException(http.StatusBadRequest, e.AuthError)
	}
	if !user.IsActive {
		return nil, e.RaiseApiException(http.StatusForbidden, e.ErrorUserForbid)
	}
	_, err = s.data.db.User.Update().Where(u.ID(user.ID)).SetLastLogin(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return biz.NewUser(user), nil
}

func (s *userData) List(ctx context.Context, page, pageSize int, ordering []ent.OrderFunc, filters ...predicate.User) ([]*biz.User, error) {
	page, pageSize = s.data.parserPage(page, pageSize)
	users, err := s.data.db.User.Query().Where(filters...).WithGroups().WithPermissions().Order(ordering...).Limit(pageSize).Offset((page - 1) * pageSize).All(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]*biz.User, len(users))
	for i, user := range users {
		ret[i] = biz.NewUser(user)
	}
	return ret, nil
}

func (s *userData) All(ctx context.Context) ([]*biz.User, error) {
	users, err := s.data.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]*biz.User, len(users))
	for i, user := range users {
		ret[i] = biz.NewUser(user)
	}
	return ret, nil
}

func (s *userData) Count(ctx context.Context, filters ...predicate.User) (int, error) {
	return s.data.db.User.Query().Where(filters...).Count(ctx)
}

func (s *userData) Create(ctx context.Context, user *biz.User) error {
	s.lock.Lock()

	defer func() {
		s.lock.Unlock()
	}()

	us1, err := s.Filter(ctx, u.Username(user.Username))
	if len(us1) > 0 || err != nil {
		return e.RaiseApiException(http.StatusBadRequest, e.ErrorUserExist)
	}

	generatePassword, err := s.generatePassword(user.Password)
	if err != nil {
		return err
	}
	_uc := s.data.db.User.Create().SetUsername(user.Username).SetEmail(user.Email).SetPhone(user.Phone).SetPassword(generatePassword).SetName(user.Name).SetIsActive(user.IsActive).SetIsSuperuser(user.IsSuperuser)
	if len(user.Groups) > 0 {
		gIds := make([]int, len(user.Groups))
		for i, g := range user.Groups {
			gIds[i] = g.Id
		}
		_uc.AddGroupIDs(gIds...)
	}
	if len(user.Permissions) > 0 {
		pIds := make([]int, len(user.Permissions))
		for i, p := range user.Permissions {
			pIds[i] = p.Id
		}
		_uc.AddPermissionIDs(pIds...)
	}
	_, err = _uc.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *userData) Update(ctx context.Context, user *biz.User) error {
	//tx, err := s.data.db.Tx(ctx)
	//if err != nil {
	//	return err
	//}
	//defer func() {
	//	_ = s.rollback(tx, err)
	//}()
	us := s.data.db.User.UpdateOneID(user.Id).SetUsername(user.Username).SetEmail(user.Email).SetPhone(user.Phone).SetName(user.Name).SetIsActive(user.IsActive).SetIsSuperuser(user.IsSuperuser)

	if user.Password != "" {
		generatePwd, err := s.generatePassword(user.Password)
		if err != nil {
			return err
		}
		us.SetPassword(generatePwd)
	}

	us.ClearGroups()
	if len(user.Groups) > 0 {
		gIds := make([]int, len(user.Groups))
		for i, g := range user.Groups {
			gIds[i] = g.Id
		}
		us.AddGroupIDs(gIds...)
	}

	us.ClearPermissions()
	if len(user.Permissions) > 0 {
		pIds := make([]int, len(user.Permissions))
		for i, p := range user.Permissions {
			pIds[i] = p.Id
		}
		us.AddPermissionIDs(pIds...)
	}
	_, err := us.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *userData) Delete(ctx context.Context, id int) error {
	tx, err := s.data.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = s.data.rollback(tx, err, recover())
	}()
	err = tx.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *userData) Get(ctx context.Context, id int) (*biz.User, error) {
	user, err := s.data.db.User.Query().WithGroups(func(q *ent.GroupQuery) {
		q.WithPermissions()
	}).WithPermissions().Where(u.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return biz.NewUser(user), nil
}

func (s *userData) Filter(ctx context.Context, filters ...predicate.User) ([]*biz.User, error) {
	users, err := s.data.db.User.Query().Where(filters...).All(ctx)
	if err != nil {
		return nil, err
	}
	ret := make([]*biz.User, len(users))
	for i, user := range users {

		ret[i] = biz.NewUser(user)
	}
	return ret, nil
}

func (s *userData) GetUserAllPermissions(ctx context.Context, id int) ([]string, error) {
	user, err := s.data.db.User.Query().WithGroups(func(q *ent.GroupQuery) {
		q.WithPermissions()
	}).WithPermissions().Where(u.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}

	if !user.IsActive {
		return nil, e.RaiseApiException(http.StatusForbidden, e.ErrorUserForbid)
	}

	if user.IsSuperuser {
		ps, err := s.data.db.Permission.Query().All(ctx)
		if err != nil {
			return nil, err
		}
		permissions := make([]string, 0)
		for _, p := range ps {
			permissions = append(permissions, p.Name)
		}
		return permissions, nil
	} else {
		permissionMap := make(map[string]bool)
		for _, p := range user.Edges.Permissions {
			permissionMap[p.Name] = true
		}
		for _, g := range user.Edges.Groups {
			for _, p := range g.Edges.Permissions {
				permissionMap[p.Name] = true
			}
		}
		permissions := make([]string, 0)
		for k := range permissionMap {
			permissions = append(permissions, k)
		}
		return permissions, nil
	}
}

func (s *userData) ValidatePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *userData) generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

//func (s *userData) rollback(tx *ent.Tx, err error) error {
//	if err != nil {
//		err2 := tx.Rollback()
//		if err2 != nil {
//			s.log.Errorf("rollback transaction fail: %v", err2)
//		}
//		return err2
//	}
//	if v := recover(); v != nil {
//		err2 := tx.Rollback()
//		if err2 != nil {
//			s.log.Errorf("rollback transaction fail: %v", err2)
//		}
//		return err2
//	}
//	if err == nil {
//		err2 := tx.Commit()
//		if err2 != nil {
//			s.log.Errorf("commit transaction fail: %v", err2)
//		}
//		return err2
//	}
//	return nil
//}
