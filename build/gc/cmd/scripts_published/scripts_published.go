package scripts_published

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
	Description = utils.FormatUsageDescription("scripts_published", "SWAGGER_OVERRIDE_/api/v2/scripts/published", "SWAGGER_OVERRIDE_/api/v2/scripts/published", )
	scripts_publishedCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scripts_published"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scripts_publishedCmd)
}

func Cmdscripts_published() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "scriptDataVersion", "", "Advanced usage - controls the data version of the script")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/scripts/published/{scriptId}", utils.FormatPermissions([]string{ "scripter:publishedScript:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Script&quot;
  }
}`)
	scripts_publishedCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "expand", "", "Expand")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name filter")
	utils.AddFlag(listCmd.Flags(), "string", "feature", "", "Feature filter")
	utils.AddFlag(listCmd.Flags(), "string", "flowId", "", "Secure flow id filter")
	utils.AddFlag(listCmd.Flags(), "string", "scriptDataVersion", "", "Advanced usage - controls the data version of the script")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/scripts/published", utils.FormatPermissions([]string{ "scripter:publishedScript:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	scripts_publishedCmd.AddCommand(listCmd)
	
	return scripts_publishedCmd
}

var getCmd = &cobra.Command{
	Use:   "get [scriptId]",
	Short: "Get the published script.",
	Long:  "Get the published script.",
	Args:  utils.DetermineArgs([]string{ "scriptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/scripts/published/{scriptId}"
		scriptId, args := args[0], args[1:]
		path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)

		scriptDataVersion := utils.GetFlag(cmd.Flags(), "string", "scriptDataVersion")
		if scriptDataVersion != "" {
			queryParams["scriptDataVersion"] = scriptDataVersion
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the published scripts.",
	Long:  "Get the published scripts.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/scripts/published"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		feature := utils.GetFlag(cmd.Flags(), "string", "feature")
		if feature != "" {
			queryParams["feature"] = feature
		}
		flowId := utils.GetFlag(cmd.Flags(), "string", "flowId")
		if flowId != "" {
			queryParams["flowId"] = flowId
		}
		scriptDataVersion := utils.GetFlag(cmd.Flags(), "string", "scriptDataVersion")
		if scriptDataVersion != "" {
			queryParams["scriptDataVersion"] = scriptDataVersion
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
