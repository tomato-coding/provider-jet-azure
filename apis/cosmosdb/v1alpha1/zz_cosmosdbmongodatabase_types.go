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

type CosmosdbMongoDatabaseAutoscaleSettingsObservation struct {
}

type CosmosdbMongoDatabaseAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput"`
}

type CosmosdbMongoDatabaseObservation struct {
}

type CosmosdbMongoDatabaseParameters struct {

	// +crossplane:generate:reference:type=CosmosdbAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []CosmosdbMongoDatabaseAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
}

// CosmosdbMongoDatabaseSpec defines the desired state of CosmosdbMongoDatabase
type CosmosdbMongoDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CosmosdbMongoDatabaseParameters `json:"forProvider"`
}

// CosmosdbMongoDatabaseStatus defines the observed state of CosmosdbMongoDatabase.
type CosmosdbMongoDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CosmosdbMongoDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbMongoDatabase is the Schema for the CosmosdbMongoDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CosmosdbMongoDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbMongoDatabaseSpec   `json:"spec"`
	Status            CosmosdbMongoDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbMongoDatabaseList contains a list of CosmosdbMongoDatabases
type CosmosdbMongoDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosdbMongoDatabase `json:"items"`
}

// Repository type metadata.
var (
	CosmosdbMongoDatabaseKind             = "CosmosdbMongoDatabase"
	CosmosdbMongoDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: CosmosdbMongoDatabaseKind}.String()
	CosmosdbMongoDatabaseKindAPIVersion   = CosmosdbMongoDatabaseKind + "." + GroupVersion.String()
	CosmosdbMongoDatabaseGroupVersionKind = GroupVersion.WithKind(CosmosdbMongoDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&CosmosdbMongoDatabase{}, &CosmosdbMongoDatabaseList{})
}
