package billing_contracts_billingperiods

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
	Description = utils.FormatUsageDescription("billing_contracts_billingperiods", "SWAGGER_OVERRIDE_/api/v2/billing/contracts/{contractId}/billingperiods", )
	billing_contracts_billingperiodsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("billing_contracts_billingperiods"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(billing_contracts_billingperiodsCmd)
}

func Cmdbilling_contracts_billingperiods() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/billing/contracts/{contractId}/billingperiods/{billingPeriodId}", utils.FormatPermissions([]string{ "billing:subscription:view",  }), utils.GenerateDevCentreLink("GET", "Billing", "/api/v2/billing/contracts/{contractId}/billingperiods/{billingPeriodId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BillingContractPeriodDetail"
      }
    }
  }
}`)
	billing_contracts_billingperiodsCmd.AddCommand(getCmd)
	return billing_contracts_billingperiodsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [contractId] [billingPeriodId]",
	Short: "Get contract billing period",
	Long:  "Get contract billing period",
	Args:  utils.DetermineArgs([]string{ "contractId", "billingPeriodId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/billing/contracts/{contractId}/billingperiods/{billingPeriodId}"
		contractId, args := args[0], args[1:]
		path = strings.Replace(path, "{contractId}", fmt.Sprintf("%v", contractId), -1)
		billingPeriodId, args := args[0], args[1:]
		path = strings.Replace(path, "{billingPeriodId}", fmt.Sprintf("%v", billingPeriodId), -1)

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
