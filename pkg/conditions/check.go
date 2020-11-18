package conditions

import (
	corev1 "k8s.io/api/core/v1"
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
)

func IsTrue(condition *capi.Condition) bool {
	return condition != nil && condition.Status == corev1.ConditionTrue
}

func IsFalse(condition *capi.Condition) bool {
	return condition != nil && condition.Status == corev1.ConditionFalse
}

func IsUnknown(condition *capi.Condition) bool {
	return condition == nil || condition.Status == corev1.ConditionUnknown
}

func WithReason(reason string) CheckOption {
	return func(condition *capi.Condition) bool {
		return condition != nil && condition.Reason == reason
	}
}

func WithCreationCompletedReason() CheckOption {
	return WithReason(CreationCompletedReason)
}

func WithExistingObjectReason() CheckOption {
	return WithReason(ExistingObjectReason)
}

func WithSeverity(severity capi.ConditionSeverity) CheckOption {
	return func(condition *capi.Condition) bool {
		return condition != nil && condition.Severity == severity
	}
}

func WithSeverityInfo() CheckOption {
	return WithSeverity(capi.ConditionSeverityInfo)
}

func WithSeverityWarning() CheckOption {
	return WithSeverity(capi.ConditionSeverityWarning)
}

func WithSeverityError() CheckOption {
	return WithSeverity(capi.ConditionSeverityError)
}

func WithoutSeverity() CheckOption {
	return WithSeverity(capi.ConditionSeverityNone)
}
