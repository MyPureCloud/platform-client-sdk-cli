package knowledge_guest_sessions_documents_copies

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
	Description = utils.FormatUsageDescription("knowledge_guest_sessions_documents_copies", "SWAGGER_OVERRIDE_/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/copies", )
	knowledge_guest_sessions_documents_copiesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_guest_sessions_documents_copies"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_guest_sessions_documents_copiesCmd)
}

func Cmdknowledge_guest_sessions_documents_copies() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/copies", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Knowledge", "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/copies")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/KnowledgeGuestDocumentCopy"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Operation was successful",
  "content" : { }
}`)
	knowledge_guest_sessions_documents_copiesCmd.AddCommand(createCmd)
	return knowledge_guest_sessions_documents_copiesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [sessionId] [documentId]",
	Short: "Indicate that the document was copied by the user.",
	Long:  "Indicate that the document was copied by the user.",
	Args:  utils.DetermineArgs([]string{ "sessionId", "documentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Knowledgeguestdocumentcopy{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/copies"
		sessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
		documentId, args := args[0], args[1:]
		path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)

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
