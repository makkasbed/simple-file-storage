package models

type Account struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	Password     string `json:"password"`
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
}
