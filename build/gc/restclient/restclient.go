package restclient

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"net/http"
	"time"

	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/tidwall/pretty"
)

var (
	Client             retryablehttp.Client
	ClientDo         = Client.Do
	RestClient         *RESTClient
	UpdateOAuthToken = config.UpdateOAuthToken
)

type RESTClient struct {
	environment string
	token       string
}

//Get executes an HTTP Get
func (r *RESTClient) Get(uri string) (string, error) {
	return r.callAPI(http.MethodGet, uri, "")
}

//Put executes an HTTP Put
func (r *RESTClient) Put(uri string, data string) (string, error) {
	return r.callAPI(http.MethodPut, uri, data)
}

//Post executes an HTTP Post
func (r *RESTClient) Post(uri string, data string) (string, error) {
	return r.callAPI(http.MethodPost, uri, data)
}

//Patch executes an HTTP Patch
func (r *RESTClient) Patch(uri string, data string) (string, error) {
	return r.callAPI(http.MethodPatch, uri, data)
}

//Delete executes an HTTP Delete
func (r *RESTClient) Delete(uri string) (string, error) {
	return r.callAPI(http.MethodDelete, uri, "")
}

func (r *RESTClient) callAPI(method string, uri string, data string) (string, error) {
	var apiURI *url.URL
	if !strings.HasPrefix(uri, "/") {
		uri = fmt.Sprintf("/%v", uri)
	}
	if !strings.Contains(r.environment, "localhost") {
		apiURI, _ = url.Parse(fmt.Sprintf("https://api.%s%s", r.environment, uri))
	} else {
		apiURI, _ = url.Parse(fmt.Sprintf("http://%s%s", r.environment, uri))
	}

	logger.Infof("Calling API with method: %v, URI: %v, data: %v", method, uri, data)

	request := &retryablehttp.Request{
		Request: &http.Request{
			URL:    apiURI,
			Close:  true,
			Method: strings.ToUpper(method),
			Header: make(map[string][]string),
		},
	}

	//Setting up the auth header
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Cache-Control", "no-cache")

	//User-Agent and SDK version headers
	request.Header.Set("User-Agent", "PureCloud SDK/go-cli")
	request.Header.Set("purecloud-sdk", "2.0.2")

	if data != "" {
		request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(data)))
	}

	retryConfiguration := retry.GetRetryConfiguration()
	if retryConfiguration == nil {
		Client.RetryMax = 0
		Client.RetryWaitMax = 0
	} else {
		Client.RetryWaitMax = retryConfiguration.RetryWaitMax
		Client.RetryWaitMin = retryConfiguration.RetryWaitMin
		Client.RetryMax = retryConfiguration.RetryMax
		if retryConfiguration.RequestLogHook == nil {
			Client.RequestLogHook = func(_ retryablehttp.Logger, req *http.Request, retryNumber int) {
				logger.Warnf("%v %v request failed. Retry count: %v\n", req.Method, req.URL, retryNumber)
			}
		} else {
			Client.RequestLogHook = func(_ retryablehttp.Logger, req *http.Request, retryNumber int) {
				retryConfiguration.RequestLogHook(req, retryNumber)
			}
		}
	}
	
	//Executing the request
	resp, err := ClientDo(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	correlationId := resp.Header["Inin-Correlation-Id"]
	if correlationId != nil {
		logger.Info("API response Correlation ID:", correlationId[0])
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responseData := string(response)

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		httpError := models.HttpStatusError{Verb: method, Path: uri, StatusCode: resp.StatusCode, Headers: resp.Header, Body: fmt.Sprintf("%s", pretty.Pretty([]byte(responseData)))}
		logger.Warn("Error from API:", httpError.ErrorDescriptive())
		return "", httpError
	}

	return responseData, nil
}

//Authorize authenticates the user using the client credentials in their profile.
func Authorize(c config.Configuration) (models.OAuthTokenData, error) {
	if strings.Contains(c.Environment(), "localhost") {
		return models.OAuthTokenData{}, nil
	}

	oAuthTokenData := &models.OAuthTokenData{}
	if c.OAuthTokenData() != "" {
		err := json.Unmarshal([]byte(c.OAuthTokenData()), oAuthTokenData)
		if err != nil {
			return models.OAuthTokenData{}, err
		}
		oauthTokenExpiry, _ := time.Parse(time.RFC3339, oAuthTokenData.OAuthTokenExpiry)
		if oauthTokenExpiry.After(time.Now()) {
			return *oAuthTokenData, nil
		}
		logger.Info("Authorizing because token has expired")
	}

	return authorize(c)
}

//ReAuthenticate re-authenticates the user using the client credentials in their profile.
func ReAuthenticate(c config.Configuration) (models.OAuthTokenData, error) {
	oAuthToken, err := authorize(c)
	if err == nil {
		RestClient.token = oAuthToken.AccessToken
	}

	return oAuthToken, err
}

func authorize(c config.Configuration) (models.OAuthTokenData, error) {
	loginURI, _ := url.Parse(fmt.Sprintf("https://login.%s/oauth/token", c.Environment()))

	request := &retryablehttp.Request{
		Request: &http.Request{
			URL:    loginURI,
			Close:  true,
			Method: http.MethodPost,
			Header: make(map[string][]string),
		},
	}

	//Setting up the basic auth headers for the call
	authHeaderString := fmt.Sprintf("%s:%s", c.ClientID(), c.ClientSecret())
	authHeader := base64.StdEncoding.EncodeToString([]byte(authHeaderString))
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", authHeader))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//User-Agent and SDK version headers
	request.Header.Set("User-Agent", "PureCloud SDK/go-cli")
	request.Header.Set("purecloud-sdk", "2.0.2")

	//Setting up the form data
	form := url.Values{}
	form["grant_type"] = []string{"client_credentials"}
	request.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))

	//Executing the request
	resp, err := ClientDo(request)
	if err != nil {
		logger.Fatal(err)
	}

	defer resp.Body.Close()

	oAuthTokenResponse := &models.OAuthTokenData{}
	// Read Response Body
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return *oAuthTokenResponse, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		httpError := models.HttpStatusError{Verb: http.MethodPost, Path: loginURI.Path, StatusCode: resp.StatusCode, Headers: resp.Header, Body: fmt.Sprintf("%s", pretty.Pretty(responseData))}
		return *oAuthTokenResponse, httpError
	}

	oAuthToken := &models.OAuthToken{}
	err = json.Unmarshal(responseData, &oAuthToken)
	if err != nil {
		return *oAuthTokenResponse, err
	}
	oAuthTokenExpiry := time.Now().Add(utils.SecondsToNanoSeconds(oAuthToken.ExpiresIn))
	oAuthTokenResponse = &models.OAuthTokenData{
		OAuthToken:       *oAuthToken,
		OAuthTokenExpiry: oAuthTokenExpiry.Format(time.RFC3339),
	}

	err = UpdateOAuthToken(c, oAuthTokenResponse)
	if err != nil {
		return *oAuthTokenResponse, err
	}

	return *oAuthTokenResponse, nil
}

//NewRESTClient is a constructor function to build an APIClient
func NewRESTClient(config config.Configuration) *RESTClient {
	if RestClient == nil {
		oAuthToken, err := Authorize(config)
		if err != nil {
			logger.Fatal(err)
		}

		RestClient = &RESTClient{environment: config.Environment(), token: oAuthToken.AccessToken}
	}

	return RestClient
}

func init() {
	Client = *retryablehttp.NewClient()
	Client.Logger = nil
	ClientDo = Client.Do
}
