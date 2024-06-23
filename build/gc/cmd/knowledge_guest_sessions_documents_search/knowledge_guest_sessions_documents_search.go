package knowledge_guest_sessions_documents_search

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
	Description = utils.FormatUsageDescription("knowledge_guest_sessions_documents_search", "SWAGGER_OVERRIDE_/api/v2/knowledge/guest/sessions/{sessionId}/documents/search", "SWAGGER_OVERRIDE_/api/v2/knowledge/guest/sessions/{sessionId}/documents/search", )
	knowledge_guest_sessions_documents_searchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_guest_sessions_documents_search"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_guest_sessions_documents_searchCmd)
}

func Cmdknowledge_guest_sessions_documents_search() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "[]string", "expand", "", "Fields, if any, to expand for each document in the search result matching the query. Valid values: documentVariations, documentAlternatives, knowledgeBaseLanguageCode")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Knowledge", "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/KnowledgeDocumentGuestSearchRequest"
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
        "$ref" : "#/components/schemas/KnowledgeDocumentGuestSearch"
      }
    }
  }
}`)
	knowledge_guest_sessions_documents_searchCmd.AddCommand(createCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "Knowledge", "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SearchUpdateRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "Search updated successfully.",
  "content" : { }
}`)
	knowledge_guest_sessions_documents_searchCmd.AddCommand(updateCmd)
	return knowledge_guest_sessions_documents_searchCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [sessionId]",
	Short: "Search the documents in a guest session.",
	Long:  "Search the documents in a guest session.",
	Args:  utils.DetermineArgs([]string{ "sessionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Knowledgedocumentguestsearchrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search"
		sessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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
var updateCmd = &cobra.Command{
	Use:   "update [sessionId] [searchId]",
	Short: "Update search result.",
	Long:  "Update search result.",
	Args:  utils.DetermineArgs([]string{ "sessionId", "searchId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Searchupdaterequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}"
		sessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
		searchId, args := args[0], args[1:]
		path = strings.Replace(path, "{searchId}", fmt.Sprintf("%v", searchId), -1)

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
