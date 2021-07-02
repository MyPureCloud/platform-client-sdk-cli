package learning_modules_rule

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
	Description = utils.FormatUsageDescription("learning_modules_rule", "SWAGGER_OVERRIDE_/api/v2/learning/modules/{moduleId}/rule", "SWAGGER_OVERRIDE_/api/v2/learning/modules/{moduleId}/rule", )
	learning_modules_ruleCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_modules_rule"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_modules_ruleCmd)
}

func Cmdlearning_modules_rule() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/learning/modules/{moduleId}/rule", utils.FormatPermissions([]string{ "learning:rule:view",  }), utils.GenerateDevCentreLink("GET", "Learning", "/api/v2/learning/modules/{moduleId}/rule")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LearningModuleRule&quot;
  }
}`)
	learning_modules_ruleCmd.AddCommand(getCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/learning/modules/{moduleId}/rule", utils.FormatPermissions([]string{ "learning:rule:edit",  }), utils.GenerateDevCentreLink("PUT", "Learning", "/api/v2/learning/modules/{moduleId}/rule")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The learning module rule to be updated&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LearningModuleRule&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LearningModuleRule&quot;
  }
}`)
	learning_modules_ruleCmd.AddCommand(updateCmd)
	
	return learning_modules_ruleCmd
}

var getCmd = &cobra.Command{
	Use:   "get [moduleId]",
	Short: "Get a learning module rule",
	Long:  "Get a learning module rule",
	Args:  utils.DetermineArgs([]string{ "moduleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/learning/modules/{moduleId}/rule"
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
var updateCmd = &cobra.Command{
	Use:   "update [moduleId]",
	Short: "Update a learning module rule",
	Long:  "Update a learning module rule",
	Args:  utils.DetermineArgs([]string{ "moduleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/learning/modules/{moduleId}/rule"
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

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
