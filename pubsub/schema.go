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

const (
	EnvKeyPubSub = "PUBSUB_KEY"
)

var (
	singleton *ClientWrapper
	once      sync.Once
)

type ClientWrapper struct {
	client *pubsub.SchemaClient
}

func GetInstance() *ClientWrapper {

	once.Do(func() {
		conf, err := google.JWTConfigFromJSON(env.GetToken(), pubsub.ScopePubSub)
		if err != nil {
			log.Fatalf("failed to config pubsub JWT with %q", err)
		}

		ctx := context.Background()
		client, err := pubsub.NewSchemaClient(ctx, env.GetProject(), option.WithTokenSource(conf.TokenSource(ctx)))
		if err != nil {
			log.Fatalf("failed to create pubsub schema client with '%v'", err)
		}

		singleton = &ClientWrapper{client: client}
	})

	return singleton
}

// CreateProtoSchema creates a schema resource from a schema proto file.
func (c *ClientWrapper) CreateProtoSchema(schema, protoFile string) error {

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

func (c *ClientWrapper) Close() error {

	return c.client.Close()
}
