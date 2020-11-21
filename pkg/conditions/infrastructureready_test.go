package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func Test_IsInfrastructureReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsInfrastructureReadyTrue returns true for CR with condition InfrastructureReady with status True",
			object:         machinePoolWith(InfrastructureReady, corev1.ConditionTrue),
			expectedOutput: true,
		},
		{
			name:           "case 1: IsInfrastructureReadyTrue returns false for CR with condition InfrastructureReady with status False",
			object:         clusterWith(InfrastructureReady, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsInfrastructureReadyTrue returns false for CR with condition InfrastructureReady with status Unknown",
			object:         machinePoolWith(InfrastructureReady, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsInfrastructureReadyTrue returns false for CR without condition InfrastructureReady",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsInfrastructureReadyTrue returns false for CR with condition InfrastructureReady with unsupported status",
			object:         machinePoolWith(InfrastructureReady, "AnotherUnsupportedValue"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsInfrastructureReadyTrue(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsInfrastructureReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsInfrastructureReadyFalse returns false for CR with condition InfrastructureReady with status True",
			object:         machinePoolWith(InfrastructureReady, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsInfrastructureReadyFalse returns true for CR with condition InfrastructureReady with status False",
			object:         clusterWith(InfrastructureReady, corev1.ConditionFalse),
			expectedOutput: true,
		},
		{
			name:           "case 2: IsInfrastructureReadyFalse returns false for CR with condition InfrastructureReady with status Unknown",
			object:         machinePoolWith(InfrastructureReady, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsInfrastructureReadyFalse returns false for CR without condition InfrastructureReady",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsInfrastructureReadyFalse returns false for CR with condition InfrastructureReady with unsupported status",
			object:         machinePoolWith(InfrastructureReady, ""),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsInfrastructureReadyFalse(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsInfrastructureReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsInfrastructureReadyUnknown returns false for CR with condition InfrastructureReady with status True",
			object:         clusterWith(InfrastructureReady, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsInfrastructureReadyUnknown returns false for CR with condition InfrastructureReady with status False",
			object:         machinePoolWith(InfrastructureReady, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsInfrastructureReadyUnknown returns true for CR with condition InfrastructureReady with status Unknown",
			object:         clusterWith(InfrastructureReady, corev1.ConditionUnknown),
			expectedOutput: true,
		},
		{
			name:           "case 3: IsInfrastructureReadyUnknown returns true for CR without condition InfrastructureReady",
			object:         machinePoolWithoutConditions(),
			expectedOutput: true,
		},
		{
			name:           "case 4: IsInfrastructureReadyUnknown returns false for CR with condition InfrastructureReady with unsupported status",
			object:         clusterWith(InfrastructureReady, "BrandNewStatusHere"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsInfrastructureReadyUnknown(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}
