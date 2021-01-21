package utils

import (
	"encoding/json"
	"fmt"
	"gc/models"
	"github.com/tidwall/pretty"
	"io/ioutil"
	"net/http"
)

var swaggerURL = "https://api.mypurecloud.com/api/v2/docs/Swagger"

func GetResourceDefinition(path string) (*models.Path, error) {
	swaggerFile, err := DownloadFile(swaggerURL)
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

func DownloadFile(uri string) (string, error) {
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