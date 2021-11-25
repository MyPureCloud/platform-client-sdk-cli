package analytics_conversations_details_jobs_availability

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
	Description = utils.FormatUsageDescription("analytics_conversations_details_jobs_availability", "SWAGGER_OVERRIDE_/api/v2/analytics/conversations/details/jobs/availability", )
	analytics_conversations_details_jobs_availabilityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_conversations_details_jobs_availability"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_conversations_details_jobs_availabilityCmd)
}

func Cmdanalytics_conversations_details_jobs_availability() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/analytics/conversations/details/jobs/availability", utils.FormatPermissions([]string{ "analytics:conversationDetail:view",  }), utils.GenerateDevCentreLink("GET", "Analytics", "/api/v2/analytics/conversations/details/jobs/availability")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DataAvailabilityResponse"
  }
}`)
	analytics_conversations_details_jobs_availabilityCmd.AddCommand(getCmd)
	
	return analytics_conversations_details_jobs_availabilityCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Lookup the datalake availability date and time",
	Long:  "Lookup the datalake availability date and time",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/conversations/details/jobs/availability"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
