package bq_test

import (
	"testing"

	"github.com/democracy-tools/countmein-infra/bq"
	"github.com/stretchr/testify/require"
)

func init() {

	bq.LoadEnv()
}

func TestOnBoarding(t *testing.T) {

	require.NoError(t, bq.CreateTableAnnouncement())
}

func TestOffBoarding(t *testing.T) {

	require.NoError(t, bq.DeleteTableAnnouncement())
}
