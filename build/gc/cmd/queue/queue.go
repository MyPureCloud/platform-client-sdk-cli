package queue
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

var queueCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("queue"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud queue"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud queue`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(queueCmd)
}

func Cmdqueue() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(getCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(getCmd.Flags(), "bool", "joined", "true", "Is joined to the queue")
	utils.AddFlag(getCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/users/{userId}/queues", utils.FormatPermissions([]string{ "routing:queue:view", "routing:queue:join", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	queueCmd.AddCommand(getCmd)
	
	joinCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", joinCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/queues/{queueId}", utils.FormatPermissions([]string{ "routing:queue:join", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(joinCmd.Flags(), "PATCH")
	queueCmd.AddCommand(joinCmd)
	
	utils.AddFlag(joinsetCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	joinsetCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", joinsetCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/queues", utils.FormatPermissions([]string{ "routing:queue:join", "routing:queueMember:manage",  })))
	utils.AddFileFlagIfUpsert(joinsetCmd.Flags(), "PATCH")
	queueCmd.AddCommand(joinsetCmd)
	
	return queueCmd
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

		retryFunc := CommandService.DetermineAction("GET", "get", urlString, "/api/v2/users/{userId}/queues")
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

		retryFunc := CommandService.DetermineAction("PATCH", "join", urlString, "/api/v2/users/{userId}/queues/{queueId}")
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

		retryFunc := CommandService.DetermineAction("PATCH", "joinset", urlString, "/api/v2/users/{userId}/queues")
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
