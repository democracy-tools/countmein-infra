package bq

import (
	"cloud.google.com/go/bigquery"
)

func CreateTableAnnouncement() error {

	return GetInstance().CreateTable(TableAnnouncement,
		&bigquery.TableMetadata{
			Schema: bigquery.Schema{
				{Name: "id", Type: bigquery.StringFieldType},
				{Name: "user_id", Type: bigquery.StringFieldType},
				{Name: "user_device_id", Type: bigquery.StringFieldType},
				{Name: "user_device_type", Type: bigquery.StringFieldType},
				{Name: "seen_device_id", Type: bigquery.StringFieldType},
				{Name: "seen_device_type", Type: bigquery.StringFieldType},
				{Name: "location_latitude", Type: bigquery.FloatFieldType},
				{Name: "location_longitude", Type: bigquery.FloatFieldType},
				{Name: "user_time", Type: bigquery.IntegerFieldType},
				{Name: "server_time", Type: bigquery.IntegerFieldType},
				// *** pubsub publisher fields ***
				{Name: "attributes", Type: bigquery.StringFieldType},
				{Name: "publish_time", Type: bigquery.TimestampFieldType},
				{Name: "subscription_name", Type: bigquery.StringFieldType},
				// pubsub might send duplicate messages due to its at least once delivery property,
				// this will create duplicate records in BigQuery which can be identified
				// using the message_id column
				{Name: "message_id", Type: bigquery.StringFieldType},
			}})
}

func DeleteTableAnnouncement() error {

	return GetInstance().DeleteTable(TableAnnouncement)
}
