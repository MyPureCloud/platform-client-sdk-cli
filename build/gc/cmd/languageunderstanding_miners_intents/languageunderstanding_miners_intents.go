package languageunderstanding_miners_intents

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("languageunderstanding_miners_intents", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/miners/{minerId}/intents", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/miners/{minerId}/intents", )
	languageunderstanding_miners_intentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_miners_intents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_miners_intentsCmd)
}

func Cmdlanguageunderstanding_miners_intents() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "Option to fetch utterances")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}", utils.FormatPermissions([]string{ "languageUnderstanding:minerIntent:view",  }), utils.GenerateDevCentreLink("GET", "Language Understanding", "/api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/MinerIntent&quot;
  }
}`)
	languageunderstanding_miners_intentsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "expand", "", "Option to fetch utterances.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/miners/{minerId}/intents", utils.FormatPermissions([]string{ "languageUnderstanding:minerIntent:view",  }), utils.GenerateDevCentreLink("GET", "Language Understanding", "/api/v2/languageunderstanding/miners/{minerId}/intents")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	languageunderstanding_miners_intentsCmd.AddCommand(listCmd)
	
	return languageunderstanding_miners_intentsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [minerId] [intentId]",
	Short: "Get information about a mined intent",
	Long:  "Get information about a mined intent",
	Args:  utils.DetermineArgs([]string{ "minerId", "intentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/miners/{minerId}/intents/{intentId}"
		minerId, args := args[0], args[1:]
		path = strings.Replace(path, "{minerId}", fmt.Sprintf("%v", minerId), -1)
		intentId, args := args[0], args[1:]
		path = strings.Replace(path, "{intentId}", fmt.Sprintf("%v", intentId), -1)

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list [minerId]",
	Short: "Retrieve a list of mined intents.",
	Long:  "Retrieve a list of mined intents.",
	Args:  utils.DetermineArgs([]string{ "minerId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/miners/{minerId}/intents"
		minerId, args := args[0], args[1:]
		path = strings.Replace(path, "{minerId}", fmt.Sprintf("%v", minerId), -1)

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
