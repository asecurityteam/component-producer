package producer

import "context"

type nullProducer struct{}

func (p *nullProducer) Produce(ctx context.Context, event interface{}) (interface{}, error) {
	return event, nil
}

// NullConfig contains settings for the Null producer.
type NullConfig struct{}

// Name of the configuration section.
func (*NullConfig) Name() string {
	return "nullproducer"
}

// NullComponent is a component for creating an HTTP POST producer.
type NullComponent struct{}

// NewNullComponent populates a NullComponentt with defaults.
func NewNullComponent() *NullComponent {
	return &NullComponent{}
}

// Settings returns the default configuration.
func (c *NullComponent) Settings() *NullConfig {
	return &NullConfig{}
}

// New constructs a benthos configuration.
func (c *NullComponent) New(ctx context.Context, conf *NullConfig) (Producer, error) {
	return &nullProducer{}, nil
}
