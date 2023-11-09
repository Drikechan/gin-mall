package types

type BasePageTypes struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

type DataListResp struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}
