package conditions

import (
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

func TestGetCreating(t *testing.T) {
	testTime := time.Now()

	testCases := []struct {
		name              string
		expectedCondition *capi.Condition
	}{
		{
			name: "case 0: Creating with Status=True is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               Creating,
				Status:             corev1.ConditionTrue,
				LastTransitionTime: metav1.NewTime(testTime),
			},
		},
		{
			name: "case 1: Creating with Status=False is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               Creating,
				Status:             corev1.ConditionFalse,
				LastTransitionTime: metav1.NewTime(testTime),
				Severity:           capi.ConditionSeverityInfo,
				Reason:             CreationCompletedReason,
				Message:            "All good!",
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
			var cluster *capi.Cluster
			if tc.expectedCondition != nil {
				cluster = &capi.Cluster{
					Status: capi.ClusterStatus{
						Conditions: capi.Conditions{*tc.expectedCondition},
					},
				}
			} else {
				cluster = &capi.Cluster{}
			}

			// act
			outputCondition, conditionWasSet := GetCreating(cluster)

			// assert
			if tc.expectedCondition != nil && conditionWasSet {
				areEqual := AreEqual(&outputCondition, tc.expectedCondition)

				if !areEqual {
					t.Logf(
						"Creating was not set correctly, got %s, expected %s",
						sprintCondition(&outputCondition),
						sprintCondition(tc.expectedCondition))
					t.Fail()
				}
			} else if tc.expectedCondition == nil && !conditionWasSet {
				// all good
			} else if tc.expectedCondition != nil && !conditionWasSet {
				t.Logf("Creating was not set, expected %s", sprintCondition(tc.expectedCondition))
				t.Fail()
			} else if tc.expectedCondition == nil && conditionWasSet {
				t.Logf("Creating was not set to %s, expected nil", sprintCondition(&outputCondition))
				t.Fail()
			}
		})
	}
}

func TestIsCreatingTrue(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsCreatingTrue returns true for CR with condition Creating with status True",
			object:         machinePoolWith(Creating, corev1.ConditionTrue),
			expectedOutput: true,
		},
		{
			name:           "case 1: IsCreatingTrue returns false for CR with condition Creating with status False",
			object:         clusterWith(Creating, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsCreatingTrue returns false for CR with condition Creating with status Unknown",
			object:         machinePoolWith(Creating, corev1.ConditionUnknown),
			expectedOutput: false,
		},
		{
			name:           "case 3: IsCreatingTrue returns false for CR without condition Creating",
			object:         clusterWithoutConditions(),
			expectedOutput: false,
		},
		{
			name:           "case 4: IsCreatingTrue returns false for CR with condition Creating with unsupported status",
			object:         machinePoolWith(Creating, "AnotherUnsupportedValue"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsCreatingTrue(tc.object)
			if result != tc.expectedOutput {
				t.Logf(
					"expected IsCreatingTrue to return %t, got %t for %s",
					tc.expectedOutput,
					result,
					sprintConditionForObject(tc.object, Creating))
				t.Fail()
			}
		})
	}
}

func TestIsCreatingFalseReturnsTrue(t *testing.T) {
	testCases := []struct {
		name         string
		object       Object
		checkOptions []CheckOption
	}{
		{
			name:         "case 0: CR with condition Creating with Status=False",
			object:       clusterWith(Creating, corev1.ConditionFalse),
			checkOptions: []CheckOption{},
		},
		{
			name: "case 1: CR with condition Creating with Status=False, Reason=CreationCompleted with check option WithCreationCompletedReason()",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionFalse,
							Reason: CreationCompletedReason,
						},
					},
				},
			},
			checkOptions: []CheckOption{
				WithCreationCompletedReason(),
			},
		},
		{
			name: "case 2: CR with condition Creating with Status=False, Reason=ExistingObject with check option WithExistingObjectReason()",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionFalse,
							Reason: ExistingObjectReason,
						},
					},
				},
			},
			checkOptions: []CheckOption{
				WithExistingObjectReason(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsCreatingFalse(tc.object, tc.checkOptions...)
			if result != true {
				t.Logf(
					"expected IsCreatingFalse to return true, got false for %s",
					sprintConditionForObject(tc.object, Creating))

				t.Fail()
			}
		})
	}
}

func TestIsCreatingFalseReturnsFalse(t *testing.T) {
	testCases := []struct {
		name         string
		object       Object
		checkOptions []CheckOption
	}{
		{
			name:   "case 0: IsCreatingFalse returns false for CR with condition Creating with status True",
			object: machinePoolWith(Creating, corev1.ConditionTrue),
		},
		{
			name:   "case 1: IsCreatingFalse returns false for CR with condition Creating with status Unknown",
			object: machinePoolWith(Creating, corev1.ConditionUnknown),
		},
		{
			name:   "case 2: IsCreatingFalse returns false for CR without condition Creating",
			object: clusterWithoutConditions(),
		},
		{
			name:   "case 3: IsCreatingFalse returns false for CR with condition Creating with unsupported status",
			object: machinePoolWith(Creating, ""),
		},
		{
			name: "case 4: CR with condition Creating with Status=False, Reason=\"Whatever\" fails for check option WithCreationCompletedReason",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionFalse,
							Reason: "Whatever",
						},
					},
				},
			},
			checkOptions: []CheckOption{
				WithCreationCompletedReason(),
			},
		},
		{
			name: "case 4: CR with condition Creating with Status=False, Reason=\"ForReasons\" fails for check option WithExistingObjectReason",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionFalse,
							Reason: "ForReasons",
						},
					},
				},
			},
			checkOptions: []CheckOption{
				WithExistingObjectReason(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsCreatingFalse(tc.object, tc.checkOptions...)
			if result != false {
				t.Logf(
					"expected IsCreatingFalse to return false, got true for %s",
					sprintConditionForObject(tc.object, Creating))
				t.Fail()
			}
		})
	}
}

func TestIsCreatingUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		object         Object
		expectedOutput bool
	}{
		{
			name:           "case 0: IsCreatingUnknown returns false for CR with condition Creating with status True",
			object:         clusterWith(Creating, corev1.ConditionTrue),
			expectedOutput: false,
		},
		{
			name:           "case 1: IsCreatingUnknown returns false for CR with condition Creating with status False",
			object:         clusterWith(Creating, corev1.ConditionFalse),
			expectedOutput: false,
		},
		{
			name:           "case 2: IsCreatingUnknown returns true for CR with condition Creating with status Unknown",
			object:         clusterWith(Creating, corev1.ConditionUnknown),
			expectedOutput: true,
		},
		{
			name:           "case 3: IsCreatingUnknown returns true for CR without condition Creating",
			object:         machinePoolWithoutConditions(),
			expectedOutput: true,
		},
		{
			name:           "case 4: IsCreatingUnknown returns false for CR with condition Creating with unsupported status",
			object:         clusterWith(Creating, "BrandNewStatusHere"),
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsCreatingUnknown(tc.object)
			if result != tc.expectedOutput {
				t.Logf(
					"expected IsCreatingUnknown to return %t, got %t for %s",
					tc.expectedOutput,
					result,
					sprintConditionForObject(tc.object, Creating))
				t.Fail()
			}
		})
	}
}
