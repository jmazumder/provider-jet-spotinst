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

package controller

import (
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/logging"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/terraform"

	aws "github.com/jmazumder/provider-jet-spotinst/internal/controller/elastigroup/aws"
	awsbeanstalk "github.com/jmazumder/provider-jet-spotinst/internal/controller/elastigroup/awsbeanstalk"
	awssuspension "github.com/jmazumder/provider-jet-spotinst/internal/controller/elastigroup/awssuspension"
	azure "github.com/jmazumder/provider-jet-spotinst/internal/controller/elastigroup/azure"
	azurev3 "github.com/jmazumder/provider-jet-spotinst/internal/controller/elastigroup/azurev3"
	gcp "github.com/jmazumder/provider-jet-spotinst/internal/controller/elastigroup/gcp"
	gke "github.com/jmazumder/provider-jet-spotinst/internal/controller/elastigroup/gke"
	check "github.com/jmazumder/provider-jet-spotinst/internal/controller/health/check"
	instanceaws "github.com/jmazumder/provider-jet-spotinst/internal/controller/managed/instanceaws"
	awsmrscaler "github.com/jmazumder/provider-jet-spotinst/internal/controller/mrscaler/aws"
	balancer "github.com/jmazumder/provider-jet-spotinst/internal/controller/multai/balancer"
	deployment "github.com/jmazumder/provider-jet-spotinst/internal/controller/multai/deployment"
	listener "github.com/jmazumder/provider-jet-spotinst/internal/controller/multai/listener"
	routingrule "github.com/jmazumder/provider-jet-spotinst/internal/controller/multai/routingrule"
	target "github.com/jmazumder/provider-jet-spotinst/internal/controller/multai/target"
	targetset "github.com/jmazumder/provider-jet-spotinst/internal/controller/multai/targetset"
	aks "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/aks"
	aksvirtualnodegroup "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/aksvirtualnodegroup"
	awsocean "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/aws"
	awslaunchspec "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/awslaunchspec"
	ecs "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/ecs"
	ecslaunchspec "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/ecslaunchspec"
	gkeimport "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/gkeimport"
	gkelaunchspec "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/gkelaunchspec"
	gkelaunchspecimport "github.com/jmazumder/provider-jet-spotinst/internal/controller/ocean/gkelaunchspecimport"
	providerconfig "github.com/jmazumder/provider-jet-spotinst/internal/controller/providerconfig"
	subscription "github.com/jmazumder/provider-jet-spotinst/internal/controller/spotinst/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, l logging.Logger, wl workqueue.RateLimiter, ps terraform.SetupFn, ws *terraform.WorkspaceStore, cfg *tjconfig.Provider, concurrency int) error {
	for _, setup := range []func(ctrl.Manager, logging.Logger, workqueue.RateLimiter, terraform.SetupFn, *terraform.WorkspaceStore, *tjconfig.Provider, int) error{
		aws.Setup,
		awsbeanstalk.Setup,
		awssuspension.Setup,
		azure.Setup,
		azurev3.Setup,
		gcp.Setup,
		gke.Setup,
		check.Setup,
		instanceaws.Setup,
		awsmrscaler.Setup,
		balancer.Setup,
		deployment.Setup,
		listener.Setup,
		routingrule.Setup,
		target.Setup,
		targetset.Setup,
		aks.Setup,
		aksvirtualnodegroup.Setup,
		awsocean.Setup,
		awslaunchspec.Setup,
		ecs.Setup,
		ecslaunchspec.Setup,
		gkeimport.Setup,
		gkelaunchspec.Setup,
		gkelaunchspecimport.Setup,
		providerconfig.Setup,
		subscription.Setup,
	} {
		if err := setup(mgr, l, wl, ps, ws, cfg, concurrency); err != nil {
			return err
		}
	}
	return nil
}
