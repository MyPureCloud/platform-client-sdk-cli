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
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/date"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/fieldconfig"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/identityproviders"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/ipranges"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/locations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/mobiledevices"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chat"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orphanrecordings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presencedefinitions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scripts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/stations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/systempresences"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/timezones"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/tokens"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/userrecordings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/autopagination"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/completion"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/experimental"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/logging"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/profiles"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/alerting"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/audits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/billing"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/certificate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/contentmanagement"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/uploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dataextensions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/profile"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/organizations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/fax"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gdpr"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/geolocations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/languageunderstanding"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/license"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/documentationfile"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/notifications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/oauth"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/responsemanagement"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/scim"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/speechandtextanalytics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/configuration"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webmessaging"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/widgets"
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
		fmt.Println("Current version: 30.0.0")
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

	if versionsAreEqual("30.0.0", latestVersion) {
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
	rootCmd.PersistentFlags().BoolP("indicateprogress", "i", false, "Trace progress indicators to stderr")

	rootCmd.PersistentFlags().StringVar(&config.Environment, "environment", "", "environment override. E.g. mypurecloud.com.au or ap-southeast-2")
	rootCmd.PersistentFlags().StringVar(&config.ClientId, "clientid", "", "clientId override")
	rootCmd.PersistentFlags().StringVar(&config.ClientSecret, "clientsecret", "", "clientSecret override")
	rootCmd.PersistentFlags().StringVar(&config.AccessToken, "accesstoken", "", "accessToken override")

	rootCmd.PersistentFlags().StringVar(&transform_data.TemplateFile, "transform", "", "Provide a Go template file for transforming output data")
	rootCmd.PersistentFlags().StringVar(&transform_data.TemplateStr, "transformstr", "", "Provide a Go template string for transforming output data")

	rootCmd.PersistentFlags().StringVar(&data_format.InputFormat, "inputformat", "", "Data input format. Supported formats: YAML, JSON")
	rootCmd.PersistentFlags().StringVar(&data_format.OutputFormat, "outputformat", "", "Data output format. Supported formats: YAML, JSON")

	if data_format.OutputFormat == "" {
		data_format.OutputFormat, _ = config.GetOutputFormat(getProfileName(os.Args))
	}
	if data_format.InputFormat == "" {
		data_format.InputFormat, _ = config.GetInputFormat(getProfileName(os.Args))
	}

	rootCmd.AddCommand(alternative_formats.Cmdalternative_formats())

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(date.Cmddate())
	rootCmd.AddCommand(conversations.Cmdconversations())
	rootCmd.AddCommand(fieldconfig.Cmdfieldconfig())
	rootCmd.AddCommand(flows.Cmdflows())
	rootCmd.AddCommand(greetings.Cmdgreetings())
	rootCmd.AddCommand(groups.Cmdgroups())
	rootCmd.AddCommand(identityproviders.Cmdidentityproviders())
	rootCmd.AddCommand(integrations.Cmdintegrations())
	rootCmd.AddCommand(ipranges.Cmdipranges())
	rootCmd.AddCommand(languages.Cmdlanguages())
	rootCmd.AddCommand(locations.Cmdlocations())
	rootCmd.AddCommand(mobiledevices.Cmdmobiledevices())
	rootCmd.AddCommand(chat.Cmdchat())
	rootCmd.AddCommand(orphanrecordings.Cmdorphanrecordings())
	rootCmd.AddCommand(presencedefinitions.Cmdpresencedefinitions())
	rootCmd.AddCommand(scripts.Cmdscripts())
	rootCmd.AddCommand(search.Cmdsearch())
	rootCmd.AddCommand(stations.Cmdstations())
	rootCmd.AddCommand(systempresences.Cmdsystempresences())
	rootCmd.AddCommand(timezones.Cmdtimezones())
	rootCmd.AddCommand(tokens.Cmdtokens())
	rootCmd.AddCommand(userrecordings.Cmduserrecordings())
	rootCmd.AddCommand(users.Cmdusers())
	rootCmd.AddCommand(autopagination.Cmdautopagination())
	rootCmd.AddCommand(completion.Cmdcompletion())
	rootCmd.AddCommand(experimental.Cmdexperimental())
	rootCmd.AddCommand(logging.Cmdlogging())
	rootCmd.AddCommand(profiles.Cmdprofiles())
	rootCmd.AddCommand(alerting.Cmdalerting())
	rootCmd.AddCommand(analytics.Cmdanalytics())
	rootCmd.AddCommand(audits.Cmdaudits())
	rootCmd.AddCommand(authorization.Cmdauthorization())
	rootCmd.AddCommand(billing.Cmdbilling())
	rootCmd.AddCommand(certificate.Cmdcertificate())
	rootCmd.AddCommand(recording.Cmdrecording())
	rootCmd.AddCommand(externalcontacts.Cmdexternalcontacts())
	rootCmd.AddCommand(contentmanagement.Cmdcontentmanagement())
	rootCmd.AddCommand(uploads.Cmduploads())
	rootCmd.AddCommand(dataextensions.Cmddataextensions())
	rootCmd.AddCommand(outbound.Cmdoutbound())
	rootCmd.AddCommand(profile.Cmdprofile())
	rootCmd.AddCommand(organizations.Cmdorganizations())
	rootCmd.AddCommand(fax.Cmdfax())
	rootCmd.AddCommand(gamification.Cmdgamification())
	rootCmd.AddCommand(gdpr.Cmdgdpr())
	rootCmd.AddCommand(geolocations.Cmdgeolocations())
	rootCmd.AddCommand(journey.Cmdjourney())
	rootCmd.AddCommand(knowledge.Cmdknowledge())
	rootCmd.AddCommand(languageunderstanding.Cmdlanguageunderstanding())
	rootCmd.AddCommand(license.Cmdlicense())
	rootCmd.AddCommand(quality.Cmdquality())
	rootCmd.AddCommand(documentationfile.Cmddocumentationfile())
	rootCmd.AddCommand(webchat.Cmdwebchat())
	rootCmd.AddCommand(workforcemanagement.Cmdworkforcemanagement())
	rootCmd.AddCommand(notifications.Cmdnotifications())
	rootCmd.AddCommand(oauth.Cmdoauth())
	rootCmd.AddCommand(orgauthorization.Cmdorgauthorization())
	rootCmd.AddCommand(routing.Cmdrouting())
	rootCmd.AddCommand(recordings.Cmdrecordings())
	rootCmd.AddCommand(responsemanagement.Cmdresponsemanagement())
	rootCmd.AddCommand(scim.Cmdscim())
	rootCmd.AddCommand(speechandtextanalytics.Cmdspeechandtextanalytics())
	rootCmd.AddCommand(telephony.Cmdtelephony())
	rootCmd.AddCommand(configuration.Cmdconfiguration())
	rootCmd.AddCommand(textbots.Cmdtextbots())
	rootCmd.AddCommand(usage.Cmdusage())
	rootCmd.AddCommand(voicemail.Cmdvoicemail())
	rootCmd.AddCommand(architect.Cmdarchitect())
	rootCmd.AddCommand(webdeployments.Cmdwebdeployments())
	rootCmd.AddCommand(webmessaging.Cmdwebmessaging())
	rootCmd.AddCommand(coaching.Cmdcoaching())
	rootCmd.AddCommand(learning.Cmdlearning())
	rootCmd.AddCommand(widgets.Cmdwidgets())

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

