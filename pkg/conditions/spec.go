package conditions

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capiconditions "sigs.k8s.io/cluster-api/util/conditions"
)

type Object interface {
	capiconditions.Setter
	metav1.Object
}
