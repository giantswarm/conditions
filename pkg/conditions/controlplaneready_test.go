package conditions

import (
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
)

func TestGetControlPlaneReady(t *testing.T) {
	testTime := time.Now()

	testCases := []struct {
		name              string
		expectedCondition *capi.Condition
	}{
		{
			name: "case 0: ControlPlaneReady with Status=True is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               ControlPlaneReady,
				Status:             corev1.ConditionTrue,
				LastTransitionTime: metav1.NewTime(testTime),
			},
		},
		{
			name: "case 1: ControlPlaneReady with Status=False is correctly returned",
			expectedCondition: &capi.Condition{
				Type:               ControlPlaneReady,
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
			outputCondition, conditionWasSet := GetControlPlaneReady(cluster)

			// assert
			if tc.expectedCondition != nil && conditionWasSet {
				areEqual := AreEqual(&outputCondition, tc.expectedCondition)

				if !areEqual {
					t.Logf(
						"ControlPlaneReady was not set correctly, got %s, expected %s",
						sprintCondition(&outputCondition),
						sprintCondition(tc.expectedCondition))
					t.Fail()
				}
			} else if tc.expectedCondition == nil && !conditionWasSet {
				// all good
			} else if tc.expectedCondition != nil && !conditionWasSet {
				t.Logf("ControlPlaneReady was not set, expected %s", sprintCondition(tc.expectedCondition))
				t.Fail()
			} else if tc.expectedCondition == nil && conditionWasSet {
				t.Logf("ControlPlaneReady was not set to %s, expected nil", sprintCondition(&outputCondition))
				t.Fail()
			}
		})
	}
}

func TestIsControlPlaneReadyTrue(t *testing.T) {
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

func TestIsControlPlaneReadyFalse(t *testing.T) {
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

func TestIsControlPlaneReadyUnknown(t *testing.T) {
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
