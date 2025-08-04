package cmd

import (
	"github.com/hashicorp/go-version"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alternative_formats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/data_format"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/transform_data"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"

	"flag"
	"fmt"
	"strings"
	"encoding/json"
	"io/ioutil"
	"time"
	"net/http"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dummy_command"
	{{addit.addImports}}
	"os"
)

type LatestVersion struct {
	Version string `json:"version"`
}

var (
	profileName string
	output string
	externalVersionURL = "https://sdk-cdn.mypurecloud.com/external/go-cli/latest-version.json"
	internalVersionURL = "https://sdk-cdn.mypurecloud.com/internal/go-cli/latest-version.json"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gc",
	Short: "gc is a CLI for interacting with Genesys Cloud",
	Long:  `gc is a CLI for interacting with Genesys Cloud`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gc",
	Long:  `All software has versions. This is gc version's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current version: 136.0.0")
		checkForNewVersion()
	},
}

func checkForNewVersion() {
	var url string
	if strings.Contains("https://api.mypurecloud.com", "api.mypurecloud.com") {
		url = externalVersionURL
	} else {
		url = internalVersionURL
	}

	err, latestVersion := retrieveLatestVersion(url)
	if err != nil {
		fmt.Println("An error occurred while retrieving the latest version.")
		return
	}

	if versionsAreEqual("136.0.0", latestVersion) {
		fmt.Println("You're all up to date.")
	} else {
		fmt.Printf("A new version of the CLI is available: %v\n", latestVersion)
	}
}

func retrieveLatestVersion(url string) (error, string) {
	client := http.Client{
		Timeout: time.Second * 6,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err, ""
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return getErr, ""
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr, ""
	}

	latestVersion := LatestVersion{}
	jsonErr := json.Unmarshal(body, &latestVersion)
	if jsonErr != nil {
		return jsonErr, ""
	}

	return nil, latestVersion.Version
}

func versionsAreEqual(current string, latest string) bool {
	v1, _ := version.NewVersion(current)
	v2, _ := version.NewVersion(latest)
	return v1.GreaterThanOrEqual(v2)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	flag.CommandLine.Parse(nil)
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(fmt.Sprintf("%v\n", err.Error()))
	}
}

func initViperConfig() {
	homeDir, _ := os.UserHomeDir()
	viper.SetConfigName("config")                       // name of config file (without extension)
	viper.SetConfigType("toml")                         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(fmt.Sprintf("%s/.gc", homeDir)) // call multiple times to add many search paths
	viper.AddConfigPath(".")                            // optionally look for config in the working directory

	viper.SetEnvPrefix("gc")
	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize()
	initViperConfig()

	rootCmd.PersistentFlags().StringVarP(&profileName, "profile", "p", "DEFAULT", "Name of the profile to use for configuring the cli")
	rootCmd.RegisterFlagCompletionFunc("profile", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return profiles.ListProfileNames(), cobra.ShellCompDirectiveDefault
	})
	rootCmd.PersistentFlags().BoolP("indicateprogress", "i", false, "Trace progress indicators to stderr")

	rootCmd.PersistentFlags().StringVar(&config.Environment, "environment", "", "environment override. E.g. mypurecloud.com.au or ap-southeast-2")
	rootCmd.RegisterFlagCompletionFunc("environment", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		regionMappings := &config.RegionMappings
		possibleEnvs := make([]string, 0)
		if regionMappings == nil {
			logger.Fatal("No regions found! Something is terribly wrong!")
		}
		for region, domain := range *regionMappings {
			possibleEnvs = append(possibleEnvs, region)
			possibleEnvs = append(possibleEnvs, domain)
		}
		return possibleEnvs, cobra.ShellCompDirectiveDefault
	})
	rootCmd.PersistentFlags().StringVar(&config.ClientId, "clientid", "", "clientId override")
	rootCmd.PersistentFlags().StringVar(&config.ClientSecret, "clientsecret", "", "clientSecret override")
	rootCmd.PersistentFlags().StringVar(&config.AccessToken, "accesstoken", "", "accessToken override")

	rootCmd.PersistentFlags().StringVar(&transform_data.TemplateFile, "transform", "", "Provide a Go template file for transforming output data")
	rootCmd.PersistentFlags().StringVar(&transform_data.TemplateStr, "transformstr", "", "Provide a Go template string for transforming output data")

	rootCmd.PersistentFlags().StringVar(&data_format.InputFormat, "inputformat", "", "Data input format. Supported formats: YAML, JSON")
	rootCmd.RegisterFlagCompletionFunc("inputformat", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"yaml", "json"}, cobra.ShellCompDirectiveDefault
	})
	rootCmd.PersistentFlags().StringVar(&data_format.OutputFormat, "outputformat", "", "Data output format. Supported formats: YAML, JSON")
	rootCmd.RegisterFlagCompletionFunc("outputformat", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"yaml", "json"}, cobra.ShellCompDirectiveDefault
	})

	if data_format.OutputFormat == "" {
		data_format.OutputFormat, _ = config.GetOutputFormat(getProfileName(os.Args))
	}
	if data_format.InputFormat == "" {
		data_format.InputFormat, _ = config.GetInputFormat(getProfileName(os.Args))
	}

	rootCmd.AddCommand(alternative_formats.Cmdalternative_formats())

	rootCmd.AddCommand(versionCmd)
	{{addit.addCommands}}

	if config.IsExperimentalFeatureEnabled(getProfileName(os.Args), models.DummyCommand.String()) {
		rootCmd.AddCommand(dummy_command.Cmddummy_command())
	}

	if os.Getenv("GenerateGcDocs") != "" {
		if _, err := os.Stat("/tmp/gc_docs"); os.IsNotExist(err) {
			err = os.Mkdir("/tmp/gc_docs", 0755)
			if err != nil {
				logger.Fatal(err)
			}
			err = doc.GenMarkdownTree(rootCmd, "/tmp/gc_docs")
			if err != nil {
				logger.Fatal(err)
			}
		}
	}
}

func getProfileName(args []string) string {
    name := ""
    for i, s := range args {
        if (s == "-p" || s == "--profile") && i < len(args)-1 {
            return args[i+1]
        }
        if strings.HasPrefix(s, "--profile=") {
            name = strings.Replace(s, "--profile=", "", -1)
        } else if strings.HasPrefix(s, "-p=") {
            name = strings.Replace(s, "-p=", "", -1)
        }
    }
    if name == "" {
        return "DEFAULT"
    }
    return name
}

