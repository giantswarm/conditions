module github.com/giantswarm/conditions

go 1.15

require (
	github.com/giantswarm/microerror v0.3.0
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	sigs.k8s.io/cluster-api v0.4.0
)

replace sigs.k8s.io/cluster-api v0.3.10 => github.com/giantswarm/cluster-api v0.3.10-gs
