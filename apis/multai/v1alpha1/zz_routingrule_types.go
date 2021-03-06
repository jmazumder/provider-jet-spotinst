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

type RoutingRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RoutingRuleParameters struct {

	// +kubebuilder:validation:Required
	BalancerID *string `json:"balancerId" tf:"balancer_id,omitempty"`

	// +kubebuilder:validation:Required
	ListenerID *string `json:"listenerId" tf:"listener_id,omitempty"`

	// +kubebuilder:validation:Optional
	MiddlewareIds []*string `json:"middlewareIds,omitempty" tf:"middleware_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	Route *string `json:"route" tf:"route,omitempty"`

	// +kubebuilder:validation:Optional
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []RoutingRuleTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TargetSetIds []*string `json:"targetSetIds" tf:"target_set_ids,omitempty"`
}

type RoutingRuleTagsObservation struct {
}

type RoutingRuleTagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// RoutingRuleSpec defines the desired state of RoutingRule
type RoutingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutingRuleParameters `json:"forProvider"`
}

// RoutingRuleStatus defines the observed state of RoutingRule.
type RoutingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingRule is the Schema for the RoutingRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type RoutingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoutingRuleSpec   `json:"spec"`
	Status            RoutingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingRuleList contains a list of RoutingRules
type RoutingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoutingRule `json:"items"`
}

// Repository type metadata.
var (
	RoutingRule_Kind             = "RoutingRule"
	RoutingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoutingRule_Kind}.String()
	RoutingRule_KindAPIVersion   = RoutingRule_Kind + "." + CRDGroupVersion.String()
	RoutingRule_GroupVersionKind = CRDGroupVersion.WithKind(RoutingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&RoutingRule{}, &RoutingRuleList{})
}
