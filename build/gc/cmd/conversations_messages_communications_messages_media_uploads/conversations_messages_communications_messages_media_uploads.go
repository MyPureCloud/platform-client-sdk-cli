package conversations_messages_communications_messages_media_uploads

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
	Description = utils.FormatUsageDescription("conversations_messages_communications_messages_media_uploads", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/uploads", )
	conversations_messages_communications_messages_media_uploadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_communications_messages_media_uploads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_communications_messages_media_uploadsCmd)
}

func Cmdconversations_messages_communications_messages_media_uploads() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/uploads", utils.FormatPermissions([]string{ "conversation:message:create", "conversation:webmessaging:create", "conversation:socialmedia:create",  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/uploads")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "request",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UploadMediaRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Pre-signed url successfully created.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/MessageMediaUploadData"
      }
    }
  }
}`)
	conversations_messages_communications_messages_media_uploadsCmd.AddCommand(createCmd)
	return conversations_messages_communications_messages_media_uploadsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [communicationId]",
	Short: "Create a URL to upload a message media file",
	Long:  "Create a URL to upload a message media file",
	Args:  utils.DetermineArgs([]string{ "conversationId", "communicationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Uploadmediarequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/uploads"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		communicationId, args := args[0], args[1:]
		path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)

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
