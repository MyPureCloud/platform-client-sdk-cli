package journey_views_versions_jobs

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
	Description = utils.FormatUsageDescription("journey_views_versions_jobs", "SWAGGER_OVERRIDE_/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs", "SWAGGER_OVERRIDE_/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs", )
	journey_views_versions_jobsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_views_versions_jobs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_views_versions_jobsCmd)
}

func Cmdjourney_views_versions_jobs() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs", utils.FormatPermissions([]string{ "journey:viewsJobs:add",  }), utils.GenerateDevCentreLink("POST", "Journey", "/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Job submitted successfully.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/JourneyViewJob"
      }
    }
  }
}`)
	journey_views_versions_jobsCmd.AddCommand(createCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs/{jobId}", utils.FormatPermissions([]string{ "journey:viewsJobs:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs/{jobId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/JourneyViewJob"
      }
    }
  }
}`)
	journey_views_versions_jobsCmd.AddCommand(getCmd)
	return journey_views_versions_jobsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [viewId] [journeyVersionId]",
	Short: "Submit a job request for a journey view version.",
	Long:  "Submit a job request for a journey view version.",
	Args:  utils.DetermineArgs([]string{ "viewId", "journeyVersionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs"
		viewId, args := args[0], args[1:]
		path = strings.Replace(path, "{viewId}", fmt.Sprintf("%v", viewId), -1)
		journeyVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{journeyVersionId}", fmt.Sprintf("%v", journeyVersionId), -1)

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

		const opId = "create"
		const httpMethod = "POST"
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
var getCmd = &cobra.Command{
	Use:   "get [viewId] [journeyVersionId] [jobId]",
	Short: "Get the job for a journey view version.",
	Long:  "Get the job for a journey view version.",
	Args:  utils.DetermineArgs([]string{ "viewId", "journeyVersionId", "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/views/{viewId}/versions/{journeyVersionId}/jobs/{jobId}"
		viewId, args := args[0], args[1:]
		path = strings.Replace(path, "{viewId}", fmt.Sprintf("%v", viewId), -1)
		journeyVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{journeyVersionId}", fmt.Sprintf("%v", journeyVersionId), -1)
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
