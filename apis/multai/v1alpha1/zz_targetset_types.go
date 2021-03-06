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

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {

	// +kubebuilder:validation:Required
	HealthyThreshold *int64 `json:"healthyThreshold" tf:"healthy_threshold,omitempty"`

	// +kubebuilder:validation:Required
	Interval *int64 `json:"interval" tf:"interval,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	Timeout *int64 `json:"timeout" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Required
	UnhealthyThreshold *int64 `json:"unhealthyThreshold" tf:"unhealthy_threshold,omitempty"`
}

type TargetSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TargetSetParameters struct {

	// +kubebuilder:validation:Required
	BalancerID *string `json:"balancerId" tf:"balancer_id,omitempty"`

	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// +kubebuilder:validation:Required
	HealthCheck []HealthCheckParameters `json:"healthCheck" tf:"health_check,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TargetSetTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	Weight *int64 `json:"weight" tf:"weight,omitempty"`
}

type TargetSetTagsObservation struct {
}

type TargetSetTagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// TargetSetSpec defines the desired state of TargetSet
type TargetSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetSetParameters `json:"forProvider"`
}

// TargetSetStatus defines the observed state of TargetSet.
type TargetSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetSet is the Schema for the TargetSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type TargetSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetSetSpec   `json:"spec"`
	Status            TargetSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetSetList contains a list of TargetSets
type TargetSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetSet `json:"items"`
}

// Repository type metadata.
var (
	TargetSet_Kind             = "TargetSet"
	TargetSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetSet_Kind}.String()
	TargetSet_KindAPIVersion   = TargetSet_Kind + "." + CRDGroupVersion.String()
	TargetSet_GroupVersionKind = CRDGroupVersion.WithKind(TargetSet_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetSet{}, &TargetSetList{})
}
