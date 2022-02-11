package controllers

import (
	"testing"

	"github.com/LucasSaraiva019/car-shop-api/gen/models"
	"github.com/LucasSaraiva019/car-shop-api/gen/resources"
	cars "github.com/LucasSaraiva019/car-shop-api/gen/restapi/operations/car"
	"github.com/LucasSaraiva019/car-shop-api/internal/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-test/deep"
)

func TestGetAllCars(t *testing.T) {
	carResource := resources.NewCarResourcesGormDB(utils.ConnectDB())
	carController := NewCarController(carResource)
	car := []*models.Cars{{}}

	type args struct {
		controller CarController
	}

	tests := []struct {
		name   string
		args   args
		params cars.FindAllCarsParams
		want   middleware.Responder
	}{
		{
			name:   "succes",
			args:   args{controller: *carController},
			params: cars.FindAllCarsParams{},
			want:   cars.NewFindAllCarsOK().WithPayload(car),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCarController(carResource).GetAll()
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("GetDeviceFiltersHandler() = %v\n want %v\n diff = %v", got, tt.want, diff)
			}
		})
	}
}
