package models

type File struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Bucket       string `json:"bucket_id"`
	CreatedAt    string `json:"created_at"`
	Public       string `json:"public"`
	LastAccess   string `json:"last_access"`
	FileType     string `json:"file_type"`
	FileLocation string `json:"file_location"`
}
