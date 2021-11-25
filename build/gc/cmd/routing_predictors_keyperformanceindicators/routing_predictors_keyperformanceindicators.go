package routing_predictors_keyperformanceindicators

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
	Description = utils.FormatUsageDescription("routing_predictors_keyperformanceindicators", "SWAGGER_OVERRIDE_/api/v2/routing/predictors/keyperformanceindicators", )
	routing_predictors_keyperformanceindicatorsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_predictors_keyperformanceindicators"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_predictors_keyperformanceindicatorsCmd)
}

func Cmdrouting_predictors_keyperformanceindicators() *cobra.Command { 
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/predictors/keyperformanceindicators", utils.FormatPermissions([]string{ "routing:keyPerformanceIndicator:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/predictors/keyperformanceindicators")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/KeyPerformanceIndicator"
    }
  }
}`)
	routing_predictors_keyperformanceindicatorsCmd.AddCommand(listCmd)
	
	return routing_predictors_keyperformanceindicatorsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of Key Performance Indicators available for the predictors.",
	Long:  "Get a list of Key Performance Indicators available for the predictors.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/predictors/keyperformanceindicators"

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
