package inject

import (
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	mygroupv1alpha1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
	rscheme "github.com/sdminonne/crd-webhook/pkg/client/clientset/versioned/scheme"
	"github.com/sdminonne/crd-webhook/pkg/controller/myresource"
	"github.com/sdminonne/crd-webhook/pkg/inject/args"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&mygroupv1alpha1.Myresource{}, arguments.Informers.Mygroup().V1alpha1().Myresources()); err != nil {
			return err
		}

		// Add Kubernetes informers

		if c, err := myresource.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &mygroupv1alpha1.MyresourceCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"mygroup.amadeus.io"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "mygroup.amadeus.io",
		Version: "v1alpha1",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
