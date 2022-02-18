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

type GkeLaunchSpecAutoscaleHeadroomsObservation struct {
}

type GkeLaunchSpecAutoscaleHeadroomsParameters struct {

	// +kubebuilder:validation:Optional
	CPUPerUnit *int64 `json:"cpuPerUnit,omitempty" tf:"cpu_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	GpuPerUnit *int64 `json:"gpuPerUnit,omitempty" tf:"gpu_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryPerUnit *int64 `json:"memoryPerUnit,omitempty" tf:"memory_per_unit,omitempty"`

	// +kubebuilder:validation:Required
	NumOfUnits *int64 `json:"numOfUnits" tf:"num_of_units,omitempty"`
}

type GkeLaunchSpecLabelsObservation struct {
}

type GkeLaunchSpecLabelsParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GkeLaunchSpecObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GkeLaunchSpecParameters struct {

	// +kubebuilder:validation:Optional
	AutoscaleHeadrooms []GkeLaunchSpecAutoscaleHeadroomsParameters `json:"autoscaleHeadrooms,omitempty" tf:"autoscale_headrooms,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types,omitempty"`

	// +kubebuilder:validation:Optional
	Labels []GkeLaunchSpecLabelsParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	NodePoolName *string `json:"nodePoolName,omitempty" tf:"node_pool_name,omitempty"`

	// +kubebuilder:validation:Required
	OceanID *string `json:"oceanId" tf:"ocean_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceLimits []GkeLaunchSpecResourceLimitsParameters `json:"resourceLimits,omitempty" tf:"resource_limits,omitempty"`

	// +kubebuilder:validation:Optional
	RestrictScaleDown *bool `json:"restrictScaleDown,omitempty" tf:"restrict_scale_down,omitempty"`

	// +kubebuilder:validation:Optional
	RootVolumeSize *int64 `json:"rootVolumeSize,omitempty" tf:"root_volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	RootVolumeType *string `json:"rootVolumeType,omitempty" tf:"root_volume_type,omitempty"`

	// +kubebuilder:validation:Optional
	SchedulingTask []GkeLaunchSpecSchedulingTaskParameters `json:"schedulingTask,omitempty" tf:"scheduling_task,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// +kubebuilder:validation:Optional
	ShieldedInstanceConfig []ShieldedInstanceConfigParameters `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config,omitempty"`

	// +kubebuilder:validation:Optional
	SourceImage *string `json:"sourceImage,omitempty" tf:"source_image,omitempty"`

	// +kubebuilder:validation:Optional
	Storage []StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// +kubebuilder:validation:Optional
	Strategy []GkeLaunchSpecStrategyParameters `json:"strategy,omitempty" tf:"strategy,omitempty"`

	// +kubebuilder:validation:Optional
	Taints []GkeLaunchSpecTaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatePolicy []GkeLaunchSpecUpdatePolicyParameters `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`
}

type GkeLaunchSpecResourceLimitsObservation struct {
}

type GkeLaunchSpecResourceLimitsParameters struct {

	// +kubebuilder:validation:Optional
	MaxInstanceCount *int64 `json:"maxInstanceCount,omitempty" tf:"max_instance_count,omitempty"`

	// +kubebuilder:validation:Optional
	MinInstanceCount *int64 `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
}

type GkeLaunchSpecSchedulingTaskObservation struct {
}

type GkeLaunchSpecSchedulingTaskParameters struct {

	// +kubebuilder:validation:Required
	CronExpression *string `json:"cronExpression" tf:"cron_expression,omitempty"`

	// +kubebuilder:validation:Required
	IsEnabled *bool `json:"isEnabled" tf:"is_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	TaskHeadroom []GkeLaunchSpecSchedulingTaskTaskHeadroomParameters `json:"taskHeadroom,omitempty" tf:"task_headroom,omitempty"`

	// +kubebuilder:validation:Required
	TaskType *string `json:"taskType" tf:"task_type,omitempty"`
}

type GkeLaunchSpecSchedulingTaskTaskHeadroomObservation struct {
}

type GkeLaunchSpecSchedulingTaskTaskHeadroomParameters struct {

	// +kubebuilder:validation:Optional
	CPUPerUnit *int64 `json:"cpuPerUnit,omitempty" tf:"cpu_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	GpuPerUnit *int64 `json:"gpuPerUnit,omitempty" tf:"gpu_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryPerUnit *int64 `json:"memoryPerUnit,omitempty" tf:"memory_per_unit,omitempty"`

	// +kubebuilder:validation:Required
	NumOfUnits *int64 `json:"numOfUnits" tf:"num_of_units,omitempty"`
}

type GkeLaunchSpecStrategyObservation struct {
}

type GkeLaunchSpecStrategyParameters struct {

	// +kubebuilder:validation:Optional
	PreemptiblePercentage *int64 `json:"preemptiblePercentage,omitempty" tf:"preemptible_percentage,omitempty"`
}

type GkeLaunchSpecTaintsObservation struct {
}

type GkeLaunchSpecTaintsParameters struct {

	// +kubebuilder:validation:Optional
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GkeLaunchSpecUpdatePolicyObservation struct {
}

type GkeLaunchSpecUpdatePolicyParameters struct {

	// +kubebuilder:validation:Optional
	RollConfig []GkeLaunchSpecUpdatePolicyRollConfigParameters `json:"rollConfig,omitempty" tf:"roll_config,omitempty"`

	// +kubebuilder:validation:Required
	ShouldRoll *bool `json:"shouldRoll" tf:"should_roll,omitempty"`
}

type GkeLaunchSpecUpdatePolicyRollConfigObservation struct {
}

type GkeLaunchSpecUpdatePolicyRollConfigParameters struct {

	// +kubebuilder:validation:Required
	BatchSizePercentage *int64 `json:"batchSizePercentage" tf:"batch_size_percentage,omitempty"`
}

type MetadataObservation struct {
}

type MetadataParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ShieldedInstanceConfigObservation struct {
}

type ShieldedInstanceConfigParameters struct {

	// +kubebuilder:validation:Optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring,omitempty"`

	// +kubebuilder:validation:Optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot,omitempty"`
}

type StorageObservation struct {
}

type StorageParameters struct {

	// +kubebuilder:validation:Optional
	LocalSsdCount *int64 `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`
}

// GkeLaunchSpecSpec defines the desired state of GkeLaunchSpec
type GkeLaunchSpecSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GkeLaunchSpecParameters `json:"forProvider"`
}

// GkeLaunchSpecStatus defines the observed state of GkeLaunchSpec.
type GkeLaunchSpecStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GkeLaunchSpecObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GkeLaunchSpec is the Schema for the GkeLaunchSpecs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type GkeLaunchSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GkeLaunchSpecSpec   `json:"spec"`
	Status            GkeLaunchSpecStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GkeLaunchSpecList contains a list of GkeLaunchSpecs
type GkeLaunchSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GkeLaunchSpec `json:"items"`
}

// Repository type metadata.
var (
	GkeLaunchSpec_Kind             = "GkeLaunchSpec"
	GkeLaunchSpec_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GkeLaunchSpec_Kind}.String()
	GkeLaunchSpec_KindAPIVersion   = GkeLaunchSpec_Kind + "." + CRDGroupVersion.String()
	GkeLaunchSpec_GroupVersionKind = CRDGroupVersion.WithKind(GkeLaunchSpec_Kind)
)

func init() {
	SchemeBuilder.Register(&GkeLaunchSpec{}, &GkeLaunchSpecList{})
}