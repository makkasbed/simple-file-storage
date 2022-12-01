package models

type Login struct {
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
}
