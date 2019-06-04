// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// GetSubnetsIDReader is a Reader for the GetSubnetsID structure.
type GetSubnetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubnetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubnetsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSubnetsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSubnetsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetSubnetsIDOK creates a GetSubnetsIDOK with default headers values
func NewGetSubnetsIDOK() *GetSubnetsIDOK {
	return &GetSubnetsIDOK{}
}

/*GetSubnetsIDOK handles this case with default header values.

dummy
*/
type GetSubnetsIDOK struct {
	Payload *models.Subnet
}

func (o *GetSubnetsIDOK) Error() string {
	return fmt.Sprintf("[GET /subnets/{id}][%d] getSubnetsIdOK  %+v", 200, o.Payload)
}

func (o *GetSubnetsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subnet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubnetsIDNotFound creates a GetSubnetsIDNotFound with default headers values
func NewGetSubnetsIDNotFound() *GetSubnetsIDNotFound {
	return &GetSubnetsIDNotFound{}
}

/*GetSubnetsIDNotFound handles this case with default header values.

error
*/
type GetSubnetsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetSubnetsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /subnets/{id}][%d] getSubnetsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSubnetsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubnetsIDDefault creates a GetSubnetsIDDefault with default headers values
func NewGetSubnetsIDDefault(code int) *GetSubnetsIDDefault {
	return &GetSubnetsIDDefault{
		_statusCode: code,
	}
}

/*GetSubnetsIDDefault handles this case with default header values.

unexpectederror
*/
type GetSubnetsIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get subnets ID default response
func (o *GetSubnetsIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSubnetsIDDefault) Error() string {
	return fmt.Sprintf("[GET /subnets/{id}][%d] GetSubnetsID default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubnetsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
