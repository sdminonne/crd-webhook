package main

import (
	"encoding/json"
	"log"

	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mygroupv1a1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
)

func validate(ar v1beta1.AdmissionReview) (reviewResponse *v1beta1.AdmissionResponse) {
	reviewResponse = &v1beta1.AdmissionResponse{}

	cr := struct {
		metav1.TypeMeta
		metav1.ObjectMeta
		Data map[string]string
	}{}

	if err := json.Unmarshal(ar.Request.Object.Raw, &cr); err != nil {
		log.Printf("Unable to unmarsall to a Kubernetes resource, %#%v: %v", ar.Request.Object.Raw, err)
		reviewResponse.Result = &metav1.Status{
			Message: err.Error(),
		}
		return
	}

	log.Printf("validating custom resource: %#v\n", cr)
	var mr mygroupv1a1.Myresource
	if err := json.Unmarshal(ar.Request.Object.Raw, &mr); err != nil {
		log.Printf("unable to unmarshall resource: %v", err)
		reviewResponse.Result = &metav1.Status{
			Message: err.Error(),
		}
		return
	}
	if len(mr.Spec.Afield) == len(mr.Spec.Bfield) && len(mr.Spec.Afield) == len(mr.Spec.Cfield) {
		reviewResponse.Allowed = true
		return
	}
	reviewResponse.Allowed = false
	reviewResponse.Result = &metav1.Status{
		Reason: "spec.afield, spec.bfield, spec.cfield must have same length",
	}
	return
}
