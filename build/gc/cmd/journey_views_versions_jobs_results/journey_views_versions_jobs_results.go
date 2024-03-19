package journey_views_versions_jobs_results

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
	Description = utils.FormatUsageDescription("journey_views_versions_jobs_results", "SWAGGER_OVERRIDE_/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/jobs/{jobId}/results", )
	journey_views_versions_jobs_resultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_views_versions_jobs_results"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_views_versions_jobs_resultsCmd)
}

func Cmdjourney_views_versions_jobs_results() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/jobs/{jobId}/results", utils.FormatPermissions([]string{ "journey:viewsResults:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/jobs/{jobId}/results")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "Result received successfully",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/JourneyViewResult"
      }
    }
  }
}`)
	journey_views_versions_jobs_resultsCmd.AddCommand(getCmd)
	return journey_views_versions_jobs_resultsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [viewId] [journeyViewVersion] [jobId]",
	Short: "Get the result of a job for a journey view version.",
	Long:  "Get the result of a job for a journey view version.",
	Args:  utils.DetermineArgs([]string{ "viewId", "journeyViewVersion", "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/jobs/{jobId}/results"
		viewId, args := args[0], args[1:]
		path = strings.Replace(path, "{viewId}", fmt.Sprintf("%v", viewId), -1)
		journeyViewVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{journeyViewVersion}", fmt.Sprintf("%v", journeyViewVersion), -1)
		jobId, args := args[0], args[1:]
		path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)

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
