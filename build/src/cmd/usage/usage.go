package usage
import (
	"fmt"
	"gc/retry"
	"gc/utils"
	"gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"strings"
)

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Manages Genesys Cloud usage",
	Long:  `Manages Genesys Cloud usage`,
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(usageCmd)
}

func Cmdusage() *cobra.Command { 
	utils.AddFileFlagIfUpsert(resultsCmd.Flags(), "GET")
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

		retryFunc := CommandService.DetermineAction("GET", "results", urlString)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		utils.Render(results)
	},
}
