package conditions

import (
	"fmt"

	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

const (
	conditionNotSet = "condition not set"
)

var UnexpectedConditionStatusError = &microerror.Error{
	Kind: "UnexpectedConditionStatus",
}

func UnexpectedConditionStatusErrorMessage(cr Object, t capi.ConditionType) string {
	c := capiconditions.Get(cr, t)
	var got string
	if c != nil {
		got = string(c.Status)
	} else {
		got = conditionNotSet
	}

	return fmt.Sprintf("Unexpected status for condition %s, got %s", t, got)
}

// IsUnexpectedConditionStatus asserts UnexpectedConditionStatusError.
func IsUnexpectedConditionStatus(err error) bool {
	return microerror.Cause(err) == UnexpectedConditionStatusError
}

func ExpectedTrueErrorMessage(cr Object, conditionType capi.ConditionType) string {
	return ExpectedStatusErrorMessage(cr, conditionType, corev1.ConditionTrue)
}

func ExpectedFalseErrorMessage(cr Object, conditionType capi.ConditionType) string {
	return ExpectedStatusErrorMessage(cr, conditionType, corev1.ConditionFalse)
}

func ExpectedStatusErrorMessage(cr Object, conditionType capi.ConditionType, expectedStatus corev1.ConditionStatus) string {
	c := capiconditions.Get(cr, conditionType)
	var got string
	if c != nil {
		got = string(c.Status)
	} else {
		got = conditionNotSet
	}

	return fmt.Sprintf("Expected that condition %s on Object %T has status %s, but got %s", conditionType, cr, expectedStatus, got)
}

var UnsupportedConditionStatusError = &microerror.Error{
	Kind: "UnsupportedConditionStatus",
}

func UnsupportedConditionStatusErrorMessage(cr Object, t capi.ConditionType) string {
	c := capiconditions.Get(cr, t)
	return fmt.Sprintf("Unsupported status for condition %s, got %s", t, c.Status)
}

// IsInvalidCondition asserts invalidConditionError.
func IsUnsupportedConditionStatus(err error) bool {
	return microerror.Cause(err) == UnsupportedConditionStatusError
}
