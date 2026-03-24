package casemanagement_cases_externalcontacts

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
	Description = utils.FormatUsageDescription("casemanagement_cases_externalcontacts", "SWAGGER_OVERRIDE_/api/v2/casemanagement/cases/externalcontacts", )
	casemanagement_cases_externalcontactsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("casemanagement_cases_externalcontacts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(casemanagement_cases_externalcontactsCmd)
}

func Cmdcasemanagement_cases_externalcontacts() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of cases that has been returned.")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "", "Number of cases to return. Maximum of 200.")
	utils.AddFlag(listCmd.Flags(), "string", "divisionIds", "", "Filter by Divisions")
	utils.AddFlag(listCmd.Flags(), "[]string", "expands", "", "Which fields to expand. Valid values: caseplan")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/casemanagement/cases/externalcontacts/{externalContactId}", utils.FormatPermissions([]string{ "caseManagement:caseExternalContact:view",  }), utils.GenerateDevCentreLink("GET", "Case Management", "/api/v2/casemanagement/cases/externalcontacts/{externalContactId}")))
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
	casemanagement_cases_externalcontactsCmd.AddCommand(listCmd)
	return casemanagement_cases_externalcontactsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [externalContactId]",
	Short: "Get a list of cases for provided external contact id.",
	Long:  "Get a list of cases for provided external contact id.",
	Args:  utils.DetermineArgs([]string{ "externalContactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/casemanagement/cases/externalcontacts/{externalContactId}"
		externalContactId, args := args[0], args[1:]
		path = strings.Replace(path, "{externalContactId}", fmt.Sprintf("%v", externalContactId), -1)

		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		divisionIds := utils.GetFlag(cmd.Flags(), "string", "divisionIds")
		if divisionIds != "" {
			queryParams["divisionIds"] = divisionIds
		}
		expands := utils.GetFlag(cmd.Flags(), "[]string", "expands")
		if expands != "" {
			queryParams["expands"] = expands
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

		const opId = "list"
		const httpMethod = "GET"
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
