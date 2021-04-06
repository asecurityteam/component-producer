package producer

import (
	"context"
	"fmt"
	"strings"

	"github.com/asecurityteam/settings"
)

const (
	// TypePOST is used to select an HTTP POST client.
	TypePOST = "POST"
	// TypeNull is used to disable event production.
	TypeNull = "NULL"
)

// Config wraps all producer related configuration.
type Config struct {
	Type string `default:"The type of producer. The choices are BENTHOS and POST."`
	POST *POSTConfig
	Null *NullConfig
}

// Name of the configuration.
func (*Config) Name() string {
	return "producer"
}

// Component is the top level producer component.
type Component struct {
	POST *POSTComponent
	Null *NullComponent
}

// NewComponent populates a ProducerComponent with defaults.
func NewComponent() *Component {
	return &Component{
		POST: NewPOSTComponent(),
		Null: NewNullComponent(),
	}
}

// Settings generates a default configuration.
func (c *Component) Settings() *Config {
	return &Config{
		Type: "BENTHOS",
		POST: c.POST.Settings(),
		Null: c.Null.Settings(),
	}
}

// New constructs a Producer from the given configuration.
func (c *Component) New(ctx context.Context, conf *Config) (Producer, error) {
	switch {
	case strings.EqualFold(conf.Type, TypePOST):
		return c.POST.New(ctx, conf.POST)
	case strings.EqualFold(conf.Type, TypeNull):
		return c.Null.New(ctx, conf.Null)
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
