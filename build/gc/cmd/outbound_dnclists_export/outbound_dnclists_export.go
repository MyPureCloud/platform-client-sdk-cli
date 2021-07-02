package outbound_dnclists_export

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
	Description = utils.FormatUsageDescription("outbound_dnclists_export", "SWAGGER_OVERRIDE_/api/v2/outbound/dnclists/{dncListId}/export", "SWAGGER_OVERRIDE_/api/v2/outbound/dnclists/{dncListId}/export", )
	outbound_dnclists_exportCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_dnclists_export"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_dnclists_exportCmd)
}

func Cmdoutbound_dnclists_export() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/dnclists/{dncListId}/export", utils.FormatPermissions([]string{ "outbound:dnc:view", "outbound:dncList:view",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/dnclists/{dncListId}/export")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DomainEntityRef&quot;
  }
}`)
	outbound_dnclists_exportCmd.AddCommand(createCmd)
	
	utils.AddFlag(getCmd.Flags(), "string", "download", "false", "Redirect to download uri")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/dnclists/{dncListId}/export", utils.FormatPermissions([]string{ "outbound:dnc:view", "outbound:dncList:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/dnclists/{dncListId}/export")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExportUri&quot;
  }
}`)
	outbound_dnclists_exportCmd.AddCommand(getCmd)
	
	return outbound_dnclists_exportCmd
}

var createCmd = &cobra.Command{
	Use:   "create [dncListId]",
	Short: "Initiate the export of a dnc list.",
	Long:  "Initiate the export of a dnc list.",
	Args:  utils.DetermineArgs([]string{ "dncListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/dnclists/{dncListId}/export"
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
var getCmd = &cobra.Command{
	Use:   "get [dncListId]",
	Short: "Get the URI of a DNC list export.",
	Long:  "Get the URI of a DNC list export.",
	Args:  utils.DetermineArgs([]string{ "dncListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/dnclists/{dncListId}/export"
		dncListId, args := args[0], args[1:]
		path = strings.Replace(path, "{dncListId}", fmt.Sprintf("%v", dncListId), -1)

		download := utils.GetFlag(cmd.Flags(), "string", "download")
		if download != "" {
			queryParams["download"] = download
		}
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
