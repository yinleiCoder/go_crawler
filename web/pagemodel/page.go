package pagemodel

type SearchResult struct {
	Hits int
	Start int
	PrevFrom int
	NextFrom int
	Query string
	Items []interface{}
	//Items []model.Post
}

