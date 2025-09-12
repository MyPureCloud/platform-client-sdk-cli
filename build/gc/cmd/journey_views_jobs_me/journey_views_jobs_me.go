package journey_views_jobs_me

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
	Description = utils.FormatUsageDescription("journey_views_jobs_me", "SWAGGER_OVERRIDE_/api/v2/journey/views/jobs/me", )
	journey_views_jobs_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_views_jobs_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_views_jobs_meCmd)
}

func Cmdjourney_views_jobs_me() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The number of the page to return")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Max number of entities to return")
	utils.AddFlag(listCmd.Flags(), "string", "interval", "", "An absolute timeframe for filtering the jobs, expressed as an ISO 8601 interval.")
	utils.AddFlag(listCmd.Flags(), "string", "statuses", "", "Job statuses to filter for")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/journey/views/jobs/me", utils.FormatPermissions([]string{ "journey:viewsJobs:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/views/jobs/me")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "Request completed successfully",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	journey_views_jobs_meCmd.AddCommand(listCmd)
	return journey_views_jobs_meCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get my jobs",
	Long:  "Get my jobs",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/views/jobs/me"

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		interval := utils.GetFlag(cmd.Flags(), "string", "interval")
		if interval != "" {
			queryParams["interval"] = interval
		}
		statuses := utils.GetFlag(cmd.Flags(), "string", "statuses")
		if statuses != "" {
			queryParams["statuses"] = statuses
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
