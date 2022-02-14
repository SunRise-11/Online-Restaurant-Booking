package utils

import (
	"Restobook/configs"
	"Restobook/entities"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString :=
		fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			config.Database.Username,
			config.Database.Password,
			config.Database.Address,
			config.Database.Port,
			config.Database.Name,
		)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitialMigration(db)
	return db
}
func InitialMigration(db *gorm.DB) {

	db.Migrator().DropTable(&entities.Rating{})
	db.Migrator().DropTable(&entities.Transaction{})
	db.Migrator().DropTable(&entities.Restaurant{})
	db.Migrator().DropTable(&entities.RestaurantDetail{})
	db.Migrator().DropTable(&entities.TopUp{})
	db.Migrator().DropTable(&entities.User{})

	db.AutoMigrate(entities.User{})
	db.AutoMigrate(entities.TopUp{})
	db.AutoMigrate(entities.RestaurantDetail{})
	db.AutoMigrate(entities.Restaurant{})
	db.AutoMigrate(entities.Transaction{})
	db.AutoMigrate(entities.Rating{})

}
