package kataset

import (
        "context"
        
	appsv1 "k8s.io/api/apps/v1"
        corev1 "k8s.io/api/core/v1"
        "k8s.io/apimachinery/pkg/api/errors"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/apimachinery/pkg/runtime"
        "k8s.io/apimachinery/pkg/types"
        mygroup2v1alpha1 "mydomain.com/mygroup2/kata-operator/pkg/apis/mygroup2/v1alpha1"
        "sigs.k8s.io/controller-runtime/pkg/client"
        "sigs.k8s.io/controller-runtime/pkg/controller"
        "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
        "sigs.k8s.io/controller-runtime/pkg/handler"
        logf "sigs.k8s.io/controller-runtime/pkg/log"
        "sigs.k8s.io/controller-runtime/pkg/manager"
        "sigs.k8s.io/controller-runtime/pkg/reconcile"
        "sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_kataset")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new KataSet Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
        return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
        return &ReconcileKataSet{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
        // Create a new controller
        c, err := controller.New("kataset-controller", mgr, controller.Options{Reconciler: r})
        if err != nil {
                return err
        }

        // Watch for changes to primary resource KataSet
        err = c.Watch(&source.Kind{Type: &mygroup2v1alpha1.KataSet{}}, &handler.EnqueueRequestForObject{})
        if err != nil {
                return err
        }

        // TODO(user): Modify this to be the types you create that are owned by the primary resource
        // Watch for changes to secondary resource Pods and requeue the owner KataSet
        err = c.Watch(&source.Kind{Type: &appsv1.DaemonSet{}}, &handler.EnqueueRequestForOwner{
                IsController: true,
                OwnerType:    &mygroup2v1alpha1.KataSet{},
        })
        if err != nil {
                return err
        }

        return nil
}

// blank assignment to verify that ReconcileKataSet implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileKataSet{}

// ReconcileKataSet reconciles a KataSet object
type ReconcileKataSet struct {
        // This client, initialized using mgr.Client() above, is a split client
        // that reads objects from the cache and writes to the apiserver
        client client.Client
        scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a KataSet object and makes changes based on the state read
// and what is in the KataSet.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileKataSet) Reconcile(request reconcile.Request) (reconcile.Result, error) {
        reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
        reqLogger.Info("Reconciling KataSet")

        // Fetch the KataSet instance
        instance := &mygroup2v1alpha1.KataSet{}
        err := r.client.Get(context.TODO(), request.NamespacedName, instance)
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
        // Define a new Pod object
        //pod := newDaemonset(instance)
 		ds := newDaemonset(instance)
        // Set KataSet instance as the owner and controller
        if err := controllerutil.SetControllerReference(instance, ds, r.scheme); err != nil {
                return reconcile.Result{}, err
        }

        // Check if this Pod already exists
        //found := &corev1.Pod{}
        found := &appsv1.DaemonSet{}
        err = r.client.Get(context.TODO(), types.NamespacedName{Name: ds.Name, Namespace: ds.Namespace}, found)
        if err != nil && errors.IsNotFound(err) {
                reqLogger.Info("Creating a new DaemonSet", "DaemonSet.Namespace", ds.Namespace, "DaemonSet.Name", ds.Name)
                err = r.client.Create(context.TODO(), ds)
                if err != nil {
                        return reconcile.Result{}, err
                }

                // Daemonset created successfully - don't requeue
                return reconcile.Result{}, nil
        } else if err != nil {
                return reconcile.Result{}, err
        }

        // DaemonSet already exists - don't requeue
        reqLogger.Info("Skip reconcile: DaemonSet already exists", "DaemonSet.Namespace", found.Namespace, "DaemonSet.Name", found.Name)
        return reconcile.Result{}, nil
}

// newDaemonset returns a kata-install pod with the same name/namespace as the cr
func newDaemonset(cr *mygroup2v1alpha1.KataSet) *appsv1.DaemonSet  {
        labels := map[string]string{
                "app": cr.Name,
        }
        //return &corev1.Pod{
        
        return &appsv1.DaemonSet{
    		TypeMeta: metav1.TypeMeta{
      			Kind: "DaemonSet",
      			APIVersion: "apps/v1",
    		},
    		ObjectMeta: metav1.ObjectMeta{
      			Name:      cr.Name + "-daemonset",
      			Namespace: cr.Namespace,
    		},
        	Spec: appsv1.DaemonSetSpec{
      			Selector: &metav1.LabelSelector{
        			MatchLabels: labels,
        		},
      		
        		Template: corev1.PodTemplateSpec{
                	ObjectMeta: metav1.ObjectMeta{
                        Name:      cr.Name + "-pod",
                        Namespace: cr.Namespace,
                        Labels:    labels,
                	},
                	Spec: corev1.PodSpec{
                        HostPID: true,
                        Containers: []corev1.Container{
                                {
                                        Name:    "kata-install",
                                        Image:   "sowjarao/kata-install",
                                        SecurityContext: &corev1.SecurityContext{
                                                Privileged: func() *bool { b := true; return &b }(),
                                        },
                                        VolumeMounts: []corev1.VolumeMount{{
                                                Name:  "install-script",
                                                MountPath: "/host",
                                        }},
                                },
                        },
                        Volumes: []corev1.Volume{{
                                Name: "install-script",
                                VolumeSource: corev1.VolumeSource{
                                      HostPath: &corev1.HostPathVolumeSource{
                                         Path: "/tmp/install",
                                      },
                                },
                        }},
                	},
        		},
        	},
   	 	}
}



                
