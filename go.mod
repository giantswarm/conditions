module github.com/giantswarm/conditions

go 1.15

require (
	github.com/giantswarm/microerror v0.3.0
	go.uber.org/zap v1.13.0 // indirect
	k8s.io/api v0.18.9
	k8s.io/apimachinery v0.18.19
	sigs.k8s.io/cluster-api v0.3.10
)

replace sigs.k8s.io/cluster-api v0.3.10 => github.com/giantswarm/cluster-api v0.3.10-gs
