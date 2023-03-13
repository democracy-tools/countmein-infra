package pubsub

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/democracy-tools/countmein-infra/env"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var (
	singleton *PubSubClientWrapper
	once      sync.Once
)

type PubSubClientWrapper struct {
	client  *pubsub.Client
	dataset string
}

func GetPubSubInstance() *PubSubClientWrapper {

	once.Do(func() {
		conf, err := google.JWTConfigFromJSON(env.GetToken(), pubsub.ScopePubSub)
		if err != nil {
			log.Fatalf("failed to config pubsub JWT with '%v'", err)
		}

		ctx := context.Background()
		client, err := pubsub.NewClient(ctx, env.Project, option.WithTokenSource(conf.TokenSource(ctx)))
		if err != nil {
			log.Fatalf("failed to create pubsub client with '%v'", err)
		}

		singleton = &PubSubClientWrapper{client: client, dataset: "dev"}
	})

	return singleton
}

func (c *PubSubClientWrapper) CreateTopicWithSchema(topic, schema string) (*pubsub.Topic, error) {

	res, err := c.client.CreateTopicWithConfig(context.Background(), topic, &pubsub.TopicConfig{
		SchemaSettings: &pubsub.SchemaSettings{
			Schema:   fmt.Sprintf("projects/%s/schemas/%s", env.Project, schema),
			Encoding: pubsub.EncodingJSON,
		},
	})
	if err != nil {
		log.Errorf("failed to create topic '%s' with schema '%s' with '%v'", topic, schema, err)
		return nil, err
	}
	log.Infof("topic '%s' with schema '%s' created", topic, schema)

	return res, nil
}

/*
Pub/Sub creates and maintains a service account for each project in the format service-project-number@gcp-sa-pubsub.iam.gserviceaccount.com.
To create a BigQuery subscription, the Pub/Sub service account must have permission to write to the specific BigQuery table and to read the table metadata.
Grant the BigQuery Data Editor (roles/bigquery.dataEditor) role and the BigQuery Metadata Viewer (roles/bigquery.metadataViewer) role to the Pub/Sub service account.
Steps:
1. In the Google Cloud console, go to the IAM -> Click Grant access
2. In the Add Principals section, enter the name of your Pub/Sub service account.
The format of the service account is service-project-number@gcp-sa-pubsub.iam.gserviceaccount.com.
For example, for a project with project-number=112233445566,
the service account is of the format service-112233445566@gcp-sa-pubsub.iam.gserviceaccount.com.
3. In the Assign Roles section, click Add roles, and add:
- BigQuery Data Editor role
- BigQuery Metadata Viewer role
4. Click Save
*/
func (c *PubSubClientWrapper) CreateBigQuerySubscription(id string, topic *pubsub.Topic, table string) error {

	sub, err := c.client.CreateSubscription(context.Background(), id, pubsub.SubscriptionConfig{
		Topic: topic,
		BigQueryConfig: pubsub.BigQueryConfig{
			Table:         fmt.Sprintf("%s.%s.%s", env.Project, env.Dataset, table),
			WriteMetadata: true,
		},
	})
	if err != nil {
		log.Errorf("failed to create pubsub subscription for big-query with '%v'", err)
		return err
	}
	log.Infof("created BigQuery subscription '%v' topic '%s'", sub, topic.String())

	return nil
}

func (c *PubSubClientWrapper) GetTopic(id string) *pubsub.Topic {

	return c.client.TopicInProject(id, env.Project)
}

func (c *PubSubClientWrapper) DetachSubscription(sub string) error {

	_, err := c.client.DetachSubscription(context.Background(), sub)
	if err != nil {
		log.Errorf("failed to delete subscription '%s' with '%v'", sub, err)
		return err
	}
	log.Infof("detach subscription '%s'", sub)

	return nil
}

func (c *PubSubClientWrapper) DeleteTopic(id string) error {

	log.Infof("deleting topic '%s'", id)
	return c.GetTopic(id).Delete(context.Background())
}

func (c *PubSubClientWrapper) Close() error {

	return c.client.Close()
}
