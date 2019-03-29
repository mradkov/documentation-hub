// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetOracleByPubkeyReader is a Reader for the GetOracleByPubkey structure.
type GetOracleByPubkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleByPubkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOracleByPubkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOracleByPubkeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOracleByPubkeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOracleByPubkeyOK creates a GetOracleByPubkeyOK with default headers values
func NewGetOracleByPubkeyOK() *GetOracleByPubkeyOK {
	return &GetOracleByPubkeyOK{}
}

/*GetOracleByPubkeyOK handles this case with default header values.

Successful operation
*/
type GetOracleByPubkeyOK struct {
	Payload *models.RegisteredOracle
}

func (o *GetOracleByPubkeyOK) Error() string {
	return fmt.Sprintf("[GET /oracles/{pubkey}][%d] getOracleByPubkeyOK  %+v", 200, o.Payload)
}

func (o *GetOracleByPubkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegisteredOracle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleByPubkeyBadRequest creates a GetOracleByPubkeyBadRequest with default headers values
func NewGetOracleByPubkeyBadRequest() *GetOracleByPubkeyBadRequest {
	return &GetOracleByPubkeyBadRequest{}
}

/*GetOracleByPubkeyBadRequest handles this case with default header values.

Invalid public key
*/
type GetOracleByPubkeyBadRequest struct {
	Payload *models.Error
}

func (o *GetOracleByPubkeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracles/{pubkey}][%d] getOracleByPubkeyBadRequest  %+v", 400, o.Payload)
}

func (o *GetOracleByPubkeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleByPubkeyNotFound creates a GetOracleByPubkeyNotFound with default headers values
func NewGetOracleByPubkeyNotFound() *GetOracleByPubkeyNotFound {
	return &GetOracleByPubkeyNotFound{}
}

/*GetOracleByPubkeyNotFound handles this case with default header values.

Oracle not found
*/
type GetOracleByPubkeyNotFound struct {
	Payload *models.Error
}

func (o *GetOracleByPubkeyNotFound) Error() string {
	return fmt.Sprintf("[GET /oracles/{pubkey}][%d] getOracleByPubkeyNotFound  %+v", 404, o.Payload)
}

func (o *GetOracleByPubkeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
