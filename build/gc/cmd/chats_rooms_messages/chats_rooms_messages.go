package chats_rooms_messages

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
	Description = utils.FormatUsageDescription("chats_rooms_messages", "SWAGGER_OVERRIDE_/api/v2/chats/rooms/{roomJid}/messages", "SWAGGER_OVERRIDE_/api/v2/chats/rooms/{roomJid}/messages", "SWAGGER_OVERRIDE_/api/v2/chats/rooms/{roomJid}/messages", "SWAGGER_OVERRIDE_/api/v2/chats/rooms/{roomJid}/messages", "SWAGGER_OVERRIDE_/api/v2/chats/rooms/{roomJid}/messages", )
	chats_rooms_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chats_rooms_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chats_rooms_messagesCmd)
}

func Cmdchats_rooms_messages() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/chats/rooms/{roomJid}/messages", utils.FormatPermissions([]string{ "chat:chat:access", "chat:roomMessage:add",  }), utils.GenerateDevCentreLink("POST", "Chat", "/api/v2/chats/rooms/{roomJid}/messages")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "messageBody",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SendMessageBody"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ChatSendMessageResponse"
      }
    }
  }
}`)
	chats_rooms_messagesCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/chats/rooms/{roomJid}/messages/{messageId}", utils.FormatPermissions([]string{ "chat:chat:access", "chat:roomMessage:delete",  }), utils.GenerateDevCentreLink("DELETE", "Chat", "/api/v2/chats/rooms/{roomJid}/messages/{messageId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Message deleted successfully",
  "content" : { }
}`)
	chats_rooms_messagesCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/chats/rooms/{roomJid}/messages/{messageIds}", utils.FormatPermissions([]string{ "chat:chat:access", "chat:room:view",  }), utils.GenerateDevCentreLink("GET", "Chat", "/api/v2/chats/rooms/{roomJid}/messages/{messageIds}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ChatMessageEntityListing"
      }
    }
  }
}`)
	chats_rooms_messagesCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "limit", "", "The maximum number of messages to retrieve")
	utils.AddFlag(listCmd.Flags(), "string", "before", "", "The cutoff date for messages to retrieve")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The beginning date for messages to retrieve")
	utils.AddFlag(listCmd.Flags(), "bool", "excludeMetadata", "", "Whether to exclude metadata for messages")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/chats/rooms/{roomJid}/messages", utils.FormatPermissions([]string{ "chat:chat:access", "chat:room:view",  }), utils.GenerateDevCentreLink("GET", "Chat", "/api/v2/chats/rooms/{roomJid}/messages")))
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
	chats_rooms_messagesCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/chats/rooms/{roomJid}/messages/{messageId}", utils.FormatPermissions([]string{ "chat:chat:access", "chat:roomMessage:edit",  }), utils.GenerateDevCentreLink("PATCH", "Chat", "/api/v2/chats/rooms/{roomJid}/messages/{messageId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "messageBody",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SendMessageBody"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ChatSendMessageResponse"
      }
    }
  }
}`)
	chats_rooms_messagesCmd.AddCommand(updateCmd)
	return chats_rooms_messagesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [roomJid]",
	Short: "Send a message to a room",
	Long:  "Send a message to a room",
	Args:  utils.DetermineArgs([]string{ "roomJid", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Sendmessagebody{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/rooms/{roomJid}/messages"
		roomJid, args := args[0], args[1:]
		path = strings.Replace(path, "{roomJid}", fmt.Sprintf("%v", roomJid), -1)

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

		const opId = "create"
		const httpMethod = "POST"
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
var deleteCmd = &cobra.Command{
	Use:   "delete [roomJid] [messageId]",
	Short: "Delete a message in a room",
	Long:  "Delete a message in a room",
	Args:  utils.DetermineArgs([]string{ "roomJid", "messageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/rooms/{roomJid}/messages/{messageId}"
		roomJid, args := args[0], args[1:]
		path = strings.Replace(path, "{roomJid}", fmt.Sprintf("%v", roomJid), -1)
		messageId, args := args[0], args[1:]
		path = strings.Replace(path, "{messageId}", fmt.Sprintf("%v", messageId), -1)

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

		const opId = "delete"
		const httpMethod = "DELETE"
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
var getCmd = &cobra.Command{
	Use:   "get [roomJid] [messageIds]",
	Short: "Get messages by id(s) from a room",
	Long:  "Get messages by id(s) from a room",
	Args:  utils.DetermineArgs([]string{ "roomJid", "messageIds", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/rooms/{roomJid}/messages/{messageIds}"
		roomJid, args := args[0], args[1:]
		path = strings.Replace(path, "{roomJid}", fmt.Sprintf("%v", roomJid), -1)
		messageIds, args := args[0], args[1:]
		path = strings.Replace(path, "{messageIds}", fmt.Sprintf("%v", messageIds), -1)

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
	Use:   "list [roomJid]",
	Short: "Get a room`s message history",
	Long:  "Get a room`s message history",
	Args:  utils.DetermineArgs([]string{ "roomJid", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/rooms/{roomJid}/messages"
		roomJid, args := args[0], args[1:]
		path = strings.Replace(path, "{roomJid}", fmt.Sprintf("%v", roomJid), -1)

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
		excludeMetadata := utils.GetFlag(cmd.Flags(), "bool", "excludeMetadata")
		if excludeMetadata != "" {
			queryParams["excludeMetadata"] = excludeMetadata
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
var updateCmd = &cobra.Command{
	Use:   "update [roomJid] [messageId]",
	Short: "Edit a message in a room",
	Long:  "Edit a message in a room",
	Args:  utils.DetermineArgs([]string{ "roomJid", "messageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Sendmessagebody{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/rooms/{roomJid}/messages/{messageId}"
		roomJid, args := args[0], args[1:]
		path = strings.Replace(path, "{roomJid}", fmt.Sprintf("%v", roomJid), -1)
		messageId, args := args[0], args[1:]
		path = strings.Replace(path, "{messageId}", fmt.Sprintf("%v", messageId), -1)

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

		const opId = "update"
		const httpMethod = "PATCH"
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
