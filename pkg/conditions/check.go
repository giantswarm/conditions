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
