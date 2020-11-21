package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

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
					conditionString(tc.object, Creating))
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
					conditionString(tc.object, Creating))

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
					conditionString(tc.object, Creating))
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
					conditionString(tc.object, Creating))
				t.Fail()
			}
		})
	}
}

func TestMarkCreatingTrue(t *testing.T) {
	testName := "MarkCreatingTrue sets Creating condition status to True"
	t.Run(testName, func(t *testing.T) {
		// arrange
		t.Log(testName)
		cluster := &capi.Cluster{}

		// act
		MarkCreatingTrue(cluster)

		// assert
		if !IsCreatingTrue(cluster) {
			got := getGotConditionStatusString(cluster, Creating)
			t.Logf("expected that Creating condition status is set to True, got %s", got)
			t.Fail()
		}
	})
}

func TestMarkCreatingFalseWithCreationCompleted(t *testing.T) {
	testName := "Creating condition is set with Status=False, Severity=Info, Reason=CreationCompleted"
	t.Run(testName, func(t *testing.T) {
		// arrange
		t.Log(testName)
		cluster := &capi.Cluster{}

		// act
		MarkCreatingFalseWithCreationCompleted(cluster)

		// assert
		expected := IsCreatingFalse(cluster, WithSeverityInfo(), WithCreationCompletedReason())
		if !expected {
			t.Logf(
				"expected that Creating condition is set with Status=False, Severity=Info, Reason=CreationCompleted, got %s",
				conditionString(cluster, Creating))
			t.Fail()
		}
	})
}

func TestMarkCreatingFalseForExistingObject(t *testing.T) {
	testName := "Creating condition is set with Status=False, Severity=Info, Reason=ExistingObject"
	t.Run(testName, func(t *testing.T) {
		// arrange
		t.Log(testName)
		cluster := &capi.Cluster{}

		// act
		MarkCreatingFalseForExistingObject(cluster)

		// assert
		expected := IsCreatingFalse(cluster, WithSeverityInfo(), WithExistingObjectReason())
		if !expected {
			t.Logf(
				"expected that Creating condition is set with Status=False, Severity=Info, Reason=ExistingObject, got %s",
				conditionString(cluster, Creating))
			t.Fail()
		}
	})
}
