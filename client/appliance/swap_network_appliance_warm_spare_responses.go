// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SwapNetworkApplianceWarmSpareReader is a Reader for the SwapNetworkApplianceWarmSpare structure.
type SwapNetworkApplianceWarmSpareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwapNetworkApplianceWarmSpareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwapNetworkApplianceWarmSpareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/appliance/warmSpare/swap] swapNetworkApplianceWarmSpare", response, response.Code())
	}
}

// NewSwapNetworkApplianceWarmSpareOK creates a SwapNetworkApplianceWarmSpareOK with default headers values
func NewSwapNetworkApplianceWarmSpareOK() *SwapNetworkApplianceWarmSpareOK {
	return &SwapNetworkApplianceWarmSpareOK{}
}

/*
SwapNetworkApplianceWarmSpareOK describes a response with status code 200, with default header values.

Successful operation
*/
type SwapNetworkApplianceWarmSpareOK struct {
	Payload interface{}
}

// IsSuccess returns true when this swap network appliance warm spare o k response has a 2xx status code
func (o *SwapNetworkApplianceWarmSpareOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this swap network appliance warm spare o k response has a 3xx status code
func (o *SwapNetworkApplianceWarmSpareOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this swap network appliance warm spare o k response has a 4xx status code
func (o *SwapNetworkApplianceWarmSpareOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this swap network appliance warm spare o k response has a 5xx status code
func (o *SwapNetworkApplianceWarmSpareOK) IsServerError() bool {
	return false
}

// IsCode returns true when this swap network appliance warm spare o k response a status code equal to that given
func (o *SwapNetworkApplianceWarmSpareOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the swap network appliance warm spare o k response
func (o *SwapNetworkApplianceWarmSpareOK) Code() int {
	return 200
}

func (o *SwapNetworkApplianceWarmSpareOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/appliance/warmSpare/swap][%d] swapNetworkApplianceWarmSpareOK  %+v", 200, o.Payload)
}

func (o *SwapNetworkApplianceWarmSpareOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/appliance/warmSpare/swap][%d] swapNetworkApplianceWarmSpareOK  %+v", 200, o.Payload)
}

func (o *SwapNetworkApplianceWarmSpareOK) GetPayload() interface{} {
	return o.Payload
}

func (o *SwapNetworkApplianceWarmSpareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
