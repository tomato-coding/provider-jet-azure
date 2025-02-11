/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BGPSettingsObservation struct {
}

type BGPSettingsParameters struct {

	// +kubebuilder:validation:Optional
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// +kubebuilder:validation:Optional
	PeerWeight *float64 `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`

	// +kubebuilder:validation:Optional
	PeeringAddress *string `json:"peeringAddress,omitempty" tf:"peering_address,omitempty"`

	// +kubebuilder:validation:Optional
	PeeringAddresses []PeeringAddressesParameters `json:"peeringAddresses,omitempty" tf:"peering_addresses,omitempty"`
}

type CustomRouteObservation struct {
}

type CustomRouteParameters struct {

	// +kubebuilder:validation:Optional
	AddressPrefixes []*string `json:"addressPrefixes,omitempty" tf:"address_prefixes,omitempty"`
}

type PeeringAddressesObservation struct {
	DefaultAddresses []*string `json:"defaultAddresses,omitempty" tf:"default_addresses,omitempty"`

	TunnelIPAddresses []*string `json:"tunnelIpAddresses,omitempty" tf:"tunnel_ip_addresses,omitempty"`
}

type PeeringAddressesParameters struct {

	// +kubebuilder:validation:Optional
	ApipaAddresses []*string `json:"apipaAddresses,omitempty" tf:"apipa_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`
}

type RevokedCertificateObservation struct {
}

type RevokedCertificateParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`
}

type RootCertificateObservation struct {
}

type RootCertificateParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PublicCertData *string `json:"publicCertData" tf:"public_cert_data,omitempty"`
}

type VPNClientConfigurationObservation struct {
}

type VPNClientConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AADAudience *string `json:"aadAudience,omitempty" tf:"aad_audience,omitempty"`

	// +kubebuilder:validation:Optional
	AADIssuer *string `json:"aadIssuer,omitempty" tf:"aad_issuer,omitempty"`

	// +kubebuilder:validation:Optional
	AADTenant *string `json:"aadTenant,omitempty" tf:"aad_tenant,omitempty"`

	// +kubebuilder:validation:Required
	AddressSpace []*string `json:"addressSpace" tf:"address_space,omitempty"`

	// +kubebuilder:validation:Optional
	RadiusServerAddress *string `json:"radiusServerAddress,omitempty" tf:"radius_server_address,omitempty"`

	// +kubebuilder:validation:Optional
	RadiusServerSecret *string `json:"radiusServerSecret,omitempty" tf:"radius_server_secret,omitempty"`

	// +kubebuilder:validation:Optional
	RevokedCertificate []RevokedCertificateParameters `json:"revokedCertificate,omitempty" tf:"revoked_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	RootCertificate []RootCertificateParameters `json:"rootCertificate,omitempty" tf:"root_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	VPNAuthTypes []*string `json:"vpnAuthTypes,omitempty" tf:"vpn_auth_types,omitempty"`

	// +kubebuilder:validation:Optional
	VPNClientProtocols []*string `json:"vpnClientProtocols,omitempty" tf:"vpn_client_protocols,omitempty"`
}

type VirtualNetworkGatewayIPConfigurationObservation struct {
}

type VirtualNetworkGatewayIPConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPAddressAllocation *string `json:"privateIpAddressAllocation,omitempty" tf:"private_ip_address_allocation,omitempty"`

	// +kubebuilder:validation:Required
	PublicIPAddressID *string `json:"publicIpAddressId" tf:"public_ip_address_id,omitempty"`

	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualNetworkGatewayObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualNetworkGatewayParameters struct {

	// +kubebuilder:validation:Optional
	ActiveActive *bool `json:"activeActive,omitempty" tf:"active_active,omitempty"`

	// +kubebuilder:validation:Optional
	BGPSettings []BGPSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// +kubebuilder:validation:Optional
	CustomRoute []CustomRouteParameters `json:"customRoute,omitempty" tf:"custom_route,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultLocalNetworkGatewayID *string `json:"defaultLocalNetworkGatewayId,omitempty" tf:"default_local_network_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBGP *bool `json:"enableBgp,omitempty" tf:"enable_bgp,omitempty"`

	// +kubebuilder:validation:Optional
	Generation *string `json:"generation,omitempty" tf:"generation,omitempty"`

	// +kubebuilder:validation:Required
	IPConfiguration []VirtualNetworkGatewayIPConfigurationParameters `json:"ipConfiguration" tf:"ip_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIPAddressEnabled *bool `json:"privateIpAddressEnabled,omitempty" tf:"private_ip_address_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	VPNClientConfiguration []VPNClientConfigurationParameters `json:"vpnClientConfiguration,omitempty" tf:"vpn_client_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	VPNType *string `json:"vpnType,omitempty" tf:"vpn_type,omitempty"`
}

// VirtualNetworkGatewaySpec defines the desired state of VirtualNetworkGateway
type VirtualNetworkGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkGatewayParameters `json:"forProvider"`
}

// VirtualNetworkGatewayStatus defines the observed state of VirtualNetworkGateway.
type VirtualNetworkGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGateway is the Schema for the VirtualNetworkGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VirtualNetworkGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewaySpec   `json:"spec"`
	Status            VirtualNetworkGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayList contains a list of VirtualNetworkGateways
type VirtualNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGateway `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkGateway_Kind             = "VirtualNetworkGateway"
	VirtualNetworkGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetworkGateway_Kind}.String()
	VirtualNetworkGateway_KindAPIVersion   = VirtualNetworkGateway_Kind + "." + CRDGroupVersion.String()
	VirtualNetworkGateway_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetworkGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkGateway{}, &VirtualNetworkGatewayList{})
}
