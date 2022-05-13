package routing_queues_comparisonperiods

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
	Description = utils.FormatUsageDescription("routing_queues_comparisonperiods", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/comparisonperiods", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/comparisonperiods", )
	routing_queues_comparisonperiodsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_queues_comparisonperiods"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_queues_comparisonperiodsCmd)
}

func Cmdrouting_queues_comparisonperiods() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/routing/queues/{queueId}/comparisonperiods/{comparisonPeriodId}", utils.FormatPermissions([]string{ "routing:comparisonPeriod:view", "routing:queue:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/queues/{queueId}/comparisonperiods/{comparisonPeriodId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ComparisonPeriod"
      }
    }
  }
}`)
	routing_queues_comparisonperiodsCmd.AddCommand(getCmd)

	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/queues/{queueId}/comparisonperiods", utils.FormatPermissions([]string{ "routing:comparisonPeriod:view", "routing:queue:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/queues/{queueId}/comparisonperiods")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ComparisonPeriodListing"
      }
    }
  }
}`)
	routing_queues_comparisonperiodsCmd.AddCommand(listCmd)
	return routing_queues_comparisonperiodsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [queueId] [comparisonPeriodId]",
	Short: "Get a Comparison Period.",
	Long:  "Get a Comparison Period.",
	Args:  utils.DetermineArgs([]string{ "queueId", "comparisonPeriodId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/comparisonperiods/{comparisonPeriodId}"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		comparisonPeriodId, args := args[0], args[1:]
		path = strings.Replace(path, "{comparisonPeriodId}", fmt.Sprintf("%v", comparisonPeriodId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [queueId]",
	Short: "Get list of comparison periods",
	Long:  "Get list of comparison periods",
	Args:  utils.DetermineArgs([]string{ "queueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/comparisonperiods"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)

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
