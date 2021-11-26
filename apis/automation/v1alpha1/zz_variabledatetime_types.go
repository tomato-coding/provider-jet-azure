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

type VariableDatetimeObservation struct {
}

type VariableDatetimeParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// VariableDatetimeSpec defines the desired state of VariableDatetime
type VariableDatetimeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VariableDatetimeParameters `json:"forProvider"`
}

// VariableDatetimeStatus defines the observed state of VariableDatetime.
type VariableDatetimeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VariableDatetimeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VariableDatetime is the Schema for the VariableDatetimes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VariableDatetime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VariableDatetimeSpec   `json:"spec"`
	Status            VariableDatetimeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VariableDatetimeList contains a list of VariableDatetimes
type VariableDatetimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VariableDatetime `json:"items"`
}

// Repository type metadata.
var (
	VariableDatetime_Kind             = "VariableDatetime"
	VariableDatetime_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VariableDatetime_Kind}.String()
	VariableDatetime_KindAPIVersion   = VariableDatetime_Kind + "." + CRDGroupVersion.String()
	VariableDatetime_GroupVersionKind = CRDGroupVersion.WithKind(VariableDatetime_Kind)
)

func init() {
	SchemeBuilder.Register(&VariableDatetime{}, &VariableDatetimeList{})
}
