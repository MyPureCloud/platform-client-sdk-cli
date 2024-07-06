package restclient

import (
        "context"
        cryptoTls "crypto/tls"
        "errors"
        "github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
        "github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
        "github.com/mypurecloud/platform-client-sdk-cli/build/gc/restclient/tls"
        "github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
        "github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
        "log"
        "net/http"

        "os"
        "os/exec"
        "regexp"
        "runtime"
        "strconv"
        "time"

        "bytes"
        "encoding/base64"
        "encoding/json"
        "fmt"
        "github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
        "io/ioutil"
        "net/url"
        "strings"
        "path/filepath"

        "github.com/hashicorp/go-retryablehttp"
        "github.com/tidwall/pretty"
)

var (
        Client              retryablehttp.Client
        ClientDo            = Client.Do
        RestClient          *RESTClient
        UpdateOAuthToken    = config.UpdateOAuthToken
        OverridesApplied    = config.OverridesApplied
        openBrowserForLogin = openBrowserForLoginFunc
        startLocalServer    = startLocalServerFunc
        redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)
        schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)
        invalidHeaderErrorRe = regexp.MustCompile(`invalid header`)
        notTrustedErrorRe = regexp.MustCompile(`certificate is not trusted`)
)

type RESTClient struct {
        environment   string
        token         string
        configuration config.Configuration
}

func (r *RESTClient) SetEnv(env string) {
        r.environment = env
}

func (r *RESTClient) SetConfig(c config.Configuration) {
        r.configuration = c
}

// Get executes an HTTP Get
func (r *RESTClient) Get(uri string) (string, error) {
        return r.callAPI(http.MethodGet, uri, "")
}

// Head executes an HTTP Head
func (r *RESTClient) Head(uri string) (string, error) {
        return r.callAPI(http.MethodHead, uri, "")
}

// Put executes an HTTP Put
func (r *RESTClient) Put(uri string, data string) (string, error) {
        return r.callAPI(http.MethodPut, uri, data)
}

// Post executes an HTTP Post
func (r *RESTClient) Post(uri string, data string) (string, error) {
        return r.callAPI(http.MethodPost, uri, data)
}

// Patch executes an HTTP Patch
func (r *RESTClient) Patch(uri string, data string) (string, error) {
        return r.callAPI(http.MethodPatch, uri, data)
}

// Delete executes an HTTP Delete
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
        request.Header.Set("purecloud-sdk", "104.0.0")

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
                                if retryNumber > 1 {
                                        logger.Warnf("%v %v request failed. Retry count: %v\n", req.Method, req.URL, retryNumber)
                                }
                        }
                } else {
                        Client.RequestLogHook = func(_ retryablehttp.Logger, req *http.Request, retryNumber int) {
                                retryConfiguration.RequestLogHook(req, retryNumber)
                        }
                }
        }

        Client.CheckRetry = DefaultRetryPolicy

        setProxyConf(r.configuration, "other")
        //Executing the request
        resp, err := ClientDo(request)
        if err != nil {
                return "", err
        }
        timestampMS := time.Now().UnixNano() / int64(time.Millisecond)
        defer resp.Body.Close()

        correlationId := resp.Header["Inin-Correlation-Id"]
        if correlationId != nil {
                logger.Infof("API response Correlation ID: %v, method: %v, URI: %v, timestamp(ms): %d", correlationId[0], method, uri, timestampMS)
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

func DefaultRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// don't propagate other errors
	shouldRetry, _ := newRetryPolicy(resp, err)
	return shouldRetry, nil
}



func newRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		if v, ok := err.(*url.Error); ok {
			// Don't retry if the error was due to too many redirects.
			if redirectsErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to an invalid protocol scheme.
			if schemeErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to an invalid header.
			if invalidHeaderErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to TLS cert verification failure.
			if notTrustedErrorRe.MatchString(v.Error()) {
				return false, v
			}
			if isCertError(v.Err) {
				return false, v
			}
		}

		// The error is likely recoverable so retry.
		return true, nil
	}

	// 429 Too Many Requests is recoverable. Sometimes the server puts
	// a Retry-After response header to indicate when the server is
	// available to start processing request from client.
	if resp.StatusCode == http.StatusTooManyRequests {
	    return verifyRetryAfterHeader(resp.Header["Retry-After"], 180), nil
	}

	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented) {
		return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	return false, nil
}

func isCertError(err error) bool {
	_, ok := err.(*cryptoTls.CertificateVerificationError)
	return ok
}

func verifyRetryAfterHeader(headers []string, defaultMaxRetry int64) bool {
	if len(headers) == 0 || headers[0] == "" {
		return true
	}
	header := headers[0]

	if sleep, err := strconv.ParseInt(header, 10, 64); err == nil {
		if sleep > defaultMaxRetry {
			return  false
		}
		return true
	}
	return true
}

// Authorize authenticates the user using the client credentials in their profile.
func Authorize(c config.Configuration) (models.OAuthTokenData, error) {
        if strings.Contains(c.Environment(), "localhost") {
                return models.OAuthTokenData{}, nil
        }

        oAuthTokenData := &models.OAuthTokenData{}
        if !OverridesApplied() && c.OAuthTokenData() != "" {
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

// ReAuthenticate re-authenticates the user using the client credentials in their profile.
func ReAuthenticate(c config.Configuration) (models.OAuthTokenData, error) {
        oAuthToken, err := authorize(c)
        if err == nil {
                RestClient.token = oAuthToken.AccessToken
        }

        return oAuthToken, err
}

func authorizePKCEGrant(c config.Configuration, code string, codeVerifier string) (models.OAuthTokenData, error) {
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
        request.Header.Set("purecloud-sdk", "104.0.0")

        //Setting up the form data
        form := url.Values{}
        form["grant_type"] = []string{"authorization_code"}
        form["client_id"] = []string{c.ClientID()}
        form["code"] = []string{code}
        form["redirect_uri"] = []string{c.RedirectURI()}
        form["code_verifier"] = []string{codeVerifier}
        request.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
        setProxyConf(c, "login")

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
        return createOAuthTokenResponse(c, *oAuthToken)
}

func authorize(c config.Configuration) (models.OAuthTokenData, error) {
        authChannel := make(chan models.OAuthToken)
        // Using implicit grant or pkce grant
        if c.GrantType() == "2" || c.GrantType() == "3" || c.RedirectURI() != "" {
                if c.SecureLoginEnabled() {
                        // Generate a self-signed X.509 certificate for a TLS server
                        err := tls.GenerateCerts()
                        if err != nil {
                                oAuthTokenResponse := &models.OAuthTokenData{}
                                return *oAuthTokenResponse, err
                        }
                }
                go startLocalServer(c, authChannel)
                openBrowserForLogin(c.RedirectURI())

                oAuthToken := <-authChannel
                if oAuthToken.Error != "" {
                        oAuthTokenResponse := &models.OAuthTokenData{}
                        return *oAuthTokenResponse, errors.New(oAuthToken.Error)
                }

                return createOAuthTokenResponse(c, oAuthToken)
        }

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
        request.Header.Set("purecloud-sdk", "104.0.0")

        //Setting up the form data
        form := url.Values{}
        form["grant_type"] = []string{"client_credentials"}
        request.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
        setProxyConf(c, "login")

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
        return createOAuthTokenResponse(c, *oAuthToken)
}

func setProxyConf(c config.Configuration, path string) {


            var proxyConfiguration config.ProxyConfiguration
            err := json.Unmarshal([]byte(c.ProxyConfiguration()), &proxyConfiguration)
            if err !=nil {
                logger.Warnf("Error parsing proxy URL: %v , %v", err, c.ProxyConfiguration())
            }

                if proxyConfiguration.Protocol != "" {
                        var proxyUrl *url.URL
                        pathValue := ""
                        proxyConfiguration.Host = proxyConfiguration.Host + ":" + proxyConfiguration.Port
                        if len(proxyConfiguration.PathParams) > 0 {
                           params := proxyConfiguration.PathParams

                           if tempValue, exists := params[path]; exists {
                                pathValue = tempValue
                            }

                        }

                        if proxyConfiguration.UserName != "" && proxyConfiguration.Password != "" {
                                proxyUrl = &url.URL{
                                        Scheme: proxyConfiguration.Protocol,
                                        User: url.UserPassword(proxyConfiguration.UserName,
                                                proxyConfiguration.Password),
                                        Host: proxyConfiguration.Host,
                                }
                        } else {
                                urlString := proxyConfiguration.Protocol + "://" +
                                        proxyConfiguration.Host
                                proxyUrl, _ = url.Parse(urlString)
                        }

                        if pathValue !="" {
                          proxyUrl.Path = pathValue
                        }

                        tr := &http.Transport{
                                Proxy: http.ProxyURL(proxyUrl),
                        }

                        Client.HTTPClient.Transport = tr
                }
}

func createOAuthTokenResponse(c config.Configuration, oAuthToken models.OAuthToken) (models.OAuthTokenData, error) {
        oAuthTokenResponse := &models.OAuthTokenData{}
        oAuthTokenExpiry := time.Now().Add(utils.SecondsToNanoSeconds(oAuthToken.ExpiresIn))

        oAuthTokenResponse = &models.OAuthTokenData{
                OAuthToken:       oAuthToken,
                OAuthTokenExpiry: oAuthTokenExpiry.Format(time.RFC3339),
        }

        if !OverridesApplied() {
                err := UpdateOAuthToken(c, oAuthTokenResponse)
                if err != nil {
                        return *oAuthTokenResponse, err
                }
        }

        return *oAuthTokenResponse, nil
}

// Starts the local server at the redirect URI
func startLocalServerFunc(c config.Configuration, authChannel chan models.OAuthToken) {
        u, err := url.Parse(c.RedirectURI())
        if err != nil {
                logger.Fatalf("Error parsing redirect URL: %v", err)
        }
        path := u.Path
        if path == "" {
                path = "/"
        }

        var oAuthToken models.OAuthToken
        var initialPage string
        oauthWebpage := models.OAuthWebpage{}
        codeVerifier, _ := utils.GeneratePKCECodeVerifier(128)
        if c.GrantType() == "3" {
                // Use PKCE Grant
                codeChallenge, _ := utils.ComputePKCECodeChallenge(codeVerifier)
                // Set initial web page as one containing the javascript to perform implicit login
                initialPage = oauthWebpage.GetPKCEWebpage(c.ClientID(), c.Environment(), c.RedirectURI(), codeChallenge)
        } else {
                // Otherwise use Implicit Grant (== "2") - no GrantType test for backward compatibility
                // Set initial web page as one containing the javascript to perform implicit login
                initialPage = oauthWebpage.GetImplicitWebpage(c.ClientID(), c.Environment(), c.RedirectURI())
        }
        activePage := initialPage

        http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
                _, err := fmt.Fprintf(w, activePage)
                if err != nil {
                        logger.Fatalf("Error handling request: %v", err)
                }

                // If no longer displaying initial page -
                // Give time for initialPage to refresh on the browser thereby printing the response page,
                // Then close channel & server
                if activePage != initialPage {
                        time.Sleep(4 * time.Second)
                        authChannel <- oAuthToken
                        close(authChannel)
                        return
                }

                if strings.Contains(r.URL.String(), "/error/") || strings.Contains(r.URL.String(), "/access_token/") || strings.Contains(r.URL.String(), "/code/") {
                        oAuthToken = getOAuthResponseDataFromURL(c, r.URL.String(), codeVerifier)
                        activePage = oauthWebpage.GetResponsePage(oAuthToken.Error)
                }
        })

        if c.SecureLoginEnabled() {
                log.SetOutput(ioutil.Discard)
                log.Fatal(http.ListenAndServeTLS(":"+u.Port(), filepath.Join(os.TempDir(), "cert.pem"), filepath.Join(os.TempDir(), "key.pem"), nil))
        } else {
                log.Fatal(http.ListenAndServe(":"+u.Port(), nil))
        }
}

// parse data from url in the format '/code/access_token/12345/expires_in/299/token_type/bearer'
func getOAuthResponseDataFromURL(c config.Configuration, url string, codeVerifier string) models.OAuthToken {
        var response models.OAuthToken
        splitURL := strings.Split(url, "/")
        for i, v := range splitURL {
                if v == "error" {
                        response.Error = splitURL[i+1]
                        break
                }
                if v == "access_token" {
                        response.AccessToken = splitURL[i+1]
                } else if v == "expires_in" {
                        response.ExpiresIn, _ = strconv.Atoi(splitURL[i+1])
                } else if v == "token_type" {
                        response.TokenType = splitURL[i+1]
                } else if v == "code" {
                        code := splitURL[i+1]
                        oAuthTokenData, err := authorizePKCEGrant(c, code, codeVerifier)
                        if err != nil {
                                response.Error = "PKCEGrantError"
                                break
                        }
                        response.AccessToken = oAuthTokenData.AccessToken
                        response.TokenType = oAuthTokenData.TokenType
                        response.ExpiresIn = oAuthTokenData.ExpiresIn
                        break
                }
        }
        return response
}

func openBrowserForLoginFunc(loginURL string) {
        var err error
        switch runtime.GOOS {
        case "linux":
                err = exec.Command("xdg-open", loginURL).Start()
        case "windows":
                err = exec.Command("rundll32", "url.dll,FileProtocolHandler", loginURL).Start()
        case "darwin":
                err = exec.Command("open", loginURL).Start()
        default:
                err = fmt.Errorf("unsupported platform '%s'", runtime.GOOS)
        }
        if err != nil {
                logger.Fatal(err)
        }
}

// NewRESTClient is a constructor function to build an APIClient
func NewRESTClient(config config.Configuration) *RESTClient {
        if RestClient == nil && config.AccessToken() == "" {
                oAuthToken, err := Authorize(config)
                if err != nil {
                        logger.Fatal(err)
                }

                RestClient = &RESTClient{environment: config.Environment(), token: oAuthToken.AccessToken, configuration: config}
                return RestClient
        }

        if RestClient == nil && config.AccessToken() != "" {
                RestClient = &RESTClient{environment: config.Environment(), token: config.AccessToken(), configuration: config}
        }

        return RestClient
}

func init() {
        Client = *retryablehttp.NewClient()
        Client.Logger = nil
        ClientDo = Client.Do
}
