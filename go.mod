module github.com/asecurityteam/component-producer

go 1.15

require (
	bitbucket.org/atlassian/go-asap v0.0.0-20201116174856-38f0143fcabd // indirect
	cloud.google.com/go v0.75.0 // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go v50.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.17 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.10 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/Jeffail/benthos/v4 v4.13.0
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/armon/go-metrics v0.3.6 // indirect
	github.com/asecurityteam/component-httpclient v0.2.0
	github.com/asecurityteam/component-stat v0.3.0 // indirect
	github.com/asecurityteam/httpstats v0.0.0-20201215174437-106328c66daa // indirect
	github.com/asecurityteam/logevent v1.4.0 // indirect
	github.com/asecurityteam/runhttp v0.4.0 // indirect
	github.com/asecurityteam/settings v0.4.0
	github.com/asecurityteam/transport v1.5.1 // indirect
	github.com/asecurityteam/transportd v1.2.4 // indirect
	github.com/aws/aws-sdk-go v1.36.31 // indirect
	github.com/bitly/go-hostpool v0.1.0 // indirect
	github.com/containerd/continuity v0.0.0-20201208142359-180525291bb7 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/getkin/kin-openapi v0.36.0 // indirect
	github.com/go-chi/chi v4.1.2+incompatible // indirect
	github.com/go-openapi/swag v0.19.13 // indirect
	github.com/gocql/gocql v0.0.0-20210126081024-994808f8e61e // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.5.0
	github.com/google/uuid v1.2.0 // indirect
	github.com/hashicorp/go-hclog v0.15.0 // indirect
	github.com/itchyny/gojq v0.12.1 // indirect
	github.com/jhump/protoreflect v1.8.1 // indirect
	github.com/linkedin/goavro/v2 v2.10.0 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/nats-io/nats-streaming-server v0.20.0 // indirect
	github.com/nats-io/stan.go v0.8.2 // indirect
	github.com/nxadm/tail v1.4.6 // indirect
	github.com/olivere/elastic/v7 v7.0.22 // indirect
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.4 // indirect
	github.com/pebbe/zmq4 v1.2.2 // indirect
	github.com/prometheus/client_golang v1.9.0 // indirect
	github.com/prometheus/procfs v0.3.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rs/zerolog v1.20.0 // indirect
	github.com/smira/go-statsd v1.3.2 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/vincent-petithory/dataurl v0.0.0-20191104211930-d1553a71de50 // indirect
	golang.org/x/oauth2 v0.0.0-20210125201302-af13f521f196 // indirect
	golang.org/x/text v0.3.8 // indirect
	google.golang.org/api v0.37.0 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
	google.golang.org/grpc v1.35.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1
	gotest.tools/v3 v3.0.3 // indirect
)

replace github.com/uber/jaeger-lib v1.5.0+incompatible => github.com/uber/jaeger-lib v1.5.0

// Force version due to a vulnerbility in the versions benthos currently uses
replace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.7.4
