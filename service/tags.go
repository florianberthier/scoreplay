package service

import (
	"fmt"
	"net/http"
	"scoreplay/models"

	"github.com/go-playground/validator/v10"
)

func (s *Service) CreateNewTag(request models.CreateTagRequest) *models.CustomError {
	err := s.Validator.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return models.NewCustomError(errors, http.StatusBadRequest)
	}

	if err := s.DB.CreateTag(request.Name); err != nil {
		return models.NewCustomError(err, http.StatusBadRequest)
	}

	return nil
}

func (s *Service) RetrieveTags() ([]models.RetrieveTagsResponse, *models.CustomError) {
	tags, err := s.DB.GetTags()
	if err != nil {
		models.NewCustomError(err, http.StatusInternalServerError)
	}

	res := []models.RetrieveTagsResponse{}
	for _, tag := range tags {
		res = append(res, models.RetrieveTagsResponse{
			ID:   fmt.Sprintf("%d", tag.ID),
			Name: tag.Name,
		})
	}

	return res, nil
}
