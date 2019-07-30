package producer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	httpclient "github.com/asecurityteam/component-httpclient"
)

type postProducer struct {
	Client   *http.Client
	Endpoint *url.URL
}

// Produce an event to the endpoint. Any 2xx series response is a success.
func (p *postProducer) Produce(ctx context.Context, event interface{}) (interface{}, error) {
	b, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	r, _ := http.NewRequest(http.MethodPost, p.Endpoint.String(), ioutil.NopCloser(bytes.NewReader(b)))
	res, err := p.Client.Do(r.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// Drain the body no matter what in order to allow for connection re-use
	// in HTTP/1.x systems.
	rb, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("failed to post. status(%d) reason(%s)", res.StatusCode, string(rb))
	}
	return event, nil
}

// POSTConfig contains settings for the HTTP POST producer.
type POSTConfig struct {
	Endpoint   string `description:"The URL to POST."`
	HTTPClient *httpclient.Config
}

// Name of the configuration section.
func (*POSTConfig) Name() string {
	return "post"
}

// POSTComponent is a component for creating an HTTP POST producer.
type POSTComponent struct {
	HTTP *httpclient.Component
}

// NewPOSTComponent populates a ProducerPOSTComponent with defaults.
func NewPOSTComponent() *POSTComponent {
	return &POSTComponent{HTTP: httpclient.NewComponent()}
}

// Settings returns the default configuration.
func (c *POSTComponent) Settings() *POSTConfig {
	return &POSTConfig{
		HTTPClient: c.HTTP.Settings(),
	}
}

// New constructs a benthos configuration.
func (c *POSTComponent) New(ctx context.Context, conf *POSTConfig) (Producer, error) {
	if conf.Endpoint == "" {
		return nil, fmt.Errorf("missing POST producer endpoint")
	}
	u, err := url.Parse(conf.Endpoint)
	if err != nil {
		return nil, err
	}
	tr, err := c.HTTP.New(ctx, conf.HTTPClient)
	if err != nil {
		return nil, err
	}
	cl := &http.Client{Transport: tr}
	return &postProducer{Endpoint: u, Client: cl}, nil
}
