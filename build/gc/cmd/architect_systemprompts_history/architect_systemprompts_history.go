package architect_systemprompts_history

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
	Description = utils.FormatUsageDescription("architect_systemprompts_history", "SWAGGER_OVERRIDE_/api/v2/architect/systemprompts/{promptId}/history", "SWAGGER_OVERRIDE_/api/v2/architect/systemprompts/{promptId}/history", )
	architect_systemprompts_historyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_systemprompts_history"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_systemprompts_historyCmd)
}

func Cmdarchitect_systemprompts_history() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/architect/systemprompts/{promptId}/history", utils.FormatPermissions([]string{ "architect:systemPrompt:view",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/architect/systemprompts/{promptId}/history")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Operation&quot;
  }
}`)
	architect_systemprompts_historyCmd.AddCommand(createCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "desc", "Sort order")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "timestamp", "Sort by Valid values: action, timestamp, user")
	utils.AddFlag(listCmd.Flags(), "[]string", "action", "", "Flow actions to include (omit to include all) Valid values: checkin, checkout, create, deactivate, debug, delete, publish, revert, save")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/architect/systemprompts/{promptId}/history/{historyId}", utils.FormatPermissions([]string{ "architect:systemPrompt:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/architect/systemprompts/{promptId}/history/{historyId}")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	architect_systemprompts_historyCmd.AddCommand(listCmd)
	
	return architect_systemprompts_historyCmd
}

var createCmd = &cobra.Command{
	Use:   "create [promptId]",
	Short: "Generate system prompt history",
	Long:  "Generate system prompt history",
	Args:  utils.DetermineArgs([]string{ "promptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/architect/systemprompts/{promptId}/history"
		promptId, args := args[0], args[1:]
		path = strings.Replace(path, "{promptId}", fmt.Sprintf("%v", promptId), -1)

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
	Use:   "list [promptId] [historyId]",
	Short: "Get generated prompt history",
	Long:  "Get generated prompt history",
	Args:  utils.DetermineArgs([]string{ "promptId", "historyId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/architect/systemprompts/{promptId}/history/{historyId}"
		promptId, args := args[0], args[1:]
		path = strings.Replace(path, "{promptId}", fmt.Sprintf("%v", promptId), -1)
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
