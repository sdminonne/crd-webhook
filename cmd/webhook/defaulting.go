package main

import (
	"encoding/json"

	"log"

	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func mutate(ar v1beta1.AdmissionReview) (reviewResponse *v1beta1.AdmissionResponse) {
	reviewResponse = &v1beta1.AdmissionResponse{}
	log.Printf("mutating (defaulting) custom resource")
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

	log.Printf("mutating custom resource: %#v\n", cr)
	/*
		var mr mygroupv1a1.Myresource
		if err := json.Unmarshal(ar.Request.Object.Raw, &mr); err != nil {
		  reviewResponse.Result = &metav1.Status{
		Message: err.Error(),
						}
						return
					}

	*/
	log.Printf("Mutated!")
	reviewResponse.Allowed = true
	return
}
