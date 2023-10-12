package integrations_speech_nuance_bots

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_nuance_bots", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots", )
	integrations_speech_nuance_botsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_nuance_bots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_nuance_botsCmd)
}

func Cmdintegrations_speech_nuance_bots() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "expand Valid values: variables, transferNodes, channels, locales")
	utils.AddFlag(getCmd.Flags(), "string", "targetChannel", "", "targetChannel Valid values: digital, voice")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/{botId}", utils.FormatPermissions([]string{ "integrations:integration:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/{botId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/NuanceBot"
      }
    }
  }
}`)
	integrations_speech_nuance_botsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "bool", "onlyRegisteredBots", "true", "Limit bots to the ones configured for Genesys Cloud usage")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots", utils.FormatPermissions([]string{ "integrations:integration:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	integrations_speech_nuance_botsCmd.AddCommand(listCmd)
	return integrations_speech_nuance_botsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [nuanceIntegrationId] [botId]",
	Short: "Get a Nuance bot in the specified Integration",
	Long:  "Get a Nuance bot in the specified Integration",
	Args:  utils.DetermineArgs([]string{ "nuanceIntegrationId", "botId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/{botId}"
		nuanceIntegrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{nuanceIntegrationId}", fmt.Sprintf("%v", nuanceIntegrationId), -1)
		botId, args := args[0], args[1:]
		path = strings.Replace(path, "{botId}", fmt.Sprintf("%v", botId), -1)

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		targetChannel := utils.GetFlag(cmd.Flags(), "string", "targetChannel")
		if targetChannel != "" {
			queryParams["targetChannel"] = targetChannel
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list [nuanceIntegrationId]",
	Short: "Get a list of Nuance bots available in the specified Integration",
	Long:  "Get a list of Nuance bots available in the specified Integration",
	Args:  utils.DetermineArgs([]string{ "nuanceIntegrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots"
		nuanceIntegrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{nuanceIntegrationId}", fmt.Sprintf("%v", nuanceIntegrationId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		onlyRegisteredBots := utils.GetFlag(cmd.Flags(), "bool", "onlyRegisteredBots")
		if onlyRegisteredBots != "" {
			queryParams["onlyRegisteredBots"] = onlyRegisteredBots
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "list"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
