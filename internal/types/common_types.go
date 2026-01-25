package types

// PaginationRequest 分页请求
type PaginationRequest struct {
	Page     int `json:"page" form:"page" binding:"min=1" default:"1"`
	PageSize int `json:"page_size" form:"page_size" binding:"min=1,max=100" default:"10"`
}

// PaginationResponse 分页响应
type PaginationResponse struct {
	Total    int64       `json:"total"`     // 总记录数
	Page     int         `json:"page"`      // 当前页码
	PageSize int         `json:"page_size"` // 每页记录数
	Pages    int64       `json:"pages"`     // 总页数
	Data     interface{} `json:"data"`
}
