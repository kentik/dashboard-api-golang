// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkSwitchStackRoutingInterfaceReader is a Reader for the DeleteNetworkSwitchStackRoutingInterface structure.
type DeleteNetworkSwitchStackRoutingInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSwitchStackRoutingInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkSwitchStackRoutingInterfaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkSwitchStackRoutingInterfaceNoContent creates a DeleteNetworkSwitchStackRoutingInterfaceNoContent with default headers values
func NewDeleteNetworkSwitchStackRoutingInterfaceNoContent() *DeleteNetworkSwitchStackRoutingInterfaceNoContent {
	return &DeleteNetworkSwitchStackRoutingInterfaceNoContent{}
}

/* DeleteNetworkSwitchStackRoutingInterfaceNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkSwitchStackRoutingInterfaceNoContent struct {
}

// IsSuccess returns true when this delete network switch stack routing interface no content response has a 2xx status code
func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network switch stack routing interface no content response has a 3xx status code
func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network switch stack routing interface no content response has a 4xx status code
func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network switch stack routing interface no content response has a 5xx status code
func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network switch stack routing interface no content response a status code equal to that given
func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}][%d] deleteNetworkSwitchStackRoutingInterfaceNoContent ", 204)
}

func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}][%d] deleteNetworkSwitchStackRoutingInterfaceNoContent ", 204)
}

func (o *DeleteNetworkSwitchStackRoutingInterfaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
