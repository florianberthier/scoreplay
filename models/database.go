package models

type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type Media struct {
	ID        string `gorm:"primaryKey;default:uuid_generate_v4()"`
	Name      string `gorm:"not null"`
	File      []byte `gorm:"type:bytea"`
	Extension string `gorm:"not null"`
	Tags      []Tag  `gorm:"many2many:tag_media;"`
}
