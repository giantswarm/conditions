package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

func Test_IsControlPlaneReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         *capi.Cluster
	}{
		{
			name:           "case 0: IsControlPlaneReadyTrue returns true for CR with condition ControlPlaneReady with status True",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsControlPlaneReadyTrue returns false for CR with condition ControlPlaneReady with status False",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsControlPlaneReadyTrue returns false for CR with condition ControlPlaneReady with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsControlPlaneReadyTrue returns false for CR without condition ControlPlaneReady",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsControlPlaneReadyTrue returns false for CR with condition ControlPlaneReady with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
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

			result := IsControlPlaneReadyTrue(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsControlPlaneReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         *capi.Cluster
	}{
		{
			name:           "case 0: IsControlPlaneReadyFalse returns false for CR with condition ControlPlaneReady with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsControlPlaneReadyFalse returns true for CR with condition ControlPlaneReady with status False",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsControlPlaneReadyFalse returns false for CR with condition ControlPlaneReady with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsControlPlaneReadyFalse returns false for CR without condition ControlPlaneReady",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsControlPlaneReadyFalse returns false for CR with condition ControlPlaneReady with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
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

			result := IsControlPlaneReadyFalse(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsControlPlaneReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         *capi.Cluster
	}{
		{
			name:           "case 0: IsControlPlaneReadyUnknown returns false for CR with condition ControlPlaneReady with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsControlPlaneReadyUnknown returns false for CR with condition ControlPlaneReady with status False",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsControlPlaneReadyUnknown returns true for CR with condition ControlPlaneReady with status Unknown",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsControlPlaneReadyUnknown returns true for CR without condition ControlPlaneReady",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsControlPlaneReadyUnknown returns false for CR with condition ControlPlaneReady with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   ControlPlaneReady,
							Status: corev1.ConditionStatus("YouShallNotPass"),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.name)

			result := IsControlPlaneReadyUnknown(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}
