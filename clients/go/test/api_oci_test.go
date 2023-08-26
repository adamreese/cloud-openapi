/*
Fermyon Cloud API

Testing OciApiService

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

func Test_cloud_openapi_OciApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OciApiService ApiOciGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OciApi.ApiOciGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OciApiService ApiOciNameBlobsUploadsDigestDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var digest string
		var name string

		httpRes, err := apiClient.OciApi.ApiOciNameBlobsUploadsDigestDelete(context.Background(), digest, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OciApiService ApiOciNameBlobsUploadsDigestGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var digest string
		var name string

		httpRes, err := apiClient.OciApi.ApiOciNameBlobsUploadsDigestGet(context.Background(), digest, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OciApiService ApiOciNameBlobsUploadsDigestPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var digest string
		var name string

		httpRes, err := apiClient.OciApi.ApiOciNameBlobsUploadsDigestPatch(context.Background(), digest, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OciApiService ApiOciNameBlobsUploadsDigestPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var digest string
		var name string

		httpRes, err := apiClient.OciApi.ApiOciNameBlobsUploadsDigestPut(context.Background(), digest, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OciApiService ApiOciNameBlobsUploadsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.OciApi.ApiOciNameBlobsUploadsPost(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OciApiService ApiOciNameManifestsReferenceHead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var reference string

		httpRes, err := apiClient.OciApi.ApiOciNameManifestsReferenceHead(context.Background(), name, reference).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OciApiService ApiOciNameManifestsReferencePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var reference string

		httpRes, err := apiClient.OciApi.ApiOciNameManifestsReferencePut(context.Background(), name, reference).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}