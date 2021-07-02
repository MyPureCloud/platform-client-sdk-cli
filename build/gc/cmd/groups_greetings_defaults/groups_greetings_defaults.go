package groups_greetings_defaults

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
	Description = utils.FormatUsageDescription("groups_greetings_defaults", "SWAGGER_OVERRIDE_/api/v2/groups/{groupId}/greetings/defaults", "SWAGGER_OVERRIDE_/api/v2/groups/{groupId}/greetings/defaults", )
	groups_greetings_defaultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("groups_greetings_defaults"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(groups_greetings_defaultsCmd)
}

func Cmdgroups_greetings_defaults() *cobra.Command { 
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/groups/{groupId}/greetings/defaults", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Greetings", "/api/v2/groups/{groupId}/greetings/defaults")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DefaultGreetingList&quot;
  }
}`)
	groups_greetings_defaultsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/groups/{groupId}/greetings/defaults", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PUT", "Greetings", "/api/v2/groups/{groupId}/greetings/defaults")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The updated defaultGreetingList&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DefaultGreetingList&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DefaultGreetingList&quot;
  }
}`)
	groups_greetings_defaultsCmd.AddCommand(updateCmd)
	
	return groups_greetings_defaultsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [groupId]",
	Short: "Grabs the list of Default Greetings given a Group`s ID",
	Long:  "Grabs the list of Default Greetings given a Group`s ID",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/groups/{groupId}/greetings/defaults"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

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
	Use:   "update [groupId]",
	Short: "Updates the DefaultGreetingList of the specified Group",
	Long:  "Updates the DefaultGreetingList of the specified Group",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/groups/{groupId}/greetings/defaults"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

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
