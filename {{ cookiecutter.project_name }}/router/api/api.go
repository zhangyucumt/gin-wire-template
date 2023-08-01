package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"{{ cookiecutter.project_name }}/pkg/e"
	"{{ cookiecutter.project_name }}/service"
)

type Api struct {
	userService       *service.UserService
	groupService      *service.GroupService
	permissionService *service.PermissionService
	operateLogService *service.OperateLogService
}

func NewApi(
	userService *service.UserService,
	groupService *service.GroupService,
	permissionService *service.PermissionService,
	operateLogService *service.OperateLogService,
) *Api {
	return &Api{userService: userService, groupService: groupService, permissionService: permissionService, operateLogService: operateLogService}
}

func (api *Api) GetOperateLogService() *service.OperateLogService {
	return api.operateLogService
}

func (api *Api) GetUserService() *service.UserService {
	return api.userService
}

func (api *Api) success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"errorCode":    e.SUCCESS,
		"errorMessage": e.GetMsg(e.SUCCESS),
		"data":         data,
		"showType":     0,
	})
}

func (api *Api) deleted(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"success":      true,
		"errorCode":    e.SUCCESS,
		"errorMessage": e.GetMsg(e.SUCCESS),
		"showType":     0,
	})
}

func (api *Api) created(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"success":      true,
		"errorCode":    e.SUCCESS,
		"errorMessage": e.GetMsg(e.SUCCESS),
		"showType":     0,
	})
}

func (api *Api) updated(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"success":      true,
		"errorCode":    e.SUCCESS,
		"errorMessage": e.GetMsg(e.SUCCESS),
		"showType":     0,
	})
}

func (api *Api) failure(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"success":      false,
		"errorCode":    statusCode,
		"errorMessage": message,
		"data":         nil,
		"showType":     2,
	})
}

func (api *Api) error(c *gin.Context, err error) {
	er, ok := err.(*e.ApiException)
	if ok {
		c.JSON(er.GetStatusCode(), gin.H{
			"success":      false,
			"errorCode":    er.GetCode(),
			"errorMessage": er.Error(),
			"data":         nil,
			"showType":     2,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorCode":    e.ERROR,
			"errorMessage": err.Error(),
			"data":         nil,
			"showType":     2,
		})
	}
}

//var validate *validator.Validate
//
//func (api *Api) validate(c *gin.Context, form interface{}) error {
//	validate = validator.New()
//	err := validate.Struct(form)
//
//	//err := c.ShouldBind(form)
//	if err != nil {
//		msg := api.processErr(form, err)
//		fmt.Println(msg)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"code": e.InvalidParams,
//			"msg":  msg,
//			"data": nil,
//		})
//		return err
//	}
//	return nil
//}

//func (api *Api) processErr(c *gin.Context, u interface{}, err error) {
//	if err == nil { //如果为nil 说明校验通过
//		return
//	}
//
//	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
//	if ok {
//		api.failure(c, http.StatusBadRequest, "输入参数错误："+invalid.Error())
//		return
//	}
//	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
//	//fmt.Println(validationErrs)
//	for _, validationErr := range validationErrs {
//		fieldName := validationErr.Field()                    //获取是哪个字段不符合格式
//		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
//		if ok {
//			errorInfo := field.Tag.Get("reg_error_info") //获取field对应的reg_error_info tag值
//
//			jsonFieldName := field.Tag.Get("json")
//			if jsonFieldName == "" {
//				jsonFieldName = fieldName
//			}
//
//			if errorInfo != "" {
//				api.failure(c, http.StatusBadRequest, errorInfo)
//				return
//			}
//			validateTag := field.Tag.Get("validate") //获取field对应的validate tag值
//			if validateTag != "" {
//				switch validateTag {
//				case "required":
//					api.failure(c, http.StatusBadRequest, jsonFieldName+":不能为空")
//					return
//				case "email", "url":
//					api.failure(c, http.StatusBadRequest, jsonFieldName+":格式不正确")
//					return
//				case "min":
//					api.failure(c, http.StatusBadRequest, jsonFieldName+":不能小于"+field.Tag.Get("min"))
//					return
//				case "max":
//					api.failure(c, http.StatusBadRequest, jsonFieldName+":不能大于"+field.Tag.Get("max"))
//					return
//				case "len":
//					api.failure(c, http.StatusBadRequest, jsonFieldName+":长度不正确")
//					return
//				default:
//					api.failure(c, http.StatusBadRequest, jsonFieldName+":"+validateTag)
//					return
//				}
//			}
//		} else {
//			api.failure(c, http.StatusBadRequest, "缺失reg_error_info")
//			return
//		}
//	}
//	api.failure(c, http.StatusBadRequest, err.Error())
//	return
//}

var ProviderSet = wire.NewSet(NewApi)
