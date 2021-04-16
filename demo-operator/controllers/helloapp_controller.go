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
	"fmt"

	appsv1 "github.com/anupamgogoi/memcached-operator/api/v1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// HelloAppReconciler reconciles a HelloApp object
type HelloAppReconciler struct {
	Client client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.anupam.com,resources=helloapps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.anupam.com,resources=helloapps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.anupam.com,resources=helloapps/finalizers,verbs=update
func (r *HelloAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("helloapp", req.NamespacedName)
	fmt.Println("CRD....")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HelloAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.HelloApp{}).
		Complete(r)
}
