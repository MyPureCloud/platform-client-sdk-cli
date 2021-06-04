package telephony_providers_edges_edgegroups_edgetrunkbases

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_edgegroups_edgetrunkbases", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases", )
	telephony_providers_edges_edgegroups_edgetrunkbasesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_edgegroups_edgetrunkbases"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_edgegroups_edgetrunkbasesCmd)
}

func Cmdtelephony_providers_edges_edgegroups_edgetrunkbases() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/EdgeTrunkBase&quot;
  }
}`)
	telephony_providers_edges_edgegroups_edgetrunkbasesCmd.AddCommand(getCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;EdgeTrunkBase&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/EdgeTrunkBase&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/EdgeTrunkBase&quot;
  }
}`)
	telephony_providers_edges_edgegroups_edgetrunkbasesCmd.AddCommand(updateCmd)
	
	return telephony_providers_edges_edgegroups_edgetrunkbasesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [edgegroupId] [edgetrunkbaseId]",
	Short: "Gets the edge trunk base associated with the edge group",
	Long:  "Gets the edge trunk base associated with the edge group",
	Args:  utils.DetermineArgs([]string{ "edgegroupId", "edgetrunkbaseId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}"
		edgegroupId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgegroupId}", fmt.Sprintf("%v", edgegroupId), -1)
		edgetrunkbaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgetrunkbaseId}", fmt.Sprintf("%v", edgetrunkbaseId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [edgegroupId] [edgetrunkbaseId]",
	Short: "Update the edge trunk base associated with the edge group",
	Long:  "Update the edge trunk base associated with the edge group",
	Args:  utils.DetermineArgs([]string{ "edgegroupId", "edgetrunkbaseId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}"
		edgegroupId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgegroupId}", fmt.Sprintf("%v", edgegroupId), -1)
		edgetrunkbaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgetrunkbaseId}", fmt.Sprintf("%v", edgetrunkbaseId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
