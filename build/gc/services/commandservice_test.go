package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/mocks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/restclient"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
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
	configGetConfig = mockGetConfig
	restclient.UpdateOAuthToken = mocks.UpdateOAuthToken

	c := commandService{
		cmd: &cobra.Command{},
	}

	// Check that RESTClient DoFunc is called 5 times
	maxRetriesBeforeQuitting := 5
	headers := make(map[string][]string)
	headers["Retry-After"] = []string{"10"}

	tc := apiClientTest{
		targetHeaders:      headers,
		targetStatusCode:   http.StatusTooManyRequests,
		expectedResponse:   fmt.Sprintf(`{"numRetries":"%v"}`, maxRetriesBeforeQuitting),
		expectedStatusCode: http.StatusTooManyRequests}
	setRestClientDoMockForRetry(tc, 10)

	retryFunc := retry.RetryWithData(tc.targetPath, []string{""}, c.Patch)
	retryConfig := &retry.RetryConfiguration{
		RetryWaitMax: 100 * time.Second,
		RetryMax:     maxRetriesBeforeQuitting,
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
	maxRetryTimeSec := 1 * time.Second
	tc.targetHeaders["Retry-After"] = []string{"2000"}
	expectedNumCalls := 1

	tc.expectedResponse = fmt.Sprintf(`{"numRetries":"%v"}`, expectedNumCalls)
	setRestClientDoMockForRetry(tc, 10)

	retryFunc = retry.RetryWithData(tc.targetPath, []string{""}, c.Patch)
	retryConfig.RetryWaitMax = maxRetryTimeSec
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
	tc.targetHeaders["Retry-After"] = []string{"10"}
	expectedNumCalls = 4

	tc.expectedResponse = fmt.Sprintf(`{"numRetries":"%v"}`, expectedNumCalls)
	setRestClientDoMockForRetry(tc, 3)

	retryFunc = retry.RetryWithData(tc.targetPath, []string{""}, c.Patch)
	retryConfig.RetryMax = maxRetriesBeforeQuitting
	results, err := retryFunc(retryConfig)
	if err != nil {
		t.Errorf("Error should not be nil, got: %s", err)
	}
	results = strings.ReplaceAll(strings.ReplaceAll(results, "\n", ""), " ", "")
	if results != tc.expectedResponse {
		t.Errorf("Did not get the expected results, got: %s, want: %s.", results, tc.expectedResponse)
	}
}

func TestReAuthenticationWithAccessToken(t *testing.T) {
	restclient.OverridesApplied = mocks.OverridesApplied
	configGetConfig = mockGetConfigWithAccessToken

	c := commandService{
		cmd: &cobra.Command{},
	}

	tc := apiClientTest{
		targetStatusCode:   http.StatusUnauthorized,
		expectedStatusCode: http.StatusUnauthorized,
		expectedResponse:   "",
	}
	setRestClientDoMockForReAuthenticate(tc)

	// the expected err from this GET request, when we have an access token in config, is the error msg below, and we do not care about the empty string returned
	_, err := c.Get("")

	expectedErr := "unauthorized. your access_token has either expired or is not valid. please authenticate"

	if err.Error() != expectedErr {
		t.Errorf("Did not get the right value, got: %s, want: %s.", err.Error(), expectedErr)
	}
}

func TestReAuthentication(t *testing.T) {
	restclient.OverridesApplied = mocks.OverridesApplied
	restclient.UpdateOAuthToken = mocks.UpdateOAuthToken
	configGetConfig = mockGetConfig

	c := commandService{
		cmd: &cobra.Command{},
	}

	tc := apiClientTest{
		targetStatusCode:    http.StatusUnauthorized,
		expectedStatusCode:  http.StatusUnauthorized,
		expectedResponse:    `{"numCalls": 2}`,
		expectedAccessToken: "c-Iyx66JoLCLVTkmQMXx-luHI3wpm-MQI1THRftzXEhgS4pNtBOaxfCzDSFw25LhcFZ3UjiczIlXVcwfoYvxfw",
	}
	setRestClientDoMockForReAuthenticate(tc)

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

//setRestClientDoMockForRetry sets the restclient.ClientDo method for the commandservice Retry test
func setRestClientDoMockForRetry(tc apiClientTest, numberOfFailedCalls int) {
	numCalls := 0

	//Building a Mock HTTP Functions
	restclient.ClientDo = func(request *retryablehttp.Request) (*http.Response, error) {
		//Setting up the response body
		response := &http.Response{
			Header:     tc.targetHeaders,
			StatusCode: tc.targetStatusCode,
		}

		startTime := time.Now()
		retryAfterValue := int64(0)
		for _, retryAfter := range tc.targetHeaders["Retry-After"] {
			if retryAfter != "" {
				retryAfterValue, _ = strconv.ParseInt(retryAfter, 10, 64)
				break
			}
		}

		for i := 0; i < restclient.Client.RetryMax; i++ {
			numCalls++

			stringReader := strings.NewReader(fmt.Sprintf(`{"numRetries": "%v"}`, numCalls))
			stringReadCloser := ioutil.NopCloser(stringReader)
			response.Body = stringReadCloser
			if numCalls > numberOfFailedCalls {
				response.StatusCode = http.StatusOK
				break
			}

			if retryAfterValue > 0 {
				fmt.Println("sleeping for", time.Duration(time.Duration(retryAfterValue)*time.Millisecond))
				time.Sleep(time.Duration(time.Duration(retryAfterValue) * time.Millisecond))
			}
			if time.Now().Sub(startTime) > restclient.Client.RetryWaitMax {
				break
			}
		}

		return response, nil
	}
}

//setRestClientDoMockForReAuthenticate sets the restclient.ClientDo method for the commandservice ReAuthenticate test
func setRestClientDoMockForReAuthenticate(tc apiClientTest) {
	numCalls := 0

	//Building a Mock HTTP Functions
	restclient.ClientDo = func(request *retryablehttp.Request) (*http.Response, error) {
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
}

//mockNewRESTClient returns a mock RESTClient object for the commandservice tests and sets the object in the restclient package
func mockNewRESTClient(_ config.Configuration) *restclient.RESTClient {
	restclient.RestClient = &restclient.RESTClient{}
	return restclient.RestClient
}

//mockGetConfig returns a mock MockClientConfig object with client credentials
func mockGetConfig(profileName string) (config.Configuration, error) {
	mockConfig := &mocks.MockClientConfig{}

	mockConfig.ProfileNameFunc = func() string {
		return profileName
	}

	mockConfig.EnvironmentFunc = func() string {
		return "mypurecloud.com"
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
		return utils.GenerateGuid()
	}

	mockConfig.ClientSecretFunc = func() string {
		return utils.GenerateGuid()
	}

	mockConfig.OAuthTokenDataFunc = func() string {
		return ""
	}

	mockConfig.AccessTokenFunc = func() string {
		return ""
	}

	return mockConfig, nil
}

//mockGetConfigWithAccessToken returns a mock MockClientConfig object with an access token
func mockGetConfigWithAccessToken(profileName string) (config.Configuration, error) {
	mockConfig := &mocks.MockClientConfig{}

	mockConfig.ProfileNameFunc = func() string {
		return profileName
	}

	mockConfig.EnvironmentFunc = func() string {
		return "mypurecloud.com"
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
		return ""
	}

	mockConfig.ClientSecretFunc = func() string {
		return ""
	}

	mockConfig.OAuthTokenDataFunc = func() string {
		return ""
	}

	mockConfig.AccessTokenFunc = func() string {
		return "XNiJQrSf2YQmJODySCxG6HaVIE2lfZfJ35Y4JDh5L9YEBJOTG3p6szRyUvWVM7pDmziPHHcq9NW7e0KxN_lb6w" // this is a "bad" token for testing purposes
	}

	return mockConfig, nil
}
