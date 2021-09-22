package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha4"
)

func Test_IsReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsReadyTrue returns true for CR with condition Ready with status True",
			object:         clusterWith(capi.ReadyCondition, corev1.ConditionTrue),
			expectedOutput: true,
		},
		{
			name:           "case 1: IsReadyTrue returns false for CR with condition Ready with status False",
			object:         machinePoolWith(capi.ReadyCondition, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsReadyTrue returns false for CR with condition Ready with status Unknown",
			object:         machinePoolWith(capi.ReadyCondition, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsReadyTrue returns false for CR without condition Ready",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsReadyTrue returns false for CR with condition Ready with unsupported status",
			object:         clusterWith(capi.ReadyCondition, "SomeUnsupportedValue"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsReadyTrue(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsReadyFalse returns false for CR with condition Ready with status True",
			object:         clusterWith(capi.ReadyCondition, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsReadyFalse returns true for CR with condition Ready with status False",
			object:         machinePoolWith(capi.ReadyCondition, corev1.ConditionFalse),
			expectedOutput: true,
		},
		{
			name:           "case 2: IsReadyFalse returns false for CR with condition Ready with status Unknown",
			object:         clusterWith(capi.ReadyCondition, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsReadyFalse returns false for CR without condition Ready",
			object:         machinePoolWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsReadyFalse returns false for CR with condition Ready with unsupported status",
			object:         machineWith(capi.ReadyCondition, "ShinyNewStatusHere"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsReadyFalse(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedOutput bool
		object         Object
	}{
		{
			name:           "case 0: IsReadyUnknown returns false for CR with condition Ready with status True",
			expectedOutput: false,
			object: &capi.Machine{
				Status: capi.MachineStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsReadyUnknown returns false for CR with condition Ready with status False",
			expectedOutput: false,
			object:         machinePoolWith(capi.ReadyCondition, corev1.ConditionFalse),
		},
		{
			name:           "case 2: IsReadyUnknown returns true for CR with condition Ready with status Unknown",
			object:         clusterWith(capi.ReadyCondition, corev1.ConditionUnknown),
			expectedOutput: true,
		},
		{
			name:           "case 3: IsReadyUnknown returns true for CR without condition Ready",
			object:         machineWithoutConditions(),
			expectedOutput: true,
		},
		{
			name:           "case 4: IsReadyUnknown returns false for CR with condition Ready with unsupported status",
			object:         clusterWith(capi.ReadyCondition, "IDonTKnowWhatIAmDoing"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsReadyUnknown(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}
