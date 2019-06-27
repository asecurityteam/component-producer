package producer

import (
	"context"
	"testing"

	"github.com/asecurityteam/settings"
	"github.com/stretchr/testify/require"
)

var benthosConfig = `http:
  address: 0.0.0.0:4195
  read_timeout: 5s
  root_path: /benthos
  debug_endpoints: false
pipeline:
  processors: []
  threads: 1
output:
  type: serverless_response
  serverless_response:
    name: lambda
logger:
  prefix: benthos
  level: OFF
  add_timestamp: true
  json_format: true
  static_fields:
    '@service': benthos
metrics:
  type: http_server
  http_server: {}
  prefix: benthos
tracer:
  type: none
  none: {}
shutdown_timeout: 20s
`

func TestProducer(t *testing.T) {
	src := settings.NewMapSource(map[string]interface{}{
		"producer": map[string]interface{}{
			"type": "POST",
			"post": map[string]interface{}{
				"endpoint": "http://localhost",
				"httpclient": map[string]interface{}{
					"type": "DEFAULT",
				},
			},
		},
	})
	p, err := New(context.Background(), src)
	require.Nil(t, err)
	require.NotNil(t, p)

	src = settings.NewMapSource(map[string]interface{}{
		"producer": map[string]interface{}{
			"type": "BENTHOS",
			"benthos": map[string]interface{}{
				"yaml": benthosConfig,
			},
		},
	})
	p, err = New(context.Background(), src)
	require.Nil(t, err)
	require.NotNil(t, p)

	src = settings.NewMapSource(map[string]interface{}{
		"producer": map[string]interface{}{
			"type": "null",
		},
	})
	p, err = New(context.Background(), src)
	require.Nil(t, err)
	require.NotNil(t, p)

	src = settings.NewMapSource(map[string]interface{}{
		"producer": map[string]interface{}{
			"type": "MISSING",
		},
	})
	_, err = New(context.Background(), src)
	require.NotNil(t, err)
}
