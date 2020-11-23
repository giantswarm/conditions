package conditions

import (
	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

// IsTrue checks if specified condition is not nil and has Status set to True.
func IsTrue(condition *capi.Condition) bool {
	return condition != nil && condition.Status == corev1.ConditionTrue
}

// IsFalse checks if specified condition is not nil and has Status set to False.
func IsFalse(condition *capi.Condition) bool {
	return condition != nil && condition.Status == corev1.ConditionFalse
}

// IsUnknown checks if specified condition is either nil or has Status set to
// Unknown.
func IsUnknown(condition *capi.Condition) bool {
	return condition == nil || condition.Status == corev1.ConditionUnknown
}

// WithReason returns a CheckOption that checks if condition reason is set to
// the specified value.
func WithReason(reason string) CheckOption {
	return func(condition *capi.Condition) bool {
		return condition != nil && condition.Reason == reason
	}
}

// WithSeverity returns a CheckOption that checks if condition severity is set
// to the specified value.
func WithSeverity(severity capi.ConditionSeverity) CheckOption {
	return func(condition *capi.Condition) bool {
		return condition != nil && condition.Severity == severity
	}
}

// WithSeverityInfo returns a CheckOption that checks if condition severity is
// set to Info.
func WithSeverityInfo() CheckOption {
	return WithSeverity(capi.ConditionSeverityInfo)
}

// WithSeverityWarning returns a CheckOption that checks if condition severity
// is set to Warning.
func WithSeverityWarning() CheckOption {
	return WithSeverity(capi.ConditionSeverityWarning)
}

// WithSeverityError returns a CheckOption that checks if condition severity is
// set to Error.
func WithSeverityError() CheckOption {
	return WithSeverity(capi.ConditionSeverityError)
}

// WithoutSeverity returns a CheckOption that checks if condition severity is
// not set.
func WithoutSeverity() CheckOption {
	return WithSeverity(capi.ConditionSeverityNone)
}

// AreEqual compares specified conditions.
//
// If both are nil, it returns true. If only one of them is nil, it returns
// false. If both are different than nil, it compares Type, Status, Severity
// and Reason fields, while it ignores LastTransitionTime and Message.
func AreEqual(c1, c2 *capi.Condition) bool {
	// Both are nil
	if c1 == nil && c2 == nil {
		return true
	}

	// Only one is nil
	if c1 == nil || c2 == nil {
		return false
	}

	// None is nil, compare fields
	equal := c1.Type == c2.Type &&
		c1.Status == c2.Status &&
		c1.Severity == c2.Severity &&
		c1.Reason == c2.Reason

	return equal
}
