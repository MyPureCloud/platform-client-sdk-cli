package businessrules_decisiontables_versions_copy

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
	Description = utils.FormatUsageDescription("businessrules_decisiontables_versions_copy", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/copy", )
	businessrules_decisiontables_versions_copyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("businessrules_decisiontables_versions_copy"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(businessrules_decisiontables_versions_copyCmd)
}

func Cmdbusinessrules_decisiontables_versions_copy() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/copy", utils.FormatPermissions([]string{ "businessrules:decisionTable:copy",  }), utils.GenerateDevCentreLink("POST", "Business Rules", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/copy")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Decision Table",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CopyDecisionTableRequest"
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
        "$ref" : "#/components/schemas/DecisionTableVersion"
      }
    }
  }
}`)
	businessrules_decisiontables_versions_copyCmd.AddCommand(createCmd)
	return businessrules_decisiontables_versions_copyCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [tableId] [tableVersion]",
	Short: "Copy a decision table version",
	Long:  "Copy a decision table version",
	Args:  utils.DetermineArgs([]string{ "tableId", "tableVersion", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Copydecisiontablerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/copy"
		tableId, args := args[0], args[1:]
		path = strings.Replace(path, "{tableId}", fmt.Sprintf("%v", tableId), -1)
		tableVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{tableVersion}", fmt.Sprintf("%v", tableVersion), -1)

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
