package flows_history

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
	Description = utils.FormatUsageDescription("flows_history", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/history", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/history", )
	flows_historyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_history"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_historyCmd)
}

func Cmdflows_history() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/{flowId}/history", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/{flowId}/history")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Operation&quot;
  }
}`)
	flows_historyCmd.AddCommand(createCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "desc", "Sort order")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "timestamp", "Sort by Valid values: action, timestamp, user")
	utils.AddFlag(listCmd.Flags(), "[]string", "action", "", "Flow actions to include (omit to include all) Valid values: checkin, checkout, create, deactivate, debug, delete, publish, revert, save")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}/history/{historyId}", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/{flowId}/history/{historyId}")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	flows_historyCmd.AddCommand(listCmd)
	
	return flows_historyCmd
}

var createCmd = &cobra.Command{
	Use:   "create [flowId]",
	Short: "Generate flow history",
	Long:  "Generate flow history",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/history"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [flowId] [historyId]",
	Short: "Get generated flow history",
	Long:  "Get generated flow history",
	Args:  utils.DetermineArgs([]string{ "flowId", "historyId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/history/{historyId}"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)
		historyId, args := args[0], args[1:]
		path = strings.Replace(path, "{historyId}", fmt.Sprintf("%v", historyId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		action := utils.GetFlag(cmd.Flags(), "[]string", "action")
		if action != "" {
			queryParams["action"] = action
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
