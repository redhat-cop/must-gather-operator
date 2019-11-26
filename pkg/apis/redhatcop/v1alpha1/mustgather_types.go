package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MustGatherSpec defines the desired state of MustGather
// +k8s:openapi-gen=true
type MustGatherSpec struct {
	// The is of the case this must gather will be uploaded to
	// +kubebuilder:validation:Required
	CaseID string `json:"caseid"`

	// the secret container a username and password field to be used to authenticate with red hat case management systems
	// +kubebuilder:validation:Required
	CaseManagementAccountSecretRef corev1.LocalObjectReference `json:"caseManagementSecretRef"`

	// the service account to use to run the must gather job pod, defaults to default
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="{Name:default}"
	ServiceAccountRef corev1.LocalObjectReference `json:"serviceAccountRef,omitempty"`

	// The list of must gather images to run, optional, it will default to: quay.io/openshift/origin-must-gather:latest
	// +kubebuilder:validation:Optional
	// +listType=set
	MustGatherImages []string `json:"mustGatherImages,omitempty"`
}

// MustGatherStatus defines the observed state of MustGather
// +k8s:openapi-gen=true
type MustGatherStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MustGather is the Schema for the mustgathers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=mustgathers,scope=Namespaced
type MustGather struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MustGatherSpec   `json:"spec,omitempty"`
	Status MustGatherStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MustGatherList contains a list of MustGather
type MustGatherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MustGather `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MustGather{}, &MustGatherList{})
}
