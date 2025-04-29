module github.com/asecurityteam/component-producer/v2

go 1.22

toolchain go1.23.6

require (
	github.com/asecurityteam/component-httpclient v0.6.1
	github.com/asecurityteam/settings v1.0.0
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/asecurityteam/logevent/v2 v2.0.2 // indirect
	github.com/asecurityteam/settings/v2 v2.0.2 // indirect
	github.com/asecurityteam/transport v1.7.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/uber/jaeger-lib v1.5.0+incompatible => github.com/uber/jaeger-lib v1.5.0

// Force version due to a vulnerbility in the versions benthos currently uses
replace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.7.4
