package query

const (
	DefaultPageIndex = 0
	DefaultPageNum   = 20
)

type PageParam struct {
	PageIndex int `json:"page_index"`
	PageNum   int `json:"page_num"`
}
