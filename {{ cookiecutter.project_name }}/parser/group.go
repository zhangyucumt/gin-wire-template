package parser

type CreateGroupParser struct {
	Name          string `json:"name" binding:"required"`
	PermissionIds []int  `json:"permission_ids"`
	Description   string `json:"description"`
	UserIds       []int  `json:"user_ids"`
}

type UpdateGroupParser struct {
	Id            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	PermissionIds []int  `json:"permission_ids"`
	Description   string `json:"description"`
	UserIds       []int  `json:"user_ids"`
}

type DeleteGroupParser struct {
	Id int `json:"id" binding:"required"`
}

type GetGroupParser struct {
	Id int `json:"id" binding:"required"`
}
