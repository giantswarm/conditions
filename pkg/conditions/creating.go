package conditions

import (
	"time"

	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

const (
	// Creating is a condition type that tells if a cluster, a node pool, a
	// tenant cluster control plane, or something else that needs a Creating
	// condition is currently being created.
	Creating capi.ConditionType = "Creating"

	// Below are condition reasons for Creating condition that are usually set
	// when condition status is set to False.

	// CreationCompletedReason is set when the creation has been completed
	// successfully.
	CreationCompletedReason = "CreationCompleted"

	// ExistingClusterReason is used when setting a Creating condition for the
	// first time on an object that was created before conditions support was
	// implemented.
	ExistingObjectReason = "ExistingObject"
)

// GetCreating tries to get Creating condition from the specified object. If
// the Creating condition was found, it returns a copy of the condition and
// true, otherwise it returns an empty struct and false.
func GetCreating(object Object) (capi.Condition, bool) {
	c := capiconditions.Get(object, Creating)

	if c != nil {
		return *c, true
	} else {
		return capi.Condition{}, false
	}
}

// IsCreatingTrue checks if specified object is in Creating condition (if
// Creating condition is set with status True).
func IsCreatingTrue(object Object) bool {
	return capiconditions.IsTrue(object, Creating)
}

// IsCreatingFalse checks if specified object is not in Creating condition (if
// Creating condition is set with status False) and if optionally specified
// checks are successful.
//
// Examples:
//
//    IsCreatingFalse(cluster)
//    IsCreatingFalse(cluster, WithCreationCompletedReason())
//    IsCreatingFalse(cluster, WithExistingObjectReason())
//
func IsCreatingFalse(object Object, checkOptions ...CheckOption) bool {
	creating := capiconditions.Get(object, Creating)
	if !IsFalse(creating) {
		// Creating condition is not set or it does not have status False
		return false
	}

	for _, checkOption := range checkOptions {
		if !checkOption(creating) {
			// additional check has failed
			return false
		}
	}

	return true
}

// IsCreatingUnknown checks if it is unknown whether the specified object is in
// Creating condition or not (if Creating condition is not set, or it is set
// with status Unknown).
func IsCreatingUnknown(object Object) bool {
	return capiconditions.IsUnknown(object, Creating)
}

// WithCreationCompletedReason returns a CheckOption that checks if condition
// reason is set to CreationCompleted.
func WithCreationCompletedReason() CheckOption {
	return WithReason(CreationCompletedReason)
}

// WithExistingObjectReason returns a CheckOption that checks if condition
// reason is set to ExistingObject.
func WithExistingObjectReason() CheckOption {
	return WithReason(ExistingObjectReason)
}

// MarkCreatingTrue sets Creating condition with status True.
func MarkCreatingTrue(object Object) {
	capiconditions.MarkTrue(object, Creating)
}

// MarkCreatingFalseWithCreationCompleted sets Creating condition with status
// False, reason CreationCompleted, severity Info and a message informing how
// long the creation took.
func MarkCreatingFalseWithCreationCompleted(object Object) {
	creationDuration := time.Since(object.GetCreationTimestamp().Time)
	capiconditions.MarkFalse(
		object,
		Creating,
		CreationCompletedReason,
		capi.ConditionSeverityInfo,
		"Cluster creation has been completed in %s",
		creationDuration)
}

// MarkCreatingFalseForExistingObject sets Creating condition with status
// False, reason ExistingObject, severity Info and a message informing that the
// object was already created.
func MarkCreatingFalseForExistingObject(object Object) {
	capiconditions.MarkFalse(
		object,
		Creating,
		ExistingObjectReason,
		capi.ConditionSeverityInfo,
		"Object was already created")
}
