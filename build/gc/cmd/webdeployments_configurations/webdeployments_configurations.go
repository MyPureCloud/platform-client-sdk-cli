package webdeployments_configurations

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
	Description = utils.FormatUsageDescription("webdeployments_configurations", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations", )
	webdeployments_configurationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_configurations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_configurationsCmd)
}

func Cmdwebdeployments_configurations() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/webdeployments/configurations", utils.FormatPermissions([]string{ "webDeployments:configuration:add",  }), utils.GenerateDevCentreLink("POST", "Web Deployments", "/api/v2/webdeployments/configurations")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;configurationVersion&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebDeploymentConfigurationVersion&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebDeploymentConfigurationVersion&quot;
  }
}`)
	webdeployments_configurationsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/webdeployments/configurations/{configurationId}", utils.FormatPermissions([]string{ "webDeployments:configuration:delete",  }), utils.GenerateDevCentreLink("DELETE", "Web Deployments", "/api/v2/webdeployments/configurations/{configurationId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;The configuration versions were deleted successfully&quot;
}`)
	webdeployments_configurationsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(listCmd.Flags(), "bool", "showOnlyPublished", "false", "Get only configuration drafts with published versions")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/webdeployments/configurations", utils.FormatPermissions([]string{ "webDeployments:configuration:view",  }), utils.GenerateDevCentreLink("GET", "Web Deployments", "/api/v2/webdeployments/configurations")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebDeploymentConfigurationVersionEntityListing&quot;
  }
}`)
	webdeployments_configurationsCmd.AddCommand(listCmd)
	
	return webdeployments_configurationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a configuration draft",
	Long:  "Create a configuration draft",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations"

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
var deleteCmd = &cobra.Command{
	Use:   "delete [configurationId]",
	Short: "Delete all versions of a configuration",
	Long:  "Delete all versions of a configuration",
	Args:  utils.DetermineArgs([]string{ "configurationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations/{configurationId}"
		configurationId, args := args[0], args[1:]
		path = strings.Replace(path, "{configurationId}", fmt.Sprintf("%v", configurationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
	Short: "View configuration drafts",
	Long:  "View configuration drafts",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations"

		showOnlyPublished := utils.GetFlag(cmd.Flags(), "bool", "showOnlyPublished")
		if showOnlyPublished != "" {
			queryParams["showOnlyPublished"] = showOnlyPublished
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
