package common

type PaginateResp struct {
	Total     int64 `json:"total"`      // 总记录数
	Page      int64 `json:"page"`       // 页数
	PageSize  int64 `json:"page_size"`  // 每页记录数
	TotalPage int64 `json:"total_page"` // 总页数
}
