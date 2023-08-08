// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkSwitchRoutingMulticastRendezvousPointReader is a Reader for the DeleteNetworkSwitchRoutingMulticastRendezvousPoint structure.
type DeleteNetworkSwitchRoutingMulticastRendezvousPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}] deleteNetworkSwitchRoutingMulticastRendezvousPoint", response, response.Code())
	}
}

// NewDeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent creates a DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent with default headers values
func NewDeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent() *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent {
	return &DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent{}
}

/*
DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent struct {
}

// IsSuccess returns true when this delete network switch routing multicast rendezvous point no content response has a 2xx status code
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network switch routing multicast rendezvous point no content response has a 3xx status code
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network switch routing multicast rendezvous point no content response has a 4xx status code
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network switch routing multicast rendezvous point no content response has a 5xx status code
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network switch routing multicast rendezvous point no content response a status code equal to that given
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete network switch routing multicast rendezvous point no content response
func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) Code() int {
	return 204
}

func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}][%d] deleteNetworkSwitchRoutingMulticastRendezvousPointNoContent ", 204)
}

func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId}][%d] deleteNetworkSwitchRoutingMulticastRendezvousPointNoContent ", 204)
}

func (o *DeleteNetworkSwitchRoutingMulticastRendezvousPointNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
