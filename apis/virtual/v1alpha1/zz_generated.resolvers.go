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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	rconfig "github.com/crossplane-contrib/provider-tf-azure/apis/rconfig"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-azure/apis/resource/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this VirtualDesktopHostPool.
func (mg *VirtualDesktopHostPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualDesktopWorkspace.
func (mg *VirtualDesktopWorkspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualHub.
func (mg *VirtualHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualWanID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualWanIDRef,
		Selector:     mg.Spec.ForProvider.VirtualWanIDSelector,
		To: reference.To{
			List:    &VirtualWanList{},
			Managed: &VirtualWan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualWanID")
	}
	mg.Spec.ForProvider.VirtualWanID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualWanIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualHubBgpConnection.
func (mg *VirtualHubBgpConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualHubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualHubIDRef,
		Selector:     mg.Spec.ForProvider.VirtualHubIDSelector,
		To: reference.To{
			List:    &VirtualHubList{},
			Managed: &VirtualHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualHubID")
	}
	mg.Spec.ForProvider.VirtualHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualHubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualHubConnection.
func (mg *VirtualHubConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RemoteVirtualNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RemoteVirtualNetworkIDRef,
		Selector:     mg.Spec.ForProvider.RemoteVirtualNetworkIDSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RemoteVirtualNetworkID")
	}
	mg.Spec.ForProvider.RemoteVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RemoteVirtualNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualHubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualHubIDRef,
		Selector:     mg.Spec.ForProvider.VirtualHubIDSelector,
		To: reference.To{
			List:    &VirtualHubList{},
			Managed: &VirtualHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualHubID")
	}
	mg.Spec.ForProvider.VirtualHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualHubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualHubIp.
func (mg *VirtualHubIp) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualHubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualHubIDRef,
		Selector:     mg.Spec.ForProvider.VirtualHubIDSelector,
		To: reference.To{
			List:    &VirtualHubList{},
			Managed: &VirtualHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualHubID")
	}
	mg.Spec.ForProvider.VirtualHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualHubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualHubRouteTable.
func (mg *VirtualHubRouteTable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualHubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualHubIDRef,
		Selector:     mg.Spec.ForProvider.VirtualHubIDSelector,
		To: reference.To{
			List:    &VirtualHubList{},
			Managed: &VirtualHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualHubID")
	}
	mg.Spec.ForProvider.VirtualHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualHubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualHubSecurityPartnerProvider.
func (mg *VirtualHubSecurityPartnerProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualHubID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualHubIDRef,
		Selector:     mg.Spec.ForProvider.VirtualHubIDSelector,
		To: reference.To{
			List:    &VirtualHubList{},
			Managed: &VirtualHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualHubID")
	}
	mg.Spec.ForProvider.VirtualHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualHubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetwork.
func (mg *VirtualNetwork) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkGateway.
func (mg *VirtualNetworkGateway) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkGatewayConnection.
func (mg *VirtualNetworkGatewayConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PeerVirtualNetworkGatewayID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PeerVirtualNetworkGatewayIDRef,
		Selector:     mg.Spec.ForProvider.PeerVirtualNetworkGatewayIDSelector,
		To: reference.To{
			List:    &VirtualNetworkGatewayList{},
			Managed: &VirtualNetworkGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PeerVirtualNetworkGatewayID")
	}
	mg.Spec.ForProvider.PeerVirtualNetworkGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PeerVirtualNetworkGatewayIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkGatewayID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkGatewayIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkGatewayIDSelector,
		To: reference.To{
			List:    &VirtualNetworkGatewayList{},
			Managed: &VirtualNetworkGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkGatewayID")
	}
	mg.Spec.ForProvider.VirtualNetworkGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkGatewayIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkPeering.
func (mg *VirtualNetworkPeering) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RemoteVirtualNetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RemoteVirtualNetworkIDRef,
		Selector:     mg.Spec.ForProvider.RemoteVirtualNetworkIDSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RemoteVirtualNetworkID")
	}
	mg.Spec.ForProvider.RemoteVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RemoteVirtualNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkNameRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkNameSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkName")
	}
	mg.Spec.ForProvider.VirtualNetworkName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualWan.
func (mg *VirtualWan) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      rconfig.ExtractResourceName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
