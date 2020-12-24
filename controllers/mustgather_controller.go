/*
Copyright 2020 Red Hat Community of Practice.

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
	"io/ioutil"
	"os"
	"reflect"
	"text/template"
	"time"

	"github.com/go-logr/logr"
	configv1 "github.com/openshift/api/config/v1"
	redhatcopv1alpha1 "github.com/redhat-cop/must-gather-operator/api/v1alpha1"
	"github.com/redhat-cop/operator-utils/pkg/util"
	"github.com/scylladb/go-set/strset"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const controllerName = "mustgather-controller"

const templateFileNameEnv = "JOB_TEMPLATE_FILE_NAME"
const defaultMustGatherImageEnv = "DEFAULT_MUST_GATHER_IMAGE"
const garbageCollectionElapsedEnv = "GARBAGE_COLLECTION_DELAY"

// MustGatherReconciler reconciles a MustGather object
type MustGatherReconciler struct {
	util.ReconcilerBase
	Log                       logr.Logger
	jobTemplate               *template.Template
	defaultMustGatherImage    string
	garbageCollectionDuration time.Duration
}

// +kubebuilder:rbac:groups=redhatcop.redhat.io,resources=mustgathers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=redhatcop.redhat.io,resources=mustgathers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=redhatcop.redhat.io,resources=mustgathers/finalizers,verbs=update
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=events,verbs=get;list;watch;create;patch
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups="config.openshift.io",resources=proxies,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MustGather object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *MustGatherReconciler) Reconcile(context context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("mustgather", req.NamespacedName)

	// Fetch the MustGather instance
	instance := &redhatcopv1alpha1.MustGather{}
	err := r.GetClient().Get(context, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if ok, err := r.IsValid(instance); !ok {
		return r.ManageError(context, instance, err)
	}

	if !r.IsInitialized(instance) {
		err := r.GetClient().Update(context, instance)
		if err != nil {
			log.Error(err, "unable to update instance", "instance", instance)
			return r.ManageError(context, instance, err)
		}
		return reconcile.Result{}, nil
	}

	//if job is complete and onject has been created more than 6 hrs ago delete instance
	if instance.Status.Completed && time.Since(instance.CreationTimestamp.Time).Milliseconds() > r.garbageCollectionDuration.Milliseconds() {
		err := r.DeleteResourceIfExists(context, instance)
		return reconcile.Result{}, err
	}

	job, err := r.getJobFromInstance(instance)
	if err != nil {
		log.Error(err, "unable to get job from", "instance", instance)
		return r.ManageError(context, instance, err)
	}

	job1 := &batchv1.Job{}
	err = r.GetClient().Get(context, types.NamespacedName{
		Name:      job.GetName(),
		Namespace: job.GetNamespace(),
	}, job1)

	if err != nil {
		if errors.IsNotFound(err) {
			// job is not there, create it.
			err = r.CreateResourceIfNotExists(context, instance, instance.GetNamespace(), job)
			//err = r.GetClient().Create(context.TODO(), job, &client.CreateOptions{})
			if err != nil {
				log.Error(err, "unable to create", "job", job)
				return r.ManageError(context, instance, err)
			}
			return r.ManageSuccess(context, instance)
		}
		// Error reading the object - requeue the request.
		log.Error(err, "unable to look up", "job", types.NamespacedName{
			Name:      job.GetName(),
			Namespace: job.GetNamespace(),
		})
		return r.ManageError(context, instance, err)
	}

	// if we get here it means that either
	// 1. the mustgather instance was updated, which we don't support and we are going to ignore
	// 2. the job was updated, probably the status piece. we should update the update the status of the instance, not supported yet.

	return r.updateStatus(context, instance, job1)
}

func (r *MustGatherReconciler) updateStatus(context context.Context, instance *redhatcopv1alpha1.MustGather, job *batchv1.Job) (reconcile.Result, error) {
	instance.Status.Completed = !job.Status.CompletionTime.IsZero()
	return r.ManageSuccess(context, instance)
}

// IsInitialized initializes the CR with default values if they are not specified.
func (r *MustGatherReconciler) IsInitialized(instance *redhatcopv1alpha1.MustGather) bool {
	initialized := true
	imageSet := strset.New(instance.Spec.MustGatherImages...)
	if !imageSet.Has(r.defaultMustGatherImage) {
		imageSet.Add(r.defaultMustGatherImage)
		instance.Spec.MustGatherImages = imageSet.List()
		initialized = false
	}
	if instance.Spec.ServiceAccountRef.Name == "" {
		instance.Spec.ServiceAccountRef.Name = "default"
		initialized = false
	}
	if reflect.DeepEqual(instance.Spec.ProxyConfig, configv1.ProxySpec{}) {
		platformProxy := &configv1.Proxy{}
		err := r.GetClient().Get(context.TODO(), types.NamespacedName{Name: "cluster"}, platformProxy)
		if err != nil {
			r.Log.Error(err, "unable to find cluster proxy configuration")
		} else {
			instance.Spec.ProxyConfig = redhatcopv1alpha1.ProxySpec{
				HTTPProxy:  platformProxy.Spec.HTTPProxy,
				HTTPSProxy: platformProxy.Spec.HTTPSProxy,
				NoProxy:    platformProxy.Spec.NoProxy,
			}
			initialized = false
		}
	}
	return initialized
}

func (r *MustGatherReconciler) getJobFromInstance(instance *redhatcopv1alpha1.MustGather) (*unstructured.Unstructured, error) {
	unstructuredJob, err := util.ProcessTemplate(instance, r.jobTemplate)
	if err != nil {
		r.Log.Error(err, "unable to process", "template", r.jobTemplate, "with parameter", instance)
		return &unstructured.Unstructured{}, err
	}
	return unstructuredJob, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MustGatherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.init()
	err := r.initializeTemplate()
	if err != nil {
		r.Log.Error(err, "unable to initialize job template")
		return err
	}

	isStateUpdated := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldJob, ok := e.ObjectOld.(*batchv1.Job)
			if !ok {
				return false
			}
			newJob, ok := e.ObjectNew.(*batchv1.Job)
			if !ok {
				return false
			}
			return !reflect.DeepEqual(oldJob.Status, newJob.Status)
		},
		CreateFunc: func(e event.CreateEvent) bool {
			return false
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return false
		},
		GenericFunc: func(e event.GenericEvent) bool {
			return false
		},
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&redhatcopv1alpha1.MustGather{}, builder.WithPredicates(util.ResourceGenerationOrFinalizerChangedPredicate{})).
		Owns(&batchv1.Job{}, builder.WithPredicates(isStateUpdated)).
		Complete(r)
}

func (r *MustGatherReconciler) initializeTemplate() error {
	templateFileName, ok := os.LookupEnv(templateFileNameEnv)
	if !ok {
		templateFileName = "/etc/templates/job.template.yaml"
	}
	text, err := ioutil.ReadFile(templateFileName)
	if err != nil {
		r.Log.Error(err, "Error reading job template file", "filename", templateFileName)
		return err
	}
	jobTemplate, err := template.New("MustGatherJob").Parse(string(text))
	if err != nil {
		r.Log.Error(err, "Error parsing template", "template", text)
		return err
	}
	r.jobTemplate = jobTemplate
	return err
}

func (r *MustGatherReconciler) init() {
	var ok bool
	r.defaultMustGatherImage, ok = os.LookupEnv(defaultMustGatherImageEnv)
	if !ok {
		r.defaultMustGatherImage = "quay.io/openshift/origin-must-gather:latest"
	}
	r.Log.Info("using default" + "must-gather-image: " + r.defaultMustGatherImage)
	garbageCollectionInterval, ok := os.LookupEnv(garbageCollectionElapsedEnv)
	if !ok {
		garbageCollectionInterval = "6h"
	}
	var err error
	r.garbageCollectionDuration, err = time.ParseDuration(garbageCollectionInterval)
	if err != nil {
		r.Log.Error(err, "unable to parse", "time", garbageCollectionInterval)
	}
}
