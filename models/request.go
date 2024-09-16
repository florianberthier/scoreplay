package models

type CreateTagRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateMediaRequest struct {
	Name string   `json:"name" validate:"required"`
	File []byte   `json:"file" validate:"required"`
	Tags []string `json:"tags" validate:"required"`
}
