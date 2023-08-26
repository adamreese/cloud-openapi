/*
Fermyon Cloud API

Testing SqlDatabasesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cloud_openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/fermyon/cloud-openapi"
)

func Test_cloud_openapi_SqlDatabasesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SqlDatabasesApiService ApiSqlDatabasesCreatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesCreatePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SqlDatabasesApiService ApiSqlDatabasesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SqlDatabasesApiService ApiSqlDatabasesExecutePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesExecutePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SqlDatabasesApiService ApiSqlDatabasesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SqlDatabasesApi.ApiSqlDatabasesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
