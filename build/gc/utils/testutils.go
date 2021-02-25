package utils

import (
	"encoding/json"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"net/http"
	"testing"
)

var swaggerURL = "https://api.mypurecloud.com/api/v2/docs/swagger"

func TestAgainstSwaggerDefinition(t *testing.T, command models.HandWrittenOperation) {
	path, err := getResourceDefinition(command.Path)
	if err != nil {
		t.Fatalf("Received error getting path structure from swagger, err: %v", err)
	}

	method := path.GetMethod(command.Method)
	if method == nil {
		t.Fatalf("Failed to retrieve method definition from swagger")
	}

	if command.Description != "" {
		if command.Description != method.Description {
			t.Errorf("Description doesn't match swagger. got: %v, want :%v", command.Description, method.Description)
		}
	}
}

func getResourceDefinition(path string) (*models.Path, error) {
	swaggerFile, err := downloadFile(swaggerURL)
	if err != nil {
		return nil, err
	}

	var swaggerDefinition map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(swaggerFile), &swaggerDefinition); err != nil {
		return nil, err
	}

	var swaggerDefinitionPaths map[string]*json.RawMessage
	if err := json.Unmarshal(*swaggerDefinition["paths"], &swaggerDefinitionPaths); err != nil {
		return nil, err
	}

	pathObject := &models.Path{}
	if err := json.Unmarshal(*swaggerDefinitionPaths[path], pathObject); err != nil {
		return nil, err
	}

	return pathObject, nil
}

func downloadFile(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responseData := string(response)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		httpError := models.HttpStatusError{Verb: http.MethodGet, Path: uri, StatusCode: resp.StatusCode, Headers: resp.Header, Body: fmt.Sprintf("%s", pretty.Pretty([]byte(responseData)))}
		return "", httpError
	}

	return responseData, nil
}
