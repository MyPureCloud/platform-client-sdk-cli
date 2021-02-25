package subscriptions

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
	subscriptionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("subscriptions"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud subscriptions"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud subscriptions`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(subscriptionsCmd)
}

func Cmdsubscriptions() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
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
	subscriptionsCmd.AddCommand(deleteCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ChannelTopicEntityListing&quot;
  }
}`)
	subscriptionsCmd.AddCommand(listCmd)
	
	subscribeCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", subscribeCmd.UsageTemplate(), "POST", "/api/v2/notifications/channels/{channelId}/subscriptions", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(subscribeCmd.Flags(), "POST")
	utils.AddPaginateFlagsIfListingResponse(subscribeCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ChannelTopicEntityListing&quot;
  }
}`)
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
