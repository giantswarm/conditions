package conditions

import (
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

// IsReadyTrue checks if specified object is in Ready condition (if Ready
// condition is set with status True).
func IsReadyTrue(object Object) bool {
	return capiconditions.IsTrue(object, capi.ReadyCondition)
}

// IsReadyFalse checks if specified object is not in Ready condition (if Ready
// condition is set with status False).
func IsReadyFalse(object Object, checkOptions ...CheckOption) bool {
	condition := capiconditions.Get(object, capi.ReadyCondition)
	if !IsFalse(condition) {
		// Condition is not set or it does not have status False
		return false
	}

	for _, checkOption := range checkOptions {
		if !checkOption(condition) {
			// additional check has failed
			return false
		}
	}

	return true
}

// IsReadyUnknown checks if it is unknown whether the specified object is in
// Ready condition or not (if Ready condition is not set, or it is set with
// status Unknown).
func IsReadyUnknown(object Object) bool {
	return capiconditions.IsUnknown(object, capi.ReadyCondition)
}
