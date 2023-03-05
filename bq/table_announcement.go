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
			}})
}

func DeleteTableAnnouncement() error {

	return GetInstance().DeleteTable(TableAnnouncement)
}
