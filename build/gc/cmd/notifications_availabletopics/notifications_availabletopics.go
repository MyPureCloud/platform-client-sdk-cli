package notifications_availabletopics

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
	Description = utils.FormatUsageDescription("notifications_availabletopics", "SWAGGER_OVERRIDE_/api/v2/notifications/availabletopics", )
	notifications_availabletopicsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("notifications_availabletopics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(notifications_availabletopicsCmd)
}

func Cmdnotifications_availabletopics() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: description, enforced, schema, visibility, transports, publicApiTemplateUriPaths, requiresPermissions, permissionDetails")
	utils.AddFlag(listCmd.Flags(), "bool", "includePreview", "true", "Whether or not to include Preview topics")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/notifications/availabletopics", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Notifications", "/api/v2/notifications/availabletopics")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AvailableTopicEntityListing&quot;
  }
}`)
	notifications_availabletopicsCmd.AddCommand(listCmd)
	
	return notifications_availabletopicsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get available notification topics.",
	Long:  "Get available notification topics.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/availabletopics"

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		includePreview := utils.GetFlag(cmd.Flags(), "bool", "includePreview")
		if includePreview != "" {
			queryParams["includePreview"] = includePreview
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
