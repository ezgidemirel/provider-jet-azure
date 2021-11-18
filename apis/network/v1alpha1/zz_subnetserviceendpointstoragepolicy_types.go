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

type DefinitionObservation struct {
}

type DefinitionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceResources []*string `json:"serviceResources" tf:"service_resources,omitempty"`
}

type SubnetServiceEndpointStoragePolicyObservation struct {
}

type SubnetServiceEndpointStoragePolicyParameters struct {

	// +kubebuilder:validation:Optional
	Definition []DefinitionParameters `json:"definition,omitempty" tf:"definition,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/azure/v1alpha1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SubnetServiceEndpointStoragePolicySpec defines the desired state of SubnetServiceEndpointStoragePolicy
type SubnetServiceEndpointStoragePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetServiceEndpointStoragePolicyParameters `json:"forProvider"`
}

// SubnetServiceEndpointStoragePolicyStatus defines the observed state of SubnetServiceEndpointStoragePolicy.
type SubnetServiceEndpointStoragePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetServiceEndpointStoragePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetServiceEndpointStoragePolicy is the Schema for the SubnetServiceEndpointStoragePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type SubnetServiceEndpointStoragePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetServiceEndpointStoragePolicySpec   `json:"spec"`
	Status            SubnetServiceEndpointStoragePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetServiceEndpointStoragePolicyList contains a list of SubnetServiceEndpointStoragePolicys
type SubnetServiceEndpointStoragePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetServiceEndpointStoragePolicy `json:"items"`
}

// Repository type metadata.
var (
	SubnetServiceEndpointStoragePolicy_Kind             = "SubnetServiceEndpointStoragePolicy"
	SubnetServiceEndpointStoragePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetServiceEndpointStoragePolicy_Kind}.String()
	SubnetServiceEndpointStoragePolicy_KindAPIVersion   = SubnetServiceEndpointStoragePolicy_Kind + "." + CRDGroupVersion.String()
	SubnetServiceEndpointStoragePolicy_GroupVersionKind = CRDGroupVersion.WithKind(SubnetServiceEndpointStoragePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetServiceEndpointStoragePolicy{}, &SubnetServiceEndpointStoragePolicyList{})
}
