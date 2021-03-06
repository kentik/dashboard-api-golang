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

// GetNetworkHealthAlertsReader is a Reader for the GetNetworkHealthAlerts structure.
type GetNetworkHealthAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkHealthAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkHealthAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkHealthAlertsOK creates a GetNetworkHealthAlertsOK with default headers values
func NewGetNetworkHealthAlertsOK() *GetNetworkHealthAlertsOK {
	return &GetNetworkHealthAlertsOK{}
}

/* GetNetworkHealthAlertsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkHealthAlertsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network health alerts o k response has a 2xx status code
func (o *GetNetworkHealthAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network health alerts o k response has a 3xx status code
func (o *GetNetworkHealthAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network health alerts o k response has a 4xx status code
func (o *GetNetworkHealthAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network health alerts o k response has a 5xx status code
func (o *GetNetworkHealthAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network health alerts o k response a status code equal to that given
func (o *GetNetworkHealthAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkHealthAlertsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/health/alerts][%d] getNetworkHealthAlertsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkHealthAlertsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/health/alerts][%d] getNetworkHealthAlertsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkHealthAlertsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkHealthAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
