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

package v1alpha2

import (
	"context"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2"
	v1alpha21 "github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2"
	rconfig "github.com/crossplane-contrib/provider-jet-azure/apis/rconfig"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ActiveDirectoryAdministrator.
func (mg *ActiveDirectoryAdministrator) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerNameRef,
		Selector:     mg.Spec.ForProvider.ServerNameSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerName")
	}
	mg.Spec.ForProvider.ServerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Configuration.
func (mg *Configuration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerNameRef,
		Selector:     mg.Spec.ForProvider.ServerNameSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerName")
	}
	mg.Spec.ForProvider.ServerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Database.
func (mg *Database) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerNameRef,
		Selector:     mg.Spec.ForProvider.ServerNameSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerName")
	}
	mg.Spec.ForProvider.ServerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FirewallRule.
func (mg *FirewallRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerNameRef,
		Selector:     mg.Spec.ForProvider.ServerNameSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerName")
	}
	mg.Spec.ForProvider.ServerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FlexibleServer.
func (mg *FlexibleServer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DelegatedSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.DelegatedSubnetIDRef,
		Selector:     mg.Spec.ForProvider.DelegatedSubnetIDSelector,
		To: reference.To{
			List:    &v1alpha21.SubnetList{},
			Managed: &v1alpha21.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DelegatedSubnetID")
	}
	mg.Spec.ForProvider.DelegatedSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DelegatedSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FlexibleServerConfiguration.
func (mg *FlexibleServerConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &FlexibleServerList{},
			Managed: &FlexibleServer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FlexibleServerDatabase.
func (mg *FlexibleServerDatabase) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &FlexibleServerList{},
			Managed: &FlexibleServer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this FlexibleServerFirewallRule.
func (mg *FlexibleServerFirewallRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ServerIDRef,
		Selector:     mg.Spec.ForProvider.ServerIDSelector,
		To: reference.To{
			List:    &FlexibleServerList{},
			Managed: &FlexibleServer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerID")
	}
	mg.Spec.ForProvider.ServerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Server.
func (mg *Server) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkRule.
func (mg *VirtualNetworkRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha2.ResourceGroupList{},
			Managed: &v1alpha2.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServerName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServerNameRef,
		Selector:     mg.Spec.ForProvider.ServerNameSelector,
		To: reference.To{
			List:    &ServerList{},
			Managed: &Server{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServerName")
	}
	mg.Spec.ForProvider.ServerName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServerNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1alpha21.SubnetList{},
			Managed: &v1alpha21.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}
