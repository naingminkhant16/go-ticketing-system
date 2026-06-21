package migration

import (
	"ticketing-system/entity"

	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	migrateUserTable(db)
	migrateEventTables(db)
	migrateSeatsTable(db)
}

func migrateUserTable(db *gorm.DB) {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
		return
	}
}

func migrateEventTables(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Event{}, entity.EventDay{})
	if err != nil {
		panic(err)
		return
	}
}

func migrateSeatsTable(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Seat{})
	if err != nil {
		panic(err)
		return
	}
}
