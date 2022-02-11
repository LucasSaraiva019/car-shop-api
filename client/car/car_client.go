// Code generated by go-swagger; DO NOT EDIT.

package car

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new car API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for car API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCar(params *CreateCarParams, opts ...ClientOption) (*CreateCarCreated, error)

	FindAllCars(params *FindAllCarsParams, opts ...ClientOption) (*FindAllCarsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCar create car API
*/
func (a *Client) CreateCar(params *CreateCarParams, opts ...ClientOption) (*CreateCarCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCarParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create_car",
		Method:             "POST",
		PathPattern:        "/cars",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCarReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCarCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateCarDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindAllCars find all cars API
*/
func (a *Client) FindAllCars(params *FindAllCarsParams, opts ...ClientOption) (*FindAllCarsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindAllCarsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "find_all_cars",
		Method:             "GET",
		PathPattern:        "/cars",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindAllCarsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindAllCarsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindAllCarsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}