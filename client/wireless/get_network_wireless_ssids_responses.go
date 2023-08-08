// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetNetworkWirelessSsidsReader is a Reader for the GetNetworkWirelessSsids structure.
type GetNetworkWirelessSsidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSsidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSsidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/wireless/ssids] getNetworkWirelessSsids", response, response.Code())
	}
}

// NewGetNetworkWirelessSsidsOK creates a GetNetworkWirelessSsidsOK with default headers values
func NewGetNetworkWirelessSsidsOK() *GetNetworkWirelessSsidsOK {
	return &GetNetworkWirelessSsidsOK{}
}

/*
GetNetworkWirelessSsidsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSsidsOK struct {
	Payload []*GetNetworkWirelessSsidsOKBodyItems0
}

// IsSuccess returns true when this get network wireless ssids o k response has a 2xx status code
func (o *GetNetworkWirelessSsidsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless ssids o k response has a 3xx status code
func (o *GetNetworkWirelessSsidsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless ssids o k response has a 4xx status code
func (o *GetNetworkWirelessSsidsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless ssids o k response has a 5xx status code
func (o *GetNetworkWirelessSsidsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless ssids o k response a status code equal to that given
func (o *GetNetworkWirelessSsidsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless ssids o k response
func (o *GetNetworkWirelessSsidsOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessSsidsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids][%d] getNetworkWirelessSsidsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids][%d] getNetworkWirelessSsidsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidsOK) GetPayload() []*GetNetworkWirelessSsidsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkWirelessSsidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessSsidsOKBodyItems0 get network wireless ssids o k body items0
swagger:model GetNetworkWirelessSsidsOKBodyItems0
*/
type GetNetworkWirelessSsidsOKBodyItems0 struct {

	// URL for the admin splash page
	AdminSplashURL string `json:"adminSplashUrl,omitempty"`

	// The association control method for the SSID
	// Enum: [8021x-google 8021x-localradius 8021x-meraki 8021x-nac 8021x-radius ipsk-with-nac ipsk-with-radius ipsk-without-radius open open-enhanced open-with-nac open-with-radius psk]
	AuthMode string `json:"authMode,omitempty"`

	// List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list
	AvailabilityTags []string `json:"availabilityTags"`

	// Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags
	AvailableOnAllAps bool `json:"availableOnAllAps,omitempty"`

	// The client-serving radio frequencies of this SSID in the default indoor RF profile
	// Enum: [5 GHz band only Dual band operation Dual band operation with Band Steering]
	BandSelection string `json:"bandSelection,omitempty"`

	// Whether or not the SSID is enabled
	Enabled bool `json:"enabled,omitempty"`

	// The psk encryption mode for the SSID
	// Enum: [wep wpa]
	EncryptionMode string `json:"encryptionMode,omitempty"`

	// The client IP assignment mode
	// Enum: [Bridge mode Ethernet over GRE Layer 3 roaming Layer 3 roaming with a concentrator NAT mode VPN]
	IPAssignmentMode string `json:"ipAssignmentMode,omitempty"`

	// Whether clients connecting to this SSID must use the IP address assigned by the DHCP server
	MandatoryDhcpEnabled bool `json:"mandatoryDhcpEnabled,omitempty"`

	// The minimum bitrate in Mbps of this SSID in the default indoor RF profile
	MinBitrate int64 `json:"minBitrate,omitempty"`

	// The name of the SSID
	Name string `json:"name,omitempty"`

	// Unique identifier of the SSID
	Number int64 `json:"number,omitempty"`

	// The download bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitDown int64 `json:"perClientBandwidthLimitDown,omitempty"`

	// The upload bandwidth limit in Kbps. (0 represents no limit.)
	PerClientBandwidthLimitUp int64 `json:"perClientBandwidthLimitUp,omitempty"`

	// The total download bandwidth limit in Kbps (0 represents no limit)
	PerSsidBandwidthLimitDown int64 `json:"perSsidBandwidthLimitDown,omitempty"`

	// The total upload bandwidth limit in Kbps (0 represents no limit)
	PerSsidBandwidthLimitUp int64 `json:"perSsidBandwidthLimitUp,omitempty"`

	// Whether or not RADIUS accounting is enabled
	RadiusAccountingEnabled bool `json:"radiusAccountingEnabled,omitempty"`

	// List of RADIUS accounting 802.1X servers to be used for authentication
	RadiusAccountingServers []*GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0 `json:"radiusAccountingServers"`

	// RADIUS attribute used to look up group policies
	// Enum: [Airespace-ACL-Name Aruba-User-Role Filter-Id Reply-Message]
	RadiusAttributeForGroupPolicies string `json:"radiusAttributeForGroupPolicies,omitempty"`

	// Whether RADIUS authentication is enabled
	RadiusEnabled bool `json:"radiusEnabled,omitempty"`

	// Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable
	// Enum: [Allow access Deny access]
	RadiusFailoverPolicy string `json:"radiusFailoverPolicy,omitempty"`

	// Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts
	// Enum: [Round robin Strict priority order]
	RadiusLoadBalancingPolicy string `json:"radiusLoadBalancingPolicy,omitempty"`

	// List of RADIUS 802.1X servers to be used for authentication
	RadiusServers []*GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0 `json:"radiusServers"`

	// The type of splash page for the SSID
	// Enum: [Billing Cisco ISE Click-through splash page Facebook Wi-Fi Google Apps domain Google OAuth None Password-protected with Active Directory Password-protected with LDAP Password-protected with Meraki RADIUS Password-protected with custom RADIUS SMS authentication Sponsored guest Systems Manager Sentry]
	SplashPage string `json:"splashPage,omitempty"`

	// Splash page timeout
	SplashTimeout string `json:"splashTimeout,omitempty"`

	// SSID Administrator access status
	SsidAdminAccessible bool `json:"ssidAdminAccessible,omitempty"`

	// Whether the SSID is advertised or hidden by the AP
	Visible bool `json:"visible,omitempty"`

	// Allow users to access a configurable list of IP ranges prior to sign-on
	WalledGardenEnabled bool `json:"walledGardenEnabled,omitempty"`

	// Domain names and IP address ranges available in Walled Garden mode
	WalledGardenRanges []string `json:"walledGardenRanges"`

	// The types of WPA encryption
	// Enum: [WPA1 and WPA2 WPA1 only WPA2 only WPA3 192-bit Security WPA3 Transition Mode WPA3 only]
	WpaEncryptionMode string `json:"wpaEncryptionMode,omitempty"`
}

// Validate validates this get network wireless ssids o k body items0
func (o *GetNetworkWirelessSsidsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBandSelection(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEncryptionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIPAssignmentMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusAccountingServers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusAttributeForGroupPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusFailoverPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusLoadBalancingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusServers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSplashPage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWpaEncryptionMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeAuthModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8021x-google","8021x-localradius","8021x-meraki","8021x-nac","8021x-radius","ipsk-with-nac","ipsk-with-radius","ipsk-without-radius","open","open-enhanced","open-with-nac","open-with-radius","psk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeAuthModePropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeAuthModePropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashGoogle captures enum value "8021x-google"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashGoogle string = "8021x-google"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashLocalradius captures enum value "8021x-localradius"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashLocalradius string = "8021x-localradius"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashMeraki captures enum value "8021x-meraki"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashMeraki string = "8021x-meraki"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashNac captures enum value "8021x-nac"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashNac string = "8021x-nac"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashRadius captures enum value "8021x-radius"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeNr8021xDashRadius string = "8021x-radius"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeIpskDashWithDashNac captures enum value "ipsk-with-nac"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeIpskDashWithDashNac string = "ipsk-with-nac"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeIpskDashWithDashRadius captures enum value "ipsk-with-radius"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeIpskDashWithDashRadius string = "ipsk-with-radius"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeIpskDashWithoutDashRadius captures enum value "ipsk-without-radius"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeIpskDashWithoutDashRadius string = "ipsk-without-radius"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeOpen captures enum value "open"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeOpen string = "open"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeOpenDashEnhanced captures enum value "open-enhanced"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeOpenDashEnhanced string = "open-enhanced"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeOpenDashWithDashNac captures enum value "open-with-nac"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeOpenDashWithDashNac string = "open-with-nac"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModeOpenDashWithDashRadius captures enum value "open-with-radius"
	GetNetworkWirelessSsidsOKBodyItems0AuthModeOpenDashWithDashRadius string = "open-with-radius"

	// GetNetworkWirelessSsidsOKBodyItems0AuthModePsk captures enum value "psk"
	GetNetworkWirelessSsidsOKBodyItems0AuthModePsk string = "psk"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateAuthModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeAuthModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateAuthMode(formats strfmt.Registry) error {
	if swag.IsZero(o.AuthMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateAuthModeEnum("authMode", "body", o.AuthMode); err != nil {
		return err
	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeBandSelectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["5 GHz band only","Dual band operation","Dual band operation with Band Steering"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeBandSelectionPropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeBandSelectionPropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0BandSelectionNr5GHzBandOnly captures enum value "5 GHz band only"
	GetNetworkWirelessSsidsOKBodyItems0BandSelectionNr5GHzBandOnly string = "5 GHz band only"

	// GetNetworkWirelessSsidsOKBodyItems0BandSelectionDualBandOperation captures enum value "Dual band operation"
	GetNetworkWirelessSsidsOKBodyItems0BandSelectionDualBandOperation string = "Dual band operation"

	// GetNetworkWirelessSsidsOKBodyItems0BandSelectionDualBandOperationWithBandSteering captures enum value "Dual band operation with Band Steering"
	GetNetworkWirelessSsidsOKBodyItems0BandSelectionDualBandOperationWithBandSteering string = "Dual band operation with Band Steering"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateBandSelectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeBandSelectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateBandSelection(formats strfmt.Registry) error {
	if swag.IsZero(o.BandSelection) { // not required
		return nil
	}

	// value enum
	if err := o.validateBandSelectionEnum("bandSelection", "body", o.BandSelection); err != nil {
		return err
	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeEncryptionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wep","wpa"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeEncryptionModePropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeEncryptionModePropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0EncryptionModeWep captures enum value "wep"
	GetNetworkWirelessSsidsOKBodyItems0EncryptionModeWep string = "wep"

	// GetNetworkWirelessSsidsOKBodyItems0EncryptionModeWpa captures enum value "wpa"
	GetNetworkWirelessSsidsOKBodyItems0EncryptionModeWpa string = "wpa"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateEncryptionModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeEncryptionModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateEncryptionMode(formats strfmt.Registry) error {
	if swag.IsZero(o.EncryptionMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateEncryptionModeEnum("encryptionMode", "body", o.EncryptionMode); err != nil {
		return err
	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeIPAssignmentModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bridge mode","Ethernet over GRE","Layer 3 roaming","Layer 3 roaming with a concentrator","NAT mode","VPN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeIPAssignmentModePropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeIPAssignmentModePropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeBridgeMode captures enum value "Bridge mode"
	GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeBridgeMode string = "Bridge mode"

	// GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeEthernetOverGRE captures enum value "Ethernet over GRE"
	GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeEthernetOverGRE string = "Ethernet over GRE"

	// GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeLayer3Roaming captures enum value "Layer 3 roaming"
	GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeLayer3Roaming string = "Layer 3 roaming"

	// GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeLayer3RoamingWithaConcentrator captures enum value "Layer 3 roaming with a concentrator"
	GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeLayer3RoamingWithaConcentrator string = "Layer 3 roaming with a concentrator"

	// GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeNATMode captures enum value "NAT mode"
	GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeNATMode string = "NAT mode"

	// GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeVPN captures enum value "VPN"
	GetNetworkWirelessSsidsOKBodyItems0IPAssignmentModeVPN string = "VPN"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateIPAssignmentModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeIPAssignmentModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateIPAssignmentMode(formats strfmt.Registry) error {
	if swag.IsZero(o.IPAssignmentMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateIPAssignmentModeEnum("ipAssignmentMode", "body", o.IPAssignmentMode); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusAccountingServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusAccountingServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusAccountingServers); i++ {
		if swag.IsZero(o.RadiusAccountingServers[i]) { // not required
			continue
		}

		if o.RadiusAccountingServers[i] != nil {
			if err := o.RadiusAccountingServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeRadiusAttributeForGroupPoliciesPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Airespace-ACL-Name","Aruba-User-Role","Filter-Id","Reply-Message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeRadiusAttributeForGroupPoliciesPropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeRadiusAttributeForGroupPoliciesPropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesAirespaceDashACLDashName captures enum value "Airespace-ACL-Name"
	GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesAirespaceDashACLDashName string = "Airespace-ACL-Name"

	// GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesArubaDashUserDashRole captures enum value "Aruba-User-Role"
	GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesArubaDashUserDashRole string = "Aruba-User-Role"

	// GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesFilterDashID captures enum value "Filter-Id"
	GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesFilterDashID string = "Filter-Id"

	// GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesReplyDashMessage captures enum value "Reply-Message"
	GetNetworkWirelessSsidsOKBodyItems0RadiusAttributeForGroupPoliciesReplyDashMessage string = "Reply-Message"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusAttributeForGroupPoliciesEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeRadiusAttributeForGroupPoliciesPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusAttributeForGroupPolicies(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusAttributeForGroupPolicies) { // not required
		return nil
	}

	// value enum
	if err := o.validateRadiusAttributeForGroupPoliciesEnum("radiusAttributeForGroupPolicies", "body", o.RadiusAttributeForGroupPolicies); err != nil {
		return err
	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeRadiusFailoverPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Allow access","Deny access"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeRadiusFailoverPolicyPropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeRadiusFailoverPolicyPropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0RadiusFailoverPolicyAllowAccess captures enum value "Allow access"
	GetNetworkWirelessSsidsOKBodyItems0RadiusFailoverPolicyAllowAccess string = "Allow access"

	// GetNetworkWirelessSsidsOKBodyItems0RadiusFailoverPolicyDenyAccess captures enum value "Deny access"
	GetNetworkWirelessSsidsOKBodyItems0RadiusFailoverPolicyDenyAccess string = "Deny access"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusFailoverPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeRadiusFailoverPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusFailoverPolicy(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusFailoverPolicy) { // not required
		return nil
	}

	// value enum
	if err := o.validateRadiusFailoverPolicyEnum("radiusFailoverPolicy", "body", o.RadiusFailoverPolicy); err != nil {
		return err
	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeRadiusLoadBalancingPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Round robin","Strict priority order"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeRadiusLoadBalancingPolicyPropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeRadiusLoadBalancingPolicyPropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0RadiusLoadBalancingPolicyRoundRobin captures enum value "Round robin"
	GetNetworkWirelessSsidsOKBodyItems0RadiusLoadBalancingPolicyRoundRobin string = "Round robin"

	// GetNetworkWirelessSsidsOKBodyItems0RadiusLoadBalancingPolicyStrictPriorityOrder captures enum value "Strict priority order"
	GetNetworkWirelessSsidsOKBodyItems0RadiusLoadBalancingPolicyStrictPriorityOrder string = "Strict priority order"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusLoadBalancingPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeRadiusLoadBalancingPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusLoadBalancingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusLoadBalancingPolicy) { // not required
		return nil
	}

	// value enum
	if err := o.validateRadiusLoadBalancingPolicyEnum("radiusLoadBalancingPolicy", "body", o.RadiusLoadBalancingPolicy); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateRadiusServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusServers); i++ {
		if swag.IsZero(o.RadiusServers[i]) { // not required
			continue
		}

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeSplashPagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Billing","Cisco ISE","Click-through splash page","Facebook Wi-Fi","Google Apps domain","Google OAuth","None","Password-protected with Active Directory","Password-protected with LDAP","Password-protected with Meraki RADIUS","Password-protected with custom RADIUS","SMS authentication","Sponsored guest","Systems Manager Sentry"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeSplashPagePropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeSplashPagePropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageBilling captures enum value "Billing"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageBilling string = "Billing"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageCiscoISE captures enum value "Cisco ISE"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageCiscoISE string = "Cisco ISE"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageClickDashThroughSplashPage captures enum value "Click-through splash page"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageClickDashThroughSplashPage string = "Click-through splash page"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageFacebookWiDashFi captures enum value "Facebook Wi-Fi"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageFacebookWiDashFi string = "Facebook Wi-Fi"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageGoogleAppsDomain captures enum value "Google Apps domain"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageGoogleAppsDomain string = "Google Apps domain"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageGoogleOAuth captures enum value "Google OAuth"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageGoogleOAuth string = "Google OAuth"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageNone captures enum value "None"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageNone string = "None"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithActiveDirectory captures enum value "Password-protected with Active Directory"
	GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithActiveDirectory string = "Password-protected with Active Directory"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithLDAP captures enum value "Password-protected with LDAP"
	GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithLDAP string = "Password-protected with LDAP"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithMerakiRADIUS captures enum value "Password-protected with Meraki RADIUS"
	GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithMerakiRADIUS string = "Password-protected with Meraki RADIUS"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithCustomRADIUS captures enum value "Password-protected with custom RADIUS"
	GetNetworkWirelessSsidsOKBodyItems0SplashPagePasswordDashProtectedWithCustomRADIUS string = "Password-protected with custom RADIUS"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageSMSAuthentication captures enum value "SMS authentication"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageSMSAuthentication string = "SMS authentication"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageSponsoredGuest captures enum value "Sponsored guest"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageSponsoredGuest string = "Sponsored guest"

	// GetNetworkWirelessSsidsOKBodyItems0SplashPageSystemsManagerSentry captures enum value "Systems Manager Sentry"
	GetNetworkWirelessSsidsOKBodyItems0SplashPageSystemsManagerSentry string = "Systems Manager Sentry"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateSplashPageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeSplashPagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateSplashPage(formats strfmt.Registry) error {
	if swag.IsZero(o.SplashPage) { // not required
		return nil
	}

	// value enum
	if err := o.validateSplashPageEnum("splashPage", "body", o.SplashPage); err != nil {
		return err
	}

	return nil
}

var getNetworkWirelessSsidsOKBodyItems0TypeWpaEncryptionModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WPA1 and WPA2","WPA1 only","WPA2 only","WPA3 192-bit Security","WPA3 Transition Mode","WPA3 only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWirelessSsidsOKBodyItems0TypeWpaEncryptionModePropEnum = append(getNetworkWirelessSsidsOKBodyItems0TypeWpaEncryptionModePropEnum, v)
	}
}

const (

	// GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA1AndWPA2 captures enum value "WPA1 and WPA2"
	GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA1AndWPA2 string = "WPA1 and WPA2"

	// GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA1Only captures enum value "WPA1 only"
	GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA1Only string = "WPA1 only"

	// GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA2Only captures enum value "WPA2 only"
	GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA2Only string = "WPA2 only"

	// GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA3192DashBitSecurity captures enum value "WPA3 192-bit Security"
	GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA3192DashBitSecurity string = "WPA3 192-bit Security"

	// GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA3TransitionMode captures enum value "WPA3 Transition Mode"
	GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA3TransitionMode string = "WPA3 Transition Mode"

	// GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA3Only captures enum value "WPA3 only"
	GetNetworkWirelessSsidsOKBodyItems0WpaEncryptionModeWPA3Only string = "WPA3 only"
)

// prop value enum
func (o *GetNetworkWirelessSsidsOKBodyItems0) validateWpaEncryptionModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWirelessSsidsOKBodyItems0TypeWpaEncryptionModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) validateWpaEncryptionMode(formats strfmt.Registry) error {
	if swag.IsZero(o.WpaEncryptionMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateWpaEncryptionModeEnum("wpaEncryptionMode", "body", o.WpaEncryptionMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get network wireless ssids o k body items0 based on the context it is used
func (o *GetNetworkWirelessSsidsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRadiusAccountingServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadiusServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) contextValidateRadiusAccountingServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusAccountingServers); i++ {

		if o.RadiusAccountingServers[i] != nil {

			if swag.IsZero(o.RadiusAccountingServers[i]) { // not required
				return nil
			}

			if err := o.RadiusAccountingServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkWirelessSsidsOKBodyItems0) contextValidateRadiusServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusServers); i++ {

		if o.RadiusServers[i] != nil {

			if swag.IsZero(o.RadiusServers[i]) { // not required
				return nil
			}

			if err := o.RadiusServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0 get network wireless ssids o k body items0 radius accounting servers items0
swagger:model GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0
*/
type GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0 struct {

	// Certificate used for authorization for the RADSEC Server
	CaCertificate string `json:"caCertificate,omitempty"`

	// IP address (or FQDN) to which the APs will send RADIUS accounting messages
	Host string `json:"host,omitempty"`

	// The ID of the Openroaming Certificate attached to radius server
	OpenRoamingCertificateID int64 `json:"openRoamingCertificateId,omitempty"`

	// Port on the RADIUS server that is listening for accounting messages
	Port int64 `json:"port,omitempty"`
}

// Validate validates this get network wireless ssids o k body items0 radius accounting servers items0
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network wireless ssids o k body items0 radius accounting servers items0 based on context it is used
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidsOKBodyItems0RadiusAccountingServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0 get network wireless ssids o k body items0 radius servers items0
swagger:model GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0
*/
type GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0 struct {

	// Certificate used for authorization for the RADSEC Server
	CaCertificate string `json:"caCertificate,omitempty"`

	// IP address (or FQDN) of your RADIUS server
	Host string `json:"host,omitempty"`

	// The ID of the Openroaming Certificate attached to radius server
	OpenRoamingCertificateID int64 `json:"openRoamingCertificateId,omitempty"`

	// UDP port the RADIUS server listens on for Access-requests
	Port int64 `json:"port,omitempty"`
}

// Validate validates this get network wireless ssids o k body items0 radius servers items0
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network wireless ssids o k body items0 radius servers items0 based on context it is used
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidsOKBodyItems0RadiusServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
