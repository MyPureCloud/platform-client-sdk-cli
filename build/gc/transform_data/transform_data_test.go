package transform_data

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

// global test data - single json objects
var jsonObjects = [...]string {
	"{\"name\": \"michael roddy\", \"id\": \"12345\"}",
	"{\"name\": \"michael roddy\", \"id\": \"12345\", \"address\": {\"town\": \"athlone\"}}",
	"{\"name\": \"michael roddy\", \"id\": \"12345\", \"address\": [{\"town\": \"athlone\"}]}",
}
// global test data - arrays of json objects
var arrayOfJsonObjects = [...]string {
	"[{\"name\": \"michael roddy\", \"id\": \"12345\"}]",
	"[{\"name\": \"michael roddy\", \"id\": \"12345\", \"address\": {\"town\": \"athlone\"}}]",
	"[{\"name\": \"michael roddy\", \"id\": \"12345\", \"address\": [{\"town\": \"athlone\"}]}]",
}

// the ConvertJsonToMap() will return either a map[string]interface{} or a []map[string]interface{}, so we have two scenarios to test for
func TestConvertJsonToMap(t *testing.T) {

	// the outputs expected when converting a json object to map
	mapOutputs := [...]map[string]interface{} {
		{"name":  "michael roddy", "id": "12345"},
		{"name":  "michael roddy", "id": "12345", "address": map[string]interface{}{"town": "athlone"}},
		{"name":  "michael roddy", "id": "12345", "address": []map[string]interface{}{{"town": "athlone"}}},
	}
	// the outputs expected when converting arrays of json objects to map
	sliceOfMapOutputs := [...][]map[string]interface{} {
		{{"name": "michael roddy", "id": "12345"}},
		{{"name":  "michael roddy", "id": "12345", "address": map[string]interface{}{"town": "athlone"}}},
		{{"name":  "michael roddy", "id": "12345", "address": []map[string]interface{}{{"town": "athlone"}}}},
	}

	type testMap struct {
		testNum int
		input string
		expectedOutput map[string]interface{}
	}

	type testSliceOfMap struct {
		testNum int
		input string
		expectedOutput []map[string]interface{}
	}

	mapTestCases := []testMap {
		{testNum: 1, input: jsonObjects[0], expectedOutput: mapOutputs[0]},
		{testNum: 2, input: jsonObjects[1], expectedOutput: mapOutputs[1]},
		{testNum: 3, input: jsonObjects[2], expectedOutput: mapOutputs[2]},
	}

	sliceOfMapTestCases := []testSliceOfMap {
		{testNum: 1, input: arrayOfJsonObjects[0], expectedOutput: sliceOfMapOutputs[0]},
		{testNum: 2, input: arrayOfJsonObjects[1], expectedOutput: sliceOfMapOutputs[1]},
		{testNum: 3, input: arrayOfJsonObjects[2], expectedOutput: sliceOfMapOutputs[2]},
	}

	// testing if ConvertJsonToMap() returns a map[string]interface{}
	for _, test := range mapTestCases {
		actualOutput := ConvertJsonToMap(test.input)
		if fmt.Sprint(actualOutput) != fmt.Sprint(test.expectedOutput) {
			t.Errorf("Did not get the right output: expectedOutput: %v, actualOutput: %v", test.expectedOutput, actualOutput)
		}
	}

	// testing if ConvertJsonToMap() returns a []map[string]interface{}
	for _, test := range sliceOfMapTestCases {
		actualOutput := ConvertJsonToMap(test.input)
		if fmt.Sprint(actualOutput) != fmt.Sprint(test.expectedOutput) {
			t.Errorf("Did not get the right output: expectedOutput: %v, actualOutput: %v", test.expectedOutput, actualOutput)
		}
	}

}

func TestProcessTemplateFile(t *testing.T) {

	templateFiles := [...]string {
		filepath.Base("./transform_data/test_template_file_1.gotmpl"),
		filepath.Base("./transform_data/test_template_file_2.gotmpl"),
		filepath.Base("./transform_data/test_template_file_3.gotmpl"), // sprig template file
	}

	templateOutputs := [...]string {
		"id: 12345 name: michael roddy",
		"address: town: athlone",
		"ATHLONEATHLONEATHLONEATHLONEATHLONE",
	}

	type testTemplateFile struct {
		testNum int
		jsonInput string
		templateInput string
		expectedOutput string
	}

	templateFileTestCases := []testTemplateFile {
		{testNum: 1, jsonInput: jsonObjects[0], templateInput: templateFiles[0], expectedOutput: templateOutputs[0]},
		{testNum: 1, jsonInput: arrayOfJsonObjects[1], templateInput: templateFiles[1], expectedOutput: templateOutputs[1]},
		{testNum: 1, jsonInput: jsonObjects[1], templateInput: templateFiles[2], expectedOutput: templateOutputs[2]}, // sprig test case
	}

	for _, test := range templateFileTestCases {
		mp := ConvertJsonToMap(test.jsonInput)
		TemplateFile = test.templateInput
		actualOutput := ProcessTemplateFile(mp)
		if strings.TrimSpace(actualOutput) != strings.TrimSpace(test.expectedOutput) {
			t.Errorf("Did not get the right output: expectedOutput: %v, actualOutput: %v", test.expectedOutput, actualOutput)
		}
	}

}

func TestProcessTemplateStr(t *testing.T) {

	templateStrs := [...]string {
		"{{- range $key, $val := . -}}{{$key}}: {{$val}}{{\" \"}}{{- end -}}",
		"{{- range . -}}address: town: {{.address.town}}{{- end -}}",
		"{{ .address.town | upper | repeat 5 }}", // sprig template str
	}

	templateOutputs := [...]string {
		"id: 12345 name: michael roddy",
		"address: town: athlone",
		"ATHLONEATHLONEATHLONEATHLONEATHLONE",
	}

	type testTemplateStr struct {
		testNum int
		jsonInput string
		templateInput string
		expectedOutput string
	}

	templateStrTestCases := []testTemplateStr {
		{testNum: 1, jsonInput: jsonObjects[0], templateInput: templateStrs[0], expectedOutput: templateOutputs[0]},
		{testNum: 2, jsonInput: arrayOfJsonObjects[1], templateInput: templateStrs[1], expectedOutput: templateOutputs[1]},
		{testNum: 3, jsonInput: jsonObjects[1], templateInput: templateStrs[2], expectedOutput: templateOutputs[2]}, // sprig test case
	}

	for _, test := range templateStrTestCases {
		mp := ConvertJsonToMap(test.jsonInput)
		TemplateStr = test.templateInput
		actualOutput := ProcessTemplateStr(mp)
		if strings.TrimSpace(actualOutput) != strings.TrimSpace(test.expectedOutput) {
			t.Errorf("Did not get the right output: expectedOutput: %v, actualOutput: %v", test.expectedOutput, actualOutput)
		}
	}

}