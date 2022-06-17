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

// GetOrganizationApplianceVpnStatsReader is a Reader for the GetOrganizationApplianceVpnStats structure.
type GetOrganizationApplianceVpnStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationApplianceVpnStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationApplianceVpnStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationApplianceVpnStatsOK creates a GetOrganizationApplianceVpnStatsOK with default headers values
func NewGetOrganizationApplianceVpnStatsOK() *GetOrganizationApplianceVpnStatsOK {
	return &GetOrganizationApplianceVpnStatsOK{}
}

/* GetOrganizationApplianceVpnStatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationApplianceVpnStatsOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

// IsSuccess returns true when this get organization appliance vpn stats o k response has a 2xx status code
func (o *GetOrganizationApplianceVpnStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization appliance vpn stats o k response has a 3xx status code
func (o *GetOrganizationApplianceVpnStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization appliance vpn stats o k response has a 4xx status code
func (o *GetOrganizationApplianceVpnStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization appliance vpn stats o k response has a 5xx status code
func (o *GetOrganizationApplianceVpnStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization appliance vpn stats o k response a status code equal to that given
func (o *GetOrganizationApplianceVpnStatsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationApplianceVpnStatsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/appliance/vpn/stats][%d] getOrganizationApplianceVpnStatsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationApplianceVpnStatsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/appliance/vpn/stats][%d] getOrganizationApplianceVpnStatsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationApplianceVpnStatsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationApplianceVpnStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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