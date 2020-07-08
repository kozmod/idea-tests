module github.com/kozmod/idea-tests/core

go 1.13

require (
	github.com/kozmod/idea-tests/utils v0.0.0
	github.com/magiconair/properties v1.8.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.0
	golang.org/x/tools v0.0.0-20200606014950-c42cb6316fb6 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace github.com/kozmod/idea-tests/utils => ../utils
