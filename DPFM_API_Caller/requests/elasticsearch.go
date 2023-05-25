package requests

type Aggs struct {
	DuplicateAggs DuplicateAggs `json:"duplicate_aggs"`
}

type DuplicateAggs struct {
	Terms Terms `json:"terms"`
}

type Terms struct {
	Field string `json:"field"`
	Size  int    `json:"size"`
}

type Query struct {
	Match Match `json:"match"`
}

type Match struct {
	Product string `json:"Product"`
}

type DescendingOrderQuery struct {
	Query Query `json:"query"`
	Aggs  Aggs  `json:"aggs"`
}
