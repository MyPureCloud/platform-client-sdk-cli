package telephony_providers_edges_trunkbasesettings_availablemetabases

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_trunkbasesettings_availablemetabases", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/trunkbasesettings/availablemetabases", )
	telephony_providers_edges_trunkbasesettings_availablemetabasesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_trunkbasesettings_availablemetabases"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_trunkbasesettings_availablemetabasesCmd)
}

func Cmdtelephony_providers_edges_trunkbasesettings_availablemetabases() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "varType", "", " Valid values: EXTERNAL, PHONE, EDGE")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/trunkbasesettings/availablemetabases", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	telephony_providers_edges_trunkbasesettings_availablemetabasesCmd.AddCommand(listCmd)
	
	return telephony_providers_edges_trunkbasesettings_availablemetabasesCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of available makes and models to create a new Trunk Base Settings",
	Long:  "Get a list of available makes and models to create a new Trunk Base Settings",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/trunkbasesettings/availablemetabases"

		varType := utils.GetFlag(cmd.Flags(), "string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
