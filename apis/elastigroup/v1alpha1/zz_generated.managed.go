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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this Aws.
func (mg *Aws) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Aws.
func (mg *Aws) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Aws.
func (mg *Aws) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Aws.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Aws) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Aws.
func (mg *Aws) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Aws.
func (mg *Aws) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Aws.
func (mg *Aws) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Aws.
func (mg *Aws) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Aws.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Aws) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Aws.
func (mg *Aws) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AwsBeanstalk.
func (mg *AwsBeanstalk) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AwsBeanstalk.
func (mg *AwsBeanstalk) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AwsBeanstalk.
func (mg *AwsBeanstalk) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AwsBeanstalk.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AwsBeanstalk) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AwsBeanstalk.
func (mg *AwsBeanstalk) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AwsBeanstalk.
func (mg *AwsBeanstalk) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AwsBeanstalk.
func (mg *AwsBeanstalk) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AwsBeanstalk.
func (mg *AwsBeanstalk) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AwsBeanstalk.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AwsBeanstalk) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AwsBeanstalk.
func (mg *AwsBeanstalk) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AwsSuspension.
func (mg *AwsSuspension) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AwsSuspension.
func (mg *AwsSuspension) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AwsSuspension.
func (mg *AwsSuspension) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AwsSuspension.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AwsSuspension) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AwsSuspension.
func (mg *AwsSuspension) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AwsSuspension.
func (mg *AwsSuspension) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AwsSuspension.
func (mg *AwsSuspension) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AwsSuspension.
func (mg *AwsSuspension) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AwsSuspension.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AwsSuspension) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AwsSuspension.
func (mg *AwsSuspension) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Azure.
func (mg *Azure) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Azure.
func (mg *Azure) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Azure.
func (mg *Azure) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Azure.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Azure) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Azure.
func (mg *Azure) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Azure.
func (mg *Azure) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Azure.
func (mg *Azure) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Azure.
func (mg *Azure) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Azure.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Azure) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Azure.
func (mg *Azure) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this AzureV3.
func (mg *AzureV3) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AzureV3.
func (mg *AzureV3) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AzureV3.
func (mg *AzureV3) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AzureV3.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AzureV3) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AzureV3.
func (mg *AzureV3) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AzureV3.
func (mg *AzureV3) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AzureV3.
func (mg *AzureV3) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AzureV3.
func (mg *AzureV3) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AzureV3.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AzureV3) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AzureV3.
func (mg *AzureV3) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this GCP.
func (mg *GCP) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this GCP.
func (mg *GCP) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this GCP.
func (mg *GCP) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this GCP.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *GCP) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this GCP.
func (mg *GCP) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this GCP.
func (mg *GCP) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this GCP.
func (mg *GCP) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this GCP.
func (mg *GCP) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this GCP.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *GCP) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this GCP.
func (mg *GCP) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Gke.
func (mg *Gke) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Gke.
func (mg *Gke) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Gke.
func (mg *Gke) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Gke.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Gke) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Gke.
func (mg *Gke) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Gke.
func (mg *Gke) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Gke.
func (mg *Gke) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Gke.
func (mg *Gke) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Gke.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Gke) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Gke.
func (mg *Gke) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
