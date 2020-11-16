package conditions

import (
	"fmt"
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
	//    // Set ControlPlaneReady for Cluster CR by mirroring Ready conditio
	//    // from AzureMachine CR.
	//    capiconditions.SetMirror(cluster, conditions.ControlPlaneReady, azureMachine)
	//
	//    // Set ControlPlaneReady for MachinePool CR by mirroring
	//    // Ready condition from AzureMachinePool CR.
	//    capiconditions.SetMirror(machinePool, conditions.ControlPlaneReady, azureMachinePool)
	//
	ControlPlaneReady capi.ConditionType = capi.ControlPlaneReadyCondition

	// Below are condition reasons for ControlPlaneReady that are usually set
	// when condition status is set to False.

	// ControlPlaneObjectNotFoundReason is a condition reason that is set when
	// ControlPlaneReady is set with status False because control plane object
	// is not found. When using this reason, the condition severity should be
	// set to Warning.
	ControlPlaneObjectNotFoundReason = "ControlPlaneObjectNotFound"

	// Waiting time during which ControlPlaneReady is set to False with
	// severity Info. After this threshold time, severity Warning is used.
	WaitingForControlPlaneWarningThresholdTime = 10 * time.Minute
)

// IsControlPlaneReadyTrue checks if specified cluster is in ControlPlaneReady
// condition (if ControlPlaneReady condition is set with status True).
func IsControlPlaneReadyTrue(cluster *capi.Cluster) bool {
	return capiconditions.IsTrue(cluster, ControlPlaneReady)
}

// IsControlPlaneReadyFalse checks if specified cluster is not in
// ControlPlaneReady condition (if ControlPlaneReady condition is set with
// status False).
func IsControlPlaneReadyFalse(cluster *capi.Cluster) bool {
	return capiconditions.IsFalse(cluster, ControlPlaneReady)
}

// IsControlPlaneReadyUnknown checks if it is unknown whether the specified
// cluster is in ControlPlaneReady condition or not (if ControlPlaneReady
// condition is not set, or it is set with status Unknown).
func IsControlPlaneReadyUnknown(cluster *capi.Cluster) bool {
	return capiconditions.IsUnknown(cluster, ControlPlaneReady)
}

// ReconcileControlPlaneReady sets ControlPlaneReady condition on specified
// cluster by mirroring Ready condition from specified control plane object.
//
// If specified control plane object is nil, object ControlPlaneReady will
// be set with condition False and Reason ControlPlaneObjectNotFoundReason.
//
// If specified control plane object's Ready condition is not set, object
// ControlPlaneReady will be set with condition False and reason
// WaitingForControlPlane.
func ReconcileControlPlaneReady(cluster *capi.Cluster, controlPlaneObject *capi.Cluster) {
	if controlPlaneObject == nil {
		warningMessage :=
			"Control plane object of type %T is not found for specified %T object %s/%s"

		capiconditions.MarkFalse(
			cluster,
			ControlPlaneReady,
			ControlPlaneObjectNotFoundReason,
			capi.ConditionSeverityWarning,
			warningMessage,
			controlPlaneObject, cluster, cluster.GetNamespace(), cluster.GetName())
	}

	clusterAge := time.Since(cluster.GetCreationTimestamp().Time)
	var fallbackSeverity capi.ConditionSeverity
	fallbackWarningMessage := ""
	if clusterAge > WaitingForControlPlaneWarningThresholdTime {
		// Control plane should be reconciled soon after it has been created.
		// If it's Ready condition is not set within 10 minutes, that means
		// that something might be wrong, so we set object's ControlPlaneReady
		// condition with severity Warning.
		fallbackSeverity = capi.ConditionSeverityWarning
		fallbackWarningMessage = fmt.Sprintf(" for more than %s", clusterAge)
	} else {
		// Otherwise, if it has been less than 10 minutes since object's
		// creation, probably everything is good, and we just have to wait few
		// more minutes.
		// Note: Upstream Cluster API implementation always just sets severity
		// Info when control plane object's Ready condition is not set.
		fallbackSeverity = capi.ConditionSeverityInfo
	}

	fallbackToFalse := capiconditions.WithFallbackValue(
		false,
		capi.WaitingForControlPlaneFallbackReason,
		fallbackSeverity,
		fmt.Sprintf("Waiting for control plane object of type %T to have Ready condition set%s",
			controlPlaneObject, fallbackWarningMessage))

	capiconditions.SetMirror(cluster, ControlPlaneReady, controlPlaneObject, fallbackToFalse)
}
