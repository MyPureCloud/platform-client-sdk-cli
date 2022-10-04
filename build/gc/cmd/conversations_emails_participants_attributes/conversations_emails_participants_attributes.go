package conversations_emails_participants_attributes

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
	Description = utils.FormatUsageDescription("conversations_emails_participants_attributes", "SWAGGER_OVERRIDE_/api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes", )
	conversations_emails_participants_attributesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_emails_participants_attributes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_emails_participants_attributesCmd)
}

func Cmdconversations_emails_participants_attributes() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "Conversations", "/api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Participant attributes",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ParticipantAttributes"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "Accepted",
  "content" : { }
}`)
	conversations_emails_participants_attributesCmd.AddCommand(updateCmd)
	return conversations_emails_participants_attributesCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [conversationId] [participantId]",
	Short: "Update the attributes on a conversation participant.",
	Long:  "Update the attributes on a conversation participant.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Participantattributes{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/emails/{conversationId}/participants/{participantId}/attributes"
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
