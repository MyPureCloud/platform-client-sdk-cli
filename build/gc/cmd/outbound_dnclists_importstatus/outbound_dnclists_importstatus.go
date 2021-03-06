package outbound_dnclists_importstatus

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
	Description = utils.FormatUsageDescription("outbound_dnclists_importstatus", "SWAGGER_OVERRIDE_/api/v2/outbound/dnclists/{dncListId}/importstatus", )
	outbound_dnclists_importstatusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_dnclists_importstatus"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_dnclists_importstatusCmd)
}

func Cmdoutbound_dnclists_importstatus() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/dnclists/{dncListId}/importstatus", utils.FormatPermissions([]string{ "outbound:dncList:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/dnclists/{dncListId}/importstatus")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ImportStatus&quot;
  }
}`)
	outbound_dnclists_importstatusCmd.AddCommand(getCmd)
	
	return outbound_dnclists_importstatusCmd
}

var getCmd = &cobra.Command{
	Use:   "get [dncListId]",
	Short: "Get dialer dncList import status.",
	Long:  "Get dialer dncList import status.",
	Args:  utils.DetermineArgs([]string{ "dncListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/dnclists/{dncListId}/importstatus"
		dncListId, args := args[0], args[1:]
		path = strings.Replace(path, "{dncListId}", fmt.Sprintf("%v", dncListId), -1)

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
