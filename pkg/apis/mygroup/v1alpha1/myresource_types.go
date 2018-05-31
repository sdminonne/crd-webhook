package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the Myresource resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MyresourceSpec defines the desired state of Myresource
type MyresourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// MyresourceStatus defines the observed state of Myresource
type MyresourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "kubebuilder generate" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Myresource
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=myresources
type Myresource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyresourceSpec   `json:"spec,omitempty"`
	Status MyresourceStatus `json:"status,omitempty"`
}
