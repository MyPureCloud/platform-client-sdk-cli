package queues_users

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
	queues_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("queues_users"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud queues_users"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud queues_users`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(queues_usersCmd)
}

func Cmdqueues_users() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(getCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(getCmd.Flags(), "bool", "joined", "true", "Is joined to the queue")
	utils.AddFlag(getCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/users/{userId}/queues", utils.FormatPermissions([]string{ "routing:queue:view", "routing:queue:join", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserQueueEntityListing&quot;
  }
}`)
	queues_usersCmd.AddCommand(getCmd)
	
	joinCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", joinCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/queues/{queueId}", utils.FormatPermissions([]string{ "routing:queue:join", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(joinCmd.Flags(), "PATCH")
	utils.AddPaginateFlagsIfListingResponse(joinCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserQueue&quot;
  }
}`)
	queues_usersCmd.AddCommand(joinCmd)
	
	utils.AddFlag(joinsetCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	joinsetCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", joinsetCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/queues", utils.FormatPermissions([]string{ "routing:queue:join", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(joinsetCmd.Flags(), "PATCH")
	utils.AddPaginateFlagsIfListingResponse(joinsetCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserQueueEntityListing&quot;
  }
}`)
	queues_usersCmd.AddCommand(joinsetCmd)
	
	return queues_usersCmd
}

var getCmd = &cobra.Command{
	Use:   "get [userId]",
	Short: "Get queues for user",
	Long:  `Get queues for user`,
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/queues"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		joined := utils.GetFlag(cmd.Flags(), "bool", "joined")
		if joined != "" {
			queryParams["joined"] = joined
		}
		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
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
var joinCmd = &cobra.Command{
	Use:   "join [queueId] [userId]",
	Short: "Join or unjoin a queue for a user",
	Long:  `Join or unjoin a queue for a user`,
	Args:  utils.DetermineArgs([]string{ "queueId", "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/queues/{queueId}"
		queueId, args := args[0], args[1:]
		path = strings.Replace(path, "{queueId}", fmt.Sprintf("%v", queueId), -1)
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
var joinsetCmd = &cobra.Command{
	Use:   "joinset [userId]",
	Short: "Join or unjoin a set of queues for a user",
	Long:  `Join or unjoin a set of queues for a user`,
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/queues"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
