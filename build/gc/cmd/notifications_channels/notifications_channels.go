package notifications_channels

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
	Description = utils.FormatUsageDescription("notifications_channels", "SWAGGER_OVERRIDE_/api/v2/notifications/channels", "SWAGGER_OVERRIDE_/api/v2/notifications/channels", "SWAGGER_OVERRIDE_/api/v2/notifications/channels", )
	notifications_channelsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("notifications_channels"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(notifications_channelsCmd)
}

func Cmdnotifications_channels() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/notifications/channels", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Notifications", "/api/v2/notifications/channels")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Channel&quot;
  }
}`)
	notifications_channelsCmd.AddCommand(createCmd)
	
	headCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", headCmd.UsageTemplate(), "HEAD", "/api/v2/notifications/channels/{channelId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("HEAD", "Notifications", "/api/v2/notifications/channels/{channelId}")))
	utils.AddFileFlagIfUpsert(headCmd.Flags(), "HEAD", ``)
	
	utils.AddPaginateFlagsIfListingResponse(headCmd.Flags(), "HEAD", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;boolean&quot;
  }
}`)
	notifications_channelsCmd.AddCommand(headCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "includechannels", "token", "Show user`s channels for this specific token or across all tokens for this user and app.  Channel Ids for other access tokens will not be shown, but will be presented to show their existence. Valid values: token, oauthclient")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/notifications/channels", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Notifications", "/api/v2/notifications/channels")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ChannelEntityListing&quot;
  }
}`)
	notifications_channelsCmd.AddCommand(listCmd)
	
	return notifications_channelsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new channel",
	Long:  "Create a new channel",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels"

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
var headCmd = &cobra.Command{
	Use:   "head [channelId]",
	Short: "Verify a channel still exists and is valid",
	Long:  "Verify a channel still exists and is valid",
	Args:  utils.DetermineArgs([]string{ "channelId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels/{channelId}"
		channelId, args := args[0], args[1:]
		path = strings.Replace(path, "{channelId}", fmt.Sprintf("%v", channelId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("HEAD", urlString, cmd.Flags())
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
	Short: "The list of existing channels",
	Long:  "The list of existing channels",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels"

		includechannels := utils.GetFlag(cmd.Flags(), "string", "includechannels")
		if includechannels != "" {
			queryParams["includechannels"] = includechannels
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
