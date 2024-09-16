package models

type RetrieveMediaResponse struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Tags    []string `json:"tags"`
	FileURL string   `json:"fileUrl"`
}

type RetrieveTagsResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
