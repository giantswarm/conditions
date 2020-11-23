package conditions

import (
	"fmt"
	"time"

	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

const (
	// InfrastructureReady is a condition type that tells if provider
	// infrastructure for the object is in Ready condition, by mirroring Ready
	// condition from the corresponding provider-specific CR which is usually
	// referenced by InfrastructureRef property in the CR Spec.
	//
	// For example, for Cluster CR it mirrors AzureCluster Ready condition and
	// for MachinePool CR it mirrors AzureMachinePool Ready condition.
	//
	// Examples:
	//
	//    import capiconditions "sigs.k8s.io/cluster-api/util/conditions"
	//
	//    // Set InfrastructureReady for Cluster CR by mirroring Ready
	//    // condition from AzureCluster CR.
	//    capiconditions.SetMirror(cluster, conditions.InfrastructureReady, azureCluster)
	//
	//    // Set InfrastructureReady for MachinePool CR by mirroring
	//    // Ready condition from AzureMachinePool CR.
	//    capiconditions.SetMirror(machinePool, conditions.InfrastructureReady, azureMachinePool)
	//
	InfrastructureReady capi.ConditionType = capi.InfrastructureReadyCondition

	// Below are condition reasons for InfrastructureReady that are
	// usually set when condition status is set to False.

	// InfrastructureObjectNotFoundReason is a condition reason that is set
	// when InfrastructureReady is set with status False because corresponding
	// provider-specific infrastructure object is not found. When using this
	// reason, the condition severity should be set to Warning.
	InfrastructureObjectNotFoundReason = "InfrastructureObjectNotFound"

	// Waiting time during which InfrastructureReady is set to False with
	// severity Info. After this threshold time, severity Warning is used.
	WaitingForInfrastructureWarningThresholdTime = 10 * time.Minute
)

// GetInfrastructureReady tries to get InfrastructureReady condition from the
// specified Cluster CR. If the InfrastructureReady condition was found, it
// returns a copy of the condition and true, otherwise it returns an empty
// struct and false.
func GetInfrastructureReady(cluster *capi.Cluster) (capi.Condition, bool) {
	infrastructureReady := capiconditions.Get(cluster, InfrastructureReady)

	if infrastructureReady != nil {
		return *infrastructureReady, true
	} else {
		return capi.Condition{}, false
	}
}

// IsInfrastructureReadyTrue checks if specified object is in InfrastructureReady
// condition (if InfrastructureReady condition is set with status True).
func IsInfrastructureReadyTrue(object Object) bool {
	return capiconditions.IsTrue(object, InfrastructureReady)
}

// IsInfrastructureReadyFalse checks if specified object is not in
// InfrastructureReady condition (if InfrastructureReady condition is set with
// status False).
func IsInfrastructureReadyFalse(object Object) bool {
	return capiconditions.IsFalse(object, InfrastructureReady)
}

// IsInfrastructureReadyUnknown checks if it is unknown whether the specified
// object is in InfrastructureReady condition or not (if InfrastructureReady
// condition is not set, or it is set with status Unknown).
func IsInfrastructureReadyUnknown(object Object) bool {
	return capiconditions.IsUnknown(object, InfrastructureReady)
}

// ReconcileInfrastructureReady sets InfrastructureReady condition on specified
// object by mirroring Ready condition from specified infrastructure object.
//
// If specified infrastructure object is nil, object InfrastructureReady will
// be set with condition False and Reason InfrastructureObjectNotFoundReason.
//
// If specified infrastructure object's Ready condition is not set, object
// InfrastructureReady will be set with condition False and reason
// WaitingForInfrastructure.
func ReconcileInfrastructureReady(object Object, infrastructureObject Object) {
	if infrastructureObject == nil {
		warningMessage :=
			"Corresponding provider-specific infrastructure object of " +
				"type %T is not found for specified %T object %s/%s"

		capiconditions.MarkFalse(
			object,
			InfrastructureReady,
			InfrastructureObjectNotFoundReason,
			capi.ConditionSeverityWarning,
			warningMessage,
			infrastructureObject, object, object.GetNamespace(), object.GetName())
	}

	objectAge := time.Since(object.GetCreationTimestamp().Time)
	var fallbackSeverity capi.ConditionSeverity
	fallbackWarningMessage := ""
	if objectAge > WaitingForInfrastructureWarningThresholdTime {
		// Provider-specific infrastructure object should be reconciled soon
		// after it has been created. If it's Ready condition is not set within
		// 10 minutes, that means that something might be wrong, so we set
		// object's InfrastructureReady condition with severity Warning.
		fallbackSeverity = capi.ConditionSeverityWarning
		fallbackWarningMessage = fmt.Sprintf(" for more than %s", objectAge)
	} else {
		// Otherwise, if it has been less than 10 minutes since object's
		// creation, probably everything is good, and we just have to wait few
		// more minutes.
		// Note: Upstream Cluster API implementation always just sets severity
		// Info when infrastructure object's Ready condition is not set.
		fallbackSeverity = capi.ConditionSeverityInfo
	}

	fallbackToFalse := capiconditions.WithFallbackValue(
		false,
		capi.WaitingForInfrastructureFallbackReason,
		fallbackSeverity,
		fmt.Sprintf("Waiting for infrastructure object of type %T to have Ready condition set%s",
			infrastructureObject, fallbackWarningMessage))

	capiconditions.SetMirror(object, InfrastructureReady, infrastructureObject, fallbackToFalse)
}
