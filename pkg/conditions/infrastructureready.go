package conditions

import (
	"time"

	capi "sigs.k8s.io/cluster-api/api/v1beta1"
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
	InfrastructureReady = capi.InfrastructureReadyCondition

	// Below are condition reasons for InfrastructureReady that are
	// usually set when condition status is set to False.

	// InfrastructureReferenceNotSetReason is a condition reason that is set
	// when InfrastructureReady is set with status False because object does
	// not have infrastructure reference set. When using this reason, the
	// condition severity should be set to Warning.
	InfrastructureReferenceNotSetReason = "InfrastructureReferenceNotSet"

	// InfrastructureObjectNotFoundReason is a condition reason that is set
	// when InfrastructureReady is set with status False because corresponding
	// provider-specific infrastructure object is not found, but infrastructure
	// reference is set. When using this reason, the condition severity should
	// be set to Warning.
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
// status False) and if optionally specified checks are successful.
func IsInfrastructureReadyFalse(object Object, checkOptions ...CheckOption) bool {
	condition := capiconditions.Get(object, InfrastructureReady)
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

// IsInfrastructureReadyUnknown checks if it is unknown whether the specified
// object is in InfrastructureReady condition or not (if InfrastructureReady
// condition is not set, or it is set with status Unknown).
func IsInfrastructureReadyUnknown(object Object) bool {
	return capiconditions.IsUnknown(object, InfrastructureReady)
}
