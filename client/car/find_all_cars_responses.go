// Code generated by go-swagger; DO NOT EDIT.

package car

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/LucasSaraiva019/car-shop-api/client/models"
)

// FindAllCarsReader is a Reader for the FindAllCars structure.
type FindAllCarsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindAllCarsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindAllCarsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindAllCarsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindAllCarsOK creates a FindAllCarsOK with default headers values
func NewFindAllCarsOK() *FindAllCarsOK {
	return &FindAllCarsOK{}
}

/* FindAllCarsOK describes a response with status code 200, with default header values.

list all cars
*/
type FindAllCarsOK struct {
	Payload []*models.Cars
}

func (o *FindAllCarsOK) Error() string {
	return fmt.Sprintf("[GET /cars][%d] findAllCarsOK  %+v", 200, o.Payload)
}
func (o *FindAllCarsOK) GetPayload() []*models.Cars {
	return o.Payload
}

func (o *FindAllCarsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindAllCarsDefault creates a FindAllCarsDefault with default headers values
func NewFindAllCarsDefault(code int) *FindAllCarsDefault {
	return &FindAllCarsDefault{
		_statusCode: code,
	}
}

/* FindAllCarsDefault describes a response with status code -1, with default header values.

Generic error responses
*/
type FindAllCarsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find all cars default response
func (o *FindAllCarsDefault) Code() int {
	return o._statusCode
}

func (o *FindAllCarsDefault) Error() string {
	return fmt.Sprintf("[GET /cars][%d] find_all_cars default  %+v", o._statusCode, o.Payload)
}
func (o *FindAllCarsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindAllCarsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
