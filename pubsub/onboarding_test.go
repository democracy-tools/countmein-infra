package pubsub_test

import (
	"fmt"
	"testing"

	"github.com/democracy-tools/countmein-infra/pubsub"
	"github.com/stretchr/testify/require"
)

func TestOnBoarding(t *testing.T) {

	const id = "announcement"

	require.NoError(t, pubsub.CreateProtoBufSchema(id, fmt.Sprintf("%s.pb", id)))
	require.NoError(t, pubsub.CreateTopicWithSchema(id, id))
}
