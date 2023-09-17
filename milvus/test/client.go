package milvus_test

import (
	"github.com/mufassa12/milvus-sdk-go/milvus"
)

const (
	test_api_key = "this-is-my-secure-apikey-do-not-steal!!"
	test_region  = "this-is-my-cloud"
)

func GetTestAuthToken() string {
	return test_api_key
}

func NewTestClient(ts *TestServer) *milvus.Client {
	client := milvus.NewClient(test_api_key, test_region)
	if ts != nil {
		client.Region = ""
		client.BaseUrl = ts.HTTPServer.URL
	}
	return client
}
