package conditions

import (
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

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
