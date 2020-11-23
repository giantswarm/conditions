package conditions

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// AzureCluster is the Schema for the azureclusters API
type MockAzureCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MockAzureClusterSpec   `json:"spec,omitempty"`
	Status MockAzureClusterStatus `json:"status,omitempty"`
}

// AzureClusterSpec defines the desired state of AzureCluster
type MockAzureClusterSpec struct{}

// AzureClusterStatus defines the observed state of AzureCluster
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
