package conversations_participants_flaggedreason

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
	Description = utils.FormatUsageDescription("conversations_participants_flaggedreason", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason", )
	conversations_participants_flaggedreasonCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_participants_flaggedreason"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_participants_flaggedreasonCmd)
}

func Cmdconversations_participants_flaggedreason() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Conversations", "/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The flagged reason was removed successfully.",
  "content" : { }
}`)
	conversations_participants_flaggedreasonCmd.AddCommand(deleteCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PUT", "Conversations", "/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", ``)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "The flagged reason was set successfully.",
  "content" : { }
}`)
	conversations_participants_flaggedreasonCmd.AddCommand(updateCmd)
	return conversations_participants_flaggedreasonCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [conversationId] [participantId]",
	Short: "Remove flagged reason from conversation participant.",
	Long:  "Remove flagged reason from conversation participant.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason"
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
var updateCmd = &cobra.Command{
	Use:   "update [conversationId] [participantId]",
	Short: "Set flagged reason on conversation participant to indicate bad conversation quality.",
	Long:  "Set flagged reason on conversation participant to indicate bad conversation quality.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason"
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
		const httpMethod = "PUT"
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
