package conditions

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

func Test_IsNodePoolsReadyTrue(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         *capi.Cluster
	}{
		{
			name:           "case 0: IsNodePoolsReadyTrue returns true for Cluster with condition NodePoolsReady with status True",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsNodePoolsReadyTrue returns false for Cluster with condition NodePoolsReady with status False",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsNodePoolsReadyTrue returns false for Cluster with condition NodePoolsReady with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsNodePoolsReadyTrue returns false for Cluster without condition NodePoolsReady",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsNodePoolsReadyTrue returns false for Cluster with condition NodePoolsReady with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
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

			result := IsNodePoolsReadyTrue(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsNodePoolsReadyFalse(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         *capi.Cluster
	}{
		{
			name:           "case 0: IsNodePoolsReadyFalse returns false for Cluster with condition NodePoolsReady with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsNodePoolsReadyFalse returns true for Cluster with condition NodePoolsReady with status False",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsNodePoolsReadyFalse returns false for Cluster with condition NodePoolsReady with status Unknown",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsNodePoolsReadyFalse returns false for Cluster without condition NodePoolsReady",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsNodePoolsReadyFalse returns false for Cluster with condition NodePoolsReady with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
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

			result := IsNodePoolsReadyFalse(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}

func Test_IsNodePoolsReadyUnknown(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult bool
		object         *capi.Cluster
	}{
		{
			name:           "case 0: IsNodePoolsReadyUnknown returns false for Cluster with condition NodePoolsReady with status True",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			name:           "case 1: IsNodePoolsReadyUnknown returns false for Cluster with condition NodePoolsReady with status False",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			name:           "case 2: IsNodePoolsReadyUnknown returns true for Cluster with condition NodePoolsReady with status Unknown",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			name:           "case 3: IsNodePoolsReadyUnknown returns true for Cluster without condition NodePoolsReady",
			expectedResult: true,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{},
				},
			},
		},
		{
			name:           "case 4: IsNodePoolsReadyUnknown returns false for Cluster with condition NodePoolsReady with unsupported status",
			expectedResult: false,
			object: &capi.Cluster{
				Status: capi.ClusterStatus{
					Conditions: capi.Conditions{
						{
							Type:   NodePoolsReady,
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

			result := IsNodePoolsReadyUnknown(tc.object)
			if result != tc.expectedResult {
				t.Logf("expected %t, got %t", tc.expectedResult, result)
				t.Fail()
			}
		})
	}
}
