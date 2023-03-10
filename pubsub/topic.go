package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/democracy-tools/countmein-infra/env"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func CreateTopicWithSchema(topic, schema string) error {

	conf, err := google.JWTConfigFromJSON(env.GetToken(), pubsub.ScopePubSub)
	if err != nil {
		log.Fatalf("failed to config pubsub JWT with %q", err)
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, env.GetProject(), option.WithTokenSource(conf.TokenSource(ctx)))
	if err != nil {
		log.Fatalf("failed to create pubsub client with '%v'", err)
	}
	defer client.Close()

	_, err = client.CreateTopicWithConfig(ctx, topic, &pubsub.TopicConfig{
		SchemaSettings: &pubsub.SchemaSettings{
			Schema:   fmt.Sprintf("projects/%s/schemas/%s", env.GetProject(), schema),
			Encoding: pubsub.EncodingJSON,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create topic '%s' with schema '%s' with '%v'", topic, schema, err)
	}
	log.Infof("topic '%s' with schema '%s' created", topic, schema)

	return nil
}
