package webchat_guest_conversations_messages

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
	Description = utils.FormatUsageDescription("webchat_guest_conversations_messages", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations/{conversationId}/messages", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations/{conversationId}/messages", )
	webchat_guest_conversations_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webchat_guest_conversations_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webchat_guest_conversations_messagesCmd)
}

func Cmdwebchat_guest_conversations_messages() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/webchat/guest/conversations/{conversationId}/messages/{messageId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "WebChat", "/api/v2/webchat/guest/conversations/{conversationId}/messages/{messageId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WebChatMessage"
      }
    }
  }
}`)
	webchat_guest_conversations_messagesCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "after", "", "If available, get the messages chronologically after the id of this message")
	utils.AddFlag(listCmd.Flags(), "string", "before", "", "If available, get the messages chronologically before the id of this message")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ascending", "Sort order Valid values: ascending, descending")
	utils.AddFlag(listCmd.Flags(), "int", "maxResults", "100", "Limit the returned number of messages, up to a maximum of 100")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/webchat/guest/conversations/{conversationId}/messages", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "WebChat", "/api/v2/webchat/guest/conversations/{conversationId}/messages")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WebChatMessageEntityList"
      }
    }
  }
}`)
	webchat_guest_conversations_messagesCmd.AddCommand(listCmd)
	return webchat_guest_conversations_messagesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [conversationId] [messageId]",
	Short: "Get a web chat conversation message",
	Long:  "Get a web chat conversation message",
	Args:  utils.DetermineArgs([]string{ "conversationId", "messageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations/{conversationId}/messages/{messageId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		messageId, args := args[0], args[1:]
		path = strings.Replace(path, "{messageId}", fmt.Sprintf("%v", messageId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
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
var listCmd = &cobra.Command{
	Use:   "list [conversationId]",
	Short: "Get the messages of a chat conversation.",
	Long:  "Get the messages of a chat conversation.",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations/{conversationId}/messages"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		before := utils.GetFlag(cmd.Flags(), "string", "before")
		if before != "" {
			queryParams["before"] = before
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		maxResults := utils.GetFlag(cmd.Flags(), "int", "maxResults")
		if maxResults != "" {
			queryParams["maxResults"] = maxResults
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
