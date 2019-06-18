package producer

import (
	"context"

	"github.com/Jeffail/benthos/lib/config"
	"github.com/Jeffail/benthos/lib/serverless"
	"github.com/Jeffail/benthos/lib/util/text"
	"gopkg.in/yaml.v2"
)

type benthosProducer struct {
	Handler *serverless.Handler
}

// Produce an event using the Benthos serverless handler. The handler is
// exactly a producer as we want it. We only adapt the function name here.
func (p *benthosProducer) Produce(ctx context.Context, event interface{}) (interface{}, error) {
	return p.Handler.Handle(ctx, event)
}

// BenthosConfig adapts the Benthos configuration system to this one.
type BenthosConfig struct {
	YAML string `description:"The YAML or JSON text of a Benthos configuration."`
}

// Name of the configuration section.
func (*BenthosConfig) Name() string {
	return "benthos"
}

// BenthosComponent is a component for creating a Benthos producer.
type BenthosComponent struct{}

// NewBenthosComponent generates a ProducerBenthosComponent and
// populates it with defaults
func NewBenthosComponent() *BenthosComponent {
	return &BenthosComponent{}
}

// Settings returns the default configuration.
func (*BenthosComponent) Settings() *BenthosConfig {
	return &BenthosConfig{}
}

// newBenthosConfig parses a Benthos YAML config file and replaces all
// environment variables. This is exactly what is done in the Benthos project
// when starting one of the main() functions.
func newBenthosConfig(b []byte) (config.Type, error) {
	conf := config.New()
	err := yaml.Unmarshal(text.ReplaceEnvVariables(b), &conf)
	return conf, err
}

// New constructs a benthos producer.
func (*BenthosComponent) New(ctx context.Context, conf *BenthosConfig) (Producer, error) {
	cfg, err := newBenthosConfig([]byte(conf.YAML))
	if err != nil {
		return nil, err
	}
	handler, err := serverless.NewHandler(cfg)
	if err != nil {
		return nil, err
	}
	return &benthosProducer{Handler: handler}, nil
}
