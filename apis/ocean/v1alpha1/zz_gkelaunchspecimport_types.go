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

type GkeLaunchSpecImportObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GkeLaunchSpecImportParameters struct {

	// +kubebuilder:validation:Required
	NodePoolName *string `json:"nodePoolName" tf:"node_pool_name,omitempty"`

	// +kubebuilder:validation:Required
	OceanID *string `json:"oceanId" tf:"ocean_id,omitempty"`
}

// GkeLaunchSpecImportSpec defines the desired state of GkeLaunchSpecImport
type GkeLaunchSpecImportSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GkeLaunchSpecImportParameters `json:"forProvider"`
}

// GkeLaunchSpecImportStatus defines the observed state of GkeLaunchSpecImport.
type GkeLaunchSpecImportStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GkeLaunchSpecImportObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GkeLaunchSpecImport is the Schema for the GkeLaunchSpecImports API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type GkeLaunchSpecImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GkeLaunchSpecImportSpec   `json:"spec"`
	Status            GkeLaunchSpecImportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GkeLaunchSpecImportList contains a list of GkeLaunchSpecImports
type GkeLaunchSpecImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GkeLaunchSpecImport `json:"items"`
}

// Repository type metadata.
var (
	GkeLaunchSpecImport_Kind             = "GkeLaunchSpecImport"
	GkeLaunchSpecImport_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GkeLaunchSpecImport_Kind}.String()
	GkeLaunchSpecImport_KindAPIVersion   = GkeLaunchSpecImport_Kind + "." + CRDGroupVersion.String()
	GkeLaunchSpecImport_GroupVersionKind = CRDGroupVersion.WithKind(GkeLaunchSpecImport_Kind)
)

func init() {
	SchemeBuilder.Register(&GkeLaunchSpecImport{}, &GkeLaunchSpecImportList{})
}
