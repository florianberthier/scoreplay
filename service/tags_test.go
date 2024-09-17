package service

import (
	"scoreplay/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNewTag(t *testing.T) {
	s := SetupMock(t)

	t.Run("should return error on invalid request", func(t *testing.T) {
		req := models.CreateTagRequest{}
		sErr := s.CreateNewTag(req)
		require.NotNil(t, sErr)

		s.CleanupMock()
	})

	t.Run("should return success on valid request", func(t *testing.T) {
		req := models.CreateTagRequest{
			Name: "test",
		}
		sErr := s.CreateNewTag(req)
		require.Nil(t, sErr)

		tags, sErr := s.RetrieveTags()
		require.Nil(t, sErr)
		require.Equal(t, 1, len(tags))
		require.Equal(t, "test", tags[0].Name)

		s.CleanupMock()
	})

	t.Run("should return success on valid request with multiple tags", func(t *testing.T) {
		req := models.CreateTagRequest{
			Name: "test",
		}
		sErr := s.CreateNewTag(req)
		require.Nil(t, sErr)

		req2 := models.CreateTagRequest{
			Name: "test2",
		}
		sErr = s.CreateNewTag(req2)
		require.Nil(t, sErr)

		tags, sErr := s.RetrieveTags()
		require.Nil(t, sErr)
		require.Equal(t, 2, len(tags))
		require.Equal(t, "test", tags[0].Name)
		require.Equal(t, "test2", tags[1].Name)

		s.CleanupMock()
	})
}

func TestRetrieveTags(t *testing.T) {
	s := SetupMock(t)

	t.Run("should return success with no tags", func(t *testing.T) {
		tags, sErr := s.RetrieveTags()
		require.Nil(t, sErr)
		require.Equal(t, 0, len(tags))
	})

	t.Run("should return success with 1 tag", func(t *testing.T) {
		req := models.CreateTagRequest{
			Name: "test",
		}
		sErr := s.CreateNewTag(req)
		require.Nil(t, sErr)

		tags, sErr := s.RetrieveTags()
		require.Nil(t, sErr)
		require.Equal(t, 1, len(tags))
		require.Equal(t, "test", tags[0].Name)

		s.CleanupMock()
	})

	t.Run("should return success with 2 tags", func(t *testing.T) {
		req := models.CreateTagRequest{
			Name: "test",
		}
		sErr := s.CreateNewTag(req)
		require.Nil(t, sErr)

		req2 := models.CreateTagRequest{
			Name: "test2",
		}
		sErr = s.CreateNewTag(req2)
		require.Nil(t, sErr)

		tags, sErr := s.RetrieveTags()
		require.Nil(t, sErr)
		require.Equal(t, 2, len(tags))
		require.Equal(t, "test", tags[0].Name)
		require.Equal(t, "test2", tags[1].Name)

		s.CleanupMock()
	})
}
