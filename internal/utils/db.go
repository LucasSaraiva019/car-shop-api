package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Cars struct {
	gorm.Model    `json:"model"`
	Description   string  `json:"description,omitempty"`
	Name          *string `json:"name"`
	PurchasePrice float64 `json:"purchasePrice,omitempty"`
	SalePrice     float64 `json:"salePrice,omitempty"`
}

func ConnectDB() *gorm.DB {
	databaseDetails := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBNAME"), os.Getenv("SSLMODE"))
	gormDB, err := gorm.Open(postgres.Open(databaseDetails), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := gormDB.AutoMigrate(Cars{}); err != nil {
		panic(err)
	}

	return gormDB
}
