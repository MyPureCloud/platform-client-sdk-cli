package integrations_actions_draft_templates

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
	Description = utils.FormatUsageDescription("integrations_actions_draft_templates", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/draft/templates", )
	integrations_actions_draft_templatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_actions_draft_templates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_actions_draft_templatesCmd)
}

func Cmdintegrations_actions_draft_templates() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/actions/{actionId}/draft/templates/{fileName}", utils.FormatPermissions([]string{ "integrations:action:view", "bridge:actions:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;string&quot;
  }
}`)
	integrations_actions_draft_templatesCmd.AddCommand(getCmd)
	
	return integrations_actions_draft_templatesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [actionId] [fileName]",
	Short: "Retrieve templates for a Draft based on filename.",
	Long:  "Retrieve templates for a Draft based on filename.",
	Args:  utils.DetermineArgs([]string{ "actionId", "fileName", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/draft/templates/{fileName}"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
		fileName, args := args[0], args[1:]
		path = strings.Replace(path, "{fileName}", fmt.Sprintf("%v", fileName), -1)

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
