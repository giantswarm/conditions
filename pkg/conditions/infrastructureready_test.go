package conditions

import (
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

func TestGetInfrastructureReady(t *testing.T) {
	testTime := time.Now()

	testCases := []struct {
		name              string
		expectedCondition *capi.Condition
	}{
		{
			name: "case 0: InfrastructureReady with Status=True is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               InfrastructureReady,
				Status:             corev1.ConditionTrue,
				LastTransitionTime: metav1.NewTime(testTime),
			},
		},
		{
			name: "case 1: InfrastructureReady with Status=False is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               InfrastructureReady,
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
			infrastructureReady, conditionWasSet := GetInfrastructureReady(cluster)

			// assert
			if tc.expectedCondition != nil && conditionWasSet {
				areEqual := infrastructureReady.Type == tc.expectedCondition.Type &&
					infrastructureReady.Status == tc.expectedCondition.Status &&
					infrastructureReady.Severity == tc.expectedCondition.Severity &&
					infrastructureReady.Reason == tc.expectedCondition.Reason &&
					infrastructureReady.LastTransitionTime.Equal(&tc.expectedCondition.LastTransitionTime)

				if !areEqual {
					t.Logf(
						"InfrastructureReady was not set correctly, got %s, expected %s",
						sprintCondition(&infrastructureReady),
						sprintCondition(tc.expectedCondition))
					t.Fail()
				}
			} else if tc.expectedCondition == nil && !conditionWasSet {
				// all good
			} else if tc.expectedCondition != nil && !conditionWasSet {
				t.Logf("InfrastructureReady was not set, expected %s", sprintCondition(tc.expectedCondition))
				t.Fail()
			} else if tc.expectedCondition == nil && conditionWasSet {
				t.Logf("InfrastructureReady was not set to %s, expected nil", sprintCondition(&infrastructureReady))
				t.Fail()
			}
		})
	}
}

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

func TestUpdateInfrastructureReady(t *testing.T) {
	testCases := []struct {
		name                 string
		cluster              *capi.Cluster
		infrastructureObject Object
		expectedCondition    capi.Condition
	}{
		{
			name:                 "case 0: For nil infrastructure object, it sets InfrastructureReady(Status=False, Severity=Warning, Reason=InfrastructureObjectNotFound)",
			cluster:              &capi.Cluster{},
			infrastructureObject: nil,
			expectedCondition: capi.Condition{
				Type:     InfrastructureReady,
				Status:   corev1.ConditionFalse,
				Severity: capi.ConditionSeverityWarning,
				Reason:   InfrastructureObjectNotFoundReason,
			},
		},
		{
			name: "case 1: For 5min old Cluster and infrastructure object w/o Ready, it sets InfrastructureReady(Status=False, Severity=Info, Reason=WaitingForInfrastructureFallback)",
			cluster: &capi.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					CreationTimestamp: metav1.NewTime(time.Now().Add(-WaitingForInfrastructureWarningThresholdTime / 2)),
				},
			},
			infrastructureObject: &MockAzureCluster{
				Status: MockAzureClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
			expectedCondition: capi.Condition{
				Type:     InfrastructureReady,
				Status:   corev1.ConditionFalse,
				Severity: capi.ConditionSeverityInfo,
				Reason:   capi.WaitingForInfrastructureFallbackReason,
			},
		},
		{
			name: "case 2: For 20min old Cluster and infrastructure object w/o Ready, it sets InfrastructureReady status to False, Severity=Warning, Reason=WaitingForInfrastructureFallback",
			cluster: &capi.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					CreationTimestamp: metav1.NewTime(time.Now().Add(-2 * WaitingForInfrastructureWarningThresholdTime)),
				},
			},
			infrastructureObject: &MockAzureCluster{
				Status: MockAzureClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
			expectedCondition: capi.Condition{
				Type:     InfrastructureReady,
				Status:   corev1.ConditionFalse,
				Severity: capi.ConditionSeverityWarning,
				Reason:   capi.WaitingForInfrastructureFallbackReason,
			},
		},
		{
			name:    "case 3: For infrastructure object w/ Ready(Status=False), it sets InfrastructureReady(Status=False)",
			cluster: &capi.Cluster{},
			infrastructureObject: &MockAzureCluster{
				Status: MockAzureClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
			expectedCondition: capi.Condition{
				Type:   InfrastructureReady,
				Status: corev1.ConditionFalse,
			},
		},
		{
			name:    "case 4: For infrastructure object w/ Ready(Status=True), it sets InfrastructureReady(Status=True)",
			cluster: &capi.Cluster{},
			infrastructureObject: &MockAzureCluster{
				Status: MockAzureClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
			expectedCondition: capi.Condition{
				Type:   InfrastructureReady,
				Status: corev1.ConditionTrue,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// act
			UpdateInfrastructureReady(tc.cluster, tc.infrastructureObject)

			// assert
			infrastructureReady, ok := GetInfrastructureReady(tc.cluster)
			if ok {
				if !AreEqual(&infrastructureReady, &tc.expectedCondition) {
					t.Logf(
						"InfrastructureReady was not set correctly, got %s, expected %s",
						sprintCondition(&infrastructureReady),
						sprintCondition(&tc.expectedCondition))
					t.Fail()
				}
			} else {
				t.Log("InfrastructureReady was not set")
				t.Fail()
			}
		})
	}
}
