package conversations_communications_agentchecklists_agentaction

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
	Description = utils.FormatUsageDescription("conversations_communications_agentchecklists_agentaction", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/communications/{communicationId}/agentchecklists/{agentChecklistId}/agentaction", )
	conversations_communications_agentchecklists_agentactionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_communications_agentchecklists_agentaction"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_communications_agentchecklists_agentactionCmd)
}

func Cmdconversations_communications_agentchecklists_agentaction() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/{conversationId}/communications/{communicationId}/agentchecklists/{agentChecklistId}/agentaction", utils.FormatPermissions([]string{ "conversation:agentchecklist:edit",  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/{conversationId}/communications/{communicationId}/agentchecklists/{agentChecklistId}/agentaction")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Agent action payload",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AgentActionPayload"
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
        "$ref" : "#/components/schemas/AgentChecklistResponse"
      }
    }
  }
}`)
	conversations_communications_agentchecklists_agentactionCmd.AddCommand(createCmd)
	return conversations_communications_agentchecklists_agentactionCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [communicationId] [agentChecklistId]",
	Short: "API invoked to capture an agent action.",
	Long:  "API invoked to capture an agent action.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "communicationId", "agentChecklistId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Agentactionpayload{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/communications/{communicationId}/agentchecklists/{agentChecklistId}/agentaction"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		communicationId, args := args[0], args[1:]
		path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)
		agentChecklistId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentChecklistId}", fmt.Sprintf("%v", agentChecklistId), -1)

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
