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

// GetNetworkFloorPlansReader is a Reader for the GetNetworkFloorPlans structure.
type GetNetworkFloorPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkFloorPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkFloorPlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkFloorPlansOK creates a GetNetworkFloorPlansOK with default headers values
func NewGetNetworkFloorPlansOK() *GetNetworkFloorPlansOK {
	return &GetNetworkFloorPlansOK{}
}

/* GetNetworkFloorPlansOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkFloorPlansOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network floor plans o k response has a 2xx status code
func (o *GetNetworkFloorPlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network floor plans o k response has a 3xx status code
func (o *GetNetworkFloorPlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network floor plans o k response has a 4xx status code
func (o *GetNetworkFloorPlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network floor plans o k response has a 5xx status code
func (o *GetNetworkFloorPlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network floor plans o k response a status code equal to that given
func (o *GetNetworkFloorPlansOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkFloorPlansOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/floorPlans][%d] getNetworkFloorPlansOK  %+v", 200, o.Payload)
}

func (o *GetNetworkFloorPlansOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/floorPlans][%d] getNetworkFloorPlansOK  %+v", 200, o.Payload)
}

func (o *GetNetworkFloorPlansOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkFloorPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}