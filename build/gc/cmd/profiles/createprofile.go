package profiles

import (
	"encoding/json"
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
	PKCEGrant                   = "3"
)

func isValidGrantType(t GrantType) bool {
	return t == None || t == ClientCredentials || t == ImplicitGrant || t == PKCEGrant
}

func constructConfig(profileName string, environment string, grantType GrantType, clientID string, clientSecret string, redirectURI string, secureLoginEnabled bool, accessToken string, proxyConf *config.ProxyConfiguration, gconf *config.GateWayConfiguration) config.Configuration {
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

	c.ProxyConfigurationFunc = func() string {
		jsonData, _ := json.MarshalIndent(proxyConf, "", "  ")
		return string(jsonData)
	}

	c.GateWayConfigurationFunc = func() string {
		jsonData, _ := json.MarshalIndent(gconf, "", "  ")
		return string(jsonData)
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
		gateWayChoice      string
	)

	fmt.Print("Profile Name [DEFAULT]: ")
	_, _ = fmt.Scanln(&name)

	if name == "" {
		name = "DEFAULT"
	}

	fmt.Printf("Environment [mypurecloud.com]: ")
	_, _ = fmt.Scanln(&environment)

	if environment == "" {
		environment = "mypurecloud.com"
	}

	fmt.Print("Note: If you provide an access token, this will take precedence over any authorization grant type.\n")
	fmt.Print("Access Token (Optional): ")
	accessToken = readSensitiveInput()

	for {
		fmt.Print("Select your authorization grant type.\n")
		fmt.Print("\t0. None\n\t1. Client Credentials\n\t2. Implicit Grant\n\t3. PKCE Grant\nGrant Type: ")
		_, _ = fmt.Scanln(&grantType)

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
		for {
			fmt.Print("Would you like to use a secure HTTP connection? [Y/N]: ")
			_, _ = fmt.Scanln(&authChoice)
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

	var proxyConfig *config.ProxyConfiguration
	for {
		fmt.Print("Would you like to use a proxy server? [Y/N]: ")
		_, _ = fmt.Scanln(&proxyChoice)
		if strings.ToUpper(proxyChoice) == "Y" {
			proxyConfig = requestProxyDetails()
			break
		} else if strings.ToUpper(proxyChoice) == "N" {
			proxyConfig = nil
			break
		} else {
			fmt.Print("Provide valid option.\n")
			continue
		}
	}

	var gateWayConfig *config.GateWayConfiguration
	for {
		fmt.Print("Would you like to use a gateway server? [Y/N]: ")
		_, _ = fmt.Scanln(&gateWayChoice)
		if strings.ToUpper(gateWayChoice) == "Y" {
			gateWayConfig = requestGateWayDetails()
			break
		} else if strings.ToUpper(gateWayChoice) == "N" {
			gateWayConfig = nil
			break
		} else {
			fmt.Print("Provide valid option.\n")
			continue
		}
	}

	return constructConfig(name, environment, grantType, clientID, clientSecret, redirectURL.String(), secureLoginEnabled, accessToken, proxyConfig, gateWayConfig)
}

func requestClientCreds(accessToken string, grantType GrantType) (string, string) {
	id := ""
	secret := ""

	if grantType == ClientCredentials {
		if accessToken != "" {
			fmt.Print("Client ID: ")
			_, _ = fmt.Scanln(&id)
			fmt.Print("Client Secret: ")
			secret = readSensitiveInput()
		} else {
			for id == "" {
				fmt.Print("Client ID: ")
				_, _ = fmt.Scanln(&id)
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
			_, _ = fmt.Scanln(&id)
		}

		fmt.Print("Client Secret (Optional): ")
		secret = readSensitiveInput()
	} else if grantType == PKCEGrant {
		// PKCE Grant
		for id == "" {
			fmt.Print("Client ID: ")
			_, _ = fmt.Scanln(&id)
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
	pathParamRequired := ""
	pathParams := make(map[string]string)
	for protocol == "" {
		fmt.Print("Protocol (http/https): ")
		_, _ = fmt.Scanln(&protocol)
	}
	for port == "" {
		fmt.Print("Port for the Proxy: ")
		_, _ = fmt.Scanln(&port)
	}
	for host == "" {
		fmt.Print("Host name for the Proxy server: ")
		_, _ = fmt.Scanln(&host)
	}

	for {
		fmt.Print("Do we require Authorisation to use the proxy server? [Y/N]: ")
		_, _ = fmt.Scanln(&isAuthRequired)
		if strings.ToUpper(isAuthRequired) == "Y" {
			_, _ = fmt.Scanln(&isAuthRequired)
			for username == "" {
				fmt.Print("username for the Proxy: ")
				_, _ = fmt.Scanln(&username)
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

outerLoopLabel:
	for {
		fmt.Print("Do we have different proxy paths for login and Others? [Y/N]: ")
		_, _ = fmt.Scanln(&pathParamRequired)
		if strings.ToUpper(pathParamRequired) == "Y" {
			for {
				pathAPI := ""
				pathValue := ""
				pathParamRequired = ""
				for pathAPI == "" {
					fmt.Print("Enter API flow for which you want to add the path? Ex [login/other] : ")
					_, _ = fmt.Scanln(&pathAPI)
				}
				for pathValue == "" {
					fmt.Print("Enter Path : ")
					_, _ = fmt.Scanln(&pathValue)
				}
				pathParams[pathAPI] = pathValue
				for pathParamRequired == "" {
					fmt.Print("Do you want to enter more Paths? [Y/N]: ")
					_, _ = fmt.Scanln(&pathParamRequired)
				}
				if strings.ToUpper(pathParamRequired) == "Y" {
					continue
				} else {
					break outerLoopLabel
				}
			}
		} else if strings.ToUpper(pathParamRequired) == "N" {
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
	if len(pathParams) > 0 {
		proxyconf.PathParams = pathParams
	}

	return &proxyconf

}

func requestGateWayDetails() *config.GateWayConfiguration {
	protocol := ""
	host := ""
	port := ""
	isAuthRequired := ""
	username := ""
	password := ""
	pathParamRequired := ""
	pathParams := make(map[string]string)
	for protocol == "" {
		fmt.Print("Protocol (http/https): ")
		_, _ = fmt.Scanln(&protocol)
	}
	for port == "" {
		fmt.Print("Port for the GateWay: ")
		_, _ = fmt.Scanln(&port)
	}
	for host == "" {
		fmt.Print("Host name for the GateWay server: ")
		_, _ = fmt.Scanln(&host)
	}

	for {
		fmt.Print("Do we require Authorisation to use the GateWay server? [Y/N]: ")
		_, _ = fmt.Scanln(&isAuthRequired)
		if strings.ToUpper(isAuthRequired) == "Y" {
			_, _ = fmt.Scanln(&isAuthRequired)
			for username == "" {
				fmt.Print("username for the GateWay: ")
				_, _ = fmt.Scanln(&username)
			}
			for password == "" {
				fmt.Print("Password for the GateWay server: ")
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

outerLoopLabel:
	for {
		fmt.Print("Do we have different GateWay paths for login and api? [Y/N]: ")
		_, _ = fmt.Scanln(&pathParamRequired)
		if strings.ToUpper(pathParamRequired) == "Y" {
			for {
				pathAPI := ""
				pathValue := ""
				pathParamRequired = ""
				for pathAPI == "" {
					fmt.Print("Enter API flow for which you want to add the path? Ex [login/api] : ")
					_, _ = fmt.Scanln(&pathAPI)
				}
				for pathValue == "" {
					fmt.Print("Enter Path : ")
					_, _ = fmt.Scanln(&pathValue)
				}
				pathParams[pathAPI] = pathValue
				for pathParamRequired == "" {
					fmt.Print("Do you want to enter more Paths? [Y/N]: ")
					_, _ = fmt.Scanln(&pathParamRequired)
				}
				if strings.ToUpper(pathParamRequired) == "Y" {
					continue
				} else {
					break outerLoopLabel
				}
			}
		} else if strings.ToUpper(pathParamRequired) == "N" {
			break
		} else {
			fmt.Print("Provide valid option.\n")
			continue
		}
	}

	gconf := config.GateWayConfiguration{}
	gconf.Port = port
	gconf.Protocol = protocol
	gconf.Host = host
	if username != "" {
		gconf.UserName = username
		gconf.Password = password
	}
	if len(pathParams) > 0 {
		gconf.PathParams = pathParams
	}

	return &gconf

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
	_, _ = fmt.Scanln(&inputPort)
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

	getConfig, err := config.GetConfig(name)

	//profile already exists we will get a nil back and we must resolve the results
	if err == nil && getConfig.ProfileName() != "" {
		for {
			var overwrite string
			fmt.Printf("Profile name %s already exists in the getConfig file. Overwrite (Y/N): ", name)
			_, _ = fmt.Scanln(&overwrite)

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
		//Check to see if it is an HTTP error and if so, check to see if it is what we are expecting
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
