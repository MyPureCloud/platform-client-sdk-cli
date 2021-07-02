package conversations_participants_codes

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
	Description = utils.FormatUsageDescription("conversations_participants_codes", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/codes", )
	conversations_participants_codesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_participants_codes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_participants_codesCmd)
}

func Cmdconversations_participants_codes() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Conversations", "/api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ErrorBody&quot;
  },
  &quot;x-inin-error-codes&quot; : {
    &quot;bad.request&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
    &quot;response.entity.too.large&quot; : &quot;The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable&quot;
  }
}`)
	conversations_participants_codesCmd.AddCommand(deleteCmd)
	
	return conversations_participants_codesCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [conversationId] [participantId] [addCommunicationCode]",
	Short: "Delete a code used to add a communication to this participant",
	Long:  "Delete a code used to add a communication to this participant",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", "addCommunicationCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)
		addCommunicationCode, args := args[0], args[1:]
		path = strings.Replace(path, "{addCommunicationCode}", fmt.Sprintf("%v", addCommunicationCode), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
