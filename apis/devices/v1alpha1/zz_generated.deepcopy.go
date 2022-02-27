// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointEventHub) DeepCopyInto(out *IOTHubEndpointEventHub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointEventHub.
func (in *IOTHubEndpointEventHub) DeepCopy() *IOTHubEndpointEventHub {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointEventHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEndpointEventHub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointEventHubList) DeepCopyInto(out *IOTHubEndpointEventHubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTHubEndpointEventHub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointEventHubList.
func (in *IOTHubEndpointEventHubList) DeepCopy() *IOTHubEndpointEventHubList {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointEventHubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEndpointEventHubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointEventHubObservation) DeepCopyInto(out *IOTHubEndpointEventHubObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointEventHubObservation.
func (in *IOTHubEndpointEventHubObservation) DeepCopy() *IOTHubEndpointEventHubObservation {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointEventHubObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointEventHubParameters) DeepCopyInto(out *IOTHubEndpointEventHubParameters) {
	*out = *in
	out.ConnectionStringSecretRef = in.ConnectionStringSecretRef
	if in.IOTHubName != nil {
		in, out := &in.IOTHubName, &out.IOTHubName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointEventHubParameters.
func (in *IOTHubEndpointEventHubParameters) DeepCopy() *IOTHubEndpointEventHubParameters {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointEventHubParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointEventHubSpec) DeepCopyInto(out *IOTHubEndpointEventHubSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointEventHubSpec.
func (in *IOTHubEndpointEventHubSpec) DeepCopy() *IOTHubEndpointEventHubSpec {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointEventHubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointEventHubStatus) DeepCopyInto(out *IOTHubEndpointEventHubStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointEventHubStatus.
func (in *IOTHubEndpointEventHubStatus) DeepCopy() *IOTHubEndpointEventHubStatus {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointEventHubStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusQueue) DeepCopyInto(out *IOTHubEndpointServiceBusQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusQueue.
func (in *IOTHubEndpointServiceBusQueue) DeepCopy() *IOTHubEndpointServiceBusQueue {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEndpointServiceBusQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusQueueList) DeepCopyInto(out *IOTHubEndpointServiceBusQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTHubEndpointServiceBusQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusQueueList.
func (in *IOTHubEndpointServiceBusQueueList) DeepCopy() *IOTHubEndpointServiceBusQueueList {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEndpointServiceBusQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusQueueObservation) DeepCopyInto(out *IOTHubEndpointServiceBusQueueObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusQueueObservation.
func (in *IOTHubEndpointServiceBusQueueObservation) DeepCopy() *IOTHubEndpointServiceBusQueueObservation {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusQueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusQueueParameters) DeepCopyInto(out *IOTHubEndpointServiceBusQueueParameters) {
	*out = *in
	out.ConnectionStringSecretRef = in.ConnectionStringSecretRef
	if in.IOTHubName != nil {
		in, out := &in.IOTHubName, &out.IOTHubName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusQueueParameters.
func (in *IOTHubEndpointServiceBusQueueParameters) DeepCopy() *IOTHubEndpointServiceBusQueueParameters {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusQueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusQueueSpec) DeepCopyInto(out *IOTHubEndpointServiceBusQueueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusQueueSpec.
func (in *IOTHubEndpointServiceBusQueueSpec) DeepCopy() *IOTHubEndpointServiceBusQueueSpec {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusQueueStatus) DeepCopyInto(out *IOTHubEndpointServiceBusQueueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusQueueStatus.
func (in *IOTHubEndpointServiceBusQueueStatus) DeepCopy() *IOTHubEndpointServiceBusQueueStatus {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusTopic) DeepCopyInto(out *IOTHubEndpointServiceBusTopic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusTopic.
func (in *IOTHubEndpointServiceBusTopic) DeepCopy() *IOTHubEndpointServiceBusTopic {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusTopic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEndpointServiceBusTopic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusTopicList) DeepCopyInto(out *IOTHubEndpointServiceBusTopicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTHubEndpointServiceBusTopic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusTopicList.
func (in *IOTHubEndpointServiceBusTopicList) DeepCopy() *IOTHubEndpointServiceBusTopicList {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusTopicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEndpointServiceBusTopicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusTopicObservation) DeepCopyInto(out *IOTHubEndpointServiceBusTopicObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusTopicObservation.
func (in *IOTHubEndpointServiceBusTopicObservation) DeepCopy() *IOTHubEndpointServiceBusTopicObservation {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusTopicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusTopicParameters) DeepCopyInto(out *IOTHubEndpointServiceBusTopicParameters) {
	*out = *in
	out.ConnectionStringSecretRef = in.ConnectionStringSecretRef
	if in.IOTHubName != nil {
		in, out := &in.IOTHubName, &out.IOTHubName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusTopicParameters.
func (in *IOTHubEndpointServiceBusTopicParameters) DeepCopy() *IOTHubEndpointServiceBusTopicParameters {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusTopicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusTopicSpec) DeepCopyInto(out *IOTHubEndpointServiceBusTopicSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusTopicSpec.
func (in *IOTHubEndpointServiceBusTopicSpec) DeepCopy() *IOTHubEndpointServiceBusTopicSpec {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusTopicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEndpointServiceBusTopicStatus) DeepCopyInto(out *IOTHubEndpointServiceBusTopicStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEndpointServiceBusTopicStatus.
func (in *IOTHubEndpointServiceBusTopicStatus) DeepCopy() *IOTHubEndpointServiceBusTopicStatus {
	if in == nil {
		return nil
	}
	out := new(IOTHubEndpointServiceBusTopicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEnrichment) DeepCopyInto(out *IOTHubEnrichment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEnrichment.
func (in *IOTHubEnrichment) DeepCopy() *IOTHubEnrichment {
	if in == nil {
		return nil
	}
	out := new(IOTHubEnrichment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEnrichment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEnrichmentList) DeepCopyInto(out *IOTHubEnrichmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTHubEnrichment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEnrichmentList.
func (in *IOTHubEnrichmentList) DeepCopy() *IOTHubEnrichmentList {
	if in == nil {
		return nil
	}
	out := new(IOTHubEnrichmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubEnrichmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEnrichmentObservation) DeepCopyInto(out *IOTHubEnrichmentObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEnrichmentObservation.
func (in *IOTHubEnrichmentObservation) DeepCopy() *IOTHubEnrichmentObservation {
	if in == nil {
		return nil
	}
	out := new(IOTHubEnrichmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEnrichmentParameters) DeepCopyInto(out *IOTHubEnrichmentParameters) {
	*out = *in
	if in.EndpointNames != nil {
		in, out := &in.EndpointNames, &out.EndpointNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IOTHubName != nil {
		in, out := &in.IOTHubName, &out.IOTHubName
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEnrichmentParameters.
func (in *IOTHubEnrichmentParameters) DeepCopy() *IOTHubEnrichmentParameters {
	if in == nil {
		return nil
	}
	out := new(IOTHubEnrichmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEnrichmentSpec) DeepCopyInto(out *IOTHubEnrichmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEnrichmentSpec.
func (in *IOTHubEnrichmentSpec) DeepCopy() *IOTHubEnrichmentSpec {
	if in == nil {
		return nil
	}
	out := new(IOTHubEnrichmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubEnrichmentStatus) DeepCopyInto(out *IOTHubEnrichmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubEnrichmentStatus.
func (in *IOTHubEnrichmentStatus) DeepCopy() *IOTHubEnrichmentStatus {
	if in == nil {
		return nil
	}
	out := new(IOTHubEnrichmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubFallbackRoute) DeepCopyInto(out *IOTHubFallbackRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubFallbackRoute.
func (in *IOTHubFallbackRoute) DeepCopy() *IOTHubFallbackRoute {
	if in == nil {
		return nil
	}
	out := new(IOTHubFallbackRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubFallbackRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubFallbackRouteList) DeepCopyInto(out *IOTHubFallbackRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTHubFallbackRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubFallbackRouteList.
func (in *IOTHubFallbackRouteList) DeepCopy() *IOTHubFallbackRouteList {
	if in == nil {
		return nil
	}
	out := new(IOTHubFallbackRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubFallbackRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubFallbackRouteObservation) DeepCopyInto(out *IOTHubFallbackRouteObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubFallbackRouteObservation.
func (in *IOTHubFallbackRouteObservation) DeepCopy() *IOTHubFallbackRouteObservation {
	if in == nil {
		return nil
	}
	out := new(IOTHubFallbackRouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubFallbackRouteParameters) DeepCopyInto(out *IOTHubFallbackRouteParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EndpointNames != nil {
		in, out := &in.EndpointNames, &out.EndpointNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IOTHubName != nil {
		in, out := &in.IOTHubName, &out.IOTHubName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubFallbackRouteParameters.
func (in *IOTHubFallbackRouteParameters) DeepCopy() *IOTHubFallbackRouteParameters {
	if in == nil {
		return nil
	}
	out := new(IOTHubFallbackRouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubFallbackRouteSpec) DeepCopyInto(out *IOTHubFallbackRouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubFallbackRouteSpec.
func (in *IOTHubFallbackRouteSpec) DeepCopy() *IOTHubFallbackRouteSpec {
	if in == nil {
		return nil
	}
	out := new(IOTHubFallbackRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubFallbackRouteStatus) DeepCopyInto(out *IOTHubFallbackRouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubFallbackRouteStatus.
func (in *IOTHubFallbackRouteStatus) DeepCopy() *IOTHubFallbackRouteStatus {
	if in == nil {
		return nil
	}
	out := new(IOTHubFallbackRouteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubRoute) DeepCopyInto(out *IOTHubRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubRoute.
func (in *IOTHubRoute) DeepCopy() *IOTHubRoute {
	if in == nil {
		return nil
	}
	out := new(IOTHubRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubRouteList) DeepCopyInto(out *IOTHubRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IOTHubRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubRouteList.
func (in *IOTHubRouteList) DeepCopy() *IOTHubRouteList {
	if in == nil {
		return nil
	}
	out := new(IOTHubRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IOTHubRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubRouteObservation) DeepCopyInto(out *IOTHubRouteObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubRouteObservation.
func (in *IOTHubRouteObservation) DeepCopy() *IOTHubRouteObservation {
	if in == nil {
		return nil
	}
	out := new(IOTHubRouteObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubRouteParameters) DeepCopyInto(out *IOTHubRouteParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EndpointNames != nil {
		in, out := &in.EndpointNames, &out.EndpointNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IOTHubName != nil {
		in, out := &in.IOTHubName, &out.IOTHubName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubRouteParameters.
func (in *IOTHubRouteParameters) DeepCopy() *IOTHubRouteParameters {
	if in == nil {
		return nil
	}
	out := new(IOTHubRouteParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubRouteSpec) DeepCopyInto(out *IOTHubRouteSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubRouteSpec.
func (in *IOTHubRouteSpec) DeepCopy() *IOTHubRouteSpec {
	if in == nil {
		return nil
	}
	out := new(IOTHubRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IOTHubRouteStatus) DeepCopyInto(out *IOTHubRouteStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IOTHubRouteStatus.
func (in *IOTHubRouteStatus) DeepCopy() *IOTHubRouteStatus {
	if in == nil {
		return nil
	}
	out := new(IOTHubRouteStatus)
	in.DeepCopyInto(out)
	return out
}
