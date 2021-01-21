package services

import (
	"fmt"
	"gc/config"
	"gc/mocks"
	"gc/models"
	"gc/restclient"
	"gc/retry"
	"gc/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type apiClientTest struct {
	targetPath          string
	targetStatusCode    int
	targetHeaders       map[string][]string
	expectedResponse    string
	expectedStatusCode  int
	expectedAccessToken string
}

func TestRetryWithData(t *testing.T) {
	restclientNewRESTClient = mockNewRESTClient
	configGetConfig = mockGetConfigRetry
	restclient.UpdateOAuthToken = mocks.UpdateOAuthToken

	c := commandService{
		cmd: &cobra.Command{},
	}

	// Check that RESTClient DoFunc is called 5 times
	maxRetriesBeforeQuitting := 5
	headers := make(map[string][]string)
	headers["Retry-After"] = []string{"10"}

	tc := apiClientTest{
		targetHeaders: headers,
		targetStatusCode: http.StatusTooManyRequests,
		expectedResponse: fmt.Sprintf(`{"numRetries":"%v"}`, maxRetriesBeforeQuitting - 1),
		expectedStatusCode: http.StatusTooManyRequests}
	restclient.Client = buildRestClientDoMockForRetry(tc, 10)

	retryFunc := retry.RetryWithData(tc.targetPath, "", c.Patch)
	retryConfig := &retry.RetryConfiguration{
		MaxRetryTimeSec:          100,
		MaxRetriesBeforeQuitting: maxRetriesBeforeQuitting,
	}
	_, err := retryFunc(retryConfig)
	if err != nil {
		//Check to see if its an HTTP error and if its check to see if its what we are expecting
		e, _ := err.(models.HttpStatusError)
		if e.StatusCode != tc.expectedStatusCode {
			t.Errorf("Did not get the right HttpStatus Code for the error expected, got: %d, want: %d.", e.StatusCode, tc.expectedStatusCode)
		}
		e.Body = strings.ReplaceAll(strings.ReplaceAll(e.Body, "\n", ""), " ", "")
		if e.Body != tc.expectedResponse {
			t.Errorf("Did not get the right Body for the error expected, got: %s, want: %s.", e.Body, tc.expectedResponse)
		}
	}

	// Check that RESTClient DoFunc is called 1 time when the retry-after value is increased and MaxRetryTimeSec is reduced below it
	maxRetryTimeSec := 1
	headers["Retry-After"] = []string{"2000"}
	expectedNumCalls := 1

	tc.expectedResponse = fmt.Sprintf(`{"numRetries":"%v"}`, expectedNumCalls)
	restclient.Client = buildRestClientDoMockForRetry(tc, 10)

	retryFunc = retry.RetryWithData(tc.targetPath, "", c.Patch)
	retryConfig.MaxRetryTimeSec = maxRetryTimeSec
	_, err = retryFunc(retryConfig)
	if err != nil {
		//Check to see if its an HTTP error and if its check to see if its what we are expecting
		e, _ := err.(models.HttpStatusError)
		if e.StatusCode != tc.expectedStatusCode {
			t.Errorf("Did not get the right HttpStatus Code for the error expected, got: %d, want: %d.", e.StatusCode, tc.expectedStatusCode)
		}
		e.Body = strings.ReplaceAll(strings.ReplaceAll(e.Body, "\n", ""), " ", "")
		if e.Body != tc.expectedResponse {
			t.Errorf("Did not get the right Body for the error expected, got: %s, want: %s.", e.Body, tc.expectedResponse)
		}
	}

	// Check that RESTClient DoFunc is called 4 times when it fails the first 3 times
	maxRetriesBeforeQuitting = 5
	headers["Retry-After"] = []string{"10"}
	expectedNumCalls = 4

	tc.expectedResponse = fmt.Sprintf(`{"numRetries":"%v"}`, expectedNumCalls)
	restclient.Client = buildRestClientDoMockForRetry(tc, 3)

	retryFunc = retry.RetryWithData(tc.targetPath, "", c.Patch)
	retryConfig.MaxRetryTimeSec = maxRetryTimeSec
	retryConfig.MaxRetriesBeforeQuitting = maxRetriesBeforeQuitting
	results, err := retryFunc(retryConfig)
	if err != nil {
		t.Errorf("Error should not be nil, got: %s", err)
	}
	results = strings.ReplaceAll(strings.ReplaceAll(results, "\n", ""), " ", "")
	if results != tc.expectedResponse {
		t.Errorf("Did not get the expected results, got: %s, want: %s.", results, tc.expectedResponse)
	}
}

func TestReAuthentication(t *testing.T) {
	restclientNewRESTClient = mockNewRESTClient
	restclient.UpdateOAuthToken = mocks.UpdateOAuthToken
	configGetConfig = mockGetConfigReAuthenticate

	c := commandService{
		cmd: &cobra.Command{},
	}

	tc := apiClientTest{
		targetStatusCode: http.StatusUnauthorized,
		expectedStatusCode: http.StatusUnauthorized,
		expectedResponse: `{"numCalls": 2}`,
		expectedAccessToken: "c-Iyx66JoLCLVTkmQMXx-luHI3wpm-MQI1THRftzXEhgS4pNtBOaxfCzDSFw25LhcFZ3UjiczIlXVcwfoYvxfw",
	}
	restclient.Client = buildRestClientDoMockForReAuthenticate(tc)

	value, err := c.Get("")
	if err != nil {
		t.Fatalf("err should be nil, got: %s", err)
	}
	if value != tc.expectedResponse {
		t.Errorf("Did not get the right value, got: %s, want: %s.", value, tc.expectedResponse)
	}
	if mocks.UpdatedAccessToken != tc.expectedAccessToken {
		t.Errorf("The access token was not updated as expected, got: %s, want: %s.", mocks.UpdatedAccessToken, tc.expectedAccessToken)
	}
}

//buildRestClientDoMockForRetry returns a mock HttpClient object for the commandservice Retry test
func buildRestClientDoMockForRetry(tc apiClientTest, numberOfFailedCalls int) *mocks.MockHttpClient {
	mock := &mocks.MockHttpClient{}
	numCalls := 0

	//Building a Mock HTTP Functions
	mock.DoFunc = func(request *http.Request) (*http.Response, error) {
		//Setting up the response body
		stringReader := strings.NewReader(fmt.Sprintf(`{"numRetries": "%v"}`, numCalls))
		stringReadCloser := ioutil.NopCloser(stringReader)

		response := &http.Response{
			Header:     tc.targetHeaders,
			StatusCode: tc.targetStatusCode,
			Body:       stringReadCloser,
		}
		if numCalls > numberOfFailedCalls {
			response.StatusCode = http.StatusOK
		}
		numCalls++

		return response, nil
	}

	return mock
}

//buildRestClientDoMockForReAuthenticate returns a mock HttpClient object for the commandservice ReAuthenticate test
func buildRestClientDoMockForReAuthenticate(tc apiClientTest) *mocks.MockHttpClient {
	mock := &mocks.MockHttpClient{}
	numCalls := 0

	//Building a Mock HTTP Functions
	mock.DoFunc = func(request *http.Request) (*http.Response, error) {
		//Setting up the response body
		responseString := ""
		statusCode := 0

		switch numCalls {
		case 0:
			// First call will return unauthorized for a GET request
			statusCode = tc.targetStatusCode
			responseString = ""
		case 1:
			// Second call will return a new access token for a login POST request
			statusCode = http.StatusOK
			responseString = fmt.Sprintf(`
				{
					"access_token": "%s",
					"token_type": "%s", 
					"expires_in": %d,
					"error":      ""
				}
				`, tc.expectedAccessToken, "Bearer", 1234)
		case 2:
			// Third call will return OK for the recursively called GET request following re-authentication
			statusCode = http.StatusOK
			responseString = fmt.Sprintf(`{"numCalls": %v}`, numCalls)
		}

		stringReader := strings.NewReader(responseString)
		stringReadCloser := ioutil.NopCloser(stringReader)

		response := &http.Response{
			Header:     tc.targetHeaders,
			StatusCode: statusCode,
			Body:       stringReadCloser,
		}

		numCalls++

		return response, nil
	}

	return mock
}

//mockNewRESTClient returns a mock RESTClient object for the commandservice tests and sets the object in the restclient package
func mockNewRESTClient(_ config.Configuration) *restclient.RESTClient {
	restclient.RestClient = &restclient.RESTClient{}
	return restclient.RestClient
}

//mockGetConfigRetry returns a mock MockClientConfig object for the commandservice Retry test
func mockGetConfigRetry(profileName string) (config.Configuration, error) {
	mockConfig := &mocks.MockClientConfig{}

	mockConfig.ProfileNameFunc = func() string {
		return profileName
	}

	mockConfig.EnvironmentFunc = func() string {
		return "mypurecloud.com"
	}

	mockConfig.ClientIDFunc = func() string {
		return utils.GenerateGuid()
	}

	mockConfig.ClientSecretFunc = func() string {
		return utils.GenerateGuid()
	}

	mockConfig.OAuthTokenDataFunc = func() string {
		return ""
	}

	return mockConfig, nil
}

//mockGetConfigReAuthenticate returns a mock MockClientConfig object for the commandservice ReAuthenticate test
func mockGetConfigReAuthenticate(profileName string) (config.Configuration, error) {
	mockConfig := &mocks.MockClientConfig{}

	mockConfig.ProfileNameFunc = func() string {
		return profileName
	}

	mockConfig.EnvironmentFunc = func() string {
		return "mypurecloud.com"
	}

	mockConfig.ClientIDFunc = func() string {
		return utils.GenerateGuid()
	}

	mockConfig.ClientSecretFunc = func() string {
		return utils.GenerateGuid()
	}

	mockConfig.OAuthTokenDataFunc = func() string {
		return ""
	}

	return mockConfig, nil
}