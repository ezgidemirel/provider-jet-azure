/*
Copyright 2020 The Crossplane Authors.

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

package flexibleserverfirewallrule

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	tjcontroller "github.com/crossplane/terrajet/pkg/controller"
	"github.com/crossplane/terrajet/pkg/terraform"
	ctrl "sigs.k8s.io/controller-runtime"

	v1alpha2 "github.com/crossplane-contrib/provider-jet-azure/apis/dbforpostgresql/v1alpha2"
)

// Setup adds a controller that reconciles FlexibleServerFirewallRule managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha2.FlexibleServerFirewallRule_GroupVersionKind.String())
	var initializers managed.InitializerChain
	initializers = append(initializers, managed.NewNameAsExternalName(mgr.GetClient()))
	r := managed.NewReconciler(mgr,
		xpresource.ManagedKind(v1alpha2.FlexibleServerFirewallRule_GroupVersionKind),
		managed.WithExternalConnecter(tjcontroller.NewConnector(mgr.GetClient(), o.WorkspaceStore, o.SetupFn, o.Provider.Resources["azurerm_postgresql_flexible_server_firewall_rule"],
			tjcontroller.WithCallbackProvider(tjcontroller.NewAPICallbacks(mgr, xpresource.ManagedKind(v1alpha2.FlexibleServerFirewallRule_GroupVersionKind))),
		)),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(terraform.NewWorkspaceFinalizer(o.WorkspaceStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3*time.Minute),
		managed.WithInitializers(initializers),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha2.FlexibleServerFirewallRule{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}
