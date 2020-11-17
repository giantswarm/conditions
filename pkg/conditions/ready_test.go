package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
)

func Test_IsReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsReadyTrue returns true for CR with condition Ready with status True",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsReadyTrue returns false for CR with condition Ready with status False",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsReadyTrue returns false for CR with condition Ready with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsReadyTrue returns false for CR without condition Ready",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsReadyTrue returns false for CR with condition Ready with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionStatus("SomeUnsupportedValue"),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)
			t.Parallel()

			result := IsReadyTrue(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsReadyFalse returns false for CR with condition Ready with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsReadyFalse returns true for CR with condition Ready with status False",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsReadyFalse returns false for CR with condition Ready with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsReadyFalse returns false for CR without condition Ready",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsReadyFalse returns false for CR with condition Ready with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionStatus("ShinyNewStatusHere"),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)
			t.Parallel()

			result := IsReadyFalse(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsReadyUnknown returns false for CR with condition Ready with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsReadyUnknown returns false for CR with condition Ready with status False",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsReadyUnknown returns true for CR with condition Ready with status Unknown",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsReadyUnknown returns false for CR without condition Ready",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsReadyUnknown returns false for CR with condition Ready with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   capi.ReadyCondition,
							Status: corev1.ConditionStatus("IDontKnowWhatIAmDoing"),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)
			t.Parallel()

			result := IsReadyUnknown(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}
