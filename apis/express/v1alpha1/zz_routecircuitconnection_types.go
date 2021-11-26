/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RouteCircuitConnectionObservation struct {
}

type RouteCircuitConnectionParameters struct {

	// +kubebuilder:validation:Required
	AddressPrefixIPv4 *string `json:"addressPrefixIpv4" tf:"address_prefix_ipv4,omitempty"`

	// +kubebuilder:validation:Optional
	AddressPrefixIPv6 *string `json:"addressPrefixIpv6,omitempty" tf:"address_prefix_ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizationKeySecretRef *v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PeerPeeringID *string `json:"peerPeeringId" tf:"peer_peering_id,omitempty"`

	// +kubebuilder:validation:Required
	PeeringID *string `json:"peeringId" tf:"peering_id,omitempty"`
}

// RouteCircuitConnectionSpec defines the desired state of RouteCircuitConnection
type RouteCircuitConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteCircuitConnectionParameters `json:"forProvider"`
}

// RouteCircuitConnectionStatus defines the observed state of RouteCircuitConnection.
type RouteCircuitConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteCircuitConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteCircuitConnection is the Schema for the RouteCircuitConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type RouteCircuitConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteCircuitConnectionSpec   `json:"spec"`
	Status            RouteCircuitConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteCircuitConnectionList contains a list of RouteCircuitConnections
type RouteCircuitConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteCircuitConnection `json:"items"`
}

// Repository type metadata.
var (
	RouteCircuitConnection_Kind             = "RouteCircuitConnection"
	RouteCircuitConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteCircuitConnection_Kind}.String()
	RouteCircuitConnection_KindAPIVersion   = RouteCircuitConnection_Kind + "." + CRDGroupVersion.String()
	RouteCircuitConnection_GroupVersionKind = CRDGroupVersion.WithKind(RouteCircuitConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteCircuitConnection{}, &RouteCircuitConnectionList{})
}
