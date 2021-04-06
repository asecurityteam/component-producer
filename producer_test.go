package producer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/asecurityteam/settings"
)

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
