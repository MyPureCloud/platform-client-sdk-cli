package flows_latestconfiguration

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
	Description = utils.FormatUsageDescription("flows_latestconfiguration", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/latestconfiguration", )
	flows_latestconfigurationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_latestconfiguration"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_latestconfigurationCmd)
}

func Cmdflows_latestconfiguration() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "deleted", "false", "Deleted flows")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}/latestconfiguration", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/{flowId}/latestconfiguration")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "object"
      }
    }
  }
}`)
	flows_latestconfigurationCmd.AddCommand(getCmd)
	return flows_latestconfigurationCmd
}

var getCmd = &cobra.Command{
	Use:   "get [flowId]",
	Short: "Get the latest configuration for flow",
	Long:  "Get the latest configuration for flow",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/latestconfiguration"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)

		deleted := utils.GetFlag(cmd.Flags(), "bool", "deleted")
		if deleted != "" {
			queryParams["deleted"] = deleted
		}
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
