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

// CreateCarReader is a Reader for the CreateCar structure.
type CreateCarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCarCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateCarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCarCreated creates a CreateCarCreated with default headers values
func NewCreateCarCreated() *CreateCarCreated {
	return &CreateCarCreated{}
}

/* CreateCarCreated describes a response with status code 201, with default header values.

Created
*/
type CreateCarCreated struct {
	Payload *models.Cars
}

func (o *CreateCarCreated) Error() string {
	return fmt.Sprintf("[POST /cars][%d] createCarCreated  %+v", 201, o.Payload)
}
func (o *CreateCarCreated) GetPayload() *models.Cars {
	return o.Payload
}

func (o *CreateCarCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cars)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCarDefault creates a CreateCarDefault with default headers values
func NewCreateCarDefault(code int) *CreateCarDefault {
	return &CreateCarDefault{
		_statusCode: code,
	}
}

/* CreateCarDefault describes a response with status code -1, with default header values.

Generic erro responses
*/
type CreateCarDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create car default response
func (o *CreateCarDefault) Code() int {
	return o._statusCode
}

func (o *CreateCarDefault) Error() string {
	return fmt.Sprintf("[POST /cars][%d] create_car default  %+v", o._statusCode, o.Payload)
}
func (o *CreateCarDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
