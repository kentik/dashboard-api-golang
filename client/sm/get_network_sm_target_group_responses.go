// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSmTargetGroupReader is a Reader for the GetNetworkSmTargetGroup structure.
type GetNetworkSmTargetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmTargetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmTargetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/sm/targetGroups/{targetGroupId}] getNetworkSmTargetGroup", response, response.Code())
	}
}

// NewGetNetworkSmTargetGroupOK creates a GetNetworkSmTargetGroupOK with default headers values
func NewGetNetworkSmTargetGroupOK() *GetNetworkSmTargetGroupOK {
	return &GetNetworkSmTargetGroupOK{}
}

/*
GetNetworkSmTargetGroupOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmTargetGroupOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network sm target group o k response has a 2xx status code
func (o *GetNetworkSmTargetGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm target group o k response has a 3xx status code
func (o *GetNetworkSmTargetGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm target group o k response has a 4xx status code
func (o *GetNetworkSmTargetGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm target group o k response has a 5xx status code
func (o *GetNetworkSmTargetGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm target group o k response a status code equal to that given
func (o *GetNetworkSmTargetGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sm target group o k response
func (o *GetNetworkSmTargetGroupOK) Code() int {
	return 200
}

func (o *GetNetworkSmTargetGroupOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/targetGroups/{targetGroupId}][%d] getNetworkSmTargetGroupOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmTargetGroupOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/targetGroups/{targetGroupId}][%d] getNetworkSmTargetGroupOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmTargetGroupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSmTargetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
