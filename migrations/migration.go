package migrations

import (
	"github.com/ManuelGarreRos/templNotes/entitites"
	"gorm.io/gorm"
)

func MigrateSchema(db *gorm.DB) {
	err := db.AutoMigrate(
		&entitites.Note{},
	)
	if err != nil {
		return
	}
}
