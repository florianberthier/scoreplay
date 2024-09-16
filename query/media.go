package query

import (
	"scoreplay/models"
)

func (q *Query) CreateMedia(name, extension string, file []byte, tags []models.Tag) error {
	tag := models.Media{
		Name:      name,
		File:      file,
		Extension: extension,
	}

	result := q.DB.Table(q.Tables.Media).Create(&tag)
	if result.Error != nil {
		return result.Error
	}

	if err := q.DB.Table(q.Tables.Media).Model(&tag).Association("Tags").Append(tags); err != nil {
		return err
	}

	return nil
}

func (q *Query) GetMedia() ([]models.Media, error) {
	media := []models.Media{}
	err := q.DB.Preload("Tags").Table(q.Tables.Media).Find(&media).Error
	return media, err
}

func (q *Query) GetMediaByTag(tagID string) ([]models.Media, error) {
	media := []models.Media{}
	err := q.DB.Preload("Tags").
		Table(q.Tables.Media).
		Joins("JOIN tag_media ON media.id = tag_media.media_id").
		Joins("JOIN tags ON tag_media.tag_id = tags.id").
		Where("tags.id = ?", tagID).Find(&media).Error
	return media, err
}

func (q *Query) GetMediaFileByID(id string) (models.Media, error) {
	media := models.Media{}
	err := q.DB.Table(q.Tables.Media).Where("id = ?", id).First(&media).Error
	return media, err
}
