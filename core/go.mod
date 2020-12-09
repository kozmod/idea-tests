module github.com/kozmod/idea-tests/core

go 1.13

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kozmod/idea-tests/utils v0.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/magiconair/properties v1.8.1
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/kozmod/idea-tests/utils => ../utils
