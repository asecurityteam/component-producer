package producer

import (
	"context"
	"fmt"
	"strings"

	"github.com/asecurityteam/settings"
)

const (
	// TypeBenthos is used to select a Benthos producer.
	TypeBenthos = "BENTHOS"
	// TypePOST is used to select an HTTP POST client.
	TypePOST = "POST"
)

// Config wraps all producer related configuration.
type Config struct {
	Type    string `default:"The type of producer. The choices are BENTHOS and POST."`
	Benthos *BenthosConfig
	POST    *POSTConfig
}

// Name of the configuration.
func (*Config) Name() string {
	return "producer"
}

// Component is the top level producer component.
type Component struct {
	Benthos *BenthosComponent
	POST    *POSTComponent
}

// NewComponent populates a ProducerComponent with defaults.
func NewComponent() *Component {
	return &Component{
		Benthos: NewBenthosComponent(),
		POST:    NewPOSTComponent(),
	}
}

// Settings generates a default configuration.
func (c *Component) Settings() *Config {
	return &Config{
		Type:    "BENTHOS",
		Benthos: c.Benthos.Settings(),
		POST:    c.POST.Settings(),
	}
}

// New constructs a Producer from the given configuration.
func (c *Component) New(ctx context.Context, conf *Config) (Producer, error) {
	switch {
	case strings.EqualFold(conf.Type, TypeBenthos):
		return c.Benthos.New(ctx, conf.Benthos)
	case strings.EqualFold(conf.Type, TypePOST):
		return c.POST.New(ctx, conf.POST)
	default:
		return nil, fmt.Errorf("unknown producer type %s", conf.Type)
	}
}

// Load is a convenience method for binding the source to the component.
func Load(ctx context.Context, source settings.Source, c *Component) (Producer, error) {
	dst := new(Producer)
	err := settings.NewComponent(ctx, source, c, dst)
	if err != nil {
		return nil, err
	}
	return *dst, nil
}

// New is the top-level entry point for creating a producer client.
func New(ctx context.Context, source settings.Source) (Producer, error) {
	return Load(ctx, source, NewComponent())
}
