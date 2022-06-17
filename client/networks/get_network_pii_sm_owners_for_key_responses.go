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

// GetNetworkPiiSmOwnersForKeyReader is a Reader for the GetNetworkPiiSmOwnersForKey structure.
type GetNetworkPiiSmOwnersForKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPiiSmOwnersForKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkPiiSmOwnersForKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkPiiSmOwnersForKeyOK creates a GetNetworkPiiSmOwnersForKeyOK with default headers values
func NewGetNetworkPiiSmOwnersForKeyOK() *GetNetworkPiiSmOwnersForKeyOK {
	return &GetNetworkPiiSmOwnersForKeyOK{}
}

/* GetNetworkPiiSmOwnersForKeyOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkPiiSmOwnersForKeyOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network pii sm owners for key o k response has a 2xx status code
func (o *GetNetworkPiiSmOwnersForKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network pii sm owners for key o k response has a 3xx status code
func (o *GetNetworkPiiSmOwnersForKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network pii sm owners for key o k response has a 4xx status code
func (o *GetNetworkPiiSmOwnersForKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network pii sm owners for key o k response has a 5xx status code
func (o *GetNetworkPiiSmOwnersForKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network pii sm owners for key o k response a status code equal to that given
func (o *GetNetworkPiiSmOwnersForKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkPiiSmOwnersForKeyOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/pii/smOwnersForKey][%d] getNetworkPiiSmOwnersForKeyOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPiiSmOwnersForKeyOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/pii/smOwnersForKey][%d] getNetworkPiiSmOwnersForKeyOK  %+v", 200, o.Payload)
}

func (o *GetNetworkPiiSmOwnersForKeyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkPiiSmOwnersForKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}