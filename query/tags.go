package query

import "scoreplay/models"

func (q *Query) CreateTag(name string) error {
	tag := models.Tag{
		Name: name,
	}

	result := q.DB.Table(q.Tables.Tags).Create(&tag)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (d *Query) GetTags() ([]models.Tag, error) {
	tags := []models.Tag{}
	err := d.DB.Table(d.Tables.Tags).Find(&tags).Error
	return tags, err
}

func (q *Query) GetTagByID(id int) (models.Tag, error) {
	tag := models.Tag{}
	err := q.DB.Table(q.Tables.Tags).Where("id = ?", id).First(&tag).Error
	return tag, err
}
