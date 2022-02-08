package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LucasSaraiva019/car-shop-api/gen/models"
	"github.com/LucasSaraiva019/car-shop-api/gen/resources"
	"github.com/LucasSaraiva019/car-shop-api/gen/restapi/operations/car"
	cars "github.com/LucasSaraiva019/car-shop-api/gen/restapi/operations/car"
	"github.com/go-openapi/runtime/middleware"
)

type CarController struct {
	resource resources.CarResources
}

func NewCarController(resource resources.CarResources) *CarController {
	return &CarController{resource}
}

func (c *CarController) GetAll() middleware.Responder {
	cars, err := c.resource.GetAll()
	if err != nil {
		return middleware.Error(http.StatusNotFound, err)
	}
	return car.NewFindAllCarsOK().WithPayload(cars)
}

func (c *CarController) Create(params car.CreateCarParams) middleware.Responder {
	body, err := params.Body.MarshalBinary()
	if err != nil {
		return middleware.Error(http.StatusUnprocessableEntity, err)
	}

	var car *models.Cars
	if err = json.Unmarshal(body, &car); err != nil {
		return middleware.Error(http.StatusBadRequest, err)
	}
	if err := c.resource.Create(params.Body); err != nil {
		return middleware.Error(http.StatusBadRequest, err)
	}

	return cars.NewCreateCarCreated().WithPayload(car)
}
