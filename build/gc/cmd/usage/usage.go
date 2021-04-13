package usage

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
	usageCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud usage", "", ),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud usage`, "", ),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usageCmd)
}

func Cmdusage() *cobra.Command { 
	resultsCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", resultsCmd.UsageTemplate(), "GET", "/api/v2/usage/query/{executionId}/results", utils.FormatPermissions([]string{ "oauth:client:view",  })))
	utils.AddFileFlagIfUpsert(resultsCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(resultsCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ApiUsageQueryResult&quot;
  }
}`)
	usageCmd.AddCommand(resultsCmd)
	
	return usageCmd
}

var resultsCmd = &cobra.Command{
	Use:   "results [executionId]",
	Short: "Get the results of a usage query",
	Long:  `Get the results of a usage query`,
	Args:  utils.DetermineArgs([]string{ "executionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/usage/query/{executionId}/results"
		executionId, args := args[0], args[1:]
		path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
