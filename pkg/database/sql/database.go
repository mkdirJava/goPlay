package pkg

import (
	"os"

	"github.com/mkdirJava/goPlay/internal/sql/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ApplicationDomain = schema.Person

func MigrateDatabase() {
	database := GetDatabase()
	// Choice to mannually describe the update or let GORM autoupdate
	database.AutoMigrate(&schema.Person{})
}

func GetDatabase() *gorm.DB {
	if database := os.Getenv("DATABASE"); database == "" {
		panic("DATABASE env should be set")
	}
	// TODO Support other databases sparing sqlite

	if db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{}); err != nil {
		panic("cannot open database")
	} else {
		return db
	}
}

func Save(items []*ApplicationDomain) {
	if len(items) != 0 {
		database := GetDatabase()
		for _, item := range items {
			database.Save(item)
		}
	}
}

func Get[T ApplicationDomain](ids []uint) []T {
	returningModels := make([]T, len(ids))
	if len(ids) != 0 {
		GetDatabase().Find(&returningModels, ids)
	}
	return returningModels
}

func Delete(models []*ApplicationDomain) {
	getIdFromModel := func(models []*ApplicationDomain) []uint {
		returningIds := make([]uint, 0)
		for _, model := range models {
			returningIds = append(returningIds, model.ID)
		}
		return returningIds
	}
	if len(models) != 0 {
		GetDatabase().Delete(&models, getIdFromModel(models))
	}
}
