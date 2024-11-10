package domain

type ListBreedResponse struct {
	ID        string  `json:"id"`
	NameEN    string  `json:"name_en"`
	NameTH    string  `json:"name_th"`
	ShortName string  `json:"short_name"`
	Remark    *string `json:"remark"`
}
