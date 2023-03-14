package bq

import (
	"context"
	"sync"

	"cloud.google.com/go/bigquery"
	"github.com/democracy-tools/countmein-infra/env"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const (
	TableAnnouncement = "announcement"
)

var (
	singleton *ClientWrapper
	once      sync.Once
)

type ClientWrapper struct {
	client  *bigquery.Client
	dataset string
}

func GetInstance() *ClientWrapper {

	once.Do(func() {
		conf, err := google.JWTConfigFromJSON(env.GetToken(), bigquery.Scope)
		if err != nil {
			log.Fatalf("failed to config bigquery JWT with '%v'", err)
		}

		ctx := context.Background()
		client, err := bigquery.NewClient(ctx, env.Project, option.WithTokenSource(conf.TokenSource(ctx)))
		if err != nil {
			log.Fatalf("failed to create bigquery client with %q", err)
		}

		singleton = &ClientWrapper{client: client, dataset: "dev"}
	})

	return singleton
}

func (c *ClientWrapper) CreateTable(table string, metadata *bigquery.TableMetadata) error {

	log.Infof("creating big-query table '%s.%s'...", c.dataset, table)
	err := c.client.Dataset(c.dataset).Table(table).Create(
		context.Background(), metadata)
	if err != nil {
		log.Errorf("failed to create table '%s.%s' with %q", c.dataset, table, err)
	}

	return err
}

func (c *ClientWrapper) DeleteTable(table string) error {

	log.Infof("deleting bigquery table '%s.%s'...", c.dataset, table)
	err := c.client.Dataset(c.dataset).Table(table).Delete(context.Background())
	if err != nil {
		log.Errorf("failed to delete table '%s.%s' with %q", c.dataset, table, err)
	}

	return err
}

func (c *ClientWrapper) Close() error {

	return c.client.Close()
}
