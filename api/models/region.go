package models

type Region struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}
