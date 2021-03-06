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

type AwsBeanstalkDeploymentPreferencesObservation struct {
}

type AwsBeanstalkDeploymentPreferencesParameters struct {

	// +kubebuilder:validation:Optional
	AutomaticRoll *bool `json:"automaticRoll,omitempty" tf:"automatic_roll,omitempty"`

	// +kubebuilder:validation:Optional
	BatchSizePercentage *int64 `json:"batchSizePercentage,omitempty" tf:"batch_size_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	GracePeriod *int64 `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`

	// +kubebuilder:validation:Optional
	Strategy []DeploymentPreferencesStrategyParameters `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

type AwsBeanstalkManagedActionsObservation struct {
}

type AwsBeanstalkManagedActionsParameters struct {

	// +kubebuilder:validation:Optional
	PlatformUpdate []ManagedActionsPlatformUpdateParameters `json:"platformUpdate,omitempty" tf:"platform_update,omitempty"`
}

type AwsBeanstalkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AwsBeanstalkParameters struct {

	// +kubebuilder:validation:Optional
	BeanstalkEnvironmentID *string `json:"beanstalkEnvironmentId,omitempty" tf:"beanstalk_environment_id,omitempty"`

	// +kubebuilder:validation:Optional
	BeanstalkEnvironmentName *string `json:"beanstalkEnvironmentName,omitempty" tf:"beanstalk_environment_name,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentPreferences []AwsBeanstalkDeploymentPreferencesParameters `json:"deploymentPreferences,omitempty" tf:"deployment_preferences,omitempty"`

	// +kubebuilder:validation:Required
	DesiredCapacity *int64 `json:"desiredCapacity" tf:"desired_capacity,omitempty"`

	// +kubebuilder:validation:Required
	InstanceTypesSpot []*string `json:"instanceTypesSpot" tf:"instance_types_spot,omitempty"`

	// +kubebuilder:validation:Optional
	Maintenance *string `json:"maintenance,omitempty" tf:"maintenance,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedActions []AwsBeanstalkManagedActionsParameters `json:"managedActions,omitempty" tf:"managed_actions,omitempty"`

	// +kubebuilder:validation:Required
	MaxSize *int64 `json:"maxSize" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Required
	MinSize *int64 `json:"minSize" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduledTask []AwsBeanstalkScheduledTaskParameters `json:"scheduledTask,omitempty" tf:"scheduled_task,omitempty"`
}

type AwsBeanstalkScheduledTaskObservation struct {
}

type AwsBeanstalkScheduledTaskParameters struct {

	// +kubebuilder:validation:Optional
	Adjustment *string `json:"adjustment,omitempty" tf:"adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	AdjustmentPercentage *string `json:"adjustmentPercentage,omitempty" tf:"adjustment_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	BatchSizePercentage *string `json:"batchSizePercentage,omitempty" tf:"batch_size_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	CronExpression *string `json:"cronExpression,omitempty" tf:"cron_expression,omitempty"`

	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// +kubebuilder:validation:Optional
	GracePeriod *string `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`

	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MaxCapacity *string `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	MinCapacity *string `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleMaxCapacity *string `json:"scaleMaxCapacity,omitempty" tf:"scale_max_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleMinCapacity *string `json:"scaleMinCapacity,omitempty" tf:"scale_min_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleTargetCapacity *string `json:"scaleTargetCapacity,omitempty" tf:"scale_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// +kubebuilder:validation:Optional
	TargetCapacity *string `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`

	// +kubebuilder:validation:Required
	TaskType *string `json:"taskType" tf:"task_type,omitempty"`
}

type DeploymentPreferencesStrategyObservation struct {
}

type DeploymentPreferencesStrategyParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	ShouldDrainInstances *bool `json:"shouldDrainInstances,omitempty" tf:"should_drain_instances,omitempty"`
}

type ManagedActionsPlatformUpdateObservation struct {
}

type ManagedActionsPlatformUpdateParameters struct {

	// +kubebuilder:validation:Optional
	PerformAt *string `json:"performAt,omitempty" tf:"perform_at,omitempty"`

	// +kubebuilder:validation:Optional
	TimeWindow *string `json:"timeWindow,omitempty" tf:"time_window,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateLevel *string `json:"updateLevel,omitempty" tf:"update_level,omitempty"`
}

// AwsBeanstalkSpec defines the desired state of AwsBeanstalk
type AwsBeanstalkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsBeanstalkParameters `json:"forProvider"`
}

// AwsBeanstalkStatus defines the observed state of AwsBeanstalk.
type AwsBeanstalkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsBeanstalkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AwsBeanstalk is the Schema for the AwsBeanstalks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type AwsBeanstalk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBeanstalkSpec   `json:"spec"`
	Status            AwsBeanstalkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsBeanstalkList contains a list of AwsBeanstalks
type AwsBeanstalkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwsBeanstalk `json:"items"`
}

// Repository type metadata.
var (
	AwsBeanstalk_Kind             = "AwsBeanstalk"
	AwsBeanstalk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AwsBeanstalk_Kind}.String()
	AwsBeanstalk_KindAPIVersion   = AwsBeanstalk_Kind + "." + CRDGroupVersion.String()
	AwsBeanstalk_GroupVersionKind = CRDGroupVersion.WithKind(AwsBeanstalk_Kind)
)

func init() {
	SchemeBuilder.Register(&AwsBeanstalk{}, &AwsBeanstalkList{})
}
