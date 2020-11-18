package conditions

import (
	"fmt"

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
