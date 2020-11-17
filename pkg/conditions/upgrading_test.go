package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
)

func Test_IsUpgradingTrue(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsUpgradingTrue returns true for CR with condition Upgrading with status True",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsUpgradingTrue returns false for CR with condition Upgrading with status False",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsUpgradingTrue returns false for CR with condition Upgrading with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsUpgradingTrue returns false for CR without condition Upgrading",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsUpgradingTrue returns false for CR with condition Upgrading with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
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

			result := IsUpgradingTrue(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsUpgradingFalse(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsUpgradingFalse returns false for CR with condition Upgrading with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsUpgradingFalse returns true for CR with condition Upgrading with status False",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsUpgradingFalse returns false for CR with condition Upgrading with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsUpgradingFalse returns false for CR without condition Upgrading",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsUpgradingFalse returns false for CR with condition Upgrading with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
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

			result := IsUpgradingFalse(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsUpgradingUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         Object
	}{
		{
			name:           "case 0: IsUpgradingUnknown returns false for CR with condition Upgrading with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsUpgradingUnknown returns false for CR with condition Upgrading with status False",
			expectedResult: false,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsUpgradingUnknown returns true for CR with condition Upgrading with status Unknown",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsUpgradingUnknown returns true for CR without condition Upgrading",
			expectedResult: true,
			object: &capiexp.MachinePool{
				Status: capiexp.MachinePoolStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsUpgradingUnknown returns false for CR with condition Upgrading with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   Upgrading,
							Status: corev1.ConditionStatus("IDonTKnowWhatIAmDoing"),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsUpgradingUnknown(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}
