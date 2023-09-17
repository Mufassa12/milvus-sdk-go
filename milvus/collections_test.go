package milvus_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	milvus_test "github.com/mufassa12/milvus-sdk-go/milvus/test"
)

func TestListCollections(t *testing.T) {
	ts := milvus_test.NewTestServer()
	ts.RegisterHandler("/collections", func(w http.ResponseWriter, _ *http.Request) {
		resBytes, _ := json.Marshal([]string{
			"collection_1",
			"collection_2",
		})
		fmt.Fprintln(w, string(resBytes))
	})
	ts.HTTPServer.Start()
	defer ts.HTTPServer.Close()

	client := milvus_test.NewTestClient(ts)
	_, err := client.Collections().ListCollections()
	t.Helper()
	if err != nil {
		t.Error(err, "ListCollections error")
		t.Fail()
	}
}
