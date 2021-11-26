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

type ManagementApiPolicyObservation struct {
}

type ManagementApiPolicyParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	APIName *string `json:"apiName" tf:"api_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

// ManagementApiPolicySpec defines the desired state of ManagementApiPolicy
type ManagementApiPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementApiPolicyParameters `json:"forProvider"`
}

// ManagementApiPolicyStatus defines the observed state of ManagementApiPolicy.
type ManagementApiPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementApiPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementApiPolicy is the Schema for the ManagementApiPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagementApiPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementApiPolicySpec   `json:"spec"`
	Status            ManagementApiPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementApiPolicyList contains a list of ManagementApiPolicys
type ManagementApiPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementApiPolicy `json:"items"`
}

// Repository type metadata.
var (
	ManagementApiPolicy_Kind             = "ManagementApiPolicy"
	ManagementApiPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementApiPolicy_Kind}.String()
	ManagementApiPolicy_KindAPIVersion   = ManagementApiPolicy_Kind + "." + CRDGroupVersion.String()
	ManagementApiPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ManagementApiPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementApiPolicy{}, &ManagementApiPolicyList{})
}
