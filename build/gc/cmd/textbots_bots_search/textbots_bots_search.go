package textbots_bots_search

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
	Description = utils.FormatUsageDescription("textbots_bots_search", "SWAGGER_OVERRIDE_/api/v2/textbots/bots/search", )
	textbots_bots_searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("textbots_bots_search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(textbots_bots_searchCmd)
}

func Cmdtextbots_bots_search() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "botType", "", "Bot types Valid values: GenesysBotConnector, GenesysDialogEngine, AmazonLex, GoogleDialogFlowES, GoogleDialogFlowCX, NuanceDlg, GenesysBotFlow, GenesysDigitalBotFlow, GenesysVoiceSurveyFlow, GenesysDigitalBotConnector")
	utils.AddFlag(getCmd.Flags(), "string", "botName", "", "Bot name")
	utils.AddFlag(getCmd.Flags(), "[]string", "botId", "", "Bot IDs. Maximum of 50")
	utils.AddFlag(getCmd.Flags(), "bool", "virtualAgentEnabled", "", "Include or exclude virtual agent flows, only applies to GenesysBotFlows or GenesysDigitalBotFlows")
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "25", "The maximum results to return. Maximum of 100")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/textbots/bots/search", utils.FormatPermissions([]string{ "integrations:integration:view",  }), utils.GenerateDevCentreLink("GET", "Textbots", "/api/v2/textbots/bots/search")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BotSearchResponseEntityListing"
      }
    }
  }
}`)
	textbots_bots_searchCmd.AddCommand(getCmd)
	return textbots_bots_searchCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Find bots using the currently configured friendly name or ID.",
	Long:  "Find bots using the currently configured friendly name or ID.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/textbots/bots/search"

		botType := utils.GetFlag(cmd.Flags(), "[]string", "botType")
		if botType != "" {
			queryParams["botType"] = botType
		}
		botName := utils.GetFlag(cmd.Flags(), "string", "botName")
		if botName != "" {
			queryParams["botName"] = botName
		}
		botId := utils.GetFlag(cmd.Flags(), "[]string", "botId")
		if botId != "" {
			queryParams["botId"] = botId
		}
		virtualAgentEnabled := utils.GetFlag(cmd.Flags(), "bool", "virtualAgentEnabled")
		if virtualAgentEnabled != "" {
			queryParams["virtualAgentEnabled"] = virtualAgentEnabled
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
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
