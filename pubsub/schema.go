package pubsub

import (
	"context"
	"io/ioutil"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/democracy-tools/countmein-infra/env"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var (
	pubSubSchemaClientSingleton *PubSubSchemaClientWrapper
	pubSubSchemaClientOnce      sync.Once
)

type PubSubSchemaClientWrapper struct {
	client *pubsub.SchemaClient
}

func GetPubSubSchemaInstance() *PubSubSchemaClientWrapper {

	pubSubSchemaClientOnce.Do(func() {
		conf, err := google.JWTConfigFromJSON(env.GetToken(), pubsub.ScopePubSub)
		if err != nil {
			log.Fatalf("failed to config pubsub JWT with '%v'", err)
		}

		ctx := context.Background()
		client, err := pubsub.NewSchemaClient(ctx, env.Project, option.WithTokenSource(conf.TokenSource(ctx)))
		if err != nil {
			log.Fatalf("failed to create pubsub schema client with '%v'", err)
		}

		pubSubSchemaClientSingleton = &PubSubSchemaClientWrapper{client: client}
	})

	return pubSubSchemaClientSingleton
}

// CreateProtoBufSchema creates a schema resource from a schema proto file
func (c *PubSubSchemaClientWrapper) CreateProtoBufSchema(schema, protoFile string) error {

	protoSource, err := ioutil.ReadFile(protoFile)
	if err != nil {
		log.Errorf("failed to read file '%s' with '%v'", protoFile, err)
		return err
	}

	config := pubsub.SchemaConfig{
		Type:       pubsub.SchemaProtocolBuffer,
		Definition: string(protoSource),
	}
	s, err := c.client.CreateSchema(context.Background(), schema, config)
	if err != nil {
		log.Errorf("failed to create schema '%s' with '%v'", schema, err)
		return err
	}
	log.Infof("schema '%v' created", s)

	return nil
}

func (c *PubSubSchemaClientWrapper) DeleteSchema(id string) error {

	log.Infof("deleting schema '%s'...", id)
	return c.client.DeleteSchema(context.Background(), id)
}

func (c *PubSubSchemaClientWrapper) Close() error {

	return c.client.Close()
}
