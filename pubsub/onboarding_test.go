package pubsub_test

import (
	"testing"

	"github.com/democracy-tools/countmein-infra/pubsub"
	"github.com/stretchr/testify/require"
)

func TestOnBoarding(t *testing.T) {

	require.NoError(t, pubsub.CreateProtoSchema("democracy-tools", "announcement", "announcement.pb"))
}
