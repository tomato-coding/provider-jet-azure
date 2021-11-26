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

type ImageVersionObservation struct {
}

type ImageVersionParameters struct {

	// +kubebuilder:validation:Optional
	ExcludeFromLatest *bool `json:"excludeFromLatest,omitempty" tf:"exclude_from_latest,omitempty"`

	// +kubebuilder:validation:Required
	GalleryName *string `json:"galleryName" tf:"gallery_name,omitempty"`

	// +kubebuilder:validation:Required
	ImageName *string `json:"imageName" tf:"image_name,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedImageID *string `json:"managedImageId,omitempty" tf:"managed_image_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OsDiskSnapshotID *string `json:"osDiskSnapshotId,omitempty" tf:"os_disk_snapshot_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TargetRegion []TargetRegionParameters `json:"targetRegion" tf:"target_region,omitempty"`
}

type TargetRegionObservation struct {
}

type TargetRegionParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RegionalReplicaCount *int64 `json:"regionalReplicaCount" tf:"regional_replica_count,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`
}

// ImageVersionSpec defines the desired state of ImageVersion
type ImageVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageVersionParameters `json:"forProvider"`
}

// ImageVersionStatus defines the observed state of ImageVersion.
type ImageVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ImageVersion is the Schema for the ImageVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ImageVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageVersionSpec   `json:"spec"`
	Status            ImageVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageVersionList contains a list of ImageVersions
type ImageVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImageVersion `json:"items"`
}

// Repository type metadata.
var (
	ImageVersion_Kind             = "ImageVersion"
	ImageVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ImageVersion_Kind}.String()
	ImageVersion_KindAPIVersion   = ImageVersion_Kind + "." + CRDGroupVersion.String()
	ImageVersion_GroupVersionKind = CRDGroupVersion.WithKind(ImageVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&ImageVersion{}, &ImageVersionList{})
}
