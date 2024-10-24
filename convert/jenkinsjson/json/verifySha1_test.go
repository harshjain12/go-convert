package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvertVerifySha1(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	filePath := filepath.Join(workingDir, "../convertTestFiles/verifySha1/verifySha1Snippet.json")
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read JSON file: %v", err)
	}

	var node1 Node
	if err := json.Unmarshal(jsonData, &node1); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	tests := []struct {
		json Node
		want Node
	}{
		{
			json: node1,
			want: Node{
				AttributesMap: map[string]string{
					"ci.pipeline.run.user":                 "SYSTEM",
					"jenkins.pipeline.step.id":             "9",
					"jenkins.pipeline.step.name":           "Verify the SHA1 of a given file",
					"jenkins.pipeline.step.plugin.name":    "pipeline-utility-steps",
					"jenkins.pipeline.step.plugin.version": "2.17.0",
					"jenkins.pipeline.step.type":           "verifySha1",
					"harness-attribute":                    "{\n  \"hash\" : \"22596363b3de40b06f981fb85d82312e8c0ed511\",\n  \"file\" : \"file.txt\"\n}",
					"harness-others":                       "",
				},
				Name:         "ag-readJSON #75",
				Parent:       "ag-readJSON",
				ParentSpanId: "4906ee7f321b135e",
				SpanId:       "98c85f37933c3250",
				SpanName:     "verifySha1",
				TraceId:      "6efe56554a7922b3e92c099fc96cccc4",
				Type:         "Run Phase Span",
				ParameterMap: map[string]any{
					"file": "file.txt",
					"hash": "22596363b3de40b06f981fb85d82312e8c0ed511"},
			},
		},
	}

	for i, test := range tests {
		got := test.json
		if diff := cmp.Diff(got, test.want); diff != "" {
			t.Errorf("Unexpected parsing results for test %v", i)
			t.Log(diff)
		}
	}
}