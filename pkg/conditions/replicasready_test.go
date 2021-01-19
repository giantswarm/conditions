package conditions

import (
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
)

func Test_GetReplicasReady(t *testing.T) {
	testTime := time.Now()

	testCases := []struct {
		name              string
		expectedCondition *capi.Condition
	}{
		{
			name: "case 0: ReplicasReady with Status=True is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               capiexp.ReplicasReadyCondition,
				Status:             corev1.ConditionTrue,
				LastTransitionTime: metav1.NewTime(testTime),
			},
		},
		{
			name: "case 1: ReplicasReady with Status=False is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               capiexp.ReplicasReadyCondition,
				Status:             corev1.ConditionFalse,
				LastTransitionTime: metav1.NewTime(testTime),
				Severity:           capi.ConditionSeverityWarning,
				Reason:             "FooBar",
				Message:            "Stuff happened",
			},
		},
		{
			name:              "case 2: Condition is neither set nor returned",
			expectedCondition: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			t.Log(tc.name)
			var machinePool *capiexp.MachinePool
			if tc.expectedCondition != nil {
				machinePool = &capiexp.MachinePool{
					Status: capiexp.MachinePoolStatus{
						Conditions: capi.Conditions{*tc.expectedCondition},
					},
				}
			} else {
				machinePool = &capiexp.MachinePool{}
			}

			// act
			outputCondition, conditionWasSet := GetReplicasReady(machinePool)

			// assert
			if tc.expectedCondition != nil && conditionWasSet {
				areEqual := AreEqual(&outputCondition, tc.expectedCondition)

				if !areEqual {
					t.Logf(
						"ReplicasReady was not set correctly, got %s, expected %s",
						sprintCondition(&outputCondition),
						sprintCondition(tc.expectedCondition))
					t.Fail()
				}
			} else if tc.expectedCondition == nil && !conditionWasSet {
				// all good
			} else if tc.expectedCondition != nil && !conditionWasSet {
				t.Logf("ReplicasReady was not set, expected %s", sprintCondition(tc.expectedCondition))
				t.Fail()
			} else if tc.expectedCondition == nil && conditionWasSet {
				t.Logf("ReplicasReady was not set to %s, expected nil", sprintCondition(&outputCondition))
				t.Fail()
			}
		})
	}
}

func Test_IsReplicasReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capiexp.MachinePool
		expectedOutput bool
	}{
		{
			name:           "case 0: IsReplicasReadyTrue returns true for CR with condition ReplicasReady with status True",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionTrue),
			expectedOutput: true,
		},
		{
			name:           "case 1: IsReplicasReadyTrue returns false for CR with condition ReplicasReady with status False",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsReplicasReadyTrue returns false for CR with condition ReplicasReady with status Unknown",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsReplicasReadyTrue returns false for CR without condition ReplicasReady",
			object:         machinePoolWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsReplicasReadyTrue returns false for CR with condition ReplicasReady with unsupported status",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, "SomethingSomething"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsReplicasReadyTrue(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsReplicasReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capiexp.MachinePool
		expectedOutput bool
	}{
		{
			name:           "case 0: IsReplicasReadyFalse returns false for CR with condition ReplicasReady with status True",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsReplicasReadyFalse returns true for CR with condition ReplicasReady with status False",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionFalse),
			expectedOutput: true,
		},
		{
			name:           "case 2: IsReplicasReadyFalse returns false for CR with condition ReplicasReady with status Unknown",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsReplicasReadyFalse returns false for CR without condition ReplicasReady",
			object:         machinePoolWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsReplicasReadyFalse returns false for CR with condition ReplicasReady with unsupported status",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, ""),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsReplicasReadyFalse(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}

func Test_IsReplicasReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		object         *capiexp.MachinePool
		expectedOutput bool
	}{
		{
			name:           "case 0: IsReplicasReadyUnknown returns false for CR with condition ReplicasReady with status True",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsReplicasReadyUnknown returns false for CR with condition ReplicasReady with status False",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsReplicasReadyUnknown returns true for CR with condition ReplicasReady with status Unknown",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, corev1.ConditionUnknown),
			expectedOutput: true,
		},
		{
			name:           "case 3: IsReplicasReadyUnknown returns true for CR without condition ReplicasReady",
			object:         machinePoolWithoutConditions(),
			expectedOutput: true,
		},
		{
			name:           "case 4: IsReplicasReadyUnknown returns false for CR with condition ReplicasReady with unsupported status",
			object:         machinePoolWith(capiexp.ReplicasReadyCondition, "SuperBrandNewStatusHere"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsReplicasReadyUnknown(tc.object)
			if result != tc.expectedOutput {
				t.Logf("expected %t, got %t", tc.expectedOutput, result)
				t.Fail()
			}
		})
	}
}
