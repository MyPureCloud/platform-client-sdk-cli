package languageunderstanding_domains_feedback

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
	Description = utils.FormatUsageDescription("languageunderstanding_domains_feedback", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/feedback", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/feedback", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/feedback", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/feedback", )
	languageunderstanding_domains_feedbackCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_domains_feedback"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_domains_feedbackCmd)
}

func Cmdlanguageunderstanding_domains_feedback() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/languageunderstanding/domains/{domainId}/feedback", utils.FormatPermissions([]string{ "languageUnderstanding:feedback:add", "dialog:bot:add",  }), utils.GenerateDevCentreLink("POST", "Language Understanding", "/api/v2/languageunderstanding/domains/{domainId}/feedback")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "The Feedback to create.",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/NluFeedbackRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/NluFeedbackResponse"
  }
}`)
	languageunderstanding_domains_feedbackCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}", utils.FormatPermissions([]string{ "languageUnderstanding:feedback:delete", "dialog:bot:delete",  }), utils.GenerateDevCentreLink("DELETE", "Language Understanding", "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Feedback deleted successfully"
}`)
	languageunderstanding_domains_feedbackCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "fields", "", "Fields and properties to get, comma-separated Valid values: version, dateCreated, text, intents")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}", utils.FormatPermissions([]string{ "languageUnderstanding:feedback:view", "dialog:bot:view",  }), utils.GenerateDevCentreLink("GET", "Language Understanding", "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/NluFeedbackResponse"
  }
}`)
	languageunderstanding_domains_feedbackCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "intentName", "", "The top intent name to retrieve feedback for.")
	utils.AddFlag(listCmd.Flags(), "string", "assessment", "", "The top assessment to retrieve feedback for. Valid values: Incorrect, Correct, Unknown, Disabled")
	utils.AddFlag(listCmd.Flags(), "time.Time", "dateStart", "", "Begin of time window as ISO-8601 date.")
	utils.AddFlag(listCmd.Flags(), "time.Time", "dateEnd", "", "End of time window as ISO-8601 date.")
	utils.AddFlag(listCmd.Flags(), "bool", "includeDeleted", "", "Whether to include soft-deleted items in the result.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "bool", "enableCursorPagination", "false", "Enable Cursor Pagination")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned. This is considered only when enableCursorPagination=true")
	utils.AddFlag(listCmd.Flags(), "[]string", "fields", "", "Fields and properties to get, comma-separated Valid values: version, dateCreated, text, intents")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/domains/{domainId}/feedback", utils.FormatPermissions([]string{ "languageUnderstanding:feedback:view", "dialog:bot:view",  }), utils.GenerateDevCentreLink("GET", "Language Understanding", "/api/v2/languageunderstanding/domains/{domainId}/feedback")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	languageunderstanding_domains_feedbackCmd.AddCommand(listCmd)
	
	return languageunderstanding_domains_feedbackCmd
}

var createCmd = &cobra.Command{
	Use:   "create [domainId]",
	Short: "Create feedback for the NLU Domain Version.",
	Long:  "Create feedback for the NLU Domain Version.",
	Args:  utils.DetermineArgs([]string{ "domainId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Nlufeedbackrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/feedback"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [domainId] [feedbackId]",
	Short: "Delete the feedback on the NLU Domain Version.",
	Long:  "Delete the feedback on the NLU Domain Version.",
	Args:  utils.DetermineArgs([]string{ "domainId", "feedbackId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		feedbackId, args := args[0], args[1:]
		path = strings.Replace(path, "{feedbackId}", fmt.Sprintf("%v", feedbackId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [domainId] [feedbackId]",
	Short: "Find a Feedback",
	Long:  "Find a Feedback",
	Args:  utils.DetermineArgs([]string{ "domainId", "feedbackId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		feedbackId, args := args[0], args[1:]
		path = strings.Replace(path, "{feedbackId}", fmt.Sprintf("%v", feedbackId), -1)

		fields := utils.GetFlag(cmd.Flags(), "[]string", "fields")
		if fields != "" {
			queryParams["fields"] = fields
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list [domainId]",
	Short: "Get all feedback in the given NLU Domain Version.",
	Long:  "Get all feedback in the given NLU Domain Version.",
	Args:  utils.DetermineArgs([]string{ "domainId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/feedback"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)

		intentName := utils.GetFlag(cmd.Flags(), "string", "intentName")
		if intentName != "" {
			queryParams["intentName"] = intentName
		}
		assessment := utils.GetFlag(cmd.Flags(), "string", "assessment")
		if assessment != "" {
			queryParams["assessment"] = assessment
		}
		dateStart := utils.GetFlag(cmd.Flags(), "time.Time", "dateStart")
		if dateStart != "" {
			queryParams["dateStart"] = dateStart
		}
		dateEnd := utils.GetFlag(cmd.Flags(), "time.Time", "dateEnd")
		if dateEnd != "" {
			queryParams["dateEnd"] = dateEnd
		}
		includeDeleted := utils.GetFlag(cmd.Flags(), "bool", "includeDeleted")
		if includeDeleted != "" {
			queryParams["includeDeleted"] = includeDeleted
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		enableCursorPagination := utils.GetFlag(cmd.Flags(), "bool", "enableCursorPagination")
		if enableCursorPagination != "" {
			queryParams["enableCursorPagination"] = enableCursorPagination
		}
		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		fields := utils.GetFlag(cmd.Flags(), "[]string", "fields")
		if fields != "" {
			queryParams["fields"] = fields
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
