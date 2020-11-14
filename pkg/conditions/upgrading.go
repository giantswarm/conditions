package conditions

import (
	"fmt"
	"time"

	capi "sigs.k8s.io/cluster-api/api/v1alpha3"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

const (
	// Upgrading is a condition type that tells if a cluster, a node pool, a
	// tenant cluster control plane, or something else that needs an Upgrading
	// condition is currently being upgraded.
	Upgrading capi.ConditionType = "Upgrading"

	// Below are condition reasons for Upgrading condition that are usually set
	// when condition status is set to False.

	// UpgradeCompletedReason is set when the upgrade has been completed
	// successfully.
	UpgradeCompletedReason = "UpgradeCompleted"

	// UpgradeNotStartedReason is set when the upgrade has not started yet.
	// This is usually during or after creation, but can also be after
	// restoring a CR from the backup.
	UpgradeNotStartedReason = "UpgradeNotStarted"
)

// IsUpgradingTrue checks if specified object is in Upgrading condition (if
// Upgrading condition is set with status True).
func IsUpgradingTrue(object Object) bool {
	return capiconditions.IsTrue(object, Upgrading)
}

// IsUpgradingFalse checks if specified object is not in Upgrading condition (if
// Upgrading condition is set with status False).
func IsUpgradingFalse(object Object) bool {
	return capiconditions.IsFalse(object, Upgrading)
}

// IsUpgradingUnknown checks if it is unknown whether the specified object is in
// Upgrading condition or not (if Upgrading condition is not set, or it is set
// with status Unknown).
func IsUpgradingUnknown(object Object) bool {
	return capiconditions.IsUnknown(object, Upgrading)
}

// MarkUpgradingTrue sets Upgrading condition with status True.
func MarkUpgradingTrue(object Object) {
	capiconditions.MarkTrue(object, Upgrading)
}

// MarkUpgradingFalseWithUpgradeCompleted sets Upgrading condition with status
// False, reason UpgradeCompleted, severity Info and a message informing how
// long the upgrade took.
func MarkUpgradingFalseWithUpgradeCompleted(object Object) {
	var upgradeTimeMessage string
	currentUpgradingCondition := capiconditions.Get(object, Upgrading)
	if currentUpgradingCondition != nil {
		upgradeDuration := time.Since(currentUpgradingCondition.LastTransitionTime.Time)
		upgradeTimeMessage = fmt.Sprintf("in %s", upgradeDuration)
	} else {
		upgradeTimeMessage = ", but upgrade duration cannot be determined"
	}

	capiconditions.MarkFalse(
		object,
		Upgrading,
		UpgradeCompletedReason,
		capi.ConditionSeverityInfo,
		"Upgrade has been completed%s",
		upgradeTimeMessage)
}

// MarkUpgradingFalseWithUpgradeNotStarted sets Upgrading condition with status
// False, reason ExistingObject, severity Info and a message informing that the
// upgrade has not been started.
func MarkUpgradingFalseWithUpgradeNotStarted(object Object) {
	capiconditions.MarkFalse(
		object,
		Upgrading,
		UpgradeNotStartedReason,
		capi.ConditionSeverityInfo,
		"Upgrade has not been started")
}
