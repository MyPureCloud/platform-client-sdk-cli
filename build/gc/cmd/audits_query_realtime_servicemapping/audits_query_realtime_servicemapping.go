package audits_query_realtime_servicemapping

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
	Description = utils.FormatUsageDescription("audits_query_realtime_servicemapping", "SWAGGER_OVERRIDE_/api/v2/audits/query/realtime/servicemapping", )
	audits_query_realtime_servicemappingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("audits_query_realtime_servicemapping"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(audits_query_realtime_servicemappingCmd)
}

func Cmdaudits_query_realtime_servicemapping() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/audits/query/realtime/servicemapping", utils.FormatPermissions([]string{ "audits:audit:view",  }), utils.GenerateDevCentreLink("GET", "Audit", "/api/v2/audits/query/realtime/servicemapping")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/AuditQueryServiceMapping"
  }
}`)
	audits_query_realtime_servicemappingCmd.AddCommand(getCmd)
	
	return audits_query_realtime_servicemappingCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get service mapping information used in realtime audits.",
	Long:  "Get service mapping information used in realtime audits.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/audits/query/realtime/servicemapping"


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
