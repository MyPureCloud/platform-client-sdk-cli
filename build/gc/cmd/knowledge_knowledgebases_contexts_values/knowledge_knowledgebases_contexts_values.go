package knowledge_knowledgebases_contexts_values

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
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_contexts_values", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts/{contextId}/values", )
	knowledge_knowledgebases_contexts_valuesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_contexts_values"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_contexts_valuesCmd)
}

func Cmdknowledge_knowledgebases_contexts_values() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts/{contextId}/values/{contextValueId}", utils.FormatPermissions([]string{ "knowledge:context:edit",  }), utils.GenerateDevCentreLink("PATCH", "Knowledge", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts/{contextId}/values/{contextValueId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/KnowledgeContextValueRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/KnowledgeContextValueResponse"
  }
}`)
	knowledge_knowledgebases_contexts_valuesCmd.AddCommand(updateCmd)
	
	return knowledge_knowledgebases_contexts_valuesCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [knowledgeBaseId] [contextId] [contextValueId]",
	Short: "Update context value.",
	Long:  "Update context value.",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "contextId", "contextValueId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Knowledgecontextvaluerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/contexts/{contextId}/values/{contextValueId}"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		contextId, args := args[0], args[1:]
		path = strings.Replace(path, "{contextId}", fmt.Sprintf("%v", contextId), -1)
		contextValueId, args := args[0], args[1:]
		path = strings.Replace(path, "{contextValueId}", fmt.Sprintf("%v", contextValueId), -1)


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
