package service

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"{{ cookiecutter.project_name }}/biz"
	"{{ cookiecutter.project_name }}/data/ent"
	"{{ cookiecutter.project_name }}/data/ent/predicate"
	"{{ cookiecutter.project_name }}/data/ent/user"
	"{{ cookiecutter.project_name }}/parser"
	"{{ cookiecutter.project_name }}/pkg/e"
	"{{ cookiecutter.project_name }}/serializer"
	"{{ cookiecutter.project_name }}/setting"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type UserService struct {
	log *logger.Logger
	uc  *biz.UserUseCase
}

func NewUserService(log *logger.Logger, uc *biz.UserUseCase) *UserService {
	us := &UserService{
		log: log,
		uc:  uc,
	}
	us.initAdmin()
	return us
}

func (u *UserService) initAdmin() {
	ctx := context.Background()
	admin, _ := u.uc.Get(ctx, 1)
	if admin == nil {
		password := u.randomPassword()
		err := u.uc.Create(ctx, &biz.User{
			Username:    "admin",
			Name:        "admin",
			Password:    password,
			Email:       "admin@exaple.com",
			Phone:       "13000000000",
			IsActive:    true,
			IsSuperuser: true,
			LastLogin:   nil,
		})
		if err != nil {
			panic(err)
		}
		u.log.Infof("创建管理员账户：admin。密码为：%v。该密码只会显示一次，请记住此密码", password)
	}
}

func (u *UserService) Login(ctx *gin.Context, form *parser.UserLoginParser) (*serializer.LoginSerializer, error) {
	_user, err := u.uc.Login(ctx, form.Username, form.Password)
	if err != nil {
		return nil, err
	}
	allPermissions, err := u.uc.GetUserAllPermissions(ctx, _user.Id)
	if err != nil {
		return nil, err
	}
	nowTime := time.Now()
	jwtId, _ := uuid.NewUUID()
	claims := jwt.RegisteredClaims{
		Issuer:    "{{ cookiecutter.project_name }}",
		Subject:   strconv.Itoa(_user.Id),
		NotBefore: &jwt.NumericDate{Time: nowTime},
		IssuedAt:  &jwt.NumericDate{Time: nowTime},
		ExpiresAt: &jwt.NumericDate{
			Time: nowTime.Add(setting.AppSetting.JwtTime),
		},
		ID: jwtId.String(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(setting.AppSetting.JwtSecret))
	if err != nil {
		return nil, err
	}
	ctx.Set("userId", _user.Id)
	return serializer.NewLoginSerializer(_user, allPermissions, token), nil
}

func (u *UserService) Create(ctx *gin.Context, form *parser.CreateUserParser) error {
	_user := &biz.User{
		Username:    form.Username,
		Name:        form.Name,
		Password:    form.Password,
		Email:       form.Email,
		Phone:       form.Phone,
		IsActive:    form.IsActive,
		IsSuperuser: form.IsSuperuser,
		LastLogin:   nil,
	}

	if len(form.GroupIds) > 0 {
		gs := make([]*biz.Group, len(form.GroupIds))
		for i, id := range form.GroupIds {
			gs[i] = &biz.Group{Id: id}
		}
		_user.Groups = gs
	}
	if len(form.PermissionIds) > 0 {
		ps := make([]*biz.Permission, len(form.PermissionIds))
		for i, id := range form.PermissionIds {
			ps[i] = &biz.Permission{Id: id}
		}
		_user.Permissions = ps
	}
	return u.uc.Create(ctx, _user)
}

func (u *UserService) Update(ctx *gin.Context, form *parser.UpdateUserParser) error {
	_user := &biz.User{
		Id:          form.Id,
		Username:    form.Username,
		Name:        form.Name,
		Password:    form.Password,
		Email:       form.Email,
		Phone:       form.Phone,
		IsActive:    form.IsActive,
		IsSuperuser: form.IsSuperuser,
	}
	if len(form.GroupIds) > 0 {
		gs := make([]*biz.Group, len(form.GroupIds))
		for i, id := range form.GroupIds {
			gs[i] = &biz.Group{Id: id}
		}
		_user.Groups = gs
	}
	if len(form.PermissionIds) > 0 {
		ps := make([]*biz.Permission, len(form.PermissionIds))
		for i, id := range form.PermissionIds {
			ps[i] = &biz.Permission{Id: id}
		}
		_user.Permissions = ps
	}
	return u.uc.Update(ctx, _user)
}

func (u *UserService) Delete(ctx *gin.Context, form *parser.DeleteUserParser) error {
	if form.Id == 1 {
		return e.RaiseApiException(400, e.ErrDeleteAdmin)
	}
	return u.uc.Delete(ctx, form.Id)
}

func (u *UserService) Get(ctx *gin.Context, form *parser.GetUserParser) (*serializer.UserSerializer, error) {
	_user, err := u.uc.Get(ctx, form.Id)
	if err != nil {
		return nil, err
	}
	return serializer.NewUserSerializer(_user), nil
}

func (u *UserService) List(ctx *gin.Context, form *parser.GetUserListParser) (*serializer.PaginationSerializer, error) {
	filters := make([]predicate.User, 0)
	if form.Name != "" {
		filters = append(filters, user.Name(form.Name))
	}
	if form.Username != "" {
		filters = append(filters, user.Username(form.Username))
	}
	if form.IsSuperuser != "" {
		filters = append(filters, user.IsSuperuser(form.IsSuperuser == "true"))
	}
	if form.IsActive != "" {
		filters = append(filters, user.IsActive(form.IsActive == "true"))
	}
	ords := make([]ent.OrderFunc, 0)
	if form.Ordering != "" {
		ordings := strings.Split(form.Ordering, ",")
		for _, ording := range ordings {
			if strings.HasPrefix(ording, "-") {
				ords = append(ords, ent.Desc(ording[1:]))
			} else {
				ords = append(ords, ent.Asc(ording))
			}
		}
	}

	total, users, err := u.uc.Pagination(ctx, form.Page, form.PageSize, ords, filters...)
	if err != nil {
		return nil, err
	}
	us := make([]*serializer.UserSerializer, len(users))
	for i, _u := range users {
		us[i] = serializer.NewUserSerializer(_u)
	}
	return serializer.NewPaginationSerializer(total, us), nil
}

func (u *UserService) SimpleList(ctx *gin.Context) ([]*serializer.SimpleSerializer, error) {
	users, err := u.uc.All(ctx)
	if err != nil {
		return nil, err
	}
	return serializer.NewSimpleListSerializer(users, "Name", "Id"), nil
}

func (u *UserService) HasPerm(ctx *gin.Context, userId int, permName string) (bool, error) {
	return u.uc.HasPerm(ctx, userId, permName)
}

func (u *UserService) ResetMyPassword(ctx *gin.Context, form *parser.ResetMyPasswordParser) error {
	_user, err := u.uc.Get(ctx, form.UserId)
	if err != nil {
		return err
	}
	if ok := u.uc.ValidatePassword(_user.HashedPassword, form.OldPassword); !ok {
		return e.RaiseApiException(403, e.ErrorPassword)
	}
	_user.Password = form.Password
	err = u.uc.Update(ctx, _user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) randomPassword() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
