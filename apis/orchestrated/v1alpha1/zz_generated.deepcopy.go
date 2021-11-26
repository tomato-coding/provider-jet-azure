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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineScaleSet) DeepCopyInto(out *VirtualMachineScaleSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineScaleSet.
func (in *VirtualMachineScaleSet) DeepCopy() *VirtualMachineScaleSet {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineScaleSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineScaleSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineScaleSetList) DeepCopyInto(out *VirtualMachineScaleSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineScaleSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineScaleSetList.
func (in *VirtualMachineScaleSetList) DeepCopy() *VirtualMachineScaleSetList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineScaleSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineScaleSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineScaleSetObservation) DeepCopyInto(out *VirtualMachineScaleSetObservation) {
	*out = *in
	if in.UniqueID != nil {
		in, out := &in.UniqueID, &out.UniqueID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineScaleSetObservation.
func (in *VirtualMachineScaleSetObservation) DeepCopy() *VirtualMachineScaleSetObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineScaleSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineScaleSetParameters) DeepCopyInto(out *VirtualMachineScaleSetParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PlatformFaultDomainCount != nil {
		in, out := &in.PlatformFaultDomainCount, &out.PlatformFaultDomainCount
		*out = new(int64)
		**out = **in
	}
	if in.ProximityPlacementGroupID != nil {
		in, out := &in.ProximityPlacementGroupID, &out.ProximityPlacementGroupID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SinglePlacementGroup != nil {
		in, out := &in.SinglePlacementGroup, &out.SinglePlacementGroup
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineScaleSetParameters.
func (in *VirtualMachineScaleSetParameters) DeepCopy() *VirtualMachineScaleSetParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineScaleSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineScaleSetSpec) DeepCopyInto(out *VirtualMachineScaleSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineScaleSetSpec.
func (in *VirtualMachineScaleSetSpec) DeepCopy() *VirtualMachineScaleSetSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineScaleSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineScaleSetStatus) DeepCopyInto(out *VirtualMachineScaleSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineScaleSetStatus.
func (in *VirtualMachineScaleSetStatus) DeepCopy() *VirtualMachineScaleSetStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineScaleSetStatus)
	in.DeepCopyInto(out)
	return out
}
