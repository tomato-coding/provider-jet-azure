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

type AppIntegrationAccountAssemblyObservation struct {
}

type AppIntegrationAccountAssemblyParameters struct {

	// +kubebuilder:validation:Required
	AssemblyName *string `json:"assemblyName" tf:"assembly_name,omitempty"`

	// +kubebuilder:validation:Optional
	AssemblyVersion *string `json:"assemblyVersion,omitempty" tf:"assembly_version,omitempty"`

	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	ContentLinkURI *string `json:"contentLinkUri,omitempty" tf:"content_link_uri,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationAccountName *string `json:"integrationAccountName" tf:"integration_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// AppIntegrationAccountAssemblySpec defines the desired state of AppIntegrationAccountAssembly
type AppIntegrationAccountAssemblySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountAssemblyParameters `json:"forProvider"`
}

// AppIntegrationAccountAssemblyStatus defines the observed state of AppIntegrationAccountAssembly.
type AppIntegrationAccountAssemblyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountAssemblyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountAssembly is the Schema for the AppIntegrationAccountAssemblys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppIntegrationAccountAssembly struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppIntegrationAccountAssemblySpec   `json:"spec"`
	Status            AppIntegrationAccountAssemblyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountAssemblyList contains a list of AppIntegrationAccountAssemblys
type AppIntegrationAccountAssemblyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccountAssembly `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccountAssembly_Kind             = "AppIntegrationAccountAssembly"
	AppIntegrationAccountAssembly_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppIntegrationAccountAssembly_Kind}.String()
	AppIntegrationAccountAssembly_KindAPIVersion   = AppIntegrationAccountAssembly_Kind + "." + CRDGroupVersion.String()
	AppIntegrationAccountAssembly_GroupVersionKind = CRDGroupVersion.WithKind(AppIntegrationAccountAssembly_Kind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccountAssembly{}, &AppIntegrationAccountAssemblyList{})
}
