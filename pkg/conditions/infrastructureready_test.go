package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
)

func Test_IsInfrastructureReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsInfrastructureReadyTrue returns true for CR with condition InfrastructureReady with status True",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsInfrastructureReadyTrue returns false for CR with condition InfrastructureReady with status False",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsInfrastructureReadyTrue returns false for CR with condition InfrastructureReady with status Unknown",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsInfrastructureReadyTrue returns false for CR without condition InfrastructureReady",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsInfrastructureReadyTrue returns false for CR with condition InfrastructureReady with unsupported status",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
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

			result := IsInfrastructureReadyTrue(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsInfrastructureReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsInfrastructureReadyFalse returns false for CR with condition InfrastructureReady with status True",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsInfrastructureReadyFalse returns true for CR with condition InfrastructureReady with status False",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsInfrastructureReadyFalse returns false for CR with condition InfrastructureReady with status Unknown",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsInfrastructureReadyFalse returns false for CR without condition InfrastructureReady",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsInfrastructureReadyFalse returns false for CR with condition InfrastructureReady with unsupported status",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionStatus(""),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsInfrastructureReadyFalse(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsInfrastructureReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsInfrastructureReadyUnknown returns false for CR with condition InfrastructureReady with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsInfrastructureReadyUnknown returns false for CR with condition InfrastructureReady with status False",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsInfrastructureReadyUnknown returns true for CR with condition InfrastructureReady with status Unknown",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsInfrastructureReadyUnknown returns true for CR without condition InfrastructureReady",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsInfrastructureReadyUnknown returns false for CR with condition InfrastructureReady with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   InfrastructureReady,
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

			result := IsInfrastructureReadyUnknown(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}
