package conversations_calls_participants

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
	Description = utils.FormatUsageDescription("conversations_calls_participants", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants", )
	conversations_calls_participantsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_calls_participants"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_calls_participantsCmd)
}

func Cmdconversations_calls_participants() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/calls/{conversationId}/participants", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/calls/{conversationId}/participants")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Conversation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Conversation"
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
        "$ref" : "#/components/schemas/Conversation"
      }
    }
  }
}`)
	conversations_calls_participantsCmd.AddCommand(createCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/conversations/calls/{conversationId}/participants/{participantId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "Conversations", "/api/v2/conversations/calls/{conversationId}/participants/{participantId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Participant request",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/MediaParticipantRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "Accepted",
  "content" : { }
}`)
	conversations_calls_participantsCmd.AddCommand(updateCmd)
	return conversations_calls_participantsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId]",
	Short: "Add participants to a conversation",
	Long:  "Add participants to a conversation",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Conversation{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/calls/{conversationId}/participants"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
var updateCmd = &cobra.Command{
	Use:   "update [conversationId] [participantId]",
	Short: "Update conversation participant",
	Long:  "Update conversation participant",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Mediaparticipantrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/calls/{conversationId}/participants/{participantId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
