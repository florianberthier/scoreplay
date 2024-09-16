package query

import (
	"fmt"
	"scoreplay/env"
	"scoreplay/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tables struct {
	Tags  string
	Media string
}

type Query struct {
	DB     *gorm.DB
	Tables Tables
}

func Setup() *Query {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env.POSTGRES_HOST, env.POSTGRES_PORT, env.POSTGRES_USER, env.POSTGRES_PASSWORD, env.POSTGRES_DB)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		logrus.Fatal("failed to enable uuid-ossp extension:", err)
	}

	if err := db.AutoMigrate(&models.Tag{}, &models.Media{}); err != nil {
		logrus.Fatal(err)
	}

	return &Query{
		DB: db,
		Tables: Tables{
			Tags:  "tags",
			Media: "media",
		},
	}
}
