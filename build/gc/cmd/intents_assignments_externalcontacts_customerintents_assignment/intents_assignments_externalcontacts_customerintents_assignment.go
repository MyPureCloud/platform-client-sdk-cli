package intents_assignments_externalcontacts_customerintents_assignment

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
	Description = utils.FormatUsageDescription("intents_assignments_externalcontacts_customerintents_assignment", "SWAGGER_OVERRIDE_/api/v2/intents/assignments/externalcontacts/{externalContactId}/customerintents/{customerIntentId}/assignment", )
	intents_assignments_externalcontacts_customerintents_assignmentCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("intents_assignments_externalcontacts_customerintents_assignment"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(intents_assignments_externalcontacts_customerintents_assignmentCmd)
}

func Cmdintents_assignments_externalcontacts_customerintents_assignment() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/intents/assignments/externalcontacts/{externalContactId}/customerintents/{customerIntentId}/assignment", utils.FormatPermissions([]string{ "externalContacts:customerIntentAssignments:add",  }), utils.GenerateDevCentreLink("POST", "Intents", "/api/v2/intents/assignments/externalcontacts/{externalContactId}/customerintents/{customerIntentId}/assignment")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Customer intent assignment",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CustomerIntentAssignmentRequest"
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
        "$ref" : "#/components/schemas/CustomerIntentAssignmentResponse"
      }
    }
  }
}`)
	intents_assignments_externalcontacts_customerintents_assignmentCmd.AddCommand(createCmd)
	return intents_assignments_externalcontacts_customerintents_assignmentCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [externalContactId] [customerIntentId]",
	Short: "Create customer intent assignment for external contact",
	Long:  "Create customer intent assignment for external contact",
	Args:  utils.DetermineArgs([]string{ "externalContactId", "customerIntentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Customerintentassignmentrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/intents/assignments/externalcontacts/{externalContactId}/customerintents/{customerIntentId}/assignment"
		externalContactId, args := args[0], args[1:]
		path = strings.Replace(path, "{externalContactId}", fmt.Sprintf("%v", externalContactId), -1)
		customerIntentId, args := args[0], args[1:]
		path = strings.Replace(path, "{customerIntentId}", fmt.Sprintf("%v", customerIntentId), -1)

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
