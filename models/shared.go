package models

type Pagination struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}
