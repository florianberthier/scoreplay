package service

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"scoreplay/models"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func GetFileContent(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func TestCreateNewMedia(t *testing.T) {
	s := SetupMock(t)

	validFile, err := GetFileContent("../statics/scoreplay.png")
	if err != nil {
		t.Fatal("Error reading file")
	}

	invalidFile, err := GetFileContent("../statics/blank.pdf")
	if err != nil {
		t.Fatal("Error reading file")
	}

	t.Run("should return error on invalid request", func(t *testing.T) {
		req := models.CreateMediaRequest{}
		err := s.CreateNewMedia(req)

		require.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.Code)
	})

	t.Run("should return error on invalid request, missing name", func(t *testing.T) {
		req := models.CreateMediaRequest{
			Tags: []string{
				"10",
			},
			File: validFile,
		}
		err := s.CreateNewMedia(req)

		require.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.Code)
	})

	t.Run("should return error on invalid request, invalid file", func(t *testing.T) {
		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"10",
			},
			File: invalidFile,
		}
		err := s.CreateNewMedia(req)

		require.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.Code)
		assert.Equal(t, "invalid image file", err.Message)
	})

	t.Run("should return success, media without tags", func(t *testing.T) {
		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"10",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)

		require.Nil(t, sErr)

		media, err := s.DB.GetMedia()
		require.Nil(t, err)

		assert.Equal(t, 1, len(media))
		assert.Equal(t, "test", media[0].Name)
		assert.Equal(t, "image/png", media[0].Extension)
		assert.Equal(t, 0, len(media[0].Tags))

		s.CleanupMock()
	})

	t.Run("should return success, media with 1 tag", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)

		require.Nil(t, sErr)

		media, err := s.DB.GetMedia()
		require.Nil(t, err)

		assert.Equal(t, 1, len(media))
		assert.Equal(t, "test", media[0].Name)
		assert.Equal(t, "image/png", media[0].Extension)
		assert.Equal(t, 1, len(media[0].Tags))
		assert.Equal(t, "test", media[0].Tags[0].Name)

		s.CleanupMock()
	})

	t.Run("should return success, media with 1 tag and 1 invalid tag", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		tags, err := s.DB.GetTags()
		require.Nil(t, err)
		assert.Equal(t, 1, len(tags))

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
				"10",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)

		require.Nil(t, sErr)

		media, err := s.DB.GetMedia()
		require.Nil(t, err)

		assert.Equal(t, 1, len(media))
		assert.Equal(t, "test", media[0].Name)
		assert.Equal(t, "image/png", media[0].Extension)
		assert.Equal(t, 1, len(media[0].Tags))
		assert.Equal(t, "test", media[0].Tags[0].Name)

		s.CleanupMock()
	})

	t.Run("should return success, media with 2 tags", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		err = s.DB.CreateTag("test2")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
				"2",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)

		require.Nil(t, sErr)

		media, err := s.DB.GetMedia()
		require.Nil(t, err)

		assert.Equal(t, 1, len(media))
		assert.Equal(t, "test", media[0].Name)
		assert.Equal(t, "image/png", media[0].Extension)
		assert.Equal(t, 2, len(media[0].Tags))
		assert.Equal(t, "test", media[0].Tags[0].Name)
		assert.Equal(t, "test2", media[0].Tags[1].Name)

		s.CleanupMock()
	})

	t.Run("should return success, 2 medias", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)
		require.Nil(t, sErr)

		sErr = s.CreateNewMedia(req)
		require.Nil(t, sErr)

		media, err := s.DB.GetMedia()
		require.Nil(t, err)

		assert.Equal(t, 2, len(media))

		assert.Equal(t, "test", media[0].Name)
		assert.Equal(t, "image/png", media[0].Extension)
		assert.Equal(t, 1, len(media[0].Tags))
		assert.Equal(t, "test", media[0].Tags[0].Name)

		assert.Equal(t, "test", media[1].Name)
		assert.Equal(t, "image/png", media[1].Extension)
		assert.Equal(t, 1, len(media[1].Tags))
		assert.Equal(t, "test", media[1].Tags[0].Name)

		s.CleanupMock()
	})
}

func TestRetrieveMedia(t *testing.T) {
	s := SetupMock(t)

	validFile, err := GetFileContent("../statics/scoreplay.png")
	if err != nil {
		t.Fatal("Error reading file")
	}

	t.Run("should return success with no media", func(t *testing.T) {
		media, err := s.RetrieveMedia("")
		require.Nil(t, err)
		assert.Equal(t, 0, len(media))
	})

	t.Run("should return success with 1 media", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)
		require.Nil(t, sErr)

		media, sErr := s.RetrieveMedia("")
		require.Nil(t, sErr)
		assert.Equal(t, 1, len(media))
		assert.Equal(t, "test", media[0].Name)

		s.CleanupMock()
	})

	t.Run("should return success with 2 media", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)
		require.Nil(t, sErr)

		req2 := models.CreateMediaRequest{
			Name: "test2",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr = s.CreateNewMedia(req2)
		require.Nil(t, sErr)

		media, sErr := s.RetrieveMedia("")
		require.Nil(t, sErr)
		assert.Equal(t, 2, len(media))
		assert.Equal(t, "test", media[0].Name)
		assert.Equal(t, "test2", media[1].Name)

		s.CleanupMock()
	})

	t.Run("should return success with 1 media with tag search", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		err = s.DB.CreateTag("test2")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)
		require.Nil(t, sErr)

		req2 := models.CreateMediaRequest{
			Name: "test2",
			Tags: []string{
				"2",
			},
			File: validFile,
		}
		sErr = s.CreateNewMedia(req2)
		require.Nil(t, sErr)

		media, sErr := s.RetrieveMedia("1")
		require.Nil(t, sErr)
		assert.Equal(t, 1, len(media))
		assert.Equal(t, "test", media[0].Name)

		s.CleanupMock()
	})

	t.Run("should return success with 0 media with tag search", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		err = s.DB.CreateTag("test2")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)
		require.Nil(t, sErr)

		req2 := models.CreateMediaRequest{
			Name: "test2",
			Tags: []string{
				"2",
			},
			File: validFile,
		}
		sErr = s.CreateNewMedia(req2)
		require.Nil(t, sErr)

		media, sErr := s.RetrieveMedia("3")
		require.Nil(t, sErr)
		assert.Equal(t, 0, len(media))

		s.CleanupMock()
	})

	t.Run("should return success with 2 medias with tag search", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		err = s.DB.CreateTag("test2")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)
		require.Nil(t, sErr)

		req2 := models.CreateMediaRequest{
			Name: "test2",
			Tags: []string{
				"2",
				"1",
			},
			File: validFile,
		}
		sErr = s.CreateNewMedia(req2)
		require.Nil(t, sErr)

		media, sErr := s.RetrieveMedia("1")
		require.Nil(t, sErr)
		assert.Equal(t, 2, len(media))
		assert.Equal(t, "test", media[0].Name)
		assert.Equal(t, "test2", media[1].Name)

		s.CleanupMock()
	})
}

func TestRetrieveMediaFile(t *testing.T) {
	s := SetupMock(t)

	validFile, err := GetFileContent("../statics/scoreplay.png")
	if err != nil {
		t.Fatal("Error reading file")
	}

	t.Run("should return error on unknown media", func(t *testing.T) {
		_, _, err := s.RetrieveMediaFile("1")

		require.NotNil(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.Code)
	})

	t.Run("should return success", func(t *testing.T) {
		err := s.DB.CreateTag("test")
		require.Nil(t, err)

		req := models.CreateMediaRequest{
			Name: "test",
			Tags: []string{
				"1",
			},
			File: validFile,
		}
		sErr := s.CreateNewMedia(req)
		require.Nil(t, sErr)

		media, err := s.DB.GetMedia()
		require.Nil(t, err)
		assert.Equal(t, 1, len(media))

		file, extension, sErr := s.RetrieveMediaFile(media[0].ID)
		require.Nil(t, sErr)
		assert.Equal(t, "image/png", extension)

		areEqual := bytes.Equal(file, validFile)
		assert.Equal(t, true, areEqual)

		s.CleanupMock()
	})
}
