package parser

type ListOperateLogParser struct {
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
	UserId      int    `form:"user_id" binding:"omitempty"`
	CreatedAtLt string `form:"created_at_lt" binding:"omitempty"`
	CreatedAtGt string `form:"created_at_gt" binding:"omitempty"`
	Path        string `form:"path"`
	Method      string `form:"method"`
	Name        string `form:"name"`
	Ip          string `form:"ip"`
	Ordering    string `form:"ordering" binding:"omitempty"`
	IsLoginLog  string `form:"is_login_log" binding:"omitempty,oneof=true false"`
}

type CreateOperateLogParser struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	IP         string `json:"ip"`
	UserId     int    `json:"user_id"`
}
