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
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsCreatingFalse(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsCreatingFalse returns false for CR with condition Creating with status True",
			expectedResult: false,
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
			name:           "case 1: IsCreatingFalse returns true for CR with condition Creating with status False",
			expectedResult: true,
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
			name:           "case 2: IsCreatingFalse returns false for CR with condition Creating with status Unknown",
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
			name:           "case 3: IsCreatingFalse returns false for CR without condition Creating",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsCreatingFalse returns false for CR with condition Creating with unsupported status",
			expectedResult: false,
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsCreatingFalse(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
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
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}
