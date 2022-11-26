package models

type Bucket struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Status      string `json:"status"`
	Url         string `json:"url"`
	Region      string `json:"region"`
	CreatedBy   string `json:"account_id"`
	BucketType  string `json:"bucket_type"`
}
