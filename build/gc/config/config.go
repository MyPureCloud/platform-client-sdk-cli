package config

import (
	"fmt"
	"os"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"

	"strings"

	"github.com/spf13/viper"

	"encoding/json"
)

type Configuration interface {
	ProfileName() string
	Environment() string
	GrantType() string
	ClientID() string
	ClientSecret() string
	RedirectURI() string
	OAuthTokenData() string
	AccessToken() string
	LogFilePath() string
	SecureLoginEnabled() bool
	LoggingEnabled() bool
	AutoPaginationEnabled() bool
	ProxyConfiguration() string
	GateWayConfiguration() string
	fmt.Stringer
}

type configuration struct {
	profileName           string
	environment           string
	grantType             string
	clientID              string
	clientSecret          string
	secureLoginEnabled    bool
	redirectURI           string
	oAuthTokenData        string
	accessToken           string
	logFilePath           string
	loggingEnabled        bool
	autoPaginationEnabled bool
	proxyConfiguration    string
	gatewayConfiguration  string
}

type GateWayConfiguration struct {
	Protocol   string
	Host       string
	Port       string
	UserName   string
	Password   string
	PathParams map[string]string
}

type ProxyConfiguration struct {
	Protocol   string
	Host       string
	Port       string
	UserName   string
	Password   string
	PathParams map[string]string
}

var (
	Environment    string
	ClientId       string
	ClientSecret   string
	AccessToken    string
	regionMappings = map[string]string{
		"us-east-1":      "mypurecloud.com",
		"eu-west-1":      "mypurecloud.ie",
		"ap-southeast-2": "mypurecloud.com.au",
		"ap-northeast-1": "mypurecloud.jp",
		"eu-central-1":   "mypurecloud.de",
		"us-west-2":      "usw2.pure.cloud",
		"ca-central-1":   "cac1.pure.cloud",
		"ap-northeast-2": "apne2.pure.cloud",
		"eu-west-2":      "euw2.pure.cloud",
		"ap-south-1":     "aps1.pure.cloud",
		"us-east-2":      "use2.us-gov-pure.cloud",
		"sa-east-1":      "sae1.pure.cloud",
		"me_central_1":   "mec1.pure.cloud",
		"ap_northeast_3": "apne3.pure.cloud",
		"eu_central_2":   "euc2.pure.cloud",
	}
)

// ProfileName is the name of the profile being used to run the CLI
func (c *configuration) ProfileName() string {
	return c.profileName
}

// GrantType is the OAuth grant type used by the OAuth Client
func (c *configuration) GrantType() string {
	return viper.GetString(fmt.Sprintf("%s.grant_type", c.profileName))
}

// ClientID is the OAuth client id used by the OAuth Client
func (c *configuration) ClientID() string {
	if ClientId != "" {
		return ClientId
	}

	return viper.GetString(fmt.Sprintf("%s.client_credentials", c.profileName))
}

// ClientSecret is the OAuth client secret used by the OAuth Client
func (c *configuration) ClientSecret() string {
	if ClientSecret != "" {
		return ClientSecret
	}

	return viper.GetString(fmt.Sprintf("%s.client_secret", c.profileName))
}

func (c *configuration) AccessToken() string {
	if AccessToken != "" {
		return AccessToken
	}

	return viper.GetString(fmt.Sprintf("%s.access_token", c.profileName))
}

func (c *configuration) SecureLoginEnabled() bool {
	return viper.GetBool(fmt.Sprintf("%s.secure_login_enabled", c.profileName))
}

func (c *configuration) RedirectURI() string {
	return viper.GetString(fmt.Sprintf("%s.redirect_uri", c.profileName))
}

// OAuthTokenData is the raw OAuth token data returned from the login API call combined with the access token expiry timestamp
func (c *configuration) OAuthTokenData() string {
	return viper.GetString(fmt.Sprintf("%s.oauth_token_data", c.profileName))
}

// Environment is the Genesys Cloud Environment the CLI will talk to
func (c *configuration) Environment() string {
	if Environment != "" {
		return MapEnvironment(Environment)
	}

	return MapEnvironment(viper.GetString(fmt.Sprintf("%s.environment", c.profileName)))
}

func MapEnvironment(env string) string {
	if env != "localhost" && !strings.Contains(env, ".") {
		basePath, ok := regionMappings[env]
		if !ok {
			fmt.Println("Invalid AWS region:", env)
			os.Exit(1)
		}
		return basePath
	}

	return env
}

// LogFilePath is the path the CLI logs to if an override has been specified
func (c *configuration) LogFilePath() string {
	return viper.GetString(fmt.Sprintf("%s.log_file_path", c.profileName))
}

// LoggingEnabled shows whether logging is enabled or disabled for the CLI
func (c *configuration) LoggingEnabled() bool {
	return viper.GetBool(fmt.Sprintf("%s.logging_enabled", c.profileName))
}

// AutoPagination shows whether auto-pagination is enabled or disabled for the CLI
func (c *configuration) AutoPaginationEnabled() bool {
	return viper.GetBool(fmt.Sprintf("%s.auto_pagination_enabled", c.profileName))
}

// ProfileName is the name of the profile being used to run the CLI
func (c *configuration) ProxyConfiguration() string {
	return getProxyConfig(c.profileName)
}

func (c *configuration) GateWayConfiguration() string {
	return getGateWayConfig(c.profileName)
}

func getProxyConfig(profileName string) string {
	// proxy
	protocol := viper.Get(fmt.Sprintf("%s.proxy_protocol", profileName))

	if protocol != nil {
		proxyconf := ProxyConfiguration{}
		proxyconf.Port = viper.GetString(fmt.Sprintf("%s.proxy_port", profileName))
		proxyconf.Protocol = viper.GetString(fmt.Sprintf("%s.proxy_protocol", profileName))
		proxyconf.Host = viper.GetString(fmt.Sprintf("%s.proxy_host", profileName))
		userName := viper.Get(fmt.Sprintf("%s.proxy_username", profileName))

		pathParams := viper.Get(fmt.Sprintf("%s.proxy_pathparams", profileName))

		if pathParams != nil {
			pathParamsMap := parsePathParams(pathParams.(string))
			proxyconf.PathParams = pathParamsMap
		}

		if userName != nil {
			proxyconf.UserName = viper.GetString(fmt.Sprintf("%s.proxy_username", profileName))
			proxyconf.Password = viper.GetString(fmt.Sprintf("%s.proxy_password", profileName))
		}
		jsonData, _ := json.MarshalIndent(proxyconf, "", "")
		return string(jsonData)
	} else {
		return ""
	}
}

func getGateWayConfig(profileName string) string {
	// proxy
	protocol := viper.Get(fmt.Sprintf("%s.gateway_protocol", profileName))

	if protocol != nil {
		gconf := GateWayConfiguration{}
		gconf.Port = viper.GetString(fmt.Sprintf("%s.gateway_port", profileName))
		gconf.Protocol = viper.GetString(fmt.Sprintf("%s.gateway_protocol", profileName))
		gconf.Host = viper.GetString(fmt.Sprintf("%s.gateway_host", profileName))
		userName := viper.Get(fmt.Sprintf("%s.gateway_username", profileName))

		pathParams := viper.Get(fmt.Sprintf("%s.gateway_pathparams", profileName))

		if pathParams != nil {
			pathParamsMap := parsePathParams(pathParams.(string))
			gconf.PathParams = pathParamsMap
		}

		if userName != nil {
			gconf.UserName = viper.GetString(fmt.Sprintf("%s.gateway_username", profileName))
			gconf.Password = viper.GetString(fmt.Sprintf("%s.gateway_password", profileName))
		}
		jsonData, _ := json.MarshalIndent(gconf, "", "")
		return string(jsonData)
	} else {
		return ""
	}
}

func (c *configuration) String() string {
	return fmt.Sprintf(`{"profileName": "%s", "environment": "%s", "logFilePath": "%s", "loggingEnabled": "%v", "grantType": "%s", "clientName": "%s", "clientSecret": "%s", "secureLoginEnabled": "%v", "redirectURI": "%s", "accessToken": "%s", "autoPaginationEnabled": "%v" , "proxyConfiguration" : "%v", "gatewayConfiguration": "%v"}`, c.ProfileName(), c.Environment(), c.LogFilePath(), c.LoggingEnabled(), c.GrantType(), c.ClientID(),
		c.ClientSecret(), c.SecureLoginEnabled(), c.RedirectURI(), c.AccessToken(), c.AutoPaginationEnabled(), c.proxyString(), c.gateWayString())
}

func (c *configuration) proxyString() string {
	a := c.ProxyConfiguration()
	var proxyConfig ProxyConfiguration
	_ = json.Unmarshal([]byte(c.proxyConfiguration), &proxyConfig)

	if a == "" {
		return ""
	}
	return fmt.Sprintf(`{protocol: %s, host: %s,  port: %s, userName: %v, password: %s, pathParams: %s}`,
		proxyConfig.Protocol, proxyConfig.Host, proxyConfig.Port, proxyConfig.UserName, proxyConfig.Password, getPathParams(proxyConfig.PathParams))
}

func (c *configuration) gateWayString() string {
	a := c.GateWayConfiguration()
	var gateWayConfig GateWayConfiguration
	_ = json.Unmarshal([]byte(c.gatewayConfiguration), &gateWayConfig)

	if a == "" {
		return ""
	}
	return fmt.Sprintf(`{protocol: %s, host: %s,  port: %s, userName: %v, password: %s, pathParams: %s}`,
		gateWayConfig.Protocol, gateWayConfig.Host, gateWayConfig.Port, gateWayConfig.UserName, gateWayConfig.Password, getPathParams(gateWayConfig.PathParams))
}

func getPathParams(pathParms map[string]string) string {

	pathParamsStr := "{"
	for k, v := range pathParms {
		pathParamsStr += fmt.Sprintf("%s: %s, ", k, v)
	}
	if len(pathParms) > 0 {
		pathParamsStr = pathParamsStr[:len(pathParamsStr)-2] // Remove the trailing comma and space
	}
	pathParamsStr += "}"

	return pathParamsStr
}

func parsePathParams(pathParamsStr string) map[string]string {
	pathParamsStr = strings.Trim(pathParamsStr, "{}")

	pathParams := make(map[string]string)

	pairs := strings.Split(pathParamsStr, ", ")

	for _, pair := range pairs {
		kv := strings.SplitN(pair, ": ", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			pathParams[key] = value
		}
	}

	return pathParams
}

func getProxyConfigString(profileName string) string {
	// proxy
	proxyI := getProxyConfig(profileName)
	if &proxyI != nil {
		jsonbody, err := json.Marshal(proxyI)
		if err != nil {
			return ""
		}
		return string(jsonbody)
	} else {
		return ""
	}
}

func applyEnvironmentVariableOverrides() {
	environment := os.Getenv("GENESYSCLOUD_REGION")
	if environment != "" && Environment == "" {
		Environment = environment
	}
	clientId := os.Getenv("GENESYSCLOUD_OAUTHCLIENT_ID")
	if clientId != "" && ClientId == "" {
		ClientId = clientId
	}
	clientSecret := os.Getenv("GENESYSCLOUD_OAUTHCLIENT_SECRET")
	if clientSecret != "" && ClientSecret == "" {
		ClientSecret = clientSecret
	}
	accessToken := os.Getenv("GENESYSCLOUD_ACCESS_TOKEN")
	if accessToken != "" && AccessToken == "" {
		AccessToken = accessToken
	}
}

// GetConfig retrieves the config for the current profile
func GetConfig(profileName string) (Configuration, error) {
	applyEnvironmentVariableOverrides()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if OverridesApplied() {
				return &configuration{
					profileName:  profileName,
					clientID:     ClientId,
					clientSecret: ClientSecret,
					accessToken:  AccessToken,
					environment:  Environment,
				}, nil
			}
			homeDir, _ := os.UserHomeDir()
			configDir := fmt.Sprintf("%s/.gc/config.toml", homeDir)
			return nil, fmt.Errorf(`Failed to load config file in "%s". Please use "gc profiles" command to configure the CLI profiles`, configDir)
		} else {
			return nil, fmt.Errorf("Error reading config file, %s", err)
		}
	}

	profile := viper.GetViper().Get(profileName)

	if profile == nil {
		return nil, fmt.Errorf("The profile named %s passed can not be located in the config file.", profileName)
	}

	return &configuration{profileName: profileName,
		grantType:             viper.GetString(fmt.Sprintf("%s.grant_type", profileName)),
		clientID:              viper.GetString(fmt.Sprintf("%s.client_credentials", profileName)),
		clientSecret:          viper.GetString(fmt.Sprintf("%s.client_secret", profileName)),
		redirectURI:           viper.GetString(fmt.Sprintf("%s.redirect_uri", profileName)),
		environment:           viper.GetString(fmt.Sprintf("%s.environment", profileName)),
		oAuthTokenData:        viper.GetString(fmt.Sprintf("%s.oauth_token_data", profileName)),
		accessToken:           viper.GetString(fmt.Sprintf("%s.access_token", profileName)),
		logFilePath:           viper.GetString(fmt.Sprintf("%s.log_file_path", profileName)),
		loggingEnabled:        viper.GetBool(fmt.Sprintf("%s.logging_enabled", profileName)),
		autoPaginationEnabled: viper.GetBool(fmt.Sprintf("%s.auto_pagination_enabled", profileName)),
		secureLoginEnabled:    viper.GetBool(fmt.Sprintf("%s.secure_login_enabled", profileName)),
		proxyConfiguration:    getProxyConfig(profileName),
		gatewayConfiguration:  getGateWayConfig(profileName),
	}, nil
}

func ListConfigs() ([]configuration, error) {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			homeDir, _ := os.UserHomeDir()
			configDir := fmt.Sprintf("%s/.gc/config.toml", homeDir)
			return nil, fmt.Errorf(`Failed to load config file in "%s". Please use "gc profiles" command to configure the CLI profiles`, configDir)
		} else {
			return nil, fmt.Errorf("Error reading config file, %s", err)
		}
	}

	settings := viper.GetViper().AllSettings()
	if settings == nil || len(settings) == 0 {
		return nil, fmt.Errorf("No profiles have been found.")
	}

	configurations := make([]configuration, 0)
	for profileName, _ := range settings {
		configurations = append(configurations, configuration{
			profileName:           profileName,
			grantType:             viper.GetString(fmt.Sprintf("%s.grant_type", profileName)),
			clientID:              viper.GetString(fmt.Sprintf("%s.client_credentials", profileName)),
			clientSecret:          viper.GetString(fmt.Sprintf("%s.client_secret", profileName)),
			redirectURI:           viper.GetString(fmt.Sprintf("%s.redirect_uri", profileName)),
			environment:           viper.GetString(fmt.Sprintf("%s.environment", profileName)),
			oAuthTokenData:        viper.GetString(fmt.Sprintf("%s.oauth_token_data", profileName)),
			accessToken:           viper.GetString(fmt.Sprintf("%s.access_token", profileName)),
			logFilePath:           viper.GetString(fmt.Sprintf("%s.log_file_path", profileName)),
			loggingEnabled:        viper.GetBool(fmt.Sprintf("%s.logging_enabled", profileName)),
			autoPaginationEnabled: viper.GetBool(fmt.Sprintf("%s.auto_pagination_enabled", profileName)),
			secureLoginEnabled:    viper.GetBool(fmt.Sprintf("%s.secure_login_enabled", profileName)),
			proxyConfiguration:    getProxyConfig(profileName),
			gatewayConfiguration:  getGateWayConfig(profileName),
		})
	}

	return configurations, nil
}

func SaveConfig(c Configuration) error {
	return writeConfig(c, nil, "", nil, nil)
}

func UpdateOAuthToken(c Configuration, data *models.OAuthTokenData) error {
	return updateConfig(configuration{
		profileName:    c.ProfileName(),
		oAuthTokenData: data.String(),
	}, nil, nil, nil)
}

func UpdateGrantType(c Configuration, grantType string) error {
	return updateConfig(configuration{
		profileName: c.ProfileName(),
		grantType:   grantType,
	}, nil, nil, nil)
}

func UpdateLogFilePath(c Configuration, filePath string) error {
	return updateConfig(configuration{
		profileName: c.ProfileName(),
		logFilePath: filePath,
	}, nil, nil, nil)
}

func UpdateProxyConfiguration(c Configuration, proxyConf *ProxyConfiguration) error {
	jsonData, _ := json.MarshalIndent(proxyConf, "", "  ")

	return updateConfig(configuration{
		profileName:        c.ProfileName(),
		proxyConfiguration: string(jsonData),
	}, nil, nil, nil)
}

func UpdateGateWayConfiguration(c Configuration, gateWayConf *GateWayConfiguration) error {
	jsonData, _ := json.MarshalIndent(gateWayConf, "", "  ")

	return updateConfig(configuration{
		profileName:          c.ProfileName(),
		gatewayConfiguration: string(jsonData),
	}, nil, nil, nil)
}

func SetLoggingEnabled(c Configuration, loggingEnabled bool) error {
	return updateConfig(configuration{
		profileName: c.ProfileName(),
	}, &loggingEnabled, nil, nil)
}

func SetExperimentalFeature(profileName string, featureName string, enabled bool) error {
	viper.Set(fmt.Sprintf("%s.%s_enabled", profileName, featureName), enabled)
	return viper.WriteConfig()
}

func SetInputFormat(profileName string, format string) error {
	viper.Set(fmt.Sprintf("%s.input_format", profileName), format)
	return viper.WriteConfig()
}

func SetOutputFormat(profileName string, format string) error {
	viper.Set(fmt.Sprintf("%s.output_format", profileName), format)
	return viper.WriteConfig()
}

func GetInputFormat(profileName string) (string, error) {
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}
	return viper.GetString(fmt.Sprintf("%s.input_format", profileName)), nil
}

func GetOutputFormat(profileName string) (string, error) {
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}
	return viper.GetString(fmt.Sprintf("%s.output_format", profileName)), nil
}

func IsExperimentalFeatureEnabled(profileName string, featureName string) bool {
	err := viper.ReadInConfig()
	if err != nil {
		return false
	}
	return viper.GetBool(fmt.Sprintf("%s.%s_enabled", profileName, featureName))
}

func SetAutoPaginationEnabled(c Configuration, autoPaginationEnabled bool) error {
	return updateConfig(configuration{
		profileName: c.ProfileName(),
	}, nil, &autoPaginationEnabled, nil)
}

func GetAutoPaginationEnabled(profileName string) (bool, error) {
	err := viper.ReadInConfig()
	if err != nil {
		return false, err
	}
	return viper.GetBool(fmt.Sprintf("%s.auto_pagination_enabled", profileName)), nil
}

func OverridesApplied() bool {
	return ClientId != "" || ClientSecret != "" || Environment != "" || AccessToken != "" ||
		os.Getenv("GENESYSCLOUD_OAUTHCLIENT_ID") != "" || os.Getenv("GENESYSCLOUD_OAUTHCLIENT_SECRET") != "" || os.Getenv("GENESYSCLOUD_REGION") != "" || os.Getenv("GENESYSCLOUD_ACCESS_TOKEN") != ""
}

func updateConfig(c configuration, loggingEnabled *bool, autoPaginationEnabled *bool, secureLoginEnabled *bool) error {
	if c.grantType != "" {
		viper.Set(fmt.Sprintf("%s.grant_type", c.profileName), c.grantType)
	}
	if c.clientID != "" {
		viper.Set(fmt.Sprintf("%s.client_credentials", c.profileName), c.clientID)
	}
	if c.clientSecret != "" {
		viper.Set(fmt.Sprintf("%s.client_secret", c.profileName), c.clientSecret)
	}
	if c.redirectURI != "" {
		viper.Set(fmt.Sprintf("%s.redirect_uri", c.profileName), c.redirectURI)
	}
	if c.environment != "" {
		viper.Set(fmt.Sprintf("%s.environment", c.profileName), c.environment)
	}
	if c.oAuthTokenData != "" {
		viper.Set(fmt.Sprintf("%s.oauth_token_data", c.profileName), c.oAuthTokenData)
	}
	if c.accessToken != "" {
		viper.Set(fmt.Sprintf("%s.access_token", c.profileName), c.accessToken)
	}
	if c.logFilePath != "" {
		viper.Set(fmt.Sprintf("%s.log_file_path", c.profileName), c.logFilePath)
	}
	if loggingEnabled != nil {
		viper.Set(fmt.Sprintf("%s.logging_enabled", c.profileName), *loggingEnabled)
	}
	if autoPaginationEnabled != nil {
		viper.Set(fmt.Sprintf("%s.auto_pagination_enabled", c.profileName), *autoPaginationEnabled)
	}
	if secureLoginEnabled != nil {
		viper.Set(fmt.Sprintf("%s.secure_login_enabled", c.profileName), *secureLoginEnabled)
	}

	if c.proxyConfiguration != "" {

		var proxyConfig ProxyConfiguration
		_ = json.Unmarshal([]byte(c.proxyConfiguration), &proxyConfig)

		protocol := fmt.Sprintf("%s.proxy_protocol", c.ProfileName())
		port := fmt.Sprintf("%s.proxy_port", c.ProfileName())
		host := fmt.Sprintf("%s.proxy_host", c.ProfileName())
		username := fmt.Sprintf("%s.proxy_username", c.ProfileName())
		password := fmt.Sprintf("%s.proxy_password", c.ProfileName())

		viper.Set(protocol, proxyConfig.Protocol)
		viper.Set(port, proxyConfig.Port)
		viper.Set(host, proxyConfig.Host)
		viper.Set(username, proxyConfig.UserName)
		viper.Set(password, proxyConfig.Password)
		viper.Set(fmt.Sprintf("%s.proxy_pathparams", c.ProfileName()), getPathParams(proxyConfig.PathParams))

	}

	if c.gatewayConfiguration != "" {

		var gConfig GateWayConfiguration
		_ = json.Unmarshal([]byte(c.gatewayConfiguration), &gConfig)

		protocol := fmt.Sprintf("%s.gateway_protocol", c.ProfileName())
		port := fmt.Sprintf("%s.gateway_port", c.ProfileName())
		host := fmt.Sprintf("%s.gateway_host", c.ProfileName())
		username := fmt.Sprintf("%s.gateway_username", c.ProfileName())
		password := fmt.Sprintf("%s.gateway_password", c.ProfileName())
		viper.Set(protocol, gConfig.Protocol)
		viper.Set(port, gConfig.Port)
		viper.Set(host, gConfig.Host)
		viper.Set(username, gConfig.UserName)
		viper.Set(password, gConfig.Password)
		viper.Set(fmt.Sprintf("%s.gateway_pathparams", c.ProfileName()), getPathParams(gConfig.PathParams))
	}

	if viper.ConfigFileUsed() == "" {
		return nil
	}

	return viper.WriteConfig()
}

func writeConfig(c Configuration, data *models.OAuthTokenData, logFilePath string, loggingEnabled *bool, autoPaginationEnabled *bool) error {
	viper.Set(fmt.Sprintf("%s.grant_type", c.ProfileName()), c.GrantType())
	viper.Set(fmt.Sprintf("%s.client_credentials", c.ProfileName()), c.ClientID())
	viper.Set(fmt.Sprintf("%s.client_secret", c.ProfileName()), c.ClientSecret())
	viper.Set(fmt.Sprintf("%s.redirect_uri", c.ProfileName()), c.RedirectURI())
	viper.Set(fmt.Sprintf("%s.environment", c.ProfileName()), c.Environment())
	viper.Set(fmt.Sprintf("%s.access_token", c.ProfileName()), c.AccessToken())
	viper.Set(fmt.Sprintf("%s.secure_login_enabled", c.ProfileName()), c.SecureLoginEnabled())
	if data != nil {
		viper.Set(fmt.Sprintf("%s.oauth_token_data", c.ProfileName()), data.String())
	}
	if logFilePath != "" {
		viper.Set(fmt.Sprintf("%s.log_file_path", c.ProfileName()), logFilePath)
	}
	if loggingEnabled != nil {
		viper.Set(fmt.Sprintf("%s.logging_enabled", c.ProfileName()), *loggingEnabled)
	}
	if autoPaginationEnabled != nil {
		viper.Set(fmt.Sprintf("%s.auto_pagination_enabled", c.ProfileName()), *autoPaginationEnabled)
	}

	if c.ProxyConfiguration() != "" {

		var proxyConfig ProxyConfiguration
		_ = json.Unmarshal([]byte(c.ProxyConfiguration()), &proxyConfig)

		protocol := fmt.Sprintf("%s.proxy_protocol", c.ProfileName())
		port := fmt.Sprintf("%s.proxy_port", c.ProfileName())
		host := fmt.Sprintf("%s.proxy_host", c.ProfileName())
		username := fmt.Sprintf("%s.proxy_username", c.ProfileName())
		password := fmt.Sprintf("%s.proxy_password", c.ProfileName())
		viper.Set(protocol, proxyConfig.Protocol)
		viper.Set(port, proxyConfig.Port)
		viper.Set(host, proxyConfig.Host)
		viper.Set(username, proxyConfig.UserName)
		viper.Set(password, proxyConfig.Password)
		viper.Set(fmt.Sprintf("%s.proxy_pathparams", c.ProfileName()), getPathParams(proxyConfig.PathParams))
	}

	if c.GateWayConfiguration() != "" {

		var gConfig GateWayConfiguration
		_ = json.Unmarshal([]byte(c.GateWayConfiguration()), &gConfig)

		protocol := fmt.Sprintf("%s.gateway_protocol", c.ProfileName())
		port := fmt.Sprintf("%s.gateway_port", c.ProfileName())
		host := fmt.Sprintf("%s.gateway_host", c.ProfileName())
		username := fmt.Sprintf("%s.gateway_username", c.ProfileName())
		password := fmt.Sprintf("%s.gateway_password", c.ProfileName())
		viper.Set(protocol, gConfig.Protocol)
		viper.Set(port, gConfig.Port)
		viper.Set(host, gConfig.Host)
		viper.Set(username, gConfig.UserName)
		viper.Set(password, gConfig.Password)
		viper.Set(fmt.Sprintf("%s.proxy_pathparams", c.ProfileName()), getPathParams(gConfig.PathParams))
	}

	//Checking to see if the file does not exist.  It it doesnt we write out the config as default config.toml
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			homeDir, _ := os.UserHomeDir()
			gcDir := fmt.Sprintf("%s/.gc", homeDir)
			if _, err := os.Stat(gcDir); os.IsNotExist(err) {
				err = os.Mkdir(gcDir, 0755)
				if err != nil {
					return err
				}
			}
			return viper.WriteConfigAs(fmt.Sprintf("%s/config.toml", gcDir))
		}
	}

	return viper.WriteConfig()
}
