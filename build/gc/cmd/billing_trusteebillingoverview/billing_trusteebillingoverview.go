package billing_trusteebillingoverview

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
	Description = utils.FormatUsageDescription("billing_trusteebillingoverview", "SWAGGER_OVERRIDE_/api/v2/billing/trusteebillingoverview", )
	billing_trusteebillingoverviewCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("billing_trusteebillingoverview"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(billing_trusteebillingoverviewCmd)
}

func Cmdbilling_trusteebillingoverview() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "int", "billingPeriodIndex", "0", "0 for active period (overview data may change until period closes). 1 for prior completed billing period. 2 for two billing cycles prior, and so on.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/billing/trusteebillingoverview/{trustorOrgId}", utils.FormatPermissions([]string{ "affiliateOrganization:clientBilling:view",  }), utils.GenerateDevCentreLink("GET", "Billing", "/api/v2/billing/trusteebillingoverview/{trustorOrgId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TrusteeBillingOverview"
      }
    }
  }
}`)
	billing_trusteebillingoverviewCmd.AddCommand(getCmd)
	return billing_trusteebillingoverviewCmd
}

var getCmd = &cobra.Command{
	Use:   "get [trustorOrgId]",
	Short: "Get the billing overview for an organization that is managed by a partner.",
	Long:  "Get the billing overview for an organization that is managed by a partner.",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/billing/trusteebillingoverview/{trustorOrgId}"
		trustorOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trustorOrgId}", fmt.Sprintf("%v", trustorOrgId), -1)

		billingPeriodIndex := utils.GetFlag(cmd.Flags(), "int", "billingPeriodIndex")
		if billingPeriodIndex != "" {
			queryParams["billingPeriodIndex"] = billingPeriodIndex
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
