package milvus_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mufassa12/milvus-sdk-go/milvus"
	milvus_test "github.com/mufassa12/milvus-sdk-go/milvus/test"
)

func TestListIndexes(t *testing.T) {
	ts := milvus_test.NewTestServer()
	ts.RegisterHandler("/databases", func(w http.ResponseWriter, _ *http.Request) {
		resBytes, _ := json.Marshal([]string{
			"index_1",
			"index_2",
		})
		fmt.Fprintln(w, string(resBytes))
	})
	ts.HTTPServer.Start()
	defer ts.HTTPServer.Close()

	client := milvus_test.NewTestClient(ts)
	_, err := client.Databases().ListIndexes()
	t.Helper()
	if err != nil {
		t.Error(err, "ListIndexes error")
		t.Fail()
	}
}

func TestDescribeIndex(t *testing.T) {
	ts := milvus_test.NewTestServer()
	ts.RegisterHandler("/databases/test_index", func(w http.ResponseWriter, _ *http.Request) {
		resBytes, _ := json.Marshal(milvus.Index{
			Database: milvus.Database{
				Name:      "test_index",
				Dimension: 512,
				Metric:    milvus.IndexMetricCosine,
				Pods:      1,
				Replicas:  1,
				PodType:   "starter",
			},
			Status: milvus.Status{
				Ready: true,
				State: "Ready",
			},
		})
		fmt.Fprintln(w, string(resBytes))
	})
	ts.HTTPServer.Start()
	defer ts.HTTPServer.Close()

	client := milvus_test.NewTestClient(ts)
	_, err := client.Databases().DescribeIndex("test_index")
	t.Helper()
	if err != nil {
		t.Error(err, "DescribeIndex error")
		t.Fail()
	}
}
