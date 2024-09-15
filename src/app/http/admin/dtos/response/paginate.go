package dtos

type PaginateResponse struct {
	Data       []interface{} `json:"data"`
	Pagination Pagination    `json:"pagination"`
}

type Pagination struct {
	TotalRecord int64 `json:"totalRecord"`
	TotalPage   int   `json:"totalPage"`
	PerPage     int   `json:"perPage"`
	Page        int   `json:"page"`
}
