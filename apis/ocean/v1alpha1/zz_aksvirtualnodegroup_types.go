/*
Copyright 2021 The Crossplane Authors.

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

type AksVirtualNodeGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AksVirtualNodeGroupParameters struct {

	// +kubebuilder:validation:Optional
	Autoscale []AutoscaleParameters `json:"autoscale,omitempty" tf:"autoscale,omitempty"`

	// +kubebuilder:validation:Optional
	Label []LabelParameters `json:"label,omitempty" tf:"label,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchSpecification []LaunchSpecificationParameters `json:"launchSpecification,omitempty" tf:"launch_specification,omitempty"`

	// +kubebuilder:validation:Required
	OceanID *string `json:"oceanId" tf:"ocean_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceLimits []AksVirtualNodeGroupResourceLimitsParameters `json:"resourceLimits,omitempty" tf:"resource_limits,omitempty"`

	// +kubebuilder:validation:Optional
	Taint []TaintParameters `json:"taint,omitempty" tf:"taint,omitempty"`
}

type AksVirtualNodeGroupResourceLimitsObservation struct {
}

type AksVirtualNodeGroupResourceLimitsParameters struct {

	// +kubebuilder:validation:Optional
	MaxInstanceCount *int64 `json:"maxInstanceCount,omitempty" tf:"max_instance_count,omitempty"`
}

type AutoscaleAutoscaleHeadroomObservation struct {
}

type AutoscaleAutoscaleHeadroomParameters struct {

	// +kubebuilder:validation:Optional
	CPUPerUnit *int64 `json:"cpuPerUnit,omitempty" tf:"cpu_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	GpuPerUnit *int64 `json:"gpuPerUnit,omitempty" tf:"gpu_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryPerUnit *int64 `json:"memoryPerUnit,omitempty" tf:"memory_per_unit,omitempty"`

	// +kubebuilder:validation:Required
	NumOfUnits *int64 `json:"numOfUnits" tf:"num_of_units,omitempty"`
}

type AutoscaleObservation struct {
}

type AutoscaleParameters struct {

	// +kubebuilder:validation:Optional
	AutoscaleHeadroom []AutoscaleAutoscaleHeadroomParameters `json:"autoscaleHeadroom,omitempty" tf:"autoscale_headroom,omitempty"`
}

type LabelObservation struct {
}

type LabelParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type LaunchSpecificationObservation struct {
}

type LaunchSpecificationOsDiskObservation struct {
}

type LaunchSpecificationOsDiskParameters struct {

	// +kubebuilder:validation:Required
	SizeGb *int64 `json:"sizeGb" tf:"size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LaunchSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	OsDisk []LaunchSpecificationOsDiskParameters `json:"osDisk,omitempty" tf:"os_disk,omitempty"`

	// +kubebuilder:validation:Optional
	Tag []LaunchSpecificationTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type LaunchSpecificationTagObservation struct {
}

type LaunchSpecificationTagParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TaintObservation struct {
}

type TaintParameters struct {

	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// AksVirtualNodeGroupSpec defines the desired state of AksVirtualNodeGroup
type AksVirtualNodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AksVirtualNodeGroupParameters `json:"forProvider"`
}

// AksVirtualNodeGroupStatus defines the observed state of AksVirtualNodeGroup.
type AksVirtualNodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AksVirtualNodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AksVirtualNodeGroup is the Schema for the AksVirtualNodeGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type AksVirtualNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AksVirtualNodeGroupSpec   `json:"spec"`
	Status            AksVirtualNodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AksVirtualNodeGroupList contains a list of AksVirtualNodeGroups
type AksVirtualNodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AksVirtualNodeGroup `json:"items"`
}

// Repository type metadata.
var (
	AksVirtualNodeGroup_Kind             = "AksVirtualNodeGroup"
	AksVirtualNodeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AksVirtualNodeGroup_Kind}.String()
	AksVirtualNodeGroup_KindAPIVersion   = AksVirtualNodeGroup_Kind + "." + CRDGroupVersion.String()
	AksVirtualNodeGroup_GroupVersionKind = CRDGroupVersion.WithKind(AksVirtualNodeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&AksVirtualNodeGroup{}, &AksVirtualNodeGroupList{})
}
