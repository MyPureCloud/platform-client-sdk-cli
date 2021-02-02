package config

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"os"

	"github.com/spf13/viper"
)

type Configuration interface {
	ProfileName() string
	Environment() string
	ClientID() string
	ClientSecret() string
	OAuthTokenData() string
	fmt.Stringer
}

type configuration struct {
	profileName  string
	environment  string
	clientID     string
	clientSecret string
	oAuthTokenData   string
}

//ProfileName is the name of the profile being used to run the CLI
func (c *configuration) ProfileName() string {
	return c.profileName
}

//ClientID is the OAuth client id used by the OAuth Client
func (c *configuration) ClientID() string {
	return viper.GetString(fmt.Sprintf("%s.client_credentials", c.profileName))
}

//ClientSecret is the OAuth client secret used by the OAuth Client
func (c *configuration) ClientSecret() string {
	return viper.GetString(fmt.Sprintf("%s.client_secret", c.profileName))
}

//OAuthTokenData is the raw OAuth token data returned from the login API call combined with the access token expiry timestamp
func (c *configuration) OAuthTokenData() string {
	return viper.GetString(fmt.Sprintf("%s.oauth_token_data", c.profileName))
}

//Environment is the Genesys Cloud Environment the CLI will talk to
func (c *configuration) Environment() string {
	return viper.GetString(fmt.Sprintf("%s.environment", c.profileName))
}

func (c *configuration) String() string {
	return fmt.Sprintf(`{"profileName": "%s", "environment": "%s", "clientName": "%s", "clientSecret": "%s"}`, c.ProfileName(), c.Environment(), c.ClientID(), c.ClientSecret())
}

//GetConfig retrieves the config for the current profile
func GetConfig(profileName string) (Configuration, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file, %s", err)
	}

	profile := viper.GetViper().Get(profileName)

	if profile == nil {
		return nil, fmt.Errorf("The profile named %s passed can not be located in the config file.", profileName)
	}

	return &configuration{profileName: profileName,
		clientID:         viper.GetString(fmt.Sprintf("%s.client_credentials", profileName)),
		clientSecret:     viper.GetString(fmt.Sprintf("%s.client_secret", profileName)),
		environment:      viper.GetString(fmt.Sprintf("%s.environment", profileName)),
		oAuthTokenData:   viper.GetString(fmt.Sprintf("%s.oauth_token_data", profileName)),
	}, nil
}

func SaveConfig(config Configuration) error {
	return writeConfig(config, nil)
}

func UpdateOAuthToken(config Configuration, data *models.OAuthTokenData) error {
	return writeConfig(config, data)
}

func writeConfig(config Configuration, data *models.OAuthTokenData) error {
	viper.Set(fmt.Sprintf("%s.client_credentials", config.ProfileName()), config.ClientID())
	viper.Set(fmt.Sprintf("%s.client_secret", config.ProfileName()), config.ClientSecret())
	viper.Set(fmt.Sprintf("%s.environment", config.ProfileName()), config.Environment())
	if data != nil {
		viper.Set(fmt.Sprintf("%s.oauth_token_data", config.ProfileName()), data.String())
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
