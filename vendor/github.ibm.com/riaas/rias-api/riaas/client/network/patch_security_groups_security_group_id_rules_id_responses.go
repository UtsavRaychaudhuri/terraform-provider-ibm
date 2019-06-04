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

// PatchSecurityGroupsSecurityGroupIDRulesIDReader is a Reader for the PatchSecurityGroupsSecurityGroupIDRulesID structure.
type PatchSecurityGroupsSecurityGroupIDRulesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSecurityGroupsSecurityGroupIDRulesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPatchSecurityGroupsSecurityGroupIDRulesIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchSecurityGroupsSecurityGroupIDRulesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSecurityGroupsSecurityGroupIDRulesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPatchSecurityGroupsSecurityGroupIDRulesIDCreated creates a PatchSecurityGroupsSecurityGroupIDRulesIDCreated with default headers values
func NewPatchSecurityGroupsSecurityGroupIDRulesIDCreated() *PatchSecurityGroupsSecurityGroupIDRulesIDCreated {
	return &PatchSecurityGroupsSecurityGroupIDRulesIDCreated{}
}

/*PatchSecurityGroupsSecurityGroupIDRulesIDCreated handles this case with default header values.

dummy
*/
type PatchSecurityGroupsSecurityGroupIDRulesIDCreated struct {
	Payload *models.SecurityGroupRule
}

func (o *PatchSecurityGroupsSecurityGroupIDRulesIDCreated) Error() string {
	return fmt.Sprintf("[PATCH /security_groups/{security_group_id}/rules/{id}][%d] patchSecurityGroupsSecurityGroupIdRulesIdCreated  %+v", 201, o.Payload)
}

func (o *PatchSecurityGroupsSecurityGroupIDRulesIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSecurityGroupsSecurityGroupIDRulesIDBadRequest creates a PatchSecurityGroupsSecurityGroupIDRulesIDBadRequest with default headers values
func NewPatchSecurityGroupsSecurityGroupIDRulesIDBadRequest() *PatchSecurityGroupsSecurityGroupIDRulesIDBadRequest {
	return &PatchSecurityGroupsSecurityGroupIDRulesIDBadRequest{}
}

/*PatchSecurityGroupsSecurityGroupIDRulesIDBadRequest handles this case with default header values.

error
*/
type PatchSecurityGroupsSecurityGroupIDRulesIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchSecurityGroupsSecurityGroupIDRulesIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /security_groups/{security_group_id}/rules/{id}][%d] patchSecurityGroupsSecurityGroupIdRulesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSecurityGroupsSecurityGroupIDRulesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSecurityGroupsSecurityGroupIDRulesIDDefault creates a PatchSecurityGroupsSecurityGroupIDRulesIDDefault with default headers values
func NewPatchSecurityGroupsSecurityGroupIDRulesIDDefault(code int) *PatchSecurityGroupsSecurityGroupIDRulesIDDefault {
	return &PatchSecurityGroupsSecurityGroupIDRulesIDDefault{
		_statusCode: code,
	}
}

/*PatchSecurityGroupsSecurityGroupIDRulesIDDefault handles this case with default header values.

unexpectederror
*/
type PatchSecurityGroupsSecurityGroupIDRulesIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the patch security groups security group ID rules ID default response
func (o *PatchSecurityGroupsSecurityGroupIDRulesIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSecurityGroupsSecurityGroupIDRulesIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /security_groups/{security_group_id}/rules/{id}][%d] PatchSecurityGroupsSecurityGroupIDRulesID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSecurityGroupsSecurityGroupIDRulesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
