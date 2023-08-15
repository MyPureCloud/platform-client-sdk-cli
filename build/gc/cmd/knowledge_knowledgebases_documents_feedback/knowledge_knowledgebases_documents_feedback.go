package knowledge_knowledgebases_documents_feedback

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
	Description = utils.FormatUsageDescription("knowledge_knowledgebases_documents_feedback", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback", "SWAGGER_OVERRIDE_/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback", )
	knowledge_knowledgebases_documents_feedbackCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_knowledgebases_documents_feedback"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_knowledgebases_documents_feedbackCmd)
}

func Cmdknowledge_knowledgebases_documents_feedback() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback", utils.FormatPermissions([]string{ "knowledge:feedback:create",  }), utils.GenerateDevCentreLink("POST", "Knowledge", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/KnowledgeDocumentFeedback"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/KnowledgeDocumentFeedbackResponse"
      }
    }
  }
}`)
	knowledge_knowledgebases_documents_feedbackCmd.AddCommand(createCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback/{feedbackId}", utils.FormatPermissions([]string{ "knowledge:feedback:view",  }), utils.GenerateDevCentreLink("GET", "Knowledge", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback/{feedbackId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/KnowledgeDocumentFeedbackResponse"
      }
    }
  }
}`)
	knowledge_knowledgebases_documents_feedbackCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "before", "", "The cursor that points to the start of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "", "Number of entities to return. Maximum of 200.")
	utils.AddFlag(listCmd.Flags(), "bool", "onlyCommented", "", "If true, only feedback records that have comment are returned. If false, feedback records with and without comment are returned. Default: false.")
	utils.AddFlag(listCmd.Flags(), "string", "documentVersionId", "", "Document version ID to filter by. Supported only if onlyCommented=true is set.")
	utils.AddFlag(listCmd.Flags(), "string", "documentVariationId", "", "Document variation ID to filter by. Supported only if onlyCommented=true is set.")
	utils.AddFlag(listCmd.Flags(), "string", "appType", "", "Application type to filter by. Supported only if onlyCommented=true is set. Valid values: Assistant, BotFlow, MessengerKnowledgeApp, SmartAdvisor, SupportCenter")
	utils.AddFlag(listCmd.Flags(), "string", "queryType", "", "Query type to filter by. Supported only if onlyCommented=true is set. Valid values: Unknown, Article, AutoSearch, Category, ManualSearch, Recommendation, Suggestion")
	utils.AddFlag(listCmd.Flags(), "string", "userId", "", "The ID of the user, who created the feedback, to filter by. Supported only if onlyCommented=true is set.")
	utils.AddFlag(listCmd.Flags(), "string", "queueId", "", "Queue ID to filter by. Supported only if onlyCommented=true is set.")
	utils.AddFlag(listCmd.Flags(), "string", "state", "", "State to filter by. Supported only if onlyCommented=true is set. Default: Final Valid values: All, Draft, Final")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback", utils.FormatPermissions([]string{ "knowledge:feedback:view",  }), utils.GenerateDevCentreLink("GET", "Knowledge", "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	knowledge_knowledgebases_documents_feedbackCmd.AddCommand(listCmd)
	return knowledge_knowledgebases_documents_feedbackCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [knowledgeBaseId] [documentId]",
	Short: "Give feedback on a document",
	Long:  "Give feedback on a document",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "documentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Knowledgedocumentfeedback{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
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
var getCmd = &cobra.Command{
	Use:   "get [knowledgeBaseId] [documentId] [feedbackId]",
	Short: "Get a single feedback record given on a document",
	Long:  "Get a single feedback record given on a document",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "documentId", "feedbackId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback/{feedbackId}"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		documentId, args := args[0], args[1:]
		path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
		feedbackId, args := args[0], args[1:]
		path = strings.Replace(path, "{feedbackId}", fmt.Sprintf("%v", feedbackId), -1)

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

		const opId = "get"
		const httpMethod = "GET"
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
var listCmd = &cobra.Command{
	Use:   "list [knowledgeBaseId] [documentId]",
	Short: "Get a list of feedback records given on a document",
	Long:  "Get a list of feedback records given on a document",
	Args:  utils.DetermineArgs([]string{ "knowledgeBaseId", "documentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback"
		knowledgeBaseId, args := args[0], args[1:]
		path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
		documentId, args := args[0], args[1:]
		path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)

		before := utils.GetFlag(cmd.Flags(), "string", "before")
		if before != "" {
			queryParams["before"] = before
		}
		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		onlyCommented := utils.GetFlag(cmd.Flags(), "bool", "onlyCommented")
		if onlyCommented != "" {
			queryParams["onlyCommented"] = onlyCommented
		}
		documentVersionId := utils.GetFlag(cmd.Flags(), "string", "documentVersionId")
		if documentVersionId != "" {
			queryParams["documentVersionId"] = documentVersionId
		}
		documentVariationId := utils.GetFlag(cmd.Flags(), "string", "documentVariationId")
		if documentVariationId != "" {
			queryParams["documentVariationId"] = documentVariationId
		}
		appType := utils.GetFlag(cmd.Flags(), "string", "appType")
		if appType != "" {
			queryParams["appType"] = appType
		}
		queryType := utils.GetFlag(cmd.Flags(), "string", "queryType")
		if queryType != "" {
			queryParams["queryType"] = queryType
		}
		userId := utils.GetFlag(cmd.Flags(), "string", "userId")
		if userId != "" {
			queryParams["userId"] = userId
		}
		queueId := utils.GetFlag(cmd.Flags(), "string", "queueId")
		if queueId != "" {
			queryParams["queueId"] = queueId
		}
		state := utils.GetFlag(cmd.Flags(), "string", "state")
		if state != "" {
			queryParams["state"] = state
		}
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

		const opId = "list"
		const httpMethod = "GET"
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
