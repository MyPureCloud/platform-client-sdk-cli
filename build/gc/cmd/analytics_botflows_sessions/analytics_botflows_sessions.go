package analytics_botflows_sessions

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
	Description = utils.FormatUsageDescription("analytics_botflows_sessions", "SWAGGER_OVERRIDE_/api/v2/analytics/botflows/{botFlowId}/sessions", )
	analytics_botflows_sessionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_botflows_sessions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_botflows_sessionsCmd)
}

func Cmdanalytics_botflows_sessions() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the ID of the last item in the list of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "50", "Max number of entities to return. Maximum of 250")
	utils.AddFlag(listCmd.Flags(), "string", "interval", "", "Date range filter based on the date the individual resources were completed. UTC is the default if no TZ is supplied, however alternate timezones can be used e.g: `2022-11-22T09:11:11.111+08:00/2022-11-30T07:17:44.586-07`. . Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "string", "botResultCategories", "", "Optional case-insensitive comma separated list of Bot Result Categories to filter sessions by. Valid values: Unknown, UserExit, BotExit, Error, RecognitionFailure, UserDisconnect, BotDisconnect, SessionExpired, Transfer")
	utils.AddFlag(listCmd.Flags(), "string", "endLanguage", "", "Optional case-insensitive language code to filter sessions by the language the sessions ended in.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/analytics/botflows/{botFlowId}/sessions", utils.FormatPermissions([]string{ "analytics:botFlowSession:view",  }), utils.GenerateDevCentreLink("GET", "Analytics", "/api/v2/analytics/botflows/{botFlowId}/sessions")))
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
	analytics_botflows_sessionsCmd.AddCommand(listCmd)
	return analytics_botflows_sessionsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [botFlowId]",
	Short: "Get Bot Flow Sessions.",
	Long:  "Get Bot Flow Sessions.",
	Args:  utils.DetermineArgs([]string{ "botFlowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/botflows/{botFlowId}/sessions"
		botFlowId, args := args[0], args[1:]
		path = strings.Replace(path, "{botFlowId}", fmt.Sprintf("%v", botFlowId), -1)

		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		interval := utils.GetFlag(cmd.Flags(), "string", "interval")
		if interval != "" {
			queryParams["interval"] = interval
		}
		botResultCategories := utils.GetFlag(cmd.Flags(), "string", "botResultCategories")
		if botResultCategories != "" {
			queryParams["botResultCategories"] = botResultCategories
		}
		endLanguage := utils.GetFlag(cmd.Flags(), "string", "endLanguage")
		if endLanguage != "" {
			queryParams["endLanguage"] = endLanguage
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
