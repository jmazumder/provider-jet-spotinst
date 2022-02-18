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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/jmazumder/provider-jet-spotinst/apis/elastigroup/v1alpha1"
	v1alpha1health "github.com/jmazumder/provider-jet-spotinst/apis/health/v1alpha1"
	v1alpha1managed "github.com/jmazumder/provider-jet-spotinst/apis/managed/v1alpha1"
	v1alpha1mrscaler "github.com/jmazumder/provider-jet-spotinst/apis/mrscaler/v1alpha1"
	v1alpha1multai "github.com/jmazumder/provider-jet-spotinst/apis/multai/v1alpha1"
	v1alpha1ocean "github.com/jmazumder/provider-jet-spotinst/apis/ocean/v1alpha1"
	v1alpha1spotinst "github.com/jmazumder/provider-jet-spotinst/apis/spotinst/v1alpha1"
	v1alpha1apis "github.com/jmazumder/provider-jet-spotinst/apis/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1health.SchemeBuilder.AddToScheme,
		v1alpha1managed.SchemeBuilder.AddToScheme,
		v1alpha1mrscaler.SchemeBuilder.AddToScheme,
		v1alpha1multai.SchemeBuilder.AddToScheme,
		v1alpha1ocean.SchemeBuilder.AddToScheme,
		v1alpha1spotinst.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}