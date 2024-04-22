package typedef

type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type ModelResponse struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}
