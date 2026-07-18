package seeder

import (
	"log"
	"ticketing-system/common/helper"
	"ticketing-system/config"
	"ticketing-system/entity"
	"time"

	"gorm.io/gorm"
)

func SeedAdminUsers(db *gorm.DB) error {
	password, err := helper.HashPassword(config.GetEnvOrPanic("AMIND_USER_PASSWORD"))

	if err != nil {
		log.Println(err)
		return err
	}

	now := time.Now()
	admins := []entity.User{
		{Name: "Super Admin", Email: "super_admin@gts.com", Password: password, Role: entity.SuperAdmin, VerifiedAt: &now, Gender: "Male"},
		{Name: "Super Admin", Email: "system_admin@gts.com", Password: password, Role: entity.SystemAdmin, VerifiedAt: &now, Gender: "Male"},
	}
	log.Println("Seeding admin users...")
	for _, admin := range admins {
		err = db.Where(entity.User{Email: admin.Email}).FirstOrCreate(&admin).Error
		if err != nil {
			return err
		}
	}
	log.Println("Seeding admin users...done")
	return nil
}
