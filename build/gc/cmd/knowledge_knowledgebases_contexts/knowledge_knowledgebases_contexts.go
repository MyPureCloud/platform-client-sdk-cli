package knowledge_knowledgebases_contexts

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
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_contexts", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts", )
	knowledge_knowledgebases_contextsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_contexts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_contextsCmd)
}

func Cmdknowledge_knowledgebases_contexts() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts/{contextId}", utils.FormatPermissions([]string{ "knowledge:context:edit",  }), utils.GenerateDevCentreLink("PATCH", "Knowledge", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts/{contextId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/KnowledgeContextRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/KnowledgeContextResponse"
  }
}`)
	knowledge_knowledgebases_contextsCmd.AddCommand(updateCmd)
	
	return knowledge_knowledgebases_contextsCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [knowledgeBaseId] [contextId]",
	Short: "Update specific context data of the knowledge base.",
	Long:  "Update specific context data of the knowledge base.",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "contextId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Knowledgecontextrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts/{contextId}"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		contextId, args := args[0], args[1:]
		path = strings.Replace(path, "{contextId}", fmt.Sprintf("%v", contextId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
