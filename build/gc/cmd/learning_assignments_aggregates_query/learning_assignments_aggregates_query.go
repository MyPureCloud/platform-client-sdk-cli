package learning_assignments_aggregates_query

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("learning_assignments_aggregates_query", "SWAGGER_OVERRIDE_/api/v2/learning/assignments/aggregates/query", )
	learning_assignments_aggregates_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_assignments_aggregates_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_assignments_aggregates_queryCmd)
}

func Cmdlearning_assignments_aggregates_query() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/learning/assignments/aggregates/query", utils.FormatPermissions([]string{ "learning:assignment:view",  }), utils.GenerateDevCentreLink("POST", "Learning", "/api/v2/learning/assignments/aggregates/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Aggregate Request&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LearningAssignmentAggregateParam&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Query completed successfully&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LearningAssignmentAggregateResponse&quot;
  }
}`)
	learning_assignments_aggregates_queryCmd.AddCommand(createCmd)
	
	return learning_assignments_aggregates_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Retrieve aggregated assignment data",
	Long:  "Retrieve aggregated assignment data",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments/aggregates/query"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
