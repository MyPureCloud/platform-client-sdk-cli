package quality_evaluations_query

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
	Description = utils.FormatUsageDescription("quality_evaluations_query", "SWAGGER_OVERRIDE_/api/v2/quality/evaluations/query", )
	quality_evaluations_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_evaluations_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_evaluations_queryCmd)
}

func Cmdquality_evaluations_query() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "variable name requested by expand list")
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "next page token")
	utils.AddFlag(listCmd.Flags(), "string", "previousPage", "", "Previous page token")
	utils.AddFlag(listCmd.Flags(), "string", "conversationId", "", "conversationId specified")
	utils.AddFlag(listCmd.Flags(), "string", "agentUserId", "", "user id of the agent")
	utils.AddFlag(listCmd.Flags(), "string", "evaluatorUserId", "", "evaluator user id")
	utils.AddFlag(listCmd.Flags(), "string", "queueId", "", "queue id")
	utils.AddFlag(listCmd.Flags(), "string", "startTime", "", "start time of the evaluation query")
	utils.AddFlag(listCmd.Flags(), "string", "endTime", "", "end time of the evaluation query")
	utils.AddFlag(listCmd.Flags(), "[]string", "evaluationState", "", "")
	utils.AddFlag(listCmd.Flags(), "bool", "isReleased", "", "the evaluation has been released")
	utils.AddFlag(listCmd.Flags(), "bool", "agentHasRead", "", "agent has the evaluation")
	utils.AddFlag(listCmd.Flags(), "bool", "expandAnswerTotalScores", "", "get the total scores for evaluations")
	utils.AddFlag(listCmd.Flags(), "int", "maximum", "", "maximum")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "", "sort order options for agentUserId or evaluatorUserId query. Valid options are `a`, `asc`, `ascending`, `d`, `desc`, `descending`")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/quality/evaluations/query", utils.FormatPermissions([]string{ "quality:evaluation:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/evaluations/query")))
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
	quality_evaluations_queryCmd.AddCommand(listCmd)
	return quality_evaluations_queryCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Queries Evaluations and returns a paged list",
	Long:  "Queries Evaluations and returns a paged list",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/evaluations/query"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		nextPage := utils.GetFlag(cmd.Flags(), "string", "nextPage")
		if nextPage != "" {
			queryParams["nextPage"] = nextPage
		}
		previousPage := utils.GetFlag(cmd.Flags(), "string", "previousPage")
		if previousPage != "" {
			queryParams["previousPage"] = previousPage
		}
		conversationId := utils.GetFlag(cmd.Flags(), "string", "conversationId")
		if conversationId != "" {
			queryParams["conversationId"] = conversationId
		}
		agentUserId := utils.GetFlag(cmd.Flags(), "string", "agentUserId")
		if agentUserId != "" {
			queryParams["agentUserId"] = agentUserId
		}
		evaluatorUserId := utils.GetFlag(cmd.Flags(), "string", "evaluatorUserId")
		if evaluatorUserId != "" {
			queryParams["evaluatorUserId"] = evaluatorUserId
		}
		queueId := utils.GetFlag(cmd.Flags(), "string", "queueId")
		if queueId != "" {
			queryParams["queueId"] = queueId
		}
		startTime := utils.GetFlag(cmd.Flags(), "string", "startTime")
		if startTime != "" {
			queryParams["startTime"] = startTime
		}
		endTime := utils.GetFlag(cmd.Flags(), "string", "endTime")
		if endTime != "" {
			queryParams["endTime"] = endTime
		}
		evaluationState := utils.GetFlag(cmd.Flags(), "[]string", "evaluationState")
		if evaluationState != "" {
			queryParams["evaluationState"] = evaluationState
		}
		isReleased := utils.GetFlag(cmd.Flags(), "bool", "isReleased")
		if isReleased != "" {
			queryParams["isReleased"] = isReleased
		}
		agentHasRead := utils.GetFlag(cmd.Flags(), "bool", "agentHasRead")
		if agentHasRead != "" {
			queryParams["agentHasRead"] = agentHasRead
		}
		expandAnswerTotalScores := utils.GetFlag(cmd.Flags(), "bool", "expandAnswerTotalScores")
		if expandAnswerTotalScores != "" {
			queryParams["expandAnswerTotalScores"] = expandAnswerTotalScores
		}
		maximum := utils.GetFlag(cmd.Flags(), "int", "maximum")
		if maximum != "" {
			queryParams["maximum"] = maximum
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
