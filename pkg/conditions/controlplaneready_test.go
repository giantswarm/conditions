package conditions

import (
	"testing"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

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

type reconcileTestCase struct {
	name                      string
	cluster                   *capi.Cluster
	controlPlaneObject        Object
	expectedControlPlaneReady capi.Condition
}

func TestUpdateControlPlaneReady(t *testing.T) {
	testCases := []reconcileTestCase{
		{
			name:               "case 0: For nil control plane object, it sets ControlPlaneReady status to False, Severity=Warning, Reason=ControlPlaneObjectNotFound",
			cluster:            &capi.Cluster{},
			controlPlaneObject: nil,
			expectedControlPlaneReady: capi.Condition{
				Type:     ControlPlaneReady,
				Status:   corev1.ConditionFalse,
				Severity: capi.ConditionSeverityWarning,
				Reason:   ControlPlaneObjectNotFoundReason,
			},
		},
		{
			name: "case 1: For 5min old Cluster and control plane object w/o Ready, it sets ControlPlaneReady status to False, Severity=Info, Reason=WaitingForControlPlaneFallback",
			cluster: &capi.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					CreationTimestamp: metav1.NewTime(time.Now().Add(-WaitingForControlPlaneWarningThresholdTime / 2)),
				},
			},
			controlPlaneObject: machineWithoutConditions(),
			expectedControlPlaneReady: capi.Condition{
				Type:     ControlPlaneReady,
				Status:   corev1.ConditionFalse,
				Severity: capi.ConditionSeverityInfo,
				Reason:   capi.WaitingForControlPlaneFallbackReason,
			},
		},
		{
			name: "case 2: For 20min old Cluster and control plane object w/o Ready, it sets ControlPlaneReady status to False, Severity=Warning, Reason=WaitingForControlPlaneFallback",
			cluster: &capi.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					CreationTimestamp: metav1.NewTime(time.Now().Add(-2 * WaitingForControlPlaneWarningThresholdTime)),
				},
			},
			controlPlaneObject: machineWithoutConditions(),
			expectedControlPlaneReady: capi.Condition{
				Type:     ControlPlaneReady,
				Status:   corev1.ConditionFalse,
				Severity: capi.ConditionSeverityWarning,
				Reason:   capi.WaitingForControlPlaneFallbackReason,
			},
		},
		{
			name:    "case 3: For control plane object w/ Ready(Status=False), it sets ControlPlaneReady(Status=False)",
			cluster: &capi.Cluster{},
			controlPlaneObject: &capi.Machine{
				Status: capi.MachineStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
			expectedControlPlaneReady: capi.Condition{
				Type:   ControlPlaneReady,
				Status: corev1.ConditionFalse,
			},
		},
		{
			name:    "case 4: For control plane object w/ Ready(Status=True), it sets ControlPlaneReady(Status=True)",
			cluster: &capi.Cluster{},
			controlPlaneObject: &capi.Machine{
				Status: capi.MachineStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
			expectedControlPlaneReady: capi.Condition{
				Type:   ControlPlaneReady,
				Status: corev1.ConditionTrue,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			// act
			UpdateControlPlaneReady(tc.cluster, tc.controlPlaneObject)

			// assert
			controlPlaneReady, ok := GetControlPlaneReady(tc.cluster)
			if ok {
				if !AreEqual(&controlPlaneReady, &tc.expectedControlPlaneReady) {
					t.Logf(
						"ControlPlaneReady was not set correctly, got %s, expected %s",
						sprintCondition(&controlPlaneReady),
						sprintCondition(&tc.expectedControlPlaneReady))
					t.Fail()
				}
			} else {
				t.Log("ControlPlaneReady was not set")
				t.Fail()
			}
		})
	}
}
