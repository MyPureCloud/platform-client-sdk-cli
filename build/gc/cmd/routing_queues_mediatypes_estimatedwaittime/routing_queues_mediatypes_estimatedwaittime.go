package routing_queues_mediatypes_estimatedwaittime

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
	Description = utils.FormatUsageDescription("routing_queues_mediatypes_estimatedwaittime", "SWAGGER_OVERRIDE_/api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime", )
	routing_queues_mediatypes_estimatedwaittimeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_queues_mediatypes_estimatedwaittime"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_queues_mediatypes_estimatedwaittimeCmd)
}

func Cmdrouting_queues_mediatypes_estimatedwaittime() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime", utils.FormatPermissions([]string{ "routing:queue:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/EstimatedWaitTimePredictions"
  }
}`)
	routing_queues_mediatypes_estimatedwaittimeCmd.AddCommand(getCmd)
	
	return routing_queues_mediatypes_estimatedwaittimeCmd
}

var getCmd = &cobra.Command{
	Use:   "get [queueId] [mediaType]",
	Short: "Get Estimated Wait Time",
	Long:  "Get Estimated Wait Time",
	Args:  utils.DetermineArgs([]string{ "queueId", "mediaType", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		mediaType, args := args[0], args[1:]
		path = strings.Replace(path, "{mediaType}", fmt.Sprintf("%v", mediaType), -1)

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
