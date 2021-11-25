package analytics_botflows_reportingturns

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
	Description = utils.FormatUsageDescription("analytics_botflows_reportingturns", "SWAGGER_OVERRIDE_/api/v2/analytics/botflows/{botFlowId}/reportingturns", )
	analytics_botflows_reportingturnsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_botflows_reportingturns"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_botflows_reportingturnsCmd)
}

func Cmdanalytics_botflows_reportingturns() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the ID of the last item in the list of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "50", "Max number of entities to return. Maximum of 250")
	utils.AddFlag(listCmd.Flags(), "string", "actionId", "", "Optional action ID to get the reporting turns associated to a particular flow action")
	utils.AddFlag(listCmd.Flags(), "string", "sessionId", "", "Optional session ID to get the reporting turns for a particular session")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/analytics/botflows/{botFlowId}/reportingturns", utils.FormatPermissions([]string{ "analytics:botFlowReportingTurn:view",  }), utils.GenerateDevCentreLink("GET", "Analytics", "/api/v2/analytics/botflows/{botFlowId}/reportingturns")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	analytics_botflows_reportingturnsCmd.AddCommand(listCmd)
	
	return analytics_botflows_reportingturnsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [botFlowId]",
	Short: "Get Reporting Turns.",
	Long:  "Get Reporting Turns.",
	Args:  utils.DetermineArgs([]string{ "botFlowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/botflows/{botFlowId}/reportingturns"
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
		actionId := utils.GetFlag(cmd.Flags(), "string", "actionId")
		if actionId != "" {
			queryParams["actionId"] = actionId
		}
		sessionId := utils.GetFlag(cmd.Flags(), "string", "sessionId")
		if sessionId != "" {
			queryParams["sessionId"] = sessionId
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
