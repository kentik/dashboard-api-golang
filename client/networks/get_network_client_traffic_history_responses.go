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

// GetNetworkClientTrafficHistoryReader is a Reader for the GetNetworkClientTrafficHistory structure.
type GetNetworkClientTrafficHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkClientTrafficHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkClientTrafficHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/clients/{clientId}/trafficHistory] getNetworkClientTrafficHistory", response, response.Code())
	}
}

// NewGetNetworkClientTrafficHistoryOK creates a GetNetworkClientTrafficHistoryOK with default headers values
func NewGetNetworkClientTrafficHistoryOK() *GetNetworkClientTrafficHistoryOK {
	return &GetNetworkClientTrafficHistoryOK{}
}

/*
GetNetworkClientTrafficHistoryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkClientTrafficHistoryOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

// IsSuccess returns true when this get network client traffic history o k response has a 2xx status code
func (o *GetNetworkClientTrafficHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network client traffic history o k response has a 3xx status code
func (o *GetNetworkClientTrafficHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network client traffic history o k response has a 4xx status code
func (o *GetNetworkClientTrafficHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network client traffic history o k response has a 5xx status code
func (o *GetNetworkClientTrafficHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network client traffic history o k response a status code equal to that given
func (o *GetNetworkClientTrafficHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network client traffic history o k response
func (o *GetNetworkClientTrafficHistoryOK) Code() int {
	return 200
}

func (o *GetNetworkClientTrafficHistoryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/clients/{clientId}/trafficHistory][%d] getNetworkClientTrafficHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkClientTrafficHistoryOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/clients/{clientId}/trafficHistory][%d] getNetworkClientTrafficHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkClientTrafficHistoryOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkClientTrafficHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
