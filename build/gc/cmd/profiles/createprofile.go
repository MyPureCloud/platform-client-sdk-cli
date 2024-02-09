package profiles

import (
	"fmt"
	"net/url"
	"strings"
	"syscall"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/mocks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/restclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

type GrantType string

const (
	None              GrantType = "0"
	ClientCredentials           = "1"
	ImplicitGrant               = "2"
	PKCEGrant               	= "3"
)

func isValidGrantType(t GrantType) bool {
	return t == None || t == ClientCredentials || t == ImplicitGrant || t == PKCEGrant
}

func constructConfig(profileName string, environment string, grantType GrantType, clientID string, clientSecret string, redirectURI string, secureLoginEnabled bool, accessToken string, proxyConf *config.ProxyConfiguration) config.Configuration {
	c := &mocks.MockClientConfig{}

	c.ProfileNameFunc = func() string {
		return profileName
	}

	c.EnvironmentFunc = func() string {
		return config.MapEnvironment(environment)
	}

	c.LogFilePathFunc = func() string {
		return ""
	}

	c.LoggingEnabledFunc = func() bool {
		return false
	}

	c.AutoPaginationEnabledFunc = func() bool {
		return false
	}

	c.GrantTypeFunc = func() string {
		return fmt.Sprintf("%s", grantType)
	}

	c.ClientIDFunc = func() string {
		return clientID
	}

	c.ClientSecretFunc = func() string {
		return clientSecret
	}

	c.RedirectURIFunc = func() string {
		return redirectURI
	}

	c.SecureLoginEnabledFunc = func() bool {
		return secureLoginEnabled
	}

	c.OAuthTokenDataFunc = func() string {
		return ""
	}

	c.AccessTokenFunc = func() string {
		return accessToken
	}

	c.ProxyConfigurationFunc = func() *config.ProxyConfiguration {
		return proxyConf
	}

	return c
}

func requestUserInput() config.Configuration {
	var (
		name               string
		environment        string
		clientID           string
		clientSecret       string
		accessToken        string
		authChoice         string
		grantType          GrantType
		redirectURL        url.URL
		secureLoginEnabled = false
		proxyChoice        string
	)

	fmt.Print("Profile Name [DEFAULT]: ")
	fmt.Scanln(&name)

	if name == "" {
		name = "DEFAULT"
	}

	fmt.Printf("Environment [mypurecloud.com]: ")
	fmt.Scanln(&environment)

	if environment == "" {
		environment = "mypurecloud.com"
	}

	fmt.Print("Note: If you provide an access token, this will take precedence over any authorization grant type.\n")
	fmt.Print("Access Token (Optional): ")
	accessToken = readSensitiveInput()

	for true {
		fmt.Print("Select your authorization grant type.\n")
		fmt.Print("\t0. None\n\t1. Client Credentials\n\t2. Implicit Grant\n\t3. PKCE Grant\nGrant Type: ")
		fmt.Scanln(&grantType)

		if accessToken == "" && grantType == None {
			fmt.Print("If you have not provided an access token, you must select a grant type.\n")
			continue
		}
		if isValidGrantType(grantType) {
			break
		}
	}

	clientID, clientSecret = requestClientCreds(accessToken, grantType)

	if (grantType == ImplicitGrant) || (grantType == PKCEGrant) {
		redirectURL.Host = "localhost:" + requestRedirectURIPort()
		for true {
			fmt.Print("Would you like to use a secure HTTP connection? [Y/N]: ")
			fmt.Scanln(&authChoice)
			if strings.ToUpper(authChoice) == "Y" {
				secureLoginEnabled = true
				redirectURL.Scheme = "https"
				break
			} else if strings.ToUpper(authChoice) == "N" {
				secureLoginEnabled = false
				redirectURL.Scheme = "http"
				break
			}
		}
		fmt.Printf("Redirect URI: %s\n", redirectURL.String())
	}

	for true {
		fmt.Print("Would you like to use a proxy server? [Y/N]: ")
		fmt.Scanln(&proxyChoice)
		if strings.ToUpper(proxyChoice) == "Y" {
			proxyConf := requestProxyDetails()
			return constructConfig(name, environment, grantType, clientID, clientSecret, redirectURL.String(), secureLoginEnabled, accessToken, proxyConf)
			break
		} else if strings.ToUpper(proxyChoice) == "N" {
			return constructConfig(name, environment, grantType, clientID, clientSecret, redirectURL.String(), secureLoginEnabled, accessToken, nil)
			break
		} else {
			fmt.Print("Provide valid option.\n")
			continue
		}
	}

	return constructConfig(name, environment, grantType, clientID, clientSecret, redirectURL.String(), secureLoginEnabled, accessToken, nil)
}

func requestClientCreds(accessToken string, grantType GrantType) (string, string) {
	id := ""
	secret := ""

	if grantType == ClientCredentials {
		if accessToken != "" {
			fmt.Print("Client ID: ")
			fmt.Scanln(&id)

			fmt.Print("Client Secret: ")
			secret = readSensitiveInput()
		} else {
			for id == "" {
				fmt.Print("Client ID: ")
				fmt.Scanln(&id)
			}
			for secret == "" {
				fmt.Print("Client Secret: ")
				secret = readSensitiveInput()
			}
		}
	} else if grantType == ImplicitGrant {
		// Implicit Grant
		for id == "" {
			fmt.Print("Client ID: ")
			fmt.Scanln(&id)
		}

		fmt.Print("Client Secret (Optional): ")
		secret = readSensitiveInput()
	} else if grantType == PKCEGrant {
		// PKCE Grant
		for id == "" {
			fmt.Print("Client ID: ")
			fmt.Scanln(&id)
		}
	}

	return id, secret
}

func requestProxyDetails() *config.ProxyConfiguration {
	protocol := ""
	host := ""
	port := ""
	isAuthRequired := ""
	username := ""
	password := ""

	for protocol == "" {
		fmt.Print("Protocol (http/https): ")
		fmt.Scanln(&protocol)
	}
	for port == "" {
		fmt.Print("Port for the Proxy: ")
		fmt.Scanln(&port)
	}
	for host == "" {
		fmt.Print("Host name for the Proxy server: ")
		fmt.Scanln(&host)
	}

	for true {
		fmt.Print("Do we require Authorisation to use the proxy server? [Y/N]: ")
		fmt.Scanln(&isAuthRequired)
		if strings.ToUpper(isAuthRequired) == "Y" {
			for username == "" {
				fmt.Print("username for the Proxy: ")
				fmt.Scanln(&username)
			}
			for password == "" {
				fmt.Print("Password for the Proxy server: ")
				password = readSensitiveInput()
			}
			break
		} else if strings.ToUpper(isAuthRequired) == "N" {
			break
		} else {
			fmt.Print("Provide valid option.\n")
			continue
		}
	}

	proxyconf := config.ProxyConfiguration{}
	proxyconf.Port = port
	proxyconf.Protocol = protocol
	proxyconf.Host = host
	if username != "" {
		proxyconf.UserName = username
		proxyconf.Password = password
	}
	return &proxyconf

}

func readSensitiveInput() string {
	bytes, _ := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	return string(bytes)
}

func requestRedirectURIPort() string {
	var inputPort string
	defaultPort := "8080"

	fmt.Printf("Redirect URI port [%s]: ", defaultPort)
	fmt.Scanln(&inputPort)
	if inputPort == "" {
		inputPort = defaultPort
	}

	return inputPort
}

func overrideConfig(name string) bool {
	//The file does not exist and therefore can not be overridden
	if err := viper.ReadInConfig(); err != nil {
		return true
	}

	config, err := config.GetConfig(name)

	//profile already exists we will get a nil back and we must resolve the results
	if err == nil && config.ProfileName() != "" {
		for true {
			var overwrite string
			fmt.Printf("Profile name %s already exists in the config file. Overwrite (Y/N): ", name)
			fmt.Scanln(&overwrite)

			if strings.ToUpper(overwrite) == "N" {
				return false
			}

			if strings.ToUpper(overwrite) == "Y" {
				break
			}
		}
	}

	return true
}

func validateCredentials(config config.Configuration) bool {
	oauthToken, err := restclient.Authorize(config)
	if err != nil || oauthToken.AccessToken == "" {
		//Check to see if its an HTTP error and if its check to see if its what we are expecting
		if _, ok := err.(*models.HttpStatusError); ok {
			return false
		}
		logger.Fatal(err)
	}

	return true
}

var createProfilesCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new profile",
	Long:  `Creates a new profile`,

	Run: func(cmd *cobra.Command, args []string) {

		newConfig := requestUserInput()

		if overrideConfig(newConfig.ProfileName()) == false {
			logger.Fatal("Exiting profile creation process")
		}

		if newConfig.AccessToken() == "" && validateCredentials(newConfig) == false {
			logger.Fatal("The credentials provided are not valid.")
		}

		if err := config.SaveConfig(newConfig); err != nil {
			logger.Fatal(err)
		}

		fmt.Printf("Profile %s saved.\n", newConfig.ProfileName())
	},
}
