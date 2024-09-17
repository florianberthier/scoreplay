package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTag(t *testing.T) {
	q := SetupMock(t)

	t.Run("should return success with 1 tags", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 1, len(tags))
		require.Equal(t, "test", tags[0].Name)

		q.CleanupMock()
	})

	t.Run("should return success with multiple tags", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		err = q.CreateTag("test2")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 2, len(tags))

		require.Equal(t, "test", tags[0].Name)
		require.Equal(t, "test2", tags[1].Name)

		q.CleanupMock()
	})

}

func TestGetTags(t *testing.T) {
	q := SetupMock(t)

	t.Run("should return success with no tags", func(t *testing.T) {
		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 0, len(tags))
	})

	t.Run("should return success with 1 tags", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 1, len(tags))
		require.Equal(t, "test", tags[0].Name)

		q.CleanupMock()
	})

	t.Run("should return success with multiple tags", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		err = q.CreateTag("test2")
		require.Nil(t, err)

		tags, err := q.GetTags()
		require.Nil(t, err)
		require.Equal(t, 2, len(tags))

		require.Equal(t, "test", tags[0].Name)
		require.Equal(t, "test2", tags[1].Name)

		q.CleanupMock()
	})
}

func TestGetTagByID(t *testing.T) {
	q := SetupMock(t)

	t.Run("should return error with no tags", func(t *testing.T) {
		_, err := q.GetTagByID(1)
		require.NotNil(t, err)
	})

	t.Run("should return success with 1 tags", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		tag, err := q.GetTagByID(1)
		require.Nil(t, err)
		require.Equal(t, "test", tag.Name)

		q.CleanupMock()
	})

	t.Run("should return success with multiple tags", func(t *testing.T) {
		err := q.CreateTag("test")
		require.Nil(t, err)

		err = q.CreateTag("test2")
		require.Nil(t, err)

		tag, err := q.GetTagByID(2)
		require.Nil(t, err)
		require.Equal(t, "test2", tag.Name)

		q.CleanupMock()
	})
}
