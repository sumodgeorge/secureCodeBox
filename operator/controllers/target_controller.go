/*


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
	"time"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"

	scansv1 "experimental.securecodebox.io/api/v1"
)

// ScanInterval defines how often a Target should be scanned
var ScanInterval time.Duration = 1 * time.Minute

// TargetReconciler reconciles a Target object
type TargetReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=scans.experimental.securecodebox.io,resources=targets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=scans.experimental.securecodebox.io,resources=targets/status,verbs=get;update;patch

func (r *TargetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("target", req.NamespacedName)

	var target scansv1.Target
	err := r.Get(ctx, req.NamespacedName, &target)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var childScans scansv1.ScanList
	if err := r.List(ctx, &childScans, client.InNamespace(req.Namespace), client.MatchingFields{ownerKey: req.Name}); err != nil {
		log.Error(err, "unable to list child Scans")
		return ctrl.Result{}, err
	}

	var nextSchedule time.Time
	if target.Status.LastScheduleTime != nil {
		nextSchedule = target.Status.LastScheduleTime.Add(ScanInterval)
	} else {
		nextSchedule = time.Now().Add(-1 * time.Second)
	}

	// check if it is time to start the scans
	if !time.Now().Before(nextSchedule) {
		// It's time!
		log.Info("Should start scans here")

		var now metav1.Time = metav1.Now()
		target.Status.LastScheduleTime = &now
		if err := r.Status().Update(ctx, &target); err != nil {
			log.Error(err, "unable to update Targets status")
			return ctrl.Result{}, err
		}

		// Recalculate next schedule
		nextSchedule = time.Now().Add(ScanInterval)
	}

	return ctrl.Result{RequeueAfter: nextSchedule.Sub(time.Now())}, nil
}

func (r *TargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(&scansv1.Scan{}, ownerKey, func(rawObj runtime.Object) []string {
		// grab the job object, extract the owner...
		scan := rawObj.(*scansv1.Scan)
		owner := metav1.GetControllerOf(scan)
		if owner == nil {
			return nil
		}
		// ...make sure it's a Scan belonging to a Target...
		if owner.APIVersion != apiGVStr || owner.Kind != "Target" {
			return nil
		}

		// ...and if so, return it
		return []string{owner.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&scansv1.Target{}).
		Owns(&scansv1.Scan{}).
		Complete(r)
}