package config

import (
	"fmt"
	"os"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"

	"strings"

	"github.com/spf13/viper"
)

type Configuration interface {
	ProfileName() string
	Environment() string
	ClientID() string
	ClientSecret() string
	OAuthTokenData() string
	LogFilePath() string
	LoggingEnabled() bool
	fmt.Stringer
}

type configuration struct {
	profileName    string
	environment    string
	clientID       string
	clientSecret   string
	oAuthTokenData string
	logFilePath    string
	loggingEnabled bool
}

var (
	Environment  string
	ClientId     string
	ClientSecret string
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
	}
)

//ProfileName is the name of the profile being used to run the CLI
func (c *configuration) ProfileName() string {
	return c.profileName
}

//ClientID is the OAuth client id used by the OAuth Client
func (c *configuration) ClientID() string {
	if ClientId != "" {
		return ClientId
	}

	return viper.GetString(fmt.Sprintf("%s.client_credentials", c.profileName))
}

//ClientSecret is the OAuth client secret used by the OAuth Client
func (c *configuration) ClientSecret() string {
	if ClientSecret != "" {
		return ClientSecret
	}

	return viper.GetString(fmt.Sprintf("%s.client_secret", c.profileName))
}

//OAuthTokenData is the raw OAuth token data returned from the login API call combined with the access token expiry timestamp
func (c *configuration) OAuthTokenData() string {
	return viper.GetString(fmt.Sprintf("%s.oauth_token_data", c.profileName))
}

//Environment is the Genesys Cloud Environment the CLI will talk to
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

//LogFilePath is the path the CLI logs to if an override has been specified
func (c *configuration) LogFilePath() string {
	return viper.GetString(fmt.Sprintf("%s.log_file_path", c.profileName))
}

//LoggingEnabled shows whether logging is enabled or disabled for the CLI
func (c *configuration) LoggingEnabled() bool {
	return viper.GetBool(fmt.Sprintf("%s.logging_enabled", c.profileName))
}

func (c *configuration) String() string {
	return fmt.Sprintf(`{"profileName": "%s", "environment": "%s", "logFilePath": "%s", "loggingEnabled": "%v", "clientName": "%s", "clientSecret": "%s"}`, c.ProfileName(), c.Environment(), c.LogFilePath(), c.LoggingEnabled(), c.ClientID(), c.ClientSecret())
}

func applyEnvironmentVariableOverrides() {
	environment := os.Getenv("GENESYSCLOUD_REGION")
	if environment != "" {
		Environment = environment
	}
	clientId := os.Getenv("GENESYSCLOUD_OAUTHCLIENT_ID")
	if clientId != "" {
		ClientId = clientId
	}
	clientSecret := os.Getenv("GENESYSCLOUD_OAUTHCLIENT_SECRET")
	if clientSecret != "" {
		ClientSecret = clientSecret
	}
}

//GetConfig retrieves the config for the current profile
func GetConfig(profileName string) (Configuration, error) {
	applyEnvironmentVariableOverrides()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if OverridesApplied() {
				return &configuration{
					profileName: profileName,
					clientID:       ClientId,
					clientSecret:   ClientSecret,
					environment:    Environment,
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
		clientID:       viper.GetString(fmt.Sprintf("%s.client_credentials", profileName)),
		clientSecret:   viper.GetString(fmt.Sprintf("%s.client_secret", profileName)),
		environment:    viper.GetString(fmt.Sprintf("%s.environment", profileName)),
		oAuthTokenData: viper.GetString(fmt.Sprintf("%s.oauth_token_data", profileName)),
		logFilePath:    viper.GetString(fmt.Sprintf("%s.log_file_path", profileName)),
		loggingEnabled: viper.GetBool(fmt.Sprintf("%s.logging_enabled", profileName)),
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
			profileName:    profileName,
			clientID:       viper.GetString(fmt.Sprintf("%s.client_credentials", profileName)),
			clientSecret:   viper.GetString(fmt.Sprintf("%s.client_secret", profileName)),
			environment:    viper.GetString(fmt.Sprintf("%s.environment", profileName)),
			oAuthTokenData: viper.GetString(fmt.Sprintf("%s.oauth_token_data", profileName)),
			logFilePath:    viper.GetString(fmt.Sprintf("%s.log_file_path", profileName)),
			loggingEnabled: viper.GetBool(fmt.Sprintf("%s.logging_enabled", profileName)),
		})
	}

	return configurations, nil
}

func SaveConfig(c Configuration) error {
	return writeConfig(c, nil, "", nil)
}

func UpdateOAuthToken(c Configuration, data *models.OAuthTokenData) error {
	return updateConfig(configuration{
		profileName: c.ProfileName(),
		oAuthTokenData: data.String(),
	}, nil)
}

func UpdateLogFilePath(c Configuration, filePath string) error {
	return updateConfig(configuration{
		profileName: c.ProfileName(),
		logFilePath: filePath,
	},nil)
}

func SetLoggingEnabled(c Configuration, loggingEnabled bool) error {
	return updateConfig(configuration{
		profileName: c.ProfileName(),
	}, &loggingEnabled)
}

func OverridesApplied() bool {
	return ClientId != "" || ClientSecret != "" || Environment != "" ||
		os.Getenv("GENESYSCLOUD_OAUTHCLIENT_ID") != "" || os.Getenv("GENESYSCLOUD_OAUTHCLIENT_SECRET") != "" || os.Getenv("GENESYSCLOUD_REGION") != ""
}

func updateConfig(c configuration, loggingEnabled *bool) error {
	if c.clientID != "" {
		viper.Set(fmt.Sprintf("%s.client_credentials", c.profileName), c.clientID)
	}
	if c.clientSecret != "" {
		viper.Set(fmt.Sprintf("%s.client_secret", c.profileName), c.clientSecret)
	}
	if c.environment != "" {
		viper.Set(fmt.Sprintf("%s.environment", c.profileName), c.environment)
	}
	if c.oAuthTokenData != "" {
		viper.Set(fmt.Sprintf("%s.oauth_token_data", c.profileName), c.oAuthTokenData)
	}
	if c.logFilePath != "" {
		viper.Set(fmt.Sprintf("%s.log_file_path", c.profileName), c.logFilePath)
	}
	if loggingEnabled != nil {
		viper.Set(fmt.Sprintf("%s.logging_enabled", c.profileName), *loggingEnabled)
	}

	return viper.WriteConfig()
}

func writeConfig(c Configuration, data *models.OAuthTokenData, logFilePath string, loggingEnabled *bool) error {
	viper.Set(fmt.Sprintf("%s.client_credentials", c.ProfileName()), c.ClientID())
	viper.Set(fmt.Sprintf("%s.client_secret", c.ProfileName()), c.ClientSecret())
	viper.Set(fmt.Sprintf("%s.environment", c.ProfileName()), c.Environment())
	if data != nil {
		viper.Set(fmt.Sprintf("%s.oauth_token_data", c.ProfileName()), data.String())
	}
	if logFilePath != "" {
		viper.Set(fmt.Sprintf("%s.log_file_path", c.ProfileName()), logFilePath)
	}
	if loggingEnabled != nil {
		viper.Set(fmt.Sprintf("%s.logging_enabled", c.ProfileName()), *loggingEnabled)
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
