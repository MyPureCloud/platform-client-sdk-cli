package chats_threads_messages

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("chats_threads_messages", "SWAGGER_OVERRIDE_/api/v2/chats/threads/{threadId}/messages", )
	chats_threads_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chats_threads_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chats_threads_messagesCmd)
}

func Cmdchats_threads_messages() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "limit", "", "The maximum number of messages to retrieve")
	utils.AddFlag(listCmd.Flags(), "string", "before", "", "The cutoff date for messages to retrieve")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The beginning date for messages to retrieve")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/chats/threads/{threadId}/messages", utils.FormatPermissions([]string{ "chat:chat:access", "chat:room:view",  }), utils.GenerateDevCentreLink("GET", "Chat", "/api/v2/chats/threads/{threadId}/messages")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ChatMessageEntityListing"
      }
    }
  }
}`)
	chats_threads_messagesCmd.AddCommand(listCmd)
	return chats_threads_messagesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [threadId]",
	Short: "Get history by thread",
	Long:  "Get history by thread",
	Args:  utils.DetermineArgs([]string{ "threadId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/threads/{threadId}/messages"
		threadId, args := args[0], args[1:]
		path = strings.Replace(path, "{threadId}", fmt.Sprintf("%v", threadId), -1)

		limit := utils.GetFlag(cmd.Flags(), "string", "limit")
		if limit != "" {
			queryParams["limit"] = limit
		}
		before := utils.GetFlag(cmd.Flags(), "string", "before")
		if before != "" {
			queryParams["before"] = before
		}
		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "list"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
