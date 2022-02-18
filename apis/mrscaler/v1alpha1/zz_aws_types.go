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

type ApplicationsObservation struct {
}

type ApplicationsParameters struct {

	// +kubebuilder:validation:Optional
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AwsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OutputClusterID *string `json:"outputClusterId,omitempty" tf:"output_cluster_id,omitempty"`
}

type AwsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalInfo *string `json:"additionalInfo,omitempty" tf:"additional_info,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalPrimarySecurityGroups []*string `json:"additionalPrimarySecurityGroups,omitempty" tf:"additional_primary_security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalReplicaSecurityGroups []*string `json:"additionalReplicaSecurityGroups,omitempty" tf:"additional_replica_security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	Applications []ApplicationsParameters `json:"applications,omitempty" tf:"applications,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	BootstrapActionsFile []BootstrapActionsFileParameters `json:"bootstrapActionsFile,omitempty" tf:"bootstrap_actions_file,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigurationsFile []ConfigurationsFileParameters `json:"configurationsFile,omitempty" tf:"configurations_file,omitempty"`

	// +kubebuilder:validation:Optional
	CoreDesiredCapacity *int64 `json:"coreDesiredCapacity,omitempty" tf:"core_desired_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	CoreEBSBlockDevice []CoreEBSBlockDeviceParameters `json:"coreEbsBlockDevice,omitempty" tf:"core_ebs_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	CoreEBSOptimized *bool `json:"coreEbsOptimized,omitempty" tf:"core_ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	CoreInstanceTypes []*string `json:"coreInstanceTypes,omitempty" tf:"core_instance_types,omitempty"`

	// +kubebuilder:validation:Optional
	CoreLifecycle *string `json:"coreLifecycle,omitempty" tf:"core_lifecycle,omitempty"`

	// +kubebuilder:validation:Optional
	CoreMaxSize *int64 `json:"coreMaxSize,omitempty" tf:"core_max_size,omitempty"`

	// +kubebuilder:validation:Optional
	CoreMinSize *int64 `json:"coreMinSize,omitempty" tf:"core_min_size,omitempty"`

	// +kubebuilder:validation:Optional
	CoreScalingDownPolicy []CoreScalingDownPolicyParameters `json:"coreScalingDownPolicy,omitempty" tf:"core_scaling_down_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CoreScalingUpPolicy []CoreScalingUpPolicyParameters `json:"coreScalingUpPolicy,omitempty" tf:"core_scaling_up_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CoreUnit *string `json:"coreUnit,omitempty" tf:"core_unit,omitempty"`

	// +kubebuilder:validation:Optional
	CustomAMIID *string `json:"customAmiId,omitempty" tf:"custom_ami_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EBSRootVolumeSize *int64 `json:"ebsRootVolumeSize,omitempty" tf:"ebs_root_volume_size,omitempty"`

	// +kubebuilder:validation:Optional
	EC2KeyName *string `json:"ec2KeyName,omitempty" tf:"ec2_key_name,omitempty"`

	// +kubebuilder:validation:Optional
	ExposeClusterID *bool `json:"exposeClusterId,omitempty" tf:"expose_cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceWeights []InstanceWeightsParameters `json:"instanceWeights,omitempty" tf:"instance_weights,omitempty"`

	// +kubebuilder:validation:Optional
	JobFlowRole *string `json:"jobFlowRole,omitempty" tf:"job_flow_role,omitempty"`

	// +kubebuilder:validation:Optional
	KeepJobFlowAlive *bool `json:"keepJobFlowAlive,omitempty" tf:"keep_job_flow_alive,omitempty"`

	// +kubebuilder:validation:Optional
	LogURI *string `json:"logUri,omitempty" tf:"log_uri,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedPrimarySecurityGroup *string `json:"managedPrimarySecurityGroup,omitempty" tf:"managed_primary_security_group,omitempty"`

	// +kubebuilder:validation:Optional
	ManagedReplicaSecurityGroup *string `json:"managedReplicaSecurityGroup,omitempty" tf:"managed_replica_security_group,omitempty"`

	// +kubebuilder:validation:Optional
	MasterEBSBlockDevice []MasterEBSBlockDeviceParameters `json:"masterEbsBlockDevice,omitempty" tf:"master_ebs_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	MasterEBSOptimized *bool `json:"masterEbsOptimized,omitempty" tf:"master_ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	MasterInstanceTypes []*string `json:"masterInstanceTypes,omitempty" tf:"master_instance_types,omitempty"`

	// +kubebuilder:validation:Optional
	MasterLifecycle *string `json:"masterLifecycle,omitempty" tf:"master_lifecycle,omitempty"`

	// +kubebuilder:validation:Optional
	MasterTarget *int64 `json:"masterTarget,omitempty" tf:"master_target,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisioningTimeout []ProvisioningTimeoutParameters `json:"provisioningTimeout,omitempty" tf:"provisioning_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	ReleaseLabel *string `json:"releaseLabel,omitempty" tf:"release_label,omitempty"`

	// +kubebuilder:validation:Optional
	RepoUpgradeOnBoot *string `json:"repoUpgradeOnBoot,omitempty" tf:"repo_upgrade_on_boot,omitempty"`

	// +kubebuilder:validation:Optional
	Retries *int64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduledTask []ScheduledTaskParameters `json:"scheduledTask,omitempty" tf:"scheduled_task,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityConfig *string `json:"securityConfig,omitempty" tf:"security_config,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccessSecurityGroup *string `json:"serviceAccessSecurityGroup,omitempty" tf:"service_access_security_group,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceRole *string `json:"serviceRole,omitempty" tf:"service_role,omitempty"`

	// +kubebuilder:validation:Optional
	StepsFile []StepsFileParameters `json:"stepsFile,omitempty" tf:"steps_file,omitempty"`

	// +kubebuilder:validation:Required
	Strategy *string `json:"strategy" tf:"strategy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TaskDesiredCapacity *int64 `json:"taskDesiredCapacity,omitempty" tf:"task_desired_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	TaskEBSBlockDevice []TaskEBSBlockDeviceParameters `json:"taskEbsBlockDevice,omitempty" tf:"task_ebs_block_device,omitempty"`

	// +kubebuilder:validation:Optional
	TaskEBSOptimized *bool `json:"taskEbsOptimized,omitempty" tf:"task_ebs_optimized,omitempty"`

	// +kubebuilder:validation:Optional
	TaskInstanceTypes []*string `json:"taskInstanceTypes,omitempty" tf:"task_instance_types,omitempty"`

	// +kubebuilder:validation:Optional
	TaskLifecycle *string `json:"taskLifecycle,omitempty" tf:"task_lifecycle,omitempty"`

	// +kubebuilder:validation:Optional
	TaskMaxSize *int64 `json:"taskMaxSize,omitempty" tf:"task_max_size,omitempty"`

	// +kubebuilder:validation:Optional
	TaskMinSize *int64 `json:"taskMinSize,omitempty" tf:"task_min_size,omitempty"`

	// +kubebuilder:validation:Optional
	TaskScalingDownPolicy []TaskScalingDownPolicyParameters `json:"taskScalingDownPolicy,omitempty" tf:"task_scaling_down_policy,omitempty"`

	// +kubebuilder:validation:Optional
	TaskScalingUpPolicy []TaskScalingUpPolicyParameters `json:"taskScalingUpPolicy,omitempty" tf:"task_scaling_up_policy,omitempty"`

	// +kubebuilder:validation:Optional
	TaskUnit *string `json:"taskUnit,omitempty" tf:"task_unit,omitempty"`

	// +kubebuilder:validation:Optional
	TerminationPolicies []TerminationPoliciesParameters `json:"terminationPolicies,omitempty" tf:"termination_policies,omitempty"`

	// +kubebuilder:validation:Optional
	TerminationProtected *bool `json:"terminationProtected,omitempty" tf:"termination_protected,omitempty"`

	// +kubebuilder:validation:Optional
	VisibleToAllUsers *bool `json:"visibleToAllUsers,omitempty" tf:"visible_to_all_users,omitempty"`
}

type BootstrapActionsFileObservation struct {
}

type BootstrapActionsFileParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type ConfigurationsFileObservation struct {
}

type ConfigurationsFileParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type CoreEBSBlockDeviceObservation struct {
}

type CoreEBSBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Required
	SizeInGb *int64 `json:"sizeInGb" tf:"size_in_gb,omitempty"`

	// +kubebuilder:validation:Required
	VolumeType *string `json:"volumeType" tf:"volume_type,omitempty"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type CoreScalingDownPolicyObservation struct {
}

type CoreScalingDownPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Adjustment *string `json:"adjustment,omitempty" tf:"adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// +kubebuilder:validation:Optional
	EvaluationPeriods *int64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// +kubebuilder:validation:Optional
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty" tf:"max_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Maximum *string `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Optional
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty" tf:"min_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Minimum *string `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Period *int64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

type CoreScalingUpPolicyObservation struct {
}

type CoreScalingUpPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Adjustment *string `json:"adjustment,omitempty" tf:"adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// +kubebuilder:validation:Optional
	EvaluationPeriods *int64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// +kubebuilder:validation:Optional
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty" tf:"max_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Maximum *string `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Optional
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty" tf:"min_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Minimum *string `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Period *int64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

type InstanceWeightsObservation struct {
}

type InstanceWeightsParameters struct {

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	WeightedCapacity *int64 `json:"weightedCapacity" tf:"weighted_capacity,omitempty"`
}

type MasterEBSBlockDeviceObservation struct {
}

type MasterEBSBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Required
	SizeInGb *int64 `json:"sizeInGb" tf:"size_in_gb,omitempty"`

	// +kubebuilder:validation:Required
	VolumeType *string `json:"volumeType" tf:"volume_type,omitempty"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type ProvisioningTimeoutObservation struct {
}

type ProvisioningTimeoutParameters struct {

	// +kubebuilder:validation:Required
	Timeout *int64 `json:"timeout" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Required
	TimeoutAction *string `json:"timeoutAction" tf:"timeout_action,omitempty"`
}

type ScheduledTaskObservation struct {
}

type ScheduledTaskParameters struct {

	// +kubebuilder:validation:Required
	Cron *string `json:"cron" tf:"cron,omitempty"`

	// +kubebuilder:validation:Optional
	DesiredCapacity *string `json:"desiredCapacity,omitempty" tf:"desired_capacity,omitempty"`

	// +kubebuilder:validation:Required
	InstanceGroupType *string `json:"instanceGroupType" tf:"instance_group_type,omitempty"`

	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MaxCapacity *string `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	MinCapacity *string `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// +kubebuilder:validation:Required
	TaskType *string `json:"taskType" tf:"task_type,omitempty"`
}

type StatementsObservation struct {
}

type StatementsParameters struct {

	// +kubebuilder:validation:Optional
	EvaluationPeriods *int64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Period *int64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit,omitempty"`
}

type StepsFileObservation struct {
}

type StepsFileParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type TagsObservation struct {
}

type TagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type TaskEBSBlockDeviceObservation struct {
}

type TaskEBSBlockDeviceParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Required
	SizeInGb *int64 `json:"sizeInGb" tf:"size_in_gb,omitempty"`

	// +kubebuilder:validation:Required
	VolumeType *string `json:"volumeType" tf:"volume_type,omitempty"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance,omitempty"`
}

type TaskScalingDownPolicyObservation struct {
}

type TaskScalingDownPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Adjustment *string `json:"adjustment,omitempty" tf:"adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// +kubebuilder:validation:Optional
	EvaluationPeriods *int64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// +kubebuilder:validation:Optional
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty" tf:"max_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Maximum *string `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Optional
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty" tf:"min_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Minimum *string `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Period *int64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

type TaskScalingUpPolicyObservation struct {
}

type TaskScalingUpPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Adjustment *string `json:"adjustment,omitempty" tf:"adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	Dimensions map[string]string `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

	// +kubebuilder:validation:Optional
	EvaluationPeriods *int64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`

	// +kubebuilder:validation:Optional
	MaxTargetCapacity *string `json:"maxTargetCapacity,omitempty" tf:"max_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Maximum *string `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// +kubebuilder:validation:Required
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// +kubebuilder:validation:Optional
	MinTargetCapacity *string `json:"minTargetCapacity,omitempty" tf:"min_target_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Minimum *string `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Period *int64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

type TerminationPoliciesObservation struct {
}

type TerminationPoliciesParameters struct {

	// +kubebuilder:validation:Required
	Statements []StatementsParameters `json:"statements" tf:"statements,omitempty"`
}

// AwsSpec defines the desired state of Aws
type AwsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsParameters `json:"forProvider"`
}

// AwsStatus defines the observed state of Aws.
type AwsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Aws is the Schema for the Awss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type Aws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpec   `json:"spec"`
	Status            AwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsList contains a list of Awss
type AwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Aws `json:"items"`
}

// Repository type metadata.
var (
	Aws_Kind             = "Aws"
	Aws_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Aws_Kind}.String()
	Aws_KindAPIVersion   = Aws_Kind + "." + CRDGroupVersion.String()
	Aws_GroupVersionKind = CRDGroupVersion.WithKind(Aws_Kind)
)

func init() {
	SchemeBuilder.Register(&Aws{}, &AwsList{})
}
