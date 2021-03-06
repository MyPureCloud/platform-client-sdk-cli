package learning_modules_versions

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
	Description = utils.FormatUsageDescription("learning_modules_versions", "SWAGGER_OVERRIDE_/api/v2/learning/modules/{moduleId}/versions", )
	learning_modules_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_modules_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_modules_versionsCmd)
}

func Cmdlearning_modules_versions() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Fields to expand in response(case insensitive) Valid values: assessmentForm")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/learning/modules/{moduleId}/versions/{versionId}", utils.FormatPermissions([]string{ "learning:module:view",  }), utils.GenerateDevCentreLink("GET", "Learning", "/api/v2/learning/modules/{moduleId}/versions/{versionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LearningModule&quot;
  }
}`)
	learning_modules_versionsCmd.AddCommand(getCmd)
	
	return learning_modules_versionsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [moduleId] [versionId]",
	Short: "Get specific version of a published module",
	Long:  "Get specific version of a published module",
	Args:  utils.DetermineArgs([]string{ "moduleId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/learning/modules/{moduleId}/versions/{versionId}"
		moduleId, args := args[0], args[1:]
		path = strings.Replace(path, "{moduleId}", fmt.Sprintf("%v", moduleId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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
