package main

import (
	"encoding/json"
	"fmt"

	"log"

	mygroupv1a1 "github.com/sdminonne/crd-webhook/pkg/apis/mygroup/v1alpha1"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/mattbaird/jsonpatch"
)

func createPatch(current, defaulted *mygroupv1a1.Myresource) ([]byte, error) {
	curStream, err := json.Marshal(current)
	if err != nil {
		return []byte{}, err
	}
	defStream, err := json.Marshal(defaulted)
	if err != nil {
		return []byte{}, err
	}
	jp, err := jsonpatch.CreatePatch(curStream, defStream)
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(jp)
}

func mutate(ar v1beta1.AdmissionReview) (reviewResponse *v1beta1.AdmissionResponse) {
	reviewResponse = &v1beta1.AdmissionResponse{}
	log.Printf("mutating (defaulting) custom resource")
	cr := struct {
		metav1.TypeMeta
		metav1.ObjectMeta
		Data map[string]string
	}{}

	if err := json.Unmarshal(ar.Request.Object.Raw, &cr); err != nil {
		log.Printf("Unable to unmarsall to a Kubernetes resource, %#v: %v", ar.Request.Object.Raw, err)
		reviewResponse.Result = &metav1.Status{
			Message: err.Error(),
		}
		return
	}

	log.Printf("mutating custom resource: %#v\n", cr)

	var mr mygroupv1a1.Myresource
	if err := json.Unmarshal(ar.Request.Object.Raw, &mr); err != nil {
		log.Printf("Unable to unmarsall to %T resource: %v", mr, err)
		reviewResponse.Result = &metav1.Status{
			Message: fmt.Sprintf("Unable to unmarsahll to %T resource: %v", mr, err.Error()),
		}
		return
	}

	defaultedMyResource := mr.DeepCopy()
	if len(defaultedMyResource.Spec.Afield) == 0 {
		defaultedMyResource.Spec.Afield = "A"
	}

	if len(defaultedMyResource.Spec.Bfield) == 0 {
		defaultedMyResource.Spec.Bfield = "B"
	}

	if len(defaultedMyResource.Spec.Cfield) == 0 {
		defaultedMyResource.Spec.Cfield = "C"
	}

	//jsonpatch.CreatePatch(
	patchStream, err := createPatch(&mr, defaultedMyResource)
	if err != nil {
		log.Printf("Unable to unmarsall to %T resource: %v", mr, err)
		reviewResponse.Result = &metav1.Status{
			Message: fmt.Sprintf("Unable to unmarsahll to %T resource: %v", mr, err.Error()),
		}
		return
	}
	log.Printf("AdmissionResponse: patch=%v\n", string(patchStream))
	reviewResponse.Allowed = true
	reviewResponse.Patch = patchStream
	reviewResponse.PatchType = func() *v1beta1.PatchType {
		pt := v1beta1.PatchTypeJSONPatch
		return &pt
	}()

	return
}
