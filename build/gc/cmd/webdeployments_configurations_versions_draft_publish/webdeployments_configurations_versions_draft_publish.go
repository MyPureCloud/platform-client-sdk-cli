package webdeployments_configurations_versions_draft_publish

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
	Description = utils.FormatUsageDescription("webdeployments_configurations_versions_draft_publish", "SWAGGER_OVERRIDE_/api/v2/webdeployments/configurations/{configurationId}/versions/draft/publish", )
	webdeployments_configurations_versions_draft_publishCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webdeployments_configurations_versions_draft_publish"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webdeployments_configurations_versions_draft_publishCmd)
}

func Cmdwebdeployments_configurations_versions_draft_publish() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/webdeployments/configurations/{configurationId}/versions/draft/publish", utils.FormatPermissions([]string{ "webDeployments:configuration:edit", "webDeployments:configuration:add",  }), utils.GenerateDevCentreLink("POST", "Web Deployments", "/api/v2/webdeployments/configurations/{configurationId}/versions/draft/publish")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebDeploymentConfigurationVersion&quot;
  }
}`)
	webdeployments_configurations_versions_draft_publishCmd.AddCommand(createCmd)
	
	return webdeployments_configurations_versions_draft_publishCmd
}

var createCmd = &cobra.Command{
	Use:   "create [configurationId]",
	Short: "Publish the configuration draft and create a new version",
	Long:  "Publish the configuration draft and create a new version",
	Args:  utils.DetermineArgs([]string{ "configurationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webdeployments/configurations/{configurationId}/versions/draft/publish"
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