package conditions

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// MockAzureCluster is the Schema for the mockazureclusters API
type MockAzureCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MockAzureClusterSpec   `json:"spec,omitempty"`
	Status MockAzureClusterStatus `json:"status,omitempty"`
}

// MockAzureClusterSpec defines the desired state of MockAzureCluster
type MockAzureClusterSpec struct{}

// MockAzureClusterStatus defines the observed state of MockAzureCluster
type MockAzureClusterStatus struct {
	Conditions capi.Conditions `json:"conditions,omitempty"`
}

// GetConditions returns the list of conditions for an MockAzureCluster API object.
func (c *MockAzureCluster) GetConditions() capi.Conditions {
	return c.Status.Conditions
}

// SetConditions will set the given conditions on an MockAzureCluster object
func (c *MockAzureCluster) SetConditions(conditions capi.Conditions) {
	c.Status.Conditions = conditions
}

func (c *MockAzureCluster) GetObjectKind() schema.ObjectKind { return c }

func (c *MockAzureCluster) GroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(c.TypeMeta.APIVersion, c.TypeMeta.Kind)
}

func (c *MockAzureCluster) DeepCopyObject() runtime.Object {
	panic("MockAzureCluster does not support DeepCopy")
}
