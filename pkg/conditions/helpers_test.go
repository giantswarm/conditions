package conditions

import (
	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
)

func clusterWith(conditionType capi.ConditionType, conditionStatus corev1.ConditionStatus) *capi.Cluster {
	return &capi.Cluster{
		Status: capi.ClusterStatus{
			Conditions: capi.Conditions{
				{
					Type:   conditionType,
					Status: conditionStatus,
				},
			},
		},
	}
}

func clusterWithoutConditions() *capi.Cluster {
	return &capi.Cluster{
		Status: capi.ClusterStatus{
			Conditions: capi.Conditions{},
		},
	}
}

func machinePoolWith(conditionType capi.ConditionType, conditionStatus corev1.ConditionStatus) *capiexp.MachinePool {
	return &capiexp.MachinePool{
		Status: capiexp.MachinePoolStatus{
			Conditions: capi.Conditions{
				{
					Type:   conditionType,
					Status: conditionStatus,
				},
			},
		},
	}
}

func machinePoolWithoutConditions() *capiexp.MachinePool {
	return &capiexp.MachinePool{
		Status: capiexp.MachinePoolStatus{
			Conditions: capi.Conditions{},
		},
	}
}
