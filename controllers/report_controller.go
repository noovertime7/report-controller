/*
Copyright 2023.

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
	"github.com/noovertime7/report-controller/pkg/task"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	defaultv1 "github.com/noovertime7/report-controller/api/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// ReportReconciler reconciles a Report object
type ReportReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=default.report.tjkj.com,resources=reports,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=default.report.tjkj.com,resources=reports/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=default.report.tjkj.com,resources=reports/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Report object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *ReportReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	defer utilruntime.HandleCrash()

	// 获取report对象
	logger.Info("get report object", "name", req.String())
	report := &defaultv1.Report{}
	if err := r.Get(ctx, req.NamespacedName, report); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("report in deleting", "name", req.String())
			err := task.Mgr.Stop(req.Name)
			if err != nil {
				logger.Error(err, "停止任务失败")
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		}
		logger.Error(err, "get report error")
		return ctrl.Result{}, err
	}

	reportTask := task.NewReportTask(report)

	if err := task.Mgr.AddTaskWithJob(reportTask); err != nil {
		return ctrl.Result{}, err
	}

	if err := task.Mgr.Start(reportTask.GetName()); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ReportReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&defaultv1.Report{}).
		Complete(r)
}
