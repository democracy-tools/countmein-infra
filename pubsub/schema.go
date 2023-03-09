package pubsub

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
)

// CreateProtoSchema creates a schema resource from a schema proto file.
func CreateProtoSchema(project, schema, protoFile string) error {

	ctx := context.Background()
	client, err := pubsub.NewSchemaClient(ctx, project)
	if err != nil {
		log.Errorf("failed to create pubsub schema client with '%v'", err)
		return err
	}
	defer client.Close()

	protoSource, err := ioutil.ReadFile(protoFile)
	if err != nil {
		log.Errorf("failed to read file '%s' with '%v'", protoFile, err)
		return err
	}

	config := pubsub.SchemaConfig{
		Type:       pubsub.SchemaProtocolBuffer,
		Definition: string(protoSource),
	}
	s, err := client.CreateSchema(ctx, schema, config)
	if err != nil {
		log.Errorf("failed to create schema '%s' with '%v'", schema, err)
		return err
	}
	log.Infof("schema '%v' created", s)

	return nil
}
