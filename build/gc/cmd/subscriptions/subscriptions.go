package subscriptions
import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

var subscriptionsCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("subscriptions"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud subscriptions"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud subscriptions`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(subscriptionsCmd)
}

func Cmdsubscriptions() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	subscriptionsCmd.AddCommand(deleteCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	subscriptionsCmd.AddCommand(listCmd)
	
	subscribeCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", subscribeCmd.UsageTemplate(), "POST", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(subscribeCmd.Flags(), "POST")
	subscriptionsCmd.AddCommand(subscribeCmd)
	
	return subscriptionsCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [channelId]",
	Short: "Remove all subscriptions",
	Long:  `Remove all subscriptions`,
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

		retryFunc := CommandService.DetermineAction("DELETE", "delete", urlString, "/api/v2/notifications/channels/{channelId}/subscriptions")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
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
	Long:  `The list of all subscriptions for this channel`,
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

		retryFunc := CommandService.DetermineAction("GET", "list", urlString, "/api/v2/notifications/channels/{channelId}/subscriptions")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
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
	Long:  `Add a list of subscriptions to the existing list of subscriptions`,
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

		retryFunc := CommandService.DetermineAction("POST", "subscribe", urlString, "/api/v2/notifications/channels/{channelId}/subscriptions")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
