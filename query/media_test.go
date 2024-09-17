package query

import (
	"scoreplay/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMedia(t *testing.T) {
	q := SetupMock(t)

	t.Run("should return success with 1 media", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 1, len(tags))
		require.Equal(t, "test", tags[0].Name)

		err = q.CreateMedia("test", "image/png", []byte{}, []models.Tag{
			{
				Name: "test",
				ID:   tags[0].ID,
			},
		})
		require.Nil(t, err)

		media, err := q.GetMedia()
		require.Nil(t, err)
		require.Equal(t, 1, len(media))
		require.Equal(t, "test", media[0].Name)
		require.Equal(t, "image/png", media[0].Extension)
		require.Equal(t, 1, len(media[0].Tags))
		require.Equal(t, "test", media[0].Tags[0].Name)

		q.CleanupMock()
	})

	t.Run("should return success with 1 media with 0 tags", func(t *testing.T) {
		err := q.CreateMedia("test", "image/png", []byte{}, []models.Tag{
			{
				Name: "test",
				ID:   1,
			},
		})
		require.Nil(t, err)

		media, err := q.GetMedia()
		require.Nil(t, err)
		require.Equal(t, 1, len(media))
		require.Equal(t, "test", media[0].Name)
		require.Equal(t, "image/png", media[0].Extension)
		require.Equal(t, 1, len(media[0].Tags))
		require.Equal(t, "test", media[0].Tags[0].Name)

		q.CleanupMock()
	})
}

func TestGetMedia(t *testing.T) {
	q := SetupMock(t)

	t.Run("should return success with no media", func(t *testing.T) {
		media, err := q.GetMedia()
		require.Nil(t, err)
		require.Equal(t, 0, len(media))
	})

	t.Run("should return success with 1 media", func(t *testing.T) {
		err := q.CreateMedia("test", "image/png", []byte{}, []models.Tag{
			{
				Name: "test",
				ID:   1,
			},
		})
		require.Nil(t, err)

		media, err := q.GetMedia()
		require.Nil(t, err)
		require.Equal(t, 1, len(media))
		require.Equal(t, "test", media[0].Name)
		require.Equal(t, "image/png", media[0].Extension)
		require.Equal(t, 1, len(media[0].Tags))
		require.Equal(t, "test", media[0].Tags[0].Name)

		q.CleanupMock()
	})

	t.Run("should return success with 2 media", func(t *testing.T) {
		err := q.CreateMedia("test", "image/png", []byte{}, []models.Tag{
			{
				Name: "test",
				ID:   1,
			},
		})
		require.Nil(t, err)

		err = q.CreateMedia("test2", "image/png", []byte{}, []models.Tag{
			{
				Name: "test",
				ID:   1,
			},
		})
		require.Nil(t, err)

		media, err := q.GetMedia()
		require.Nil(t, err)
		require.Equal(t, 2, len(media))

		require.Equal(t, "test", media[0].Name)
		require.Equal(t, "test2", media[1].Name)

		q.CleanupMock()
	})
}

func TestGetMediaByTag(t *testing.T) {
	q := SetupMock(t)

	t.Run("should return success with 1 media", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 1, len(tags))
		require.Equal(t, "test", tags[0].Name)

		err = q.CreateMedia("test", "image/png", []byte{}, []models.Tag{
			{
				Name: "test",
				ID:   tags[0].ID,
			},
		})
		require.Nil(t, err)

		media, err := q.GetMediaByTag("1")
		require.Nil(t, err)
		require.Equal(t, 1, len(media))
		require.Equal(t, "test", media[0].Name)
		require.Equal(t, "image/png", media[0].Extension)
		require.Equal(t, 1, len(media[0].Tags))
		require.Equal(t, "test", media[0].Tags[0].Name)

		q.CleanupMock()
	})

	t.Run("should return success with 2 media", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		err = q.CreateTag("test2")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 2, len(tags))

		err = q.CreateMedia("test", "image/png", []byte{}, []models.Tag{
			{
				Name: "test",
				ID:   tags[0].ID,
			},
		})
		require.Nil(t, err)

		err = q.CreateMedia("test2", "image/png", []byte{}, []models.Tag{
			{
				Name: "test2",
				ID:   tags[1].ID,
			},
		})
		require.Nil(t, err)

		media, err := q.GetMediaByTag("1")
		require.Nil(t, err)
		require.Equal(t, 1, len(media))
		require.Equal(t, "test", media[0].Name)
		require.Equal(t, "image/png", media[0].Extension)
		require.Equal(t, 1, len(media[0].Tags))
		require.Equal(t, "test", media[0].Tags[0].Name)

		media, err = q.GetMediaByTag("2")
		require.Nil(t, err)
		require.Equal(t, 1, len(media))
		require.Equal(t, "test2", media[0].Name)
		require.Equal(t, "image/png", media[0].Extension)
		require.Equal(t, 1, len(media[0].Tags))
		require.Equal(t, "test2", media[0].Tags[0].Name)

		q.CleanupMock()
	})
}

func TestGetMediaFileByID(t *testing.T) {
	q := SetupMock(t)

	t.Run("should return success with 1 media", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 1, len(tags))
		require.Equal(t, "test", tags[0].Name)

		err = q.CreateMedia("test", "image/png", []byte{1, 2, 3}, []models.Tag{
			{
				Name: "test",
				ID:   tags[0].ID,
			},
		})
		require.Nil(t, err)

		medias, err := q.GetMedia()
		require.Nil(t, err)
		require.Equal(t, 1, len(medias))

		media, err := q.GetMediaFileByID(medias[0].ID)
		require.Nil(t, err)
		require.Equal(t, "test", media.Name)
		require.Equal(t, "image/png", media.Extension)
		require.Equal(t, []byte{1, 2, 3}, media.File)
		require.Equal(t, 0, len(media.Tags))

		q.CleanupMock()
	})

	t.Run("should return success with 1 media with 0 tags", func(t *testing.T) {
		err := q.CreateMedia("test", "image/png", []byte{1, 2, 3}, []models.Tag{
			{
				Name: "test",
				ID:   1,
			},
		})
		require.Nil(t, err)

		medias, err := q.GetMedia()
		require.Nil(t, err)
		require.Equal(t, 1, len(medias))

		media, err := q.GetMediaFileByID(medias[0].ID)
		require.Nil(t, err)
		require.Equal(t, "test", media.Name)
		require.Equal(t, "image/png", media.Extension)
		require.Equal(t, []byte{1, 2, 3}, media.File)
		require.Equal(t, 0, len(media.Tags))

		q.CleanupMock()
	})

	t.Run("should return error with no media", func(t *testing.T) {
		_, err := q.GetMediaFileByID("1")
		require.NotNil(t, err)
	})
}
