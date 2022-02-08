package resources

import "github.com/LucasSaraiva019/car-shop-api/gen/models"

type BaseResource interface {
	Create(interface{}) error
	GetAll() ([]*models.Cars, error)
}
