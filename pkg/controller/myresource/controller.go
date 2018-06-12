package myresource

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	mygroupv1alpha1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
	mygroupv1alpha1client "github.com/sdminonne/crd-webhook/pkg/client/clientset/versioned/typed/mygroup/v1alpha1"
	mygroupv1alpha1informer "github.com/sdminonne/crd-webhook/pkg/client/informers/externalversions/mygroup/v1alpha1"
	mygroupv1alpha1lister "github.com/sdminonne/crd-webhook/pkg/client/listers/mygroup/v1alpha1"

	"github.com/sdminonne/crd-webhook/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for Myresource resources goes here.

func (bc *MyresourceController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on myresource.MyresourceController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=mygroup,version=v1alpha1,kind=Myresource,resource=myresources
type MyresourceController struct {
	// INSERT ADDITIONAL FIELDS HERE
	myresourceLister mygroupv1alpha1lister.MyresourceLister
	myresourceclient mygroupv1alpha1client.MygroupV1alpha1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	myresourcerecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &MyresourceController{
		myresourceLister: arguments.ControllerManager.GetInformerProvider(&mygroupv1alpha1.Myresource{}).(mygroupv1alpha1informer.MyresourceInformer).Lister(),

		myresourceclient:   arguments.Clientset.MygroupV1alpha1(),
		myresourcerecorder: arguments.CreateRecorder("MyresourceController"),
	}

	// Create a new controller that will call MyresourceController.Reconcile on changes to Myresources
	gc := &controller.GenericController{
		Name:             "MyresourceController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&mygroupv1alpha1.Myresource{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a Myresource Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the MyresourceController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
