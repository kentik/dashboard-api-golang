// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkPiiPiiKeysReader is a Reader for the GetNetworkPiiPiiKeys structure.
type GetNetworkPiiPiiKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPiiPiiKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkPiiPiiKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkPiiPiiKeysOK creates a GetNetworkPiiPiiKeysOK with default headers values
func NewGetNetworkPiiPiiKeysOK() *GetNetworkPiiPiiKeysOK {
	return &GetNetworkPiiPiiKeysOK{}
}

/* GetNetworkPiiPiiKeysOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkPiiPiiKeysOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network pii pii keys o k response has a 2xx status code
func (o *GetNetworkPiiPiiKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network pii pii keys o k response has a 3xx status code
func (o *GetNetworkPiiPiiKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pii pii keys o k response has a 4xx status code
func (o *GetNetworkPiiPiiKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network pii pii keys o k response has a 5xx status code
func (o *GetNetworkPiiPiiKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network pii pii keys o k response a status code equal to that given
func (o *GetNetworkPiiPiiKeysOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkPiiPiiKeysOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/pii/piiKeys][%d] getNetworkPiiPiiKeysOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPiiPiiKeysOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/pii/piiKeys][%d] getNetworkPiiPiiKeysOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPiiPiiKeysOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkPiiPiiKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
