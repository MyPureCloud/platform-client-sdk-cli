package signals_anomalydetection_queues_analysis

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
	Description = utils.FormatUsageDescription("signals_anomalydetection_queues_analysis", "SWAGGER_OVERRIDE_/api/v2/signals/anomalydetection/queues/{queueId}/analysis", )
	signals_anomalydetection_queues_analysisCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("signals_anomalydetection_queues_analysis"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(signals_anomalydetection_queues_analysisCmd)
}

func Cmdsignals_anomalydetection_queues_analysis() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "metrics", "", "The metrics by which to filter the analysis. When not specified, returns all metrics Valid values: nOffered, tTalkCompleteAverage, oAnsweredPercent, oAbandonPercent, oServiceLevel, tAnsweredMax, tAnsweredAverage, tAcwAverage, tHeldCompleteAverage, tHandleMax, tHandleAverage, nHandled, nTransferred")
	utils.AddFlag(getCmd.Flags(), "[]string", "mediaTypes", "", "The media types by which to filter the analysis. When not specified, returns analysis for Email, Message, and Voice Valid values: Email, Message, Voice")
	utils.AddFlag(getCmd.Flags(), "string", "startTime", "", "The start of the time range for the analysis, in ISO-8601 format. If not specified, defaults to 48 hours before the most recent UTC midnight")
	utils.AddFlag(getCmd.Flags(), "string", "endTime", "", "The end of the time range for the analysis, in ISO-8601 format. If not specified, defaults to 48 hours after the most recent UTC midnight")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/signals/anomalydetection/queues/{queueId}/analysis", utils.FormatPermissions([]string{ "signals:queueAnalysis:view",  }), utils.GenerateDevCentreLink("GET", "Anomaly Detection", "/api/v2/signals/anomalydetection/queues/{queueId}/analysis")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AnomalyDetectionQueueAnalysisResponse"
      }
    }
  }
}`)
	signals_anomalydetection_queues_analysisCmd.AddCommand(getCmd)
	return signals_anomalydetection_queues_analysisCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [queueId]",
	Short: "Get forecasted bounds and analysis for the specified queue",
	Long:  "Get forecasted bounds and analysis for the specified queue",
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/signals/anomalydetection/queues/{queueId}/analysis"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

		metrics := utils.GetFlag(cmd.Flags(), "[]string", "metrics")
		if metrics != "" {
			queryParams["metrics"] = metrics
		}
		mediaTypes := utils.GetFlag(cmd.Flags(), "[]string", "mediaTypes")
		if mediaTypes != "" {
			queryParams["mediaTypes"] = mediaTypes
		}
		startTime := utils.GetFlag(cmd.Flags(), "string", "startTime")
		if startTime != "" {
			queryParams["startTime"] = startTime
		}
		endTime := utils.GetFlag(cmd.Flags(), "string", "endTime")
		if endTime != "" {
			queryParams["endTime"] = endTime
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

		headerParams := make(map[string]string)
		// to determine the Content-Type header
		localVarHttpContentTypes := []string{ "application/json",  }
		// set Content-Type header
		localVarHttpContentType := utils.SelectHeaderContentType(localVarHttpContentTypes)
		if localVarHttpContentType != "" {
			headerParams["Content-Type"] = localVarHttpContentType
		}
		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{
			"application/json",
		}
		// set Accept header
		localVarHttpHeaderAccept := utils.SelectHeaderAccept(localVarHttpHeaderAccepts)
		if localVarHttpHeaderAccept != "" {
			headerParams["Accept"] = localVarHttpHeaderAccept
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, headerParams, cmd, opId)
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
