package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
)

func Test_IsCreatingTrue(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsCreatingTrue returns true for CR with condition Creating with status True",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsCreatingTrue returns false for CR with condition Creating with status False",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsCreatingTrue returns false for CR with condition Creating with status Unknown",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsCreatingTrue returns false for CR without condition Creating",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsCreatingTrue returns false for CR with condition Creating with unsupported status",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionStatus("AnotherUnsupportedValue"),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsCreatingTrue(tc.object)
			if result != tc.expectedResult {
				t.Logf(
					"expected IsCreatingTrue to return %t, got %t for %s",
					tc.expectedResult,
					result,
					conditionString(tc.object, Creating))
				t.Fail()
			}
		})
	}
}

func Test_IsCreatingFalse_ReturnsTrue(t *testing.T) {
	testCases := []struct {
		name         string
		object       Object
		checkOptions []CheckOption
	}{
		{
			name: "case 0: CR with condition Creating with Status=False",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
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

func Test_IsCreatingFalse_ReturnsFalse(t *testing.T) {
	testCases := []struct {
		name         string
		object       Object
		checkOptions []CheckOption
	}{
		{
			name: "case 0: IsCreatingFalse returns false for CR with condition Creating with status True",
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name: "case 1: IsCreatingFalse returns false for CR with condition Creating with status Unknown",
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name: "case 2: IsCreatingFalse returns false for CR without condition Creating",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name: "case 3: IsCreatingFalse returns false for CR with condition Creating with unsupported status",
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionStatus(""),
						},
					},
				},
			},
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

func Test_IsCreatingUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsCreatingUnknown returns false for CR with condition Creating with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsCreatingUnknown returns false for CR with condition Creating with status False",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsCreatingUnknown returns true for CR with condition Creating with status Unknown",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsCreatingUnknown returns true for CR without condition Creating",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsCreatingUnknown returns false for CR with condition Creating with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Creating,
							Status: corev1.ConditionStatus("BrandNewStatusHere"),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsCreatingUnknown(tc.object)
			if result != tc.expectedResult {
				t.Logf(
					"expected IsCreatingUnknown to return %t, got %t for %s",
					tc.expectedResult,
					result,
					conditionString(tc.object, Creating))
				t.Fail()
			}
		})
	}
}

func Test_MarkCreatingTrue_SetsCreatingConditionStatusToTrue(t *testing.T) {
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

func Test_MarkCreatingFalseWithCreationCompleted(t *testing.T) {
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

func Test_MarkCreatingFalseForExistingObject(t *testing.T) {
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
