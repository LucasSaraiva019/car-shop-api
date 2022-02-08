package resources

import "gorm.io/gorm"

const (
	CarResourceName = "Product"
)

type CarResourcesGormDB struct {
	AbstractResourceGormDB
}

func NewCarResourcesGormDB(dbProvider *gorm.DB) *CarResourcesGormDB {
	return &CarResourcesGormDB{
		AbstractResourceGormDB: AbstractResourceGormDB{
			dbProvider:   dbProvider,
			resourceName: CarResourceName,
		},
	}
}
