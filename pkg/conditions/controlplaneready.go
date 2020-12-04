package conditions

import (
	"time"

	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

const (
	// ControlPlaneReady is a condition type that tells if tenant cluster's
	// control plane object is in Ready condition, by mirroring Ready condition
	// from the control plane CR which is usually referenced by ControlPlaneRef
	// property in the CR Spec.
	//
	// For example, for Cluster CR it currently mirrors AzureMachine Ready
	// condition (when we start reconciling tenant cluster control plane with
	// Cluster API types, control plane type).
	//
	// Examples:
	//
	//    import capiconditions "sigs.k8s.io/cluster-api/util/conditions"
	//
	//    // Set ControlPlaneReady for Cluster CR by mirroring Ready condition
	//    // from AzureMachine CR.
	//    capiconditions.SetMirror(cluster, conditions.ControlPlaneReady, azureMachine)
	//
	//    // Set ControlPlaneReady for MachinePool CR by mirroring
	//    // Ready condition from AzureMachinePool CR.
	//    capiconditions.SetMirror(machinePool, conditions.ControlPlaneReady, azureMachinePool)
	//
	ControlPlaneReady = capi.ControlPlaneReadyCondition

	// Below are condition reasons for ControlPlaneReady that are usually set
	// when condition status is set to False.

	// ControlPlaneReferenceNotSetReason is a condition reason that is set when
	// ControlPlaneReady is set with status False because control plane reference
	// is not set on Cluster object. When using this reason, the condition
	// severity should be set to Warning.
	ControlPlaneReferenceNotSetReason = "ControlPlaneReferenceNotSet"

	// ControlPlaneObjectNotFoundReason is a condition reason that is set when
	// ControlPlaneReady is set with status False because control plane object
	// is not found, but control plane reference is set. When using this reason,
	// the condition severity should be set to Warning (in the future this might
	// change to Error).
	ControlPlaneObjectNotFoundReason = "ControlPlaneObjectNotFound"

	// Waiting time during which ControlPlaneReady is set to False with
	// severity Info. After this threshold time, severity Warning is used.
	WaitingForControlPlaneWarningThresholdTime = 10 * time.Minute
)

// GetControlPlaneReady tries to get ControlPlaneReady condition from the
// specified Cluster CR. If the ControlPlaneReady condition was found, it
// returns a copy of the condition and true, otherwise it returns an empty
// struct and false.
func GetControlPlaneReady(cluster *capi.Cluster) (capi.Condition, bool) {
	controlPlaneReady := capiconditions.Get(cluster, ControlPlaneReady)

	if controlPlaneReady != nil {
		return *controlPlaneReady, true
	} else {
		return capi.Condition{}, false
	}
}

// IsControlPlaneReadyTrue checks if specified cluster is in ControlPlaneReady
// condition (if ControlPlaneReady condition is set with status True).
func IsControlPlaneReadyTrue(cluster *capi.Cluster) bool {
	return capiconditions.IsTrue(cluster, ControlPlaneReady)
}

// IsControlPlaneReadyFalse checks if specified cluster is not in
// ControlPlaneReady condition (if ControlPlaneReady condition is set with
// status False) and if optionally specified checks are successful.
func IsControlPlaneReadyFalse(cluster *capi.Cluster, checkOptions ...CheckOption) bool {
	condition := capiconditions.Get(cluster, ControlPlaneReady)
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

// IsControlPlaneReadyUnknown checks if it is unknown whether the specified
// cluster is in ControlPlaneReady condition or not (if ControlPlaneReady
// condition is not set, or it is set with status Unknown).
func IsControlPlaneReadyUnknown(cluster *capi.Cluster) bool {
	return capiconditions.IsUnknown(cluster, ControlPlaneReady)
}
