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

type AutoscaleSettingObservation struct {
}

type AutoscaleSettingParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Notification []NotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// +kubebuilder:validation:Required
	Profile []ProfileParameters `json:"profile" tf:"profile,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TargetResourceID *string `json:"targetResourceId" tf:"target_resource_id,omitempty"`
}

type CapacityObservation struct {
}

type CapacityParameters struct {

	// +kubebuilder:validation:Required
	Default *int64 `json:"default" tf:"default,omitempty"`

	// +kubebuilder:validation:Required
	Maximum *int64 `json:"maximum" tf:"maximum,omitempty"`

	// +kubebuilder:validation:Required
	Minimum *int64 `json:"minimum" tf:"minimum,omitempty"`
}

type DimensionsObservation struct {
}

type DimensionsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type EmailObservation struct {
}

type EmailParameters struct {

	// +kubebuilder:validation:Optional
	CustomEmails []*string `json:"customEmails,omitempty" tf:"custom_emails,omitempty"`

	// +kubebuilder:validation:Optional
	SendToSubscriptionAdministrator *bool `json:"sendToSubscriptionAdministrator,omitempty" tf:"send_to_subscription_administrator,omitempty"`

	// +kubebuilder:validation:Optional
	SendToSubscriptionCoAdministrator *bool `json:"sendToSubscriptionCoAdministrator,omitempty" tf:"send_to_subscription_co_administrator,omitempty"`
}

type FixedDateObservation struct {
}

type FixedDateParameters struct {

	// +kubebuilder:validation:Required
	End *string `json:"end" tf:"end,omitempty"`

	// +kubebuilder:validation:Required
	Start *string `json:"start" tf:"start,omitempty"`

	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type MetricTriggerObservation struct {
}

type MetricTriggerParameters struct {

	// +kubebuilder:validation:Optional
	Dimensions []DimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// +kubebuilder:validation:Optional
	DivideByInstanceCount *bool `json:"divideByInstanceCount,omitempty" tf:"divide_by_instance_count,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Optional
	MetricNamespace *string `json:"metricNamespace,omitempty" tf:"metric_namespace,omitempty"`

	// +kubebuilder:validation:Required
	MetricResourceID *string `json:"metricResourceId" tf:"metric_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Statistic *string `json:"statistic" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Required
	TimeAggregation *string `json:"timeAggregation" tf:"time_aggregation,omitempty"`

	// +kubebuilder:validation:Required
	TimeGrain *string `json:"timeGrain" tf:"time_grain,omitempty"`

	// +kubebuilder:validation:Required
	TimeWindow *string `json:"timeWindow" tf:"time_window,omitempty"`
}

type NotificationObservation struct {
}

type NotificationParameters struct {

	// +kubebuilder:validation:Optional
	Email []EmailParameters `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	Webhook []WebhookParameters `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type ProfileObservation struct {
}

type ProfileParameters struct {

	// +kubebuilder:validation:Required
	Capacity []CapacityParameters `json:"capacity" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Optional
	FixedDate []FixedDateParameters `json:"fixedDate,omitempty" tf:"fixed_date,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Recurrence []RecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RecurrenceObservation struct {
}

type RecurrenceParameters struct {

	// +kubebuilder:validation:Required
	Days []*string `json:"days" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	Hours []*int64 `json:"hours" tf:"hours,omitempty"`

	// +kubebuilder:validation:Required
	Minutes []*int64 `json:"minutes" tf:"minutes,omitempty"`

	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Required
	MetricTrigger []MetricTriggerParameters `json:"metricTrigger" tf:"metric_trigger,omitempty"`

	// +kubebuilder:validation:Required
	ScaleAction []ScaleActionParameters `json:"scaleAction" tf:"scale_action,omitempty"`
}

type ScaleActionObservation struct {
}

type ScaleActionParameters struct {

	// +kubebuilder:validation:Required
	Cooldown *string `json:"cooldown" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Value *int64 `json:"value" tf:"value,omitempty"`
}

type WebhookObservation struct {
}

type WebhookParameters struct {

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Required
	ServiceURI *string `json:"serviceUri" tf:"service_uri,omitempty"`
}

// AutoscaleSettingSpec defines the desired state of AutoscaleSetting
type AutoscaleSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AutoscaleSettingParameters `json:"forProvider"`
}

// AutoscaleSettingStatus defines the observed state of AutoscaleSetting.
type AutoscaleSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AutoscaleSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscaleSetting is the Schema for the AutoscaleSettings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscaleSettingSpec   `json:"spec"`
	Status            AutoscaleSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscaleSettingList contains a list of AutoscaleSettings
type AutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscaleSetting `json:"items"`
}

// Repository type metadata.
var (
	AutoscaleSetting_Kind             = "AutoscaleSetting"
	AutoscaleSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AutoscaleSetting_Kind}.String()
	AutoscaleSetting_KindAPIVersion   = AutoscaleSetting_Kind + "." + CRDGroupVersion.String()
	AutoscaleSetting_GroupVersionKind = CRDGroupVersion.WithKind(AutoscaleSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&AutoscaleSetting{}, &AutoscaleSettingList{})
}
