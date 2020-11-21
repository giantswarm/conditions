package conditions

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
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

func machineWith(conditionType capi.ConditionType, conditionStatus corev1.ConditionStatus) *capi.Machine {
	return &capi.Machine{
		Status: capi.MachineStatus{
			Conditions: capi.Conditions{
				{
					Type:   conditionType,
					Status: conditionStatus,
				},
			},
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

func getGotConditionStatusString(object Object, conditionType capi.ConditionType) string {
	condition := capiconditions.Get(object, conditionType)
	var got string
	if condition != nil {
		got = string(condition.Status)
	} else {
		got = "condition not set"
	}

	return got
}

func conditionString(object Object, conditionType capi.ConditionType) string {
	condition := capiconditions.Get(object, conditionType)
	var text string
	if condition != nil {
		text = fmt.Sprintf(
			"%s: Status=%q, Reason=%q, Severity=%q, Message=%q",
			condition.Type,
			condition.Status,
			condition.Reason,
			condition.Severity,
			condition.Message)
	} else {
		text = "condition not set"
	}

	return text
}
