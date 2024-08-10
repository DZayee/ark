// Code generated by go-swagger; DO NOT EDIT.

package ark_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ark-network/ark/pkg/client-sdk/client/rest/service/models"
)

// ArkServiceCreatePaymentReader is a Reader for the ArkServiceCreatePayment structure.
type ArkServiceCreatePaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArkServiceCreatePaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArkServiceCreatePaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewArkServiceCreatePaymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArkServiceCreatePaymentOK creates a ArkServiceCreatePaymentOK with default headers values
func NewArkServiceCreatePaymentOK() *ArkServiceCreatePaymentOK {
	return &ArkServiceCreatePaymentOK{}
}

/*
ArkServiceCreatePaymentOK describes a response with status code 200, with default header values.

A successful response.
*/
type ArkServiceCreatePaymentOK struct {
	Payload *models.V1CreatePaymentResponse
}

// IsSuccess returns true when this ark service create payment o k response has a 2xx status code
func (o *ArkServiceCreatePaymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ark service create payment o k response has a 3xx status code
func (o *ArkServiceCreatePaymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ark service create payment o k response has a 4xx status code
func (o *ArkServiceCreatePaymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ark service create payment o k response has a 5xx status code
func (o *ArkServiceCreatePaymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ark service create payment o k response a status code equal to that given
func (o *ArkServiceCreatePaymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ark service create payment o k response
func (o *ArkServiceCreatePaymentOK) Code() int {
	return 200
}

func (o *ArkServiceCreatePaymentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment][%d] arkServiceCreatePaymentOK %s", 200, payload)
}

func (o *ArkServiceCreatePaymentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment][%d] arkServiceCreatePaymentOK %s", 200, payload)
}

func (o *ArkServiceCreatePaymentOK) GetPayload() *models.V1CreatePaymentResponse {
	return o.Payload
}

func (o *ArkServiceCreatePaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CreatePaymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArkServiceCreatePaymentDefault creates a ArkServiceCreatePaymentDefault with default headers values
func NewArkServiceCreatePaymentDefault(code int) *ArkServiceCreatePaymentDefault {
	return &ArkServiceCreatePaymentDefault{
		_statusCode: code,
	}
}

/*
ArkServiceCreatePaymentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ArkServiceCreatePaymentDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this ark service create payment default response has a 2xx status code
func (o *ArkServiceCreatePaymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ark service create payment default response has a 3xx status code
func (o *ArkServiceCreatePaymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ark service create payment default response has a 4xx status code
func (o *ArkServiceCreatePaymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ark service create payment default response has a 5xx status code
func (o *ArkServiceCreatePaymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ark service create payment default response a status code equal to that given
func (o *ArkServiceCreatePaymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ark service create payment default response
func (o *ArkServiceCreatePaymentDefault) Code() int {
	return o._statusCode
}

func (o *ArkServiceCreatePaymentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment][%d] ArkService_CreatePayment default %s", o._statusCode, payload)
}

func (o *ArkServiceCreatePaymentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/payment][%d] ArkService_CreatePayment default %s", o._statusCode, payload)
}

func (o *ArkServiceCreatePaymentDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ArkServiceCreatePaymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
