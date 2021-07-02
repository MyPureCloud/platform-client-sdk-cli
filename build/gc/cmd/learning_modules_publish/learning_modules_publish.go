package learning_modules_publish

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
	Description = utils.FormatUsageDescription("learning_modules_publish", "SWAGGER_OVERRIDE_/api/v2/learning/modules/{moduleId}/publish", )
	learning_modules_publishCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_modules_publish"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_modules_publishCmd)
}

func Cmdlearning_modules_publish() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/learning/modules/{moduleId}/publish", utils.FormatPermissions([]string{ "learning:module:publish",  }), utils.GenerateDevCentreLink("POST", "Learning", "/api/v2/learning/modules/{moduleId}/publish")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LearningModulePublishResponse&quot;
  }
}`)
	learning_modules_publishCmd.AddCommand(createCmd)
	
	return learning_modules_publishCmd
}

var createCmd = &cobra.Command{
	Use:   "create [moduleId]",
	Short: "Publish a Learning module",
	Long:  "Publish a Learning module",
	Args:  utils.DetermineArgs([]string{ "moduleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/learning/modules/{moduleId}/publish"
		moduleId, args := args[0], args[1:]
		path = strings.Replace(path, "{moduleId}", fmt.Sprintf("%v", moduleId), -1)

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
