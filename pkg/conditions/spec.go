package conditions

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

type Object interface {
	capiconditions.Setter
	metav1.Object
}

type CheckOption func(condition *capi.Condition) bool
