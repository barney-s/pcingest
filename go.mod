module github.com/pcingest

go 1.14

require (
	github.com/GoogleContainerTools/kpt v0.36.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.9.0
	sigs.k8s.io/kustomize/kyaml v0.9.3
)

//replace github.com/GoogleContainerTools/kpt => github.com/barney-s/kpt v0.33.1-0.20200914224538-e972dc07a203
