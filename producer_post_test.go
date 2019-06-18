package producer

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	httpclient "github.com/asecurityteam/component-httpclient"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var transportdConfig = `openapi: 3.0.0
x-transportd:
  backends:
    - app
  app:
    host: "http://app:8081"
    pool:
      ttl: "24h"
      count: 1
info:
  version: 1.0.0
  title: "Example"
  description: "An example"
  contact:
    name: Security Development
    email: secdev-external@atlassian.com
  license:
    name: Apache 2.0
    url: 'https://www.apache.org/licenses/LICENSE-2.0.html'
paths:
  /healthcheck:
    get:
      description: "Liveness check."
      responses:
        "200":
          description: "Success."
      x-transportd:
        backend: app
`

func TestPostProducerCantMarshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := NewMockRoundTripper(ctrl)
	u, _ := url.Parse("http://localhost")
	event := make(chan interface{})

	p := &postProducer{
		Client:   &http.Client{Transport: tr},
		Endpoint: u,
	}
	_, err := p.Produce(context.Background(), event)
	require.NotNil(t, err)
}

func TestPostProducerClientError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := NewMockRoundTripper(ctrl)
	u, _ := url.Parse("http://localhost")
	event := make(map[string]interface{})

	p := &postProducer{
		Client:   &http.Client{Transport: tr},
		Endpoint: u,
	}

	tr.EXPECT().RoundTrip(gomock.Any()).Return(nil, errors.New("error"))
	_, err := p.Produce(context.Background(), event)
	require.NotNil(t, err)
}

func TestPostProducerBadStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := NewMockRoundTripper(ctrl)
	u, _ := url.Parse("http://localhost")
	event := make(map[string]interface{})

	p := &postProducer{
		Client:   &http.Client{Transport: tr},
		Endpoint: u,
	}

	tr.EXPECT().RoundTrip(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       http.NoBody,
	}, nil)
	_, err := p.Produce(context.Background(), event)
	require.NotNil(t, err)
}

func TestPostProducerSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := NewMockRoundTripper(ctrl)
	u, _ := url.Parse("http://localhost")
	event := make(map[string]interface{})

	p := &postProducer{
		Client:   &http.Client{Transport: tr},
		Endpoint: u,
	}

	tr.EXPECT().RoundTrip(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       http.NoBody,
	}, nil)
	_, err := p.Produce(context.Background(), event)
	require.Nil(t, err)
}

func TestProducerPOSTComponent_New(t *testing.T) {
	tests := []struct {
		name         string
		conf         *POSTConfig
		wantErr      bool
		wantProducer bool
	}{
		{
			name: "missing url",
			conf: func() *POSTConfig {
				return NewPOSTComponent().Settings()
			}(),
			wantErr: true,
		},
		{
			name: "invalid url",
			conf: func() *POSTConfig {
				conf := NewPOSTComponent().Settings()
				conf.Endpoint = ":/localhost"
				return conf
			}(),
			wantErr: true,
		},
		{
			name: "default http",
			conf: func() *POSTConfig {
				conf := NewPOSTComponent().Settings()
				conf.Endpoint = "http://localhost"
				conf.HTTPClient.Type = httpclient.TypeDefault
				return conf
			}(),
			wantProducer: true,
		},
		{
			name: "smart http",
			conf: func() *POSTConfig {
				conf := NewPOSTComponent().Settings()
				conf.Endpoint = "http://localhost"
				conf.HTTPClient.Type = httpclient.TypeSmart
				conf.HTTPClient.Smart.OpenAPI = transportdConfig
				return conf
			}(),
			wantProducer: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewPOSTComponent()
			got, err := c.New(context.Background(), tt.conf)
			if tt.wantErr {
				require.NotNil(t, err)
			}
			if tt.wantProducer {
				require.Nil(t, err)
				require.NotNil(t, got)
			}
		})
	}
}
