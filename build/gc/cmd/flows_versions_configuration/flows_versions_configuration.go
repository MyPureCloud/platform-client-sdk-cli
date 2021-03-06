package flows_versions_configuration

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
	Description = utils.FormatUsageDescription("flows_versions_configuration", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions/{versionId}/configuration", )
	flows_versions_configurationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_versions_configuration"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_versions_configurationCmd)
}

func Cmdflows_versions_configuration() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "deleted", "", "Deleted flows")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}/versions/{versionId}/configuration", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/{flowId}/versions/{versionId}/configuration")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;object&quot;,
    &quot;properties&quot; : { }
  }
}`)
	flows_versions_configurationCmd.AddCommand(getCmd)
	
	return flows_versions_configurationCmd
}

var getCmd = &cobra.Command{
	Use:   "get [flowId] [versionId]",
	Short: "Create flow version configuration",
	Long:  "Create flow version configuration",
	Args:  utils.DetermineArgs([]string{ "flowId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/versions/{versionId}/configuration"
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
