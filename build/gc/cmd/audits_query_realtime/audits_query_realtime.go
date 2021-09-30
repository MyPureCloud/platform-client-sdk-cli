package audits_query_realtime

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
	Description = utils.FormatUsageDescription("audits_query_realtime", "SWAGGER_OVERRIDE_/api/v2/audits/query/realtime", )
	audits_query_realtimeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("audits_query_realtime"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(audits_query_realtimeCmd)
}

func Cmdaudits_query_realtime() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: user")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/audits/query/realtime", utils.FormatPermissions([]string{ "audits:audit:view",  }), utils.GenerateDevCentreLink("POST", "Audit", "/api/v2/audits/query/realtime")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "query",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/AuditRealtimeQueryRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/AuditRealtimeQueryResultsResponse"
  }
}`)
	audits_query_realtimeCmd.AddCommand(createCmd)
	
	return audits_query_realtimeCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This endpoint will only retrieve 14 days worth of audits for certain services. Please use /query to get a full list and older audits.",
	Long:  "This endpoint will only retrieve 14 days worth of audits for certain services. Please use /query to get a full list and older audits.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Auditrealtimequeryrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/audits/query/realtime"


		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
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
