package flows_versions

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
	Description = utils.FormatUsageDescription("flows_versions", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions", )
	flows_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_versionsCmd)
}

func Cmdflows_versions() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/{flowId}/versions", utils.FormatPermissions([]string{ "architect:flow:edit",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;object&quot;,
    &quot;additionalProperties&quot; : {
      &quot;type&quot; : &quot;object&quot;,
      &quot;properties&quot; : { }
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/FlowVersion&quot;
  }
}`)
	flows_versionsCmd.AddCommand(createCmd)
	
	utils.AddFlag(getCmd.Flags(), "string", "deleted", "", "Deleted flows")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}/versions/{versionId}", utils.FormatPermissions([]string{ "architect:flow:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/FlowVersion&quot;
  }
}`)
	flows_versionsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "bool", "deleted", "", "Include Deleted flows")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}/versions", utils.FormatPermissions([]string{ "architect:flow:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	flows_versionsCmd.AddCommand(listCmd)
	
	return flows_versionsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [flowId]",
	Short: "Create flow version",
	Long:  "Create flow version",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/versions"
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
var getCmd = &cobra.Command{
	Use:   "get [flowId] [versionId]",
	Short: "Get flow version",
	Long:  "Get flow version",
	Args:  utils.DetermineArgs([]string{ "flowId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/versions/{versionId}"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)

		deleted := utils.GetFlag(cmd.Flags(), "string", "deleted")
		if deleted != "" {
			queryParams["deleted"] = deleted
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
	Use:   "list [flowId]",
	Short: "Get flow version list",
	Long:  "Get flow version list",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/versions"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		deleted := utils.GetFlag(cmd.Flags(), "bool", "deleted")
		if deleted != "" {
			queryParams["deleted"] = deleted
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
