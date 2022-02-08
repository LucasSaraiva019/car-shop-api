package resources

import (
	"github.com/LucasSaraiva019/car-shop-api/gen/models"
	"gorm.io/gorm"
)

type AbstractResourceGormDB struct {
	dbProvider   *gorm.DB
	resourceName string
}

func (r *AbstractResourceGormDB) Create(value interface{}) error {
	tx := r.dbProvider.Create(value)

	return tx.Error
}

func (r *AbstractResourceGormDB) GetAll() ([]*models.Cars, error) {
	var cars []*models.Cars
	tx := r.dbProvider.Find(&cars)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return cars, nil
}
