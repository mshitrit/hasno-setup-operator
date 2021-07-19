/*
Copyright 2021.

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

package controllers

import (
	"context"
	"os"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appv1alpha1 "github.com/mshitrit/hasno-setup-operator/api/v1alpha1"
)

// HALayerSetReconciler reconciles a HALayerSet object
type HALayerSetReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=app.hasno.com,resources=halayersets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.hasno.com,resources=halayersets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.hasno.com,resources=halayersets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HALayerSet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *HALayerSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//Assumptions
	//Clusters names are cluster1 & cluster2
	//Cluster domain is shared by both clusters

	//TODO mshitrit after phase1 is working consider exporting some of the assumptions

	log := r.Log.WithValues("halayerset", req.NamespacedName)
	log.Info("reconciling...")
	// your logic here
	//Get the resource
	setupResource := new(appv1alpha1.HALayerSet)
	key := client.ObjectKey{Name: req.Name, Namespace: req.Namespace}
	if err := r.Client.Get(ctx, key, setupResource); err != nil {
		//TODO mshitrit handle error
	}
	//Needed Data on CR
	r.setEnvVars(setupResource)

	//create a new manifest based on pod-cluster.yaml where the env vars are injected.
	//envsubst < pod-cluster.yaml > pod-cluster-new.yaml

	//create a pod with the new manifest

	//run make csr-cluster1 csr-cluster
	//set up fencing
	//test ?

	return ctrl.Result{}, nil
}

func (r *HALayerSetReconciler) setEnvVars(setupResource *appv1alpha1.HALayerSet) {
	//Cluster Name
	//Cluster1 IP
	//Cluster2 IP
	//AUTH (shared for both)
	//Extract & Set the Cluster Domain as env vars
	if err := os.Setenv(envVarNodeName, setupResource.Spec.NodeName); err != nil {
		//TODO mshitrit handle error
	}

	if err := os.Setenv(envVarNode1Addr, setupResource.Spec.Node1Address); err != nil {
		//TODO mshitrit handle error
	}

	if err := os.Setenv(envVarNode2Addr, setupResource.Spec.Node2Address); err != nil {
		//TODO mshitrit handle error
	}

	if err := os.Setenv(envVarAuthKey, setupResource.Spec.AuthKey); err != nil {
		//TODO mshitrit handle error
	}
}

const (
	envVarNode1Addr = "NODE1ADDR"
	envVarNode2Addr = "NODE2ADDR"
	envVarNodeName  = "NODENAME"
	envVarAuthKey   = "AUTHKEY"
)

// SetupWithManager sets up the controller with the Manager.
func (r *HALayerSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1alpha1.HALayerSet{}).
		Complete(r)
}
