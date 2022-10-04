package conversations_messages_participants_communications

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
	Description = utils.FormatUsageDescription("conversations_messages_participants_communications", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{conversationId}/participants/{participantId}/communications", )
	conversations_messages_participants_communicationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_participants_communications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_participants_communicationsCmd)
}

func Cmdconversations_messages_participants_communications() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/conversations/messages/{conversationId}/participants/{participantId}/communications/{communicationId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "Conversations", "/api/v2/conversations/messages/{conversationId}/participants/{participantId}/communications/{communicationId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Participant",
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
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Empty"
      }
    }
  }
}`)
	conversations_messages_participants_communicationsCmd.AddCommand(updateCmd)
	return conversations_messages_participants_communicationsCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [conversationId] [participantId] [communicationId]",
	Short: "Update conversation participant`s communication by disconnecting it.",
	Long:  "Update conversation participant`s communication by disconnecting it.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", "communicationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Mediaparticipantrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messages/{conversationId}/participants/{participantId}/communications/{communicationId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)
		communicationId, args := args[0], args[1:]
		path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)

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
