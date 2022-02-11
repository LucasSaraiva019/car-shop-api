package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/LucasSaraiva019/car-shop-api/client"
	"github.com/LucasSaraiva019/car-shop-api/client/car"
	"github.com/LucasSaraiva019/car-shop-api/client/models"
)

func TestFindAllCars(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8000")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := car.NewFindAllCarsParams()
	car, err := c.Car.FindAllCars(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", car.GetPayload())
}

func TestCreateCar(t *testing.T) {
	name := "name"
	cfg := client.DefaultTransportConfig().WithHost("localhost:8000")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := car.NewCreateCarParams().WithBody(&models.Cars{Description: "", Name: &name, PurchasePrice: 123, SalePrice: 1231})
	car, err := c.Car.CreateCar(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", car.GetPayload())
}
