package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

func Test_IsControlPlaneReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capi.Cluster
		expectedOutput bool
	}{
		{
			name:           "case 0: IsControlPlaneReadyTrue returns true for Cluster with condition ControlPlaneReady with status True",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionTrue),
			expectedOutput: true,
		},
		{
			name:           "case 1: IsControlPlaneReadyTrue returns false for Cluster with condition ControlPlaneReady with status False",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsControlPlaneReadyTrue returns false for Cluster with condition ControlPlaneReady with status Unknown",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsControlPlaneReadyTrue returns false for Cluster without condition ControlPlaneReady",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsControlPlaneReadyTrue returns false for Cluster with condition ControlPlaneReady with unsupported status",
			object:         clusterWith(ControlPlaneReady, "AnotherUnsupportedValue"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsControlPlaneReadyTrue(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsControlPlaneReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capi.Cluster
		expectedOutput bool
	}{
		{
			name:           "case 0: IsControlPlaneReadyFalse returns false for Cluster with condition ControlPlaneReady with status True",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsControlPlaneReadyFalse returns true for Cluster with condition ControlPlaneReady with status False",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionFalse),
			expectedOutput: true,
		},
		{
			name:           "case 2: IsControlPlaneReadyFalse returns false for Cluster with condition ControlPlaneReady with status Unknown",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsControlPlaneReadyFalse returns false for Cluster without condition ControlPlaneReady",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsControlPlaneReadyFalse returns false for Cluster with condition ControlPlaneReady with unsupported status",
			object:         clusterWith(ControlPlaneReady, ""),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsControlPlaneReadyFalse(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsControlPlaneReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capi.Cluster
		expectedOutput bool
	}{
		{
			name:           "case 0: IsControlPlaneReadyUnknown returns false for Cluster with condition ControlPlaneReady with status True",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsControlPlaneReadyUnknown returns false for Cluster with condition ControlPlaneReady with status False",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsControlPlaneReadyUnknown returns true for Cluster with condition ControlPlaneReady with status Unknown",
			object:         clusterWith(ControlPlaneReady, corev1.ConditionUnknown),
			expectedOutput: true,
		},
		{
			name:           "case 3: IsControlPlaneReadyUnknown returns true for Cluster without condition ControlPlaneReady",
			object:         clusterWithoutConditions(),
			expectedOutput: true,
		},
		{
			name:           "case 4: IsControlPlaneReadyUnknown returns false for Cluster with condition ControlPlaneReady with unsupported status",
			object:         clusterWith(ControlPlaneReady, "YouShallNotPass"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsControlPlaneReadyUnknown(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}
