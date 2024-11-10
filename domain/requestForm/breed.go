package domain

type ListBreedRequest struct {
	IDs        []string `json:"ids"`
	Keyword    string   `json:"keyword"`
	ShortNames []string `json:"shortNames"`
}
