module github.com/giantswarm/conditions

go 1.15

require (
	github.com/giantswarm/microerror v0.3.0
	k8s.io/api v0.18.9
	k8s.io/apimachinery v0.18.19
	sigs.k8s.io/cluster-api v0.3.10
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.24+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.0 => github.com/gorilla/websocket v1.4.2
	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
)
