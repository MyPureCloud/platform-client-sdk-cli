package restclient

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/mocks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
)

type apiClientTest struct {
	oAuthToken          string
	targetVerb          string
	targetPath          string
	targetEnv           string
	targetStatusCode    int
	targetBody          string
	expectedHeaderToken string
	expectedHost        string
	expectedPath        string
	expectedResponse    string
	expectedAuthHeader  string
	expectedStatusCode  int
}

func TestAuthorize(t *testing.T) {
	UpdateOAuthToken = mocks.UpdateOAuthToken

	mockConfig := buildMockConfig("DEFAULT", "mypurecloud.com", utils.GenerateGuid(), utils.GenerateGuid(), "")
	accessToken := "aJdvugb8k1kwnOovm2qX6LXTctJksYvdzcoXPrRDi-nL1phQhcKRN-bjcflq7CUDOmUCQv5OWuBSkPQr0peWhw"
	setRestClientDoMockForAuthorize(t, *mockConfig, accessToken)

	oauthData, err := Authorize(mockConfig)
	if err != nil {
		t.Fatalf("err should be nil, got: %s", err)
	}
	if oauthData.AccessToken != accessToken {
		t.Errorf("OAuth Access Token incorrect, got: %s, want: %s.", oauthData.AccessToken, accessToken)
	}

	// Check that the same token is returned when the the expiry time stamp is in the future
	mockConfig = buildMockConfig(mockConfig.ProfileName(), mockConfig.Environment(), mockConfig.ClientID(), mockConfig.ClientSecret(), oauthData.String())
	oauthData, err = Authorize(mockConfig)
	if err != nil {
		t.Fatalf("err should be nil, got: %s", err)
	}
	if oauthData.AccessToken != accessToken {
		t.Errorf("OAuth Access Token incorrect, got: %s, want: %s.", oauthData.AccessToken, accessToken)
	}

	// Check that a new token is retrieved when the the expiry time stamp is in the past
	oauthData.OAuthTokenExpiry = time.Now().AddDate(0, 0, -1).Format(time.RFC3339)
	accessToken = "aJdvugb8k1kwnOovm2qX6LXTctJksYvdzcoXPrRDi-nL1phQhcKRN-bjcflq7CUDOmUCQv5OWuBSkPQr0peWhw"
	mockConfig = buildMockConfig(mockConfig.ProfileName(), mockConfig.Environment(), mockConfig.ClientID(), mockConfig.ClientSecret(), oauthData.String())
	oauthData, err = Authorize(mockConfig)
	if err != nil {
		t.Fatalf("err should be nil, got: %s", err)
	}
	if oauthData.AccessToken != accessToken {
		t.Errorf("OAuth Access Token incorrect, got: %s, want: %s.", oauthData.AccessToken, accessToken)
	}
}

func TestLowLevelRestClient(t *testing.T) {
	tests := buildTestCaseTable()

	for _, tc := range tests {
		restClient := &RESTClient{
			environment: tc.targetEnv,
			token:       tc.oAuthToken,
		}

		setRestClientDoMock(t, tc)

		//First checking the low level API call
		results, err := restClient.callAPI(tc.targetVerb, tc.targetPath, "")

		//Testing to see if we got an http status code error.  If we got an http status code error we should get an HttpStatusError struct and the status code should match the status code on the response
		//Later on checks like this will be important when we add
		if err != nil {
			//Check to see if its an HTTP error and if its check to see if its what we are expecting
			if e, ok := err.(*models.HttpStatusError); ok && e.StatusCode != tc.expectedStatusCode {
				t.Errorf("Did not get the right HttpStatus Code for the error expected, got: %d, want: %d.", e.StatusCode, tc.expectedStatusCode)
			}
		} else {
			if tc.expectedResponse != results {
				t.Errorf("Retrieved the incorrect response calling the RestClient.Get, got: %s, want: %s.", results, tc.expectedResponse)
			}
		}
	}
}

func TestHighLevelRestClient(t *testing.T) {
	tests := buildTestCaseTable()

	for _, tc := range tests {
		restClient := &RESTClient{
			environment: tc.targetEnv,
			token:       tc.oAuthToken,
		}

		setRestClientDoMock(t, tc)
		var results string
		var err error

		//Calling the higher levl API functions
		switch tc.targetVerb {
		case http.MethodGet:
			results, err = restClient.Get(tc.targetPath)
		case http.MethodPost:
			results, err = restClient.Post(tc.targetPath, tc.targetBody)
		case http.MethodPut:
			results, err = restClient.Put(tc.targetPath, tc.targetBody)
		case http.MethodPatch:
			results, err = restClient.Patch(tc.targetPath, tc.targetBody)
		case http.MethodDelete:
			results, err = restClient.Delete(tc.targetPath)
		}

		//Rechecking the error codes for the higher level RestClient functions
		if err != nil {
			//Check to see if its an HTTP error and if its check to see if its what we are expecting
			if e, ok := err.(*models.HttpStatusError); ok && e.StatusCode != tc.expectedStatusCode {
				t.Errorf("Did not get the right HttpStatus Code for the error expected, got: %d, want: %d.", e.StatusCode, tc.expectedStatusCode)
			}
		} else {
			if tc.expectedResponse != results {
				t.Errorf("Retrieved the incorrect response calling the RestClient.Get, got: %s, want: %s.", results, tc.expectedResponse)
			}
		}
	}
}

func buildMockConfig(profileName string, environment string, clientID string, clientSecret string, oauthTokenData string) *mocks.MockClientConfig {
	mockConfig := &mocks.MockClientConfig{}

	mockConfig.ProfileNameFunc = func() string {
		return profileName
	}

	mockConfig.EnvironmentFunc = func() string {
		return environment
	}

	mockConfig.LogFilePathFunc = func() string {
		return ""
	}

	mockConfig.LoggingEnabledFunc = func() bool {
		return false
	}

	mockConfig.AutoPaginationEnabledFunc = func() bool {
		return false
	}

	mockConfig.ClientIDFunc = func() string {
		return clientID
	}

	mockConfig.ClientSecretFunc = func() string {
		return clientSecret
	}

	mockConfig.OAuthTokenDataFunc = func() string {
		return oauthTokenData
	}

	return mockConfig
}

func setRestClientDoMockForAuthorize(t *testing.T, mockConfig mocks.MockClientConfig, accessToken string) {
	ClientDo = func(request *retryablehttp.Request) (*http.Response, error) {
		authHeaderString := fmt.Sprintf("%s:%s", mockConfig.ClientID(), mockConfig.ClientSecret())
		expectedAuthHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(authHeaderString)))
		requestAuthHeader := request.Header.Get("Authorization")

		//Checking to make sure the auth header is built correctly
		if requestAuthHeader != expectedAuthHeader {
			t.Errorf("Authorization token on header not set propertly got: %s, want: %s.", requestAuthHeader, expectedAuthHeader)
		}

		//Checking to make sure the URL is built correctly
		expectedHost := fmt.Sprintf("login.%s", mockConfig.Environment())
		expectedPath := "/oauth/token"
		urlHost := request.URL.Host
		urlPath := request.URL.Path

		if expectedHost != urlHost {
			t.Errorf("Target oauth host is not correct : %s, want: %s.", expectedHost, urlHost)
		}

		if expectedPath != urlPath {
			t.Errorf("Target oauth path is not correct : %s, want: %s.", expectedPath, urlPath)
		}

		responseString := fmt.Sprintf(`
		  {
				"access_token": "%s",
				"token_type": "Bearer", 
				"expires_in": 1234,
				"error":      ""
			}
		`, accessToken)

		stringReader := strings.NewReader(responseString)
		stringReadCloser := ioutil.NopCloser(stringReader)
		response := &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       stringReadCloser,
		}
		return response, nil
	}
}

func setRestClientDoMock(t *testing.T, tc apiClientTest) {
	//Building a Mock HTTP Functions
	ClientDo = func(request *retryablehttp.Request) (*http.Response, error) {
		urlHost := request.URL.Host
		urlPath := request.URL.Path

		//Testing to see if the host and the path are being set correctly in my rest client code.
		if tc.expectedHost != urlHost {
			t.Errorf("Target host is not correct : %s, want: %s.", tc.expectedHost, urlHost)
		}

		if tc.expectedPath != urlPath {
			t.Errorf("Target path is not correct : %s, want: %s.", tc.expectedPath, urlPath)
		}

		//Check the auth header on the request to make sure is set
		if request.Header.Get("Authorization:") != tc.expectedAuthHeader {
			t.Errorf("Auth header on requests is not set properly : %s, want: %s.", tc.expectedAuthHeader, request.Header.Get("Authorization:")) //The extra colon at the head of the header string is not incorrect.  Go adds this
		}

		//Setting up the response body
		stringReader := strings.NewReader(tc.expectedResponse)
		stringReadCloser := ioutil.NopCloser(stringReader)
		response := &http.Response{
			StatusCode: tc.targetStatusCode,
			Body:       stringReadCloser,
		}
		return response, nil
	}
}

func buildTestCaseTable() []apiClientTest {
	oAuthToken := utils.GenerateGuid()
	skillId := utils.GenerateGuid()
	tests := []apiClientTest{
		{oAuthToken: oAuthToken, targetVerb: http.MethodGet, targetEnv: "mypurecloud.com", targetPath: fmt.Sprintf("api/v2/routing/skills/%s", skillId), targetStatusCode: http.StatusOK, targetBody: "", expectedHeaderToken: fmt.Sprintf("Bearer %s", oAuthToken), expectedHost: "api.mypurecloud.com", expectedPath: fmt.Sprintf("/api/v2/routing/skills/%s", skillId),
			expectedResponse: fmt.Sprintf(`{"id": "%s","name":"test2","dateModified": "2016-07-27T15:57:58Z","state": "active","version": "1","selfUri": "/api/v2/routing/skills/7eba08c7-e62b-49bb-867f-ca69bbab66f0"}`, skillId)},
		{oAuthToken: oAuthToken, targetVerb: http.MethodGet, targetEnv: "mypurecloud.de", targetPath: fmt.Sprintf("api/v2/routing/skills/%s", skillId), targetStatusCode: http.StatusOK, targetBody: "", expectedHeaderToken: fmt.Sprintf("Bearer %s", oAuthToken), expectedHost: "api.mypurecloud.de", expectedPath: fmt.Sprintf("/api/v2/routing/skills/%s", skillId),
			expectedResponse: fmt.Sprintf(`{"id": "%s","name":"test2","dateModified": "2016-07-27T15:57:58Z","state": "active","version": "1","selfUri": "/api/v2/routing/skills/7eba08c7-e62b-49bb-867f-ca69bbab66f0"}`, skillId)},
		{oAuthToken: oAuthToken, targetVerb: http.MethodPost, targetEnv: "mypurecloud.com", targetPath: fmt.Sprintf("api/v2/routing/skills/%s", skillId), targetStatusCode: http.StatusOK, targetBody: `{"id": "%s","name":"test2","state": "active""}`, expectedHeaderToken: fmt.Sprintf("Bearer %s", oAuthToken), expectedHost: "api.mypurecloud.com", expectedPath: fmt.Sprintf("/api/v2/routing/skills/%s", skillId),
			expectedResponse: fmt.Sprintf(`{"id": "%s","name":"test2","dateCreated": "2016-07-27T15:57:58Z","state": "active","version": "1","selfUri": "/api/v2/routing/skills/7eba08c7-e62b-49bb-867f-ca69bbab66f0"}`, skillId)},
		{oAuthToken: oAuthToken, targetVerb: http.MethodPut, targetEnv: "mypurecloud.com", targetPath: fmt.Sprintf("api/v2/routing/skills/%s", skillId), targetStatusCode: http.StatusOK, targetBody: `{"id": "%s","name":"test2","state": "active""}`, expectedHeaderToken: fmt.Sprintf("Bearer %s", oAuthToken), expectedHost: "api.mypurecloud.com", expectedPath: fmt.Sprintf("/api/v2/routing/skills/%s", skillId),
			expectedResponse: fmt.Sprintf(`{"id": "%s","name":"test2","dateCreated": "2016-07-27T15:57:58Z","state": "active","version": "3","selfUri": "/api/v2/routing/skills/7eba08c7-e62b-49bb-867f-ca69bbab66f0"}`, skillId)},
		{oAuthToken: oAuthToken, targetVerb: http.MethodPatch, targetEnv: "mypurecloud.com", targetPath: fmt.Sprintf("api/v2/routing/skills/%s", skillId), targetStatusCode: http.StatusOK, targetBody: `{"name":"test4"}`, expectedHeaderToken: fmt.Sprintf("Bearer %s", oAuthToken), expectedHost: "api.mypurecloud.com", expectedPath: fmt.Sprintf("/api/v2/routing/skills/%s", skillId),
			expectedResponse: fmt.Sprintf(`{"id": "%s","name":"test2","dateCreated": "2016-07-27T15:57:58Z","state": "active","version": "4","selfUri": "/api/v2/routing/skills/7eba08c7-e62b-49bb-867f-ca69bbab66f0"}`, skillId)},
		{oAuthToken: oAuthToken, targetVerb: http.MethodDelete, targetEnv: "mypurecloud.com", targetPath: fmt.Sprintf("api/v2/routing/skills/%s", skillId), targetStatusCode: http.StatusOK, targetBody: "", expectedHeaderToken: fmt.Sprintf("Bearer %s", oAuthToken), expectedHost: "api.mypurecloud.com", expectedPath: fmt.Sprintf("/api/v2/routing/skills/%s", skillId),
			expectedResponse: ""},
		//Error condition - Check to make sure we are return the httpStatusCode when there is an error
		{oAuthToken: oAuthToken, targetVerb: http.MethodGet, targetEnv: "mypurecloud.com", targetPath: fmt.Sprintf("api/v2/routing/skills/%s", skillId), targetStatusCode: http.StatusInternalServerError, targetBody: "", expectedHeaderToken: fmt.Sprintf("Bearer %s", oAuthToken), expectedHost: "api.mypurecloud.com", expectedPath: fmt.Sprintf("/api/v2/routing/skills/%s", skillId),
			expectedResponse: fmt.Sprintf(`{"id": "%s","name":"test2","dateModified": "2016-07-27T15:57:58Z","state": "active","version": "1","selfUri": "/api/v2/routing/skills/7eba08c7-e62b-49bb-867f-ca69bbab66f0"}`, skillId), expectedStatusCode: http.StatusInternalServerError},
	}

	return tests
}
