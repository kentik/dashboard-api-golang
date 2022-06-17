// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkCameraWirelessProfileReader is a Reader for the DeleteNetworkCameraWirelessProfile structure.
type DeleteNetworkCameraWirelessProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkCameraWirelessProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkCameraWirelessProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkCameraWirelessProfileNoContent creates a DeleteNetworkCameraWirelessProfileNoContent with default headers values
func NewDeleteNetworkCameraWirelessProfileNoContent() *DeleteNetworkCameraWirelessProfileNoContent {
	return &DeleteNetworkCameraWirelessProfileNoContent{}
}

/* DeleteNetworkCameraWirelessProfileNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkCameraWirelessProfileNoContent struct {
}

// IsSuccess returns true when this delete network camera wireless profile no content response has a 2xx status code
func (o *DeleteNetworkCameraWirelessProfileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network camera wireless profile no content response has a 3xx status code
func (o *DeleteNetworkCameraWirelessProfileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network camera wireless profile no content response has a 4xx status code
func (o *DeleteNetworkCameraWirelessProfileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network camera wireless profile no content response has a 5xx status code
func (o *DeleteNetworkCameraWirelessProfileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network camera wireless profile no content response a status code equal to that given
func (o *DeleteNetworkCameraWirelessProfileNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteNetworkCameraWirelessProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}][%d] deleteNetworkCameraWirelessProfileNoContent ", 204)
}

func (o *DeleteNetworkCameraWirelessProfileNoContent) String() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}][%d] deleteNetworkCameraWirelessProfileNoContent ", 204)
}

func (o *DeleteNetworkCameraWirelessProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
