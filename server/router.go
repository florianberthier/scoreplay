package server

import (
	"encoding/json"
	"io"
	"net/http"
	"scoreplay/models"
	"scoreplay/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	s := service.Setup()

	r.POST("/tags", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		var request models.CreateTagRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		sErr := s.CreateNewTag(request)
		if sErr != nil {
			c.JSON(sErr.Code, sErr.Message)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "New Tag Created",
		})
	})

	r.GET("/tags", func(c *gin.Context) {
		tags, sErr := s.RetrieveTags()
		if sErr != nil {
			c.JSON(sErr.Code, sErr.Message)
			return
		}

		c.JSON(http.StatusOK, tags)
	})

	r.POST("/media", func(c *gin.Context) {
		name := c.Request.FormValue("name")
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		tagsJSON := c.Request.FormValue("tags")

		var tags []string
		err = json.Unmarshal([]byte(tagsJSON), &tags)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		request := models.CreateMediaRequest{
			Name: name,
			File: fileBytes,
			Tags: tags,
		}

		sErr := s.CreateNewMedia(request)
		if sErr != nil {
			c.JSON(sErr.Code, sErr.Message)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "New Media Created",
		})
	})

	r.GET("/media", func(c *gin.Context) {
		tag := c.Request.URL.Query().Get("tag")

		media, sErr := s.RetrieveMedia(tag)
		if sErr != nil {
			c.JSON(sErr.Code, sErr.Message)
			return
		}

		c.JSON(http.StatusOK, media)
	})

	r.GET("/media/file/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		file, extension, sErr := s.RetrieveMediaFile(id)
		if sErr != nil {
			c.JSON(sErr.Code, sErr.Message)
			return
		}

		c.Data(http.StatusOK, extension, file)
	})
}
