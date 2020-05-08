module github.com/kozmod/idea-tests/test-operator

go 1.13

require (
	github.com/operator-framework/operator-sdk v0.17.0
	github.com/spf13/cobra v1.0.0 // indirect
	k8s.io/api v0.17.4
	k8s.io/apimachinery v0.17.4
	sigs.k8s.io/controller-runtime v0.5.2
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	k8s.io/client-go => k8s.io/client-go v0.17.4 // Required by prometheus-operator
)
