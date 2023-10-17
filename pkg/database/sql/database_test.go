package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/mkdirJava/goPlay/internal/sql/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const database string = "test.sqlite"
const migrationDatabase string = "migration_test.sqlite"

func TestMain(m *testing.M) {
	m.Run()
	if dir, err := os.Getwd(); err != nil {
		panic(fmt.Sprintf("test cannot delete directory %v", err.Error()))
	} else {
		os.Remove(filepath.Join(dir, database))
		os.Remove(filepath.Join(dir, migrationDatabase))
	}
}

func TestMigrateDatabase(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "migration should happen",
		},
	}
	for _, tt := range tests {
		os.Setenv("DATABASE", migrationDatabase)
		t.Run(tt.name, func(t *testing.T) {
			MigrateDatabase()
			if db, err := gorm.Open(sqlite.Open(migrationDatabase), &gorm.Config{}); err != nil {
				panic("cannot open database")
			} else {
				migrator := db.Migrator()
				migrator.HasTable(&schema.Person{})
			}
		})
	}
}

func TestCRUD(t *testing.T) {
	t.Run("Should be able to crud", func(t *testing.T) {
		os.Setenv("DATABASE", database)
		MigrateDatabase()
		testData := []*schema.Person{{Name: "Bob"}, {Age: 2}}
		Save(testData)
		savedDataIds := getIdsFromModel(testData)
		for _, savedId := range savedDataIds {
			if savedId == 0 {
				t.Fail()
			}
		}
		valueTestData := pointerSliceToValue[schema.Person](testData)
		getResult := Get[schema.Person](savedDataIds)
		if reflect.DeepEqual(valueTestData, getResult) {
			t.Fail()
		}
		Delete(testData)
		if !areModelsDeleted(testData) {
			t.Fail()
		}
	})
}

func areModelsDeleted(testData []*ApplicationDomain) bool {
	values := Get[ApplicationDomain](getIdsFromModel(testData))
	for _, value := range values {
		if value.DeletedAt.Valid != true {
			return false
		}
	}
	return true
}

func pointerSliceToValue[T ApplicationDomain](data []*T) []T {
	valueTestData := make([]T, 0)
	for _, item := range data {
		valueTestData = append(valueTestData, *item)
	}
	return valueTestData
}

func getIdsFromModel(testData []*ApplicationDomain) []uint {
	savedDataIds := make([]uint, 0)
	for _, testItem := range testData {
		savedDataIds = append(savedDataIds, testItem.ID)
	}
	return savedDataIds
}
