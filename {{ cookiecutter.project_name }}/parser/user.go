package parser

type UserLoginParser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserParser struct {
	Username      string `json:"username" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Email         string `json:"email" binding:"omitempty,email"`
	Phone         string `json:"phone"`
	GroupIds      []int  `json:"group_ids"`
	PermissionIds []int  `json:"permission_ids"`
	IsActive      bool   `json:"is_active"`
	IsSuperuser   bool   `json:"is_superuser"`
}

type ResetMyPasswordParser struct {
	Password    string `json:"password" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	UserId      int    `json:"user_id"`
}

type UpdateUserParser struct {
	Id            int    `json:"id"`
	Username      string `json:"username" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Password      string `json:"password" binding:"omitempty"`
	Email         string `json:"email" binding:"omitempty,email"`
	Phone         string `json:"phone"`
	GroupIds      []int  `json:"group_ids"`
	PermissionIds []int  `json:"permission_ids"`
	IsActive      bool   `json:"is_active"`
	IsSuperuser   bool   `json:"is_superuser"`
}

type DeleteUserParser struct {
	Id int `json:"id" binding:"required"`
}

type GetUserParser struct {
	Id int `json:"id" binding:"required"`
}

type GetUserListParser struct {
	Username    string `form:"username" binding:"omitempty"`
	Page        int    `form:"page" binding:"omitempty"`
	PageSize    int    `form:"page_size" binding:"omitempty"`
	Name        string `form:"name" binding:"omitempty"`
	Ordering    string `form:"ordering" binding:"omitempty"`
	IsActive    string `form:"is_active" binding:"omitempty,oneof=true false"`
	IsSuperuser string `form:"is_superuser" binding:"omitempty,oneof=true false"`
	GroupIds    string `form:"group_ids" binding:"omitempty"`
}
