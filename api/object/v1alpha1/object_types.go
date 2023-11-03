// +kubebuilder:object:generate=true
// +groupName=embedded.aerf.io
// +versionName=v1alpha1
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ObjectSpec struct {
	// Raw YAML representation of the kubernetes object to be created.
	// +kubebuilder:validation:EmbeddedResource
	// +kubebuilder:pruning:PreserveUnknownFields
	Manifest runtime.RawExtension `json:"manifest"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ObjectSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Object
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}
