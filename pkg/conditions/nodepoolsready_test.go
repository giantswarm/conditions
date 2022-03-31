package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
)

func TestIsNodePoolsReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capi.Cluster
		expectedOutput bool
	}{
		{
			name:           "case 0: IsNodePoolsReadyTrue returns true for Cluster with condition NodePoolsReady with status True",
			object:         clusterWith(NodePoolsReady, corev1.ConditionTrue),
			expectedOutput: true,
		},
		{
			name:           "case 1: IsNodePoolsReadyTrue returns false for Cluster with condition NodePoolsReady with status False",
			object:         clusterWith(NodePoolsReady, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsNodePoolsReadyTrue returns false for Cluster with condition NodePoolsReady with status Unknown",
			object:         clusterWith(NodePoolsReady, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsNodePoolsReadyTrue returns false for Cluster without condition NodePoolsReady",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsNodePoolsReadyTrue returns false for Cluster with condition NodePoolsReady with unsupported status",
			object:         clusterWith(NodePoolsReady, "AnotherUnsupportedValue"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsNodePoolsReadyTrue(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func TestIsNodePoolsReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capi.Cluster
		expectedOutput bool
	}{
		{
			name:           "case 0: IsNodePoolsReadyFalse returns false for Cluster with condition NodePoolsReady with status True",
			object:         clusterWith(NodePoolsReady, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsNodePoolsReadyFalse returns true for Cluster with condition NodePoolsReady with status False",
			object:         clusterWith(NodePoolsReady, corev1.ConditionFalse),
			expectedOutput: true,
		},
		{
			name:           "case 2: IsNodePoolsReadyFalse returns false for Cluster with condition NodePoolsReady with status Unknown",
			object:         clusterWith(NodePoolsReady, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsNodePoolsReadyFalse returns false for Cluster without condition NodePoolsReady",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsNodePoolsReadyFalse returns false for Cluster with condition NodePoolsReady with unsupported status",
			object:         clusterWith(NodePoolsReady, ""),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsNodePoolsReadyFalse(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func TestIsNodePoolsReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedOutput bool
		object         *capi.Cluster
	}{
		{
			name:           "case 0: IsNodePoolsReadyUnknown returns false for Cluster with condition NodePoolsReady with status True",
			object:         clusterWith(NodePoolsReady, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsNodePoolsReadyUnknown returns false for Cluster with condition NodePoolsReady with status False",
			object:         clusterWith(NodePoolsReady, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsNodePoolsReadyUnknown returns true for Cluster with condition NodePoolsReady with status Unknown",
			object:         clusterWith(NodePoolsReady, corev1.ConditionUnknown),
			expectedOutput: true,
		},
		{
			name:           "case 3: IsNodePoolsReadyUnknown returns true for Cluster without condition NodePoolsReady",
			object:         clusterWithoutConditions(),
			expectedOutput: true,
		},
		{
			name:           "case 4: IsNodePoolsReadyUnknown returns false for Cluster with condition NodePoolsReady with unsupported status",
			object:         clusterWith(NodePoolsReady, "YouShallNotPass"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsNodePoolsReadyUnknown(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}
