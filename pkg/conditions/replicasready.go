package conditions

import (
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
	capiexp "sigs.k8s.io/cluster-api/exp/api/v1beta1"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

// GetReplicasReady tries to get ReplicasReady condition from the specified
// MachinePool CR. If the ReplicasReady condition was found, it returns a copy
// of the condition and true, otherwise it returns an empty struct and false.
func GetReplicasReady(machinePool *capiexp.MachinePool) (capi.Condition, bool) {
	replicasReady := capiconditions.Get(machinePool, capiexp.ReplicasReadyCondition)

	if replicasReady != nil {
		return *replicasReady, true
	} else {
		return capi.Condition{}, false
	}
}

// IsReplicasReadyTrue checks if specified MachinePool is in ReplicasReady
// condition (if ReplicasReady condition is set with status True).
func IsReplicasReadyTrue(machinePool *capiexp.MachinePool) bool {
	return capiconditions.IsTrue(machinePool, capiexp.ReplicasReadyCondition)
}

// IsReplicasReadyFalse checks if specified MachinePool is not in ReplicasReady
// condition (if ReplicasReady condition is set with status False) and if
// optionally specified checks are successful.
func IsReplicasReadyFalse(machinePool *capiexp.MachinePool, checkOptions ...CheckOption) bool {
	condition := capiconditions.Get(machinePool, capiexp.ReplicasReadyCondition)
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

// IsReplicasReadyUnknown checks if it is unknown whether the specified
// MachinePool is in ReplicasReady condition or not (if ReplicasReady condition
// is not set, or it is set with status Unknown).
func IsReplicasReadyUnknown(machinePool *capiexp.MachinePool) bool {
	return capiconditions.IsUnknown(machinePool, capiexp.ReplicasReadyCondition)
}

// WithWaitingForReplicasReadyReason returns a CheckOption that checks if
// condition reason is set to WaitingForReplicasReady.
func WithWaitingForReplicasReadyReason() CheckOption {
	return WithReason(capiexp.WaitingForReplicasReadyReason)
}
