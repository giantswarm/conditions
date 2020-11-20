package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestIsUpgradingTrue(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsUpgradingTrue returns true for CR with condition Upgrading with status True",
			object:         clusterWith(Upgrading, corev1.ConditionTrue),
			expectedOutput: true,
		},
		{
			name:           "case 1: IsUpgradingTrue returns false for CR with condition Upgrading with status False",
			object:         machinePoolWith(Upgrading, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsUpgradingTrue returns false for CR with condition Upgrading with status Unknown",
			object:         clusterWith(Upgrading, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsUpgradingTrue returns false for CR without condition Upgrading",
			object:         machinePoolWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsUpgradingTrue returns false for CR with condition Upgrading with unsupported status",
			object:         clusterWith(Upgrading, corev1.ConditionStatus("SomeUnsupportedValue")),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsUpgradingTrue(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func TestIsUpgradingFalse(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		checkOptions   []CheckOption
		expectedOutput bool
	}{
		{
			name:           "case 0: IsUpgradingFalse returns false for CR with condition Upgrading with status True",
			object:         clusterWith(Upgrading, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsUpgradingFalse returns true for CR with condition Upgrading with status False",
			object:         machinePoolWith(Upgrading, corev1.ConditionFalse),
			expectedOutput: true,
		},
		{
			name:           "case 2: IsUpgradingFalse returns false for CR with condition Upgrading with status Unknown",
			object:         clusterWith(Upgrading, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsUpgradingFalse returns false for CR without condition Upgrading",
			object:         machinePoolWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsUpgradingFalse returns false for CR with condition Upgrading with unsupported status",
			object:         clusterWith(Upgrading, "ShinyNewStatusHere"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsUpgradingFalse(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func TestIsUpgradingUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsUpgradingUnknown returns false for CR with condition Upgrading with status True",
			object:         clusterWith(Upgrading, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsUpgradingUnknown returns false for CR with condition Upgrading with status False",
			object:         machinePoolWith(Upgrading, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsUpgradingUnknown returns true for CR with condition Upgrading with status Unknown",
			object:         clusterWith(Upgrading, corev1.ConditionUnknown),
			expectedOutput: true,
		},
		{
			name:           "case 3: IsUpgradingUnknown returns true for CR without condition Upgrading",
			object:         machinePoolWithoutConditions(),
			expectedOutput: true,
		},
		{
			name:           "case 4: IsUpgradingUnknown returns false for CR with condition Upgrading with unsupported status",
			object:         clusterWith(Upgrading, "IDonTKnowWhatIAmDoing"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsUpgradingUnknown(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}
