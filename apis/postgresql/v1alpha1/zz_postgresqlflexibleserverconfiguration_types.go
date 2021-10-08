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

type PostgresqlFlexibleServerConfigurationObservation struct {
}

type PostgresqlFlexibleServerConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +crossplane:generate:reference:type=PostgresqlFlexibleServer
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id"`

	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIDRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIDSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value"`
}

// PostgresqlFlexibleServerConfigurationSpec defines the desired state of PostgresqlFlexibleServerConfiguration
type PostgresqlFlexibleServerConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlFlexibleServerConfigurationParameters `json:"forProvider"`
}

// PostgresqlFlexibleServerConfigurationStatus defines the observed state of PostgresqlFlexibleServerConfiguration.
type PostgresqlFlexibleServerConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlFlexibleServerConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlFlexibleServerConfiguration is the Schema for the PostgresqlFlexibleServerConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PostgresqlFlexibleServerConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlFlexibleServerConfigurationSpec   `json:"spec"`
	Status            PostgresqlFlexibleServerConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlFlexibleServerConfigurationList contains a list of PostgresqlFlexibleServerConfigurations
type PostgresqlFlexibleServerConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlFlexibleServerConfiguration `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlFlexibleServerConfigurationKind             = "PostgresqlFlexibleServerConfiguration"
	PostgresqlFlexibleServerConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: PostgresqlFlexibleServerConfigurationKind}.String()
	PostgresqlFlexibleServerConfigurationKindAPIVersion   = PostgresqlFlexibleServerConfigurationKind + "." + GroupVersion.String()
	PostgresqlFlexibleServerConfigurationGroupVersionKind = GroupVersion.WithKind(PostgresqlFlexibleServerConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlFlexibleServerConfiguration{}, &PostgresqlFlexibleServerConfigurationList{})
}
