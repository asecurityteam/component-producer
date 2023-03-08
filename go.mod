module github.com/asecurityteam/component-producer

go 1.15

require (
	bitbucket.org/atlassian/go-asap v0.0.0-20201116174856-38f0143fcabd // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Jeffail/benthos/v3 v3.65.0
	github.com/armon/go-metrics v0.3.6 // indirect
	github.com/asecurityteam/component-httpclient v0.2.0
	github.com/asecurityteam/component-stat v0.3.0 // indirect
	github.com/asecurityteam/httpstats v0.0.0-20201215174437-106328c66daa // indirect
	github.com/asecurityteam/logevent v1.4.0 // indirect
	github.com/asecurityteam/runhttp v0.4.0 // indirect
	github.com/asecurityteam/settings v0.4.0
	github.com/asecurityteam/transport v1.5.1 // indirect
	github.com/asecurityteam/transportd v1.2.4 // indirect
	github.com/bitly/go-hostpool v0.1.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/getkin/kin-openapi v0.36.0 // indirect
	github.com/go-chi/chi v4.1.2+incompatible // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-openapi/swag v0.19.13 // indirect
	github.com/golang/mock v1.6.0
	github.com/hashicorp/raft-boltdb v0.0.0-20171010151810-6e5ba93211ea // indirect
	github.com/itchyny/astgen-go v0.0.0-20210113000433-0da0671862a3 // indirect
	github.com/itchyny/go-flags v1.5.0 // indirect
	github.com/nats-io/jwt v0.3.2 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rs/zerolog v1.20.0 // indirect
	github.com/smira/go-statsd v1.3.2 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/vincent-petithory/dataurl v0.0.0-20191104211930-d1553a71de50 // indirect
	golang.org/x/text v0.3.8 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1
	gotest.tools/v3 v3.0.3 // indirect
)

replace github.com/uber/jaeger-lib v1.5.0+incompatible => github.com/uber/jaeger-lib v1.5.0

// Force version due to a vulnerbility in the versions benthos currently uses
replace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.7.4
