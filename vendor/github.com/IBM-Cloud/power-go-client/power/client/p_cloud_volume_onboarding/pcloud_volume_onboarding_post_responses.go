// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_onboarding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVolumeOnboardingPostReader is a Reader for the PcloudVolumeOnboardingPost structure.
type PcloudVolumeOnboardingPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumeOnboardingPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudVolumeOnboardingPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumeOnboardingPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumeOnboardingPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumeOnboardingPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudVolumeOnboardingPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumeOnboardingPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVolumeOnboardingPostAccepted creates a PcloudVolumeOnboardingPostAccepted with default headers values
func NewPcloudVolumeOnboardingPostAccepted() *PcloudVolumeOnboardingPostAccepted {
	return &PcloudVolumeOnboardingPostAccepted{}
}

/* PcloudVolumeOnboardingPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudVolumeOnboardingPostAccepted struct {
	Payload *models.VolumeOnboardingCreateResponse
}

func (o *PcloudVolumeOnboardingPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingPostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudVolumeOnboardingPostAccepted) GetPayload() *models.VolumeOnboardingCreateResponse {
	return o.Payload
}

func (o *PcloudVolumeOnboardingPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeOnboardingCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingPostBadRequest creates a PcloudVolumeOnboardingPostBadRequest with default headers values
func NewPcloudVolumeOnboardingPostBadRequest() *PcloudVolumeOnboardingPostBadRequest {
	return &PcloudVolumeOnboardingPostBadRequest{}
}

/* PcloudVolumeOnboardingPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumeOnboardingPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVolumeOnboardingPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudVolumeOnboardingPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingPostUnauthorized creates a PcloudVolumeOnboardingPostUnauthorized with default headers values
func NewPcloudVolumeOnboardingPostUnauthorized() *PcloudVolumeOnboardingPostUnauthorized {
	return &PcloudVolumeOnboardingPostUnauthorized{}
}

/* PcloudVolumeOnboardingPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumeOnboardingPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVolumeOnboardingPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudVolumeOnboardingPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingPostForbidden creates a PcloudVolumeOnboardingPostForbidden with default headers values
func NewPcloudVolumeOnboardingPostForbidden() *PcloudVolumeOnboardingPostForbidden {
	return &PcloudVolumeOnboardingPostForbidden{}
}

/* PcloudVolumeOnboardingPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumeOnboardingPostForbidden struct {
	Payload *models.Error
}

func (o *PcloudVolumeOnboardingPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingPostForbidden  %+v", 403, o.Payload)
}
func (o *PcloudVolumeOnboardingPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingPostConflict creates a PcloudVolumeOnboardingPostConflict with default headers values
func NewPcloudVolumeOnboardingPostConflict() *PcloudVolumeOnboardingPostConflict {
	return &PcloudVolumeOnboardingPostConflict{}
}

/* PcloudVolumeOnboardingPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudVolumeOnboardingPostConflict struct {
	Payload *models.Error
}

func (o *PcloudVolumeOnboardingPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudVolumeOnboardingPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingPostInternalServerError creates a PcloudVolumeOnboardingPostInternalServerError with default headers values
func NewPcloudVolumeOnboardingPostInternalServerError() *PcloudVolumeOnboardingPostInternalServerError {
	return &PcloudVolumeOnboardingPostInternalServerError{}
}

/* PcloudVolumeOnboardingPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumeOnboardingPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVolumeOnboardingPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudVolumeOnboardingPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}