package global

type PaginateReq struct {
	Page     int64 `json:"page" query:"page" form:"page"`
	PageSize int64 `json:"page_size" query:"page_size" form:"page_size"`
}
type PaginatorResp[T any] struct {
	PageSize   int64 `json:"page_size"`
	Page       int64 `json:"page"`
	TotalPage  int64 `json:"total_page"`
	TotalCount int64 `json:"total_count"`
	Data       []T   `json:"data"`
}
