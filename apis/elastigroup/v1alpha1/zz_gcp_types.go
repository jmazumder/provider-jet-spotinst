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

type AccessConfigsObservation struct {
}

type AccessConfigsParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AliasIPRangesObservation struct {
}

type AliasIPRangesParameters struct {

	// +kubebuilder:validation:Required
	IPCidrRange *string `json:"ipCidrRange" tf:"ip_cidr_range,omitempty"`

	// +kubebuilder:validation:Required
	SubnetworkRangeName *string `json:"subnetworkRangeName" tf:"subnetwork_range_name,omitempty"`
}

type BackendServicesObservation struct {
}

type BackendServicesParameters struct {

	// +kubebuilder:validation:Optional
	LocationType *string `json:"locationType,omitempty" tf:"location_type,omitempty"`

	// +kubebuilder:validation:Optional
	NamedPorts []NamedPortsParameters `json:"namedPorts,omitempty" tf:"named_ports,omitempty"`

	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

type DiskObservation struct {
}

type DiskParameters struct {

	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// +kubebuilder:validation:Optional
	Boot *bool `json:"boot,omitempty" tf:"boot,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	InitializeParams []InitializeParamsParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`

	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GCPIntegrationDockerSwarmObservation struct {
}

type GCPIntegrationDockerSwarmParameters struct {

	// +kubebuilder:validation:Required
	MasterHost *string `json:"masterHost" tf:"master_host,omitempty"`

	// +kubebuilder:validation:Required
	MasterPort *int64 `json:"masterPort" tf:"master_port,omitempty"`
}

type GCPNetworkInterfaceObservation struct {
}

type GCPNetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	AccessConfigs []AccessConfigsParameters `json:"accessConfigs,omitempty" tf:"access_configs,omitempty"`

	// +kubebuilder:validation:Optional
	AliasIPRanges []AliasIPRangesParameters `json:"aliasIpRanges,omitempty" tf:"alias_ip_ranges,omitempty"`

	// +kubebuilder:validation:Required
	Network *string `json:"network" tf:"network,omitempty"`
}

type GCPObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GCPParameters struct {

	// +kubebuilder:validation:Optional
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`

	// +kubebuilder:validation:Optional
	BackendServices []BackendServicesParameters `json:"backendServices,omitempty" tf:"backend_services,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DesiredCapacity *int64 `json:"desiredCapacity" tf:"desired_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	Disk []DiskParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// +kubebuilder:validation:Optional
	DrainingTimeout *int64 `json:"drainingTimeout,omitempty" tf:"draining_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	FallbackToOndemand *bool `json:"fallbackToOndemand,omitempty" tf:"fallback_to_ondemand,omitempty"`

	// +kubebuilder:validation:Optional
	Gpu []GpuParameters `json:"gpu,omitempty" tf:"gpu,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckGracePeriod *int64 `json:"healthCheckGracePeriod,omitempty" tf:"health_check_grace_period,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// +kubebuilder:validation:Optional
	IPForwarding *bool `json:"ipForwarding,omitempty" tf:"ip_forwarding,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTypesCustom []InstanceTypesCustomParameters `json:"instanceTypesCustom,omitempty" tf:"instance_types_custom,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTypesOndemand *string `json:"instanceTypesOndemand,omitempty" tf:"instance_types_ondemand,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceTypesPreemptible []*string `json:"instanceTypesPreemptible,omitempty" tf:"instance_types_preemptible,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationDockerSwarm []GCPIntegrationDockerSwarmParameters `json:"integrationDockerSwarm,omitempty" tf:"integration_docker_swarm,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationGke []IntegrationGkeParameters `json:"integrationGke,omitempty" tf:"integration_gke,omitempty"`

	// +kubebuilder:validation:Optional
	Labels []LabelsParameters `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaxSize *int64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	MinSize *int64 `json:"minSize,omitempty" tf:"min_size,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterface []GCPNetworkInterfaceParameters `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	OndemandCount *int64 `json:"ondemandCount,omitempty" tf:"ondemand_count,omitempty"`

	// +kubebuilder:validation:Optional
	PreemptiblePercentage *int64 `json:"preemptiblePercentage,omitempty" tf:"preemptible_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisioningModel *string `json:"provisioningModel,omitempty" tf:"provisioning_model,omitempty"`

	// +kubebuilder:validation:Optional
	ScalingDownPolicy []GCPScalingDownPolicyParameters `json:"scalingDownPolicy,omitempty" tf:"scaling_down_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ScalingUpPolicy []GCPScalingUpPolicyParameters `json:"scalingUpPolicy,omitempty" tf:"scaling_up_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ScheduledTask []GCPScheduledTaskParameters `json:"scheduledTask,omitempty" tf:"scheduled_task,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// +kubebuilder:validation:Optional
	ShutdownScript *string `json:"shutdownScript,omitempty" tf:"shutdown_script,omitempty"`

	// +kubebuilder:validation:Optional
	StartupScript *string `json:"startupScript,omitempty" tf:"startup_script,omitempty"`

	// +kubebuilder:validation:Optional
	Subnets []SubnetsParameters `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UnhealthyDuration *int64 `json:"unhealthyDuration,omitempty" tf:"unhealthy_duration,omitempty"`
}

type GCPScalingDownPolicyDimensionsObservation struct {
}

type GCPScalingDownPolicyDimensionsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GCPScalingDownPolicyObservation struct {
}

type GCPScalingDownPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Adjustment *int64 `json:"adjustment,omitempty" tf:"adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	Dimensions []GCPScalingDownPolicyDimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

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

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

type GCPScalingUpPolicyDimensionsObservation struct {
}

type GCPScalingUpPolicyDimensionsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type GCPScalingUpPolicyObservation struct {
}

type GCPScalingUpPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// +kubebuilder:validation:Optional
	Adjustment *int64 `json:"adjustment,omitempty" tf:"adjustment,omitempty"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	Dimensions []GCPScalingUpPolicyDimensionsParameters `json:"dimensions,omitempty" tf:"dimensions,omitempty"`

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

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic,omitempty"`

	// +kubebuilder:validation:Required
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// +kubebuilder:validation:Required
	Unit *string `json:"unit" tf:"unit,omitempty"`
}

type GCPScheduledTaskObservation struct {
}

type GCPScheduledTaskParameters struct {

	// +kubebuilder:validation:Optional
	CronExpression *string `json:"cronExpression,omitempty" tf:"cron_expression,omitempty"`

	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MaxCapacity *string `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	MinCapacity *string `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	TargetCapacity *string `json:"targetCapacity,omitempty" tf:"target_capacity,omitempty"`

	// +kubebuilder:validation:Required
	TaskType *string `json:"taskType" tf:"task_type,omitempty"`
}

type GpuObservation struct {
}

type GpuParameters struct {

	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InitializeParamsObservation struct {
}

type InitializeParamsParameters struct {

	// +kubebuilder:validation:Optional
	DiskSizeGb *string `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Required
	SourceImage *string `json:"sourceImage" tf:"source_image,omitempty"`
}

type InstanceTypesCustomObservation struct {
}

type InstanceTypesCustomParameters struct {

	// +kubebuilder:validation:Required
	MemoryGib *int64 `json:"memoryGib" tf:"memory_gib,omitempty"`

	// +kubebuilder:validation:Required
	Vcpu *int64 `json:"vcpu" tf:"vcpu,omitempty"`
}

type IntegrationGkeAutoscaleDownObservation struct {
}

type IntegrationGkeAutoscaleDownParameters struct {

	// +kubebuilder:validation:Optional
	EvaluationPeriods *int64 `json:"evaluationPeriods,omitempty" tf:"evaluation_periods,omitempty"`
}

type IntegrationGkeAutoscaleHeadroomObservation struct {
}

type IntegrationGkeAutoscaleHeadroomParameters struct {

	// +kubebuilder:validation:Optional
	CPUPerUnit *int64 `json:"cpuPerUnit,omitempty" tf:"cpu_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	MemoryPerUnit *int64 `json:"memoryPerUnit,omitempty" tf:"memory_per_unit,omitempty"`

	// +kubebuilder:validation:Optional
	NumOfUnits *int64 `json:"numOfUnits,omitempty" tf:"num_of_units,omitempty"`
}

type IntegrationGkeAutoscaleLabelsObservation struct {
}

type IntegrationGkeAutoscaleLabelsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type IntegrationGkeObservation struct {
}

type IntegrationGkeParameters struct {

	// +kubebuilder:validation:Optional
	AutoUpdate *bool `json:"autoUpdate,omitempty" tf:"auto_update,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleCooldown *int64 `json:"autoscaleCooldown,omitempty" tf:"autoscale_cooldown,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleDown []IntegrationGkeAutoscaleDownParameters `json:"autoscaleDown,omitempty" tf:"autoscale_down,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleHeadroom []IntegrationGkeAutoscaleHeadroomParameters `json:"autoscaleHeadroom,omitempty" tf:"autoscale_headroom,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleIsAutoConfig *bool `json:"autoscaleIsAutoConfig,omitempty" tf:"autoscale_is_auto_config,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleIsEnabled *bool `json:"autoscaleIsEnabled,omitempty" tf:"autoscale_is_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscaleLabels []IntegrationGkeAutoscaleLabelsParameters `json:"autoscaleLabels,omitempty" tf:"autoscale_labels,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type LabelsObservation struct {
}

type LabelsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MetadataObservation struct {
}

type MetadataParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type NamedPortsObservation struct {
}

type NamedPortsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Ports []*string `json:"ports" tf:"ports,omitempty"`
}

type SubnetsObservation struct {
}

type SubnetsParameters struct {

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	SubnetNames []*string `json:"subnetNames" tf:"subnet_names,omitempty"`
}

// GCPSpec defines the desired state of GCP
type GCPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GCPParameters `json:"forProvider"`
}

// GCPStatus defines the observed state of GCP.
type GCPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GCPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GCP is the Schema for the GCPs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,spotinstjet}
type GCP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GCPSpec   `json:"spec"`
	Status            GCPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPList contains a list of GCPs
type GCPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCP `json:"items"`
}

// Repository type metadata.
var (
	GCP_Kind             = "GCP"
	GCP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GCP_Kind}.String()
	GCP_KindAPIVersion   = GCP_Kind + "." + CRDGroupVersion.String()
	GCP_GroupVersionKind = CRDGroupVersion.WithKind(GCP_Kind)
)

func init() {
	SchemeBuilder.Register(&GCP{}, &GCPList{})
}
