package service

import (
	"bytes"
	"fmt"
	"image"
	"net/http"
	"scoreplay/env"
	"scoreplay/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	_ "image/jpeg"
	_ "image/png"
)

var (
	EXTENSION_MAP = map[string]string{
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"jpg":  "image/jpg",
	}
)

func getImageExtension(data []byte) (string, bool) {
	imgReader := bytes.NewReader(data)

	_, format, err := image.DecodeConfig(imgReader)
	if err != nil {
		return "", false
	}

	extension, ok := EXTENSION_MAP[format]
	return extension, ok
}

func (s *Service) CreateNewMedia(request models.CreateMediaRequest) *models.CustomError {
	err := s.Validator.Struct(request)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return models.NewCustomError(errors, http.StatusBadRequest)
	}

	extension, valid := getImageExtension(request.File)
	if !valid {
		return models.NewCustomError(fmt.Errorf("invalid image file"), http.StatusBadRequest)
	}

	tags := []models.Tag{}

	for _, tagID := range request.Tags {
		iTagID, err := strconv.Atoi(tagID)
		if err != nil {
			logrus.Info("Error converting tag ID to int: ", err)
			continue
		}

		tag, err := s.DB.GetTagByID(iTagID)
		if err != nil {
			logrus.Info("Error getting tag by ID: ", err)
			continue
		}

		tags = append(tags, tag)
	}

	if err := s.DB.CreateMedia(request.Name, extension, request.File, tags); err != nil {
		return models.NewCustomError(err, http.StatusInternalServerError)
	}

	return nil
}

func (s *Service) RetrieveMedia(tag string) ([]models.RetrieveMediaResponse, *models.CustomError) {
	medias := []models.Media{}
	err := error(nil)
	switch tag {
	case "":
		medias, err = s.DB.GetMedia()
		if err != nil {
			return nil, models.NewCustomError(err, http.StatusInternalServerError)
		}
	default:
		medias, err = s.DB.GetMediaByTag(tag)
		if err != nil {
			return nil, models.NewCustomError(err, http.StatusInternalServerError)
		}
	}

	resp := []models.RetrieveMediaResponse{}

	for _, media := range medias {
		tags := []string{}
		for _, tag := range media.Tags {
			tags = append(tags, tag.Name)
		}

		resp = append(resp, models.RetrieveMediaResponse{
			ID:      media.ID,
			Name:    media.Name,
			FileURL: env.API_URL + "/media/file/" + media.ID,
			Tags:    tags,
		})
	}

	return resp, nil
}

func (s *Service) RetrieveMediaFile(id string) ([]byte, string, *models.CustomError) {
	media, err := s.DB.GetMediaFileByID(id)
	if err != nil {
		return nil, "", models.NewCustomError(err, http.StatusInternalServerError)
	}

	return media.File, media.Extension, nil
}
