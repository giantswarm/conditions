package conditions

import (
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
)

func TestGetUpgrading(t *testing.T) {
	testTime := time.Now()

	testCases := []struct {
		name              string
		expectedCondition *capi.Condition
	}{
		{
			name: "case 0: Upgrading with Status=True is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               Upgrading,
				Status:             corev1.ConditionTrue,
				LastTransitionTime: metav1.NewTime(testTime),
			},
		},
		{
			name: "case 1: Upgrading with Status=False is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               Upgrading,
				Status:             corev1.ConditionFalse,
				LastTransitionTime: metav1.NewTime(testTime),
				Severity:           capi.ConditionSeverityInfo,
				Reason:             UpgradeNotStartedReason,
				Message:            "Nothing new here!",
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
			var cr *capiexp.MachinePool
			if tc.expectedCondition != nil {
				cr = &capiexp.MachinePool{
					Status: capiexp.MachinePoolStatus{
						Conditions: capi.Conditions{*tc.expectedCondition},
					},
				}
			} else {
				cr = &capiexp.MachinePool{}
			}

			// act
			condition, conditionWasSet := GetUpgrading(cr)

			// assert
			if tc.expectedCondition != nil && conditionWasSet {
				areEqual := condition.Type == tc.expectedCondition.Type &&
					condition.Status == tc.expectedCondition.Status &&
					condition.Severity == tc.expectedCondition.Severity &&
					condition.Reason == tc.expectedCondition.Reason &&
					condition.LastTransitionTime.Equal(&tc.expectedCondition.LastTransitionTime)

				if !areEqual {
					t.Logf(
						"Upgrading was not set correctly, got %s, expected %s",
						sprintCondition(&condition),
						sprintCondition(tc.expectedCondition))
					t.Fail()
				}
			} else if tc.expectedCondition == nil && !conditionWasSet {
				// all good
			} else if tc.expectedCondition != nil && !conditionWasSet {
				t.Logf("Upgrading was not set, expected %s", sprintCondition(tc.expectedCondition))
				t.Fail()
			} else if tc.expectedCondition == nil && conditionWasSet {
				t.Logf("Upgrading was not set to %s, expected nil", sprintCondition(&condition))
				t.Fail()
			}
		})
	}
}

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
			object:         clusterWith(Upgrading, "SomeUnsupportedValue"),
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
		{
			name: "case 5: CR with condition Upgrading(Status=False, Reason=\"UpgradeCompleted\") for check options WithUpgradeCompletedReason() and WithSeverityInfo()",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:     Upgrading,
							Status:   corev1.ConditionFalse,
							Reason:   UpgradeCompletedReason,
							Severity: capi.ConditionSeverityInfo,
						},
					},
				},
			},
			checkOptions:   []CheckOption{WithUpgradeCompletedReason(), WithSeverityInfo()},
			expectedOutput: true,
		},
		{
			name: "case 6: CR with condition Upgrading(Status=False, Reason=\"ForReasons\") for check options WithUpgradeNotStartedReason() and WithSeverityInfo()",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:     Upgrading,
							Status:   corev1.ConditionFalse,
							Reason:   UpgradeNotStartedReason,
							Severity: capi.ConditionSeverityInfo,
						},
					},
				},
			},
			checkOptions:   []CheckOption{WithUpgradeNotStartedReason(), WithSeverityInfo()},
			expectedOutput: true,
		},
		{
			name: "case 7: CR with condition Upgrading(Status=False, Reason=\"Whatever\") fails for check option WithUpgradeCompletedReason()",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionFalse,
							Reason: "Whatever",
						},
					},
				},
			},
			checkOptions: []CheckOption{
				WithUpgradeCompletedReason(),
			},
			expectedOutput: false,
		},
		{
			name: "case 8: CR with condition Upgrading(Status=False, Reason=\"ForReasons\") fails for check option WithUpgradeNotStartedReason()",
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionFalse,
							Reason: "ForReasons",
						},
					},
				},
			},
			checkOptions: []CheckOption{
				WithUpgradeNotStartedReason(),
			},
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsUpgradingFalse(tc.object, tc.checkOptions...)
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
