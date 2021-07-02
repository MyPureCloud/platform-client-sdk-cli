package notifications_channels_subscriptions

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
	Description = utils.FormatUsageDescription("notifications_channels_subscriptions", "SWAGGER_OVERRIDE_/api/v2/notifications/channels/{channelId}/subscriptions", "SWAGGER_OVERRIDE_/api/v2/notifications/channels/{channelId}/subscriptions", "SWAGGER_OVERRIDE_/api/v2/notifications/channels/{channelId}/subscriptions", "SWAGGER_OVERRIDE_/api/v2/notifications/channels/{channelId}/subscriptions", )
	notifications_channels_subscriptionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("notifications_channels_subscriptions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(notifications_channels_subscriptionsCmd)
}

func Cmdnotifications_channels_subscriptions() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Notifications", "/api/v2/notifications/channels/{channelId}/subscriptions")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ErrorBody&quot;
  },
  &quot;x-inin-error-codes&quot; : {
    &quot;bad.request&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
    &quot;response.entity.too.large&quot; : &quot;The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable&quot;
  }
}`)
	notifications_channels_subscriptionsCmd.AddCommand(deleteCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Notifications", "/api/v2/notifications/channels/{channelId}/subscriptions")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ChannelTopicEntityListing&quot;
  }
}`)
	notifications_channels_subscriptionsCmd.AddCommand(listCmd)
	
	subscribeCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", subscribeCmd.UsageTemplate(), "POST", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Notifications", "/api/v2/notifications/channels/{channelId}/subscriptions")))
	utils.AddFileFlagIfUpsert(subscribeCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Body&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/ChannelTopic&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(subscribeCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ChannelTopicEntityListing&quot;
  }
}`)
	notifications_channels_subscriptionsCmd.AddCommand(subscribeCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PUT", "Notifications", "/api/v2/notifications/channels/{channelId}/subscriptions")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Body&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/ChannelTopic&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ChannelTopicEntityListing&quot;
  }
}`)
	notifications_channels_subscriptionsCmd.AddCommand(updateCmd)
	
	return notifications_channels_subscriptionsCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [channelId]",
	Short: "Remove all subscriptions",
	Long:  "Remove all subscriptions",
	Args:  utils.DetermineArgs([]string{ "channelId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels/{channelId}/subscriptions"
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
	Use:   "list [channelId]",
	Short: "The list of all subscriptions for this channel",
	Long:  "The list of all subscriptions for this channel",
	Args:  utils.DetermineArgs([]string{ "channelId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels/{channelId}/subscriptions"
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
var subscribeCmd = &cobra.Command{
	Use:   "subscribe [channelId]",
	Short: "Add a list of subscriptions to the existing list of subscriptions",
	Long:  "Add a list of subscriptions to the existing list of subscriptions",
	Args:  utils.DetermineArgs([]string{ "channelId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels/{channelId}/subscriptions"
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
var updateCmd = &cobra.Command{
	Use:   "update [channelId]",
	Short: "Replace the current list of subscriptions with a new list.",
	Long:  "Replace the current list of subscriptions with a new list.",
	Args:  utils.DetermineArgs([]string{ "channelId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/notifications/channels/{channelId}/subscriptions"
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
