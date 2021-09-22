package conditions

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha4"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1alpha4"
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

func machineWithoutConditions() *capi.Machine {
	return &capi.Machine{
		Status: capi.MachineStatus{
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

func sprintConditionForObject(object Object, conditionType capi.ConditionType) string {
	return sprintCondition(capiconditions.Get(object, conditionType))
}

func sprintCondition(condition *capi.Condition) string {
	var text string
	if condition != nil {
		text = fmt.Sprintf(
			"%s(Status=%q, Reason=%q, Severity=%q, [optional: Message=%q])",
			condition.Type,
			condition.Status,
			condition.Reason,
			condition.Severity,
			condition.Message)
	} else {
		text = "condition is nil"
	}

	return text
}
