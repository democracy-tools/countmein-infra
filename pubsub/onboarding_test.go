package pubsub_test

import (
	"fmt"
	"testing"

	"github.com/democracy-tools/countmein-infra/pubsub"
	"github.com/stretchr/testify/require"
)

func TestOnBoarding(t *testing.T) {

	const id = "announcement"

	defer pubsub.GetPubSubInstance().Close()
	defer pubsub.GetPubSubSchemaInstance().Close()

	require.NoError(t, pubsub.GetPubSubSchemaInstance().CreateProtoBufSchema(id, fmt.Sprintf("%s.pb", id)))
	topic, err := pubsub.GetPubSubInstance().CreateTopicWithSchema(id, id)
	require.NoError(t, err)
	require.NoError(t, pubsub.GetPubSubInstance().CreateBigQuerySubscription(id, topic, id))
	// require.NoError(t, pubsub.GetPubSubInstance().CreateBigQuerySubscription(
	// 	id, pubsub.GetPubSubInstance().GetTopic(id), id))
}

func TestOffBoarding(t *testing.T) {

	t.Skip("REMOVE IF YOU WANT TO DELETE PUBSUB INTEGRATION")

	const id = "announcement"

	defer pubsub.GetPubSubInstance().Close()
	defer pubsub.GetPubSubSchemaInstance().Close()

	require.NoError(t, pubsub.GetPubSubInstance().DetachSubscription(id))
	require.NoError(t, pubsub.GetPubSubInstance().DeleteTopic(id))
	require.NoError(t, pubsub.GetPubSubSchemaInstance().DeleteSchema(id))
}
