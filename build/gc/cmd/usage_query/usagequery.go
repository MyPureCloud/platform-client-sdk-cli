package usage_query

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"

	"github.com/spf13/cobra"
)

//Example of a custom API
func init() {
	createCmd = queryUsageCmd
}

var (
	usageQueryOperation = models.HandWrittenOperation{
		Path:   "/api/v2/usage/query",
		Method: http.MethodPost,
	}
	usageQueryResultsOperation = models.HandWrittenOperation{
		Path:   "/api/v2/usage/query/{executionId}/results",
		Method: http.MethodGet,
	}
)

var queryUsageCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a query for Genesys Cloud API Usage",
	Long:  `Creates a query for Genesys Cloud API Usage`,
	Args:  cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		timeout, _ := cmd.Flags().GetInt("timeout")

		retryFunc := CommandService.DetermineAction(usageQueryOperation.Method, usageQueryOperation.Path, cmd, "")
		results, err := retryFunc(nil)
		if err != nil {
			logger.Fatal(err)
		}

		usageSubmitResponse := &models.UsageSubmitResponse{}
		if err := json.Unmarshal([]byte(results), usageSubmitResponse); err != nil {
			logger.Fatal("Error occurred while unmarshalling data in usage query command (usageSubmitResponse)", err)
		}

		if timeout > 0 {
			usageQueryResponse := &models.UsageQueryResponse{}
			totalIterationsAllowed := (timeout * 60 / 5) //Number of minutes to wait * 60 for seconds/5 (which is the number of second to wait between calls)
			currentIteration := 1

			for true {
				path := usageQueryResultsOperation.Path
				targetURI := strings.Replace(path, "{executionId}", fmt.Sprintf("%v", usageSubmitResponse.ExecutionID), -1)
				retryFunc := CommandService.DetermineAction(usageQueryResultsOperation.Method, targetURI, cmd, "")
				rawData, commandErr := retryFunc(nil)
				if commandErr != nil {
					logger.Fatal(commandErr)
				}

				if err = json.Unmarshal([]byte(rawData), usageQueryResponse); err != nil {
					logger.Fatal("Error occurred while unmarshalling data from usage query command (usageQueryResponse)", err)
				}

				if usageQueryResponse.QueryStatus == "Complete" {
					utils.Render(rawData)
					break
				} else {
					if currentIteration <= totalIterationsAllowed {
						time.Sleep(5 * time.Second)
						currentIteration++
					} else {
						logger.Fatalf("Waited for %d minutes to retrieve API Usage results.  Giving up", timeout)
					}
				}
			}
		} else {
			utils.Render(results)
		}
	},
}