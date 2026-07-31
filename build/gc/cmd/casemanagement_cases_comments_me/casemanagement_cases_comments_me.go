package casemanagement_cases_comments_me

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
	Description = utils.FormatUsageDescription("casemanagement_cases_comments_me", "SWAGGER_OVERRIDE_/api/v2/casemanagement/cases/{caseId}/comments/me", )
	casemanagement_cases_comments_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("casemanagement_cases_comments_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(casemanagement_cases_comments_meCmd)
}

func Cmdcasemanagement_cases_comments_me() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/casemanagement/cases/{caseId}/comments/me/{commentId}", utils.FormatPermissions([]string{ "caseManagement:commentSelf:delete",  }), utils.GenerateDevCentreLink("DELETE", "Case Management", "/api/v2/casemanagement/cases/{caseId}/comments/me/{commentId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Empty"
      }
    }
  }
}`)
	casemanagement_cases_comments_meCmd.AddCommand(deleteCmd)
	return casemanagement_cases_comments_meCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [caseId] [commentId]",
	Short: "Delete my Comment.",
	Long:  "Delete my Comment.",
	Args:  utils.DetermineArgs([]string{ "caseId", "commentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/casemanagement/cases/{caseId}/comments/me/{commentId}"
		caseId, args := args[0], args[1:]
		path = strings.Replace(path, "{caseId}", fmt.Sprintf("%v", caseId), -1)
		commentId, args := args[0], args[1:]
		path = strings.Replace(path, "{commentId}", fmt.Sprintf("%v", commentId), -1)

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

		headerParams := make(map[string]string)
		// to determine the Content-Type header
		localVarHttpContentTypes := []string{ "application/json",  }
		// set Content-Type header
		localVarHttpContentType := utils.SelectHeaderContentType(localVarHttpContentTypes)
		if localVarHttpContentType != "" {
			headerParams["Content-Type"] = localVarHttpContentType
		}
		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{
			"application/json",
		}
		// set Accept header
		localVarHttpHeaderAccept := utils.SelectHeaderAccept(localVarHttpHeaderAccepts)
		if localVarHttpHeaderAccept != "" {
			headerParams["Accept"] = localVarHttpHeaderAccept
		}

		const opId = "delete"
		const httpMethod = "DELETE"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, headerParams, cmd, opId)
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
