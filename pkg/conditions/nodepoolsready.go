package conditions

import (
	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

const (
	// NodePoolsReady is a condition type that tells if tenant cluster node
	// pools are in Ready condition, by aggregating Ready conditions for all
	// node pool objects.
	NodePoolsReady capi.ConditionType = "NodePoolsReady"

	// NodePoolsNotFoundReason is a condition reason that is set when
	// NodePoolsReady is set with status False because node pool objects (e.g.
	// MachinePool CRs or MachineDeployment CRs) are not found.
	NodePoolsNotFoundReason = "NodePoolObjectsNotFound"
)

// IsNodePoolsReadyTrue checks if specified cluster is in NodePoolsReady
// condition (if NodePoolsReady condition is set with status True).
func IsNodePoolsReadyTrue(cluster *capi.Cluster) bool {
	return capiconditions.IsTrue(cluster, NodePoolsReady)
}

// IsNodePoolsReadyFalse checks if specified object is not in NodePoolsReady
// condition (if NodePoolsReady condition is set with status False) and if
// optionally specified checks are successful.
func IsNodePoolsReadyFalse(cluster *capi.Cluster, checkOptions ...CheckOption) bool {
	condition := capiconditions.Get(cluster, NodePoolsReady)
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

// IsNodePoolsReadyUnknown checks if it is unknown whether the specified cluster
// is in NodePoolsReady condition or not (if NodePoolsReady condition is not
// set, or it is set with status Unknown).
func IsNodePoolsReadyUnknown(cluster *capi.Cluster) bool {
	return capiconditions.IsUnknown(cluster, NodePoolsReady)
}

// UpdateNodePoolsReady sets NodePoolsReady condition on specified cluster
// by aggregating Ready conditions from specified node pool objects.
//
// If node pool objects are found, cluster NodePoolsReady is set to with status
// False and reason NodePoolsNotFoundReason.
func UpdateNodePoolsReady(cluster *capi.Cluster, nodePools []capiconditions.Getter) {
	if len(nodePools) == 0 {
		capiconditions.MarkFalse(
			cluster,
			NodePoolsReady,
			NodePoolsNotFoundReason,
			capi.ConditionSeverityWarning,
			"Node pools are not found for Cluster %s/%s",
			cluster.Namespace, cluster.Name)
		return
	}

	capiconditions.SetAggregate(
		cluster,
		NodePoolsReady,
		nodePools,
		capiconditions.WithStepCounter(),
		capiconditions.AddSourceRef())
}
