package bq_test

import (
	"testing"

	"github.com/democracy-tools/countmein-infra/bq"
	"github.com/stretchr/testify/require"
)

func TestOnBoarding(t *testing.T) {

	require.NoError(t, bq.CreateTableAnnouncement())
	require.NoError(t, bq.GetInstance().Close())
}

func TestOffBoarding(t *testing.T) {

	require.NoError(t, bq.DeleteTableAnnouncement())
	require.NoError(t, bq.GetInstance().Close())
}
