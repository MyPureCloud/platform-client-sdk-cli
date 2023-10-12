package chats_rooms_pinnedmessages

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
	Description = utils.FormatUsageDescription("chats_rooms_pinnedmessages", "SWAGGER_OVERRIDE_/api/v2/chats/rooms/{roomJid}/pinnedmessages", "SWAGGER_OVERRIDE_/api/v2/chats/rooms/{roomJid}/pinnedmessages", )
	chats_rooms_pinnedmessagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chats_rooms_pinnedmessages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chats_rooms_pinnedmessagesCmd)
}

func Cmdchats_rooms_pinnedmessages() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/chats/rooms/{roomJid}/pinnedmessages", utils.FormatPermissions([]string{ "chat:chat:access", "chat:room:edit",  }), utils.GenerateDevCentreLink("POST", "Chat", "/api/v2/chats/rooms/{roomJid}/pinnedmessages")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Pinned Message Ids",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PinnedMessageRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Pinned messages added successfully",
  "content" : { }
}`)
	chats_rooms_pinnedmessagesCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/chats/rooms/{roomJid}/pinnedmessages/{pinnedMessageId}", utils.FormatPermissions([]string{ "chat:chat:access", "chat:room:edit",  }), utils.GenerateDevCentreLink("DELETE", "Chat", "/api/v2/chats/rooms/{roomJid}/pinnedmessages/{pinnedMessageId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Pinned message removed successfully",
  "content" : { }
}`)
	chats_rooms_pinnedmessagesCmd.AddCommand(deleteCmd)
	return chats_rooms_pinnedmessagesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [roomJid]",
	Short: "Add pinned messages for a room, up to a maximum of 5 pinned messages",
	Long:  "Add pinned messages for a room, up to a maximum of 5 pinned messages",
	Args:  utils.DetermineArgs([]string{ "roomJid", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Pinnedmessagerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/rooms/{roomJid}/pinnedmessages"
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
	Use:   "delete [roomJid] [pinnedMessageId]",
	Short: "Remove a pinned message from a room",
	Long:  "Remove a pinned message from a room",
	Args:  utils.DetermineArgs([]string{ "roomJid", "pinnedMessageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/chats/rooms/{roomJid}/pinnedmessages/{pinnedMessageId}"
		roomJid, args := args[0], args[1:]
		path = strings.Replace(path, "{roomJid}", fmt.Sprintf("%v", roomJid), -1)
		pinnedMessageId, args := args[0], args[1:]
		path = strings.Replace(path, "{pinnedMessageId}", fmt.Sprintf("%v", pinnedMessageId), -1)

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
