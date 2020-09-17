module github.com/pcingest

go 1.14

require (
	github.com/GoogleContainerTools/kpt v0.33.0
	github.com/stretchr/testify v1.4.0
	sigs.k8s.io/kustomize/kyaml v0.7.2-0.20200914180048-6a0a909e7315
)

replace github.com/GoogleContainerTools/kpt => github.com/barney-s/kpt v0.33.1-0.20200914224538-e972dc07a203
