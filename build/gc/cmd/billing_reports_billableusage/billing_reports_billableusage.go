package billing_reports_billableusage

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
	Description = utils.FormatUsageDescription("billing_reports_billableusage", "SWAGGER_OVERRIDE_/api/v2/billing/reports/billableusage", )
	billing_reports_billableusageCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("billing_reports_billableusage"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(billing_reports_billableusageCmd)
}

func Cmdbilling_reports_billableusage() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "time.Time", "startDate", "", "The period start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "endDate", "", "The period end date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/billing/reports/billableusage", utils.FormatPermissions([]string{ "billing:subscription:read", "billing:subscription:view",  }), utils.GenerateDevCentreLink("GET", "Billing", "/api/v2/billing/reports/billableusage")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("startDate")
	getCmd.MarkFlagRequired("endDate")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BillingUsageReport"
      }
    }
  }
}`)
	billing_reports_billableusageCmd.AddCommand(getCmd)
	return billing_reports_billableusageCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a report of the billable license usages",
	Long:  "Get a report of the billable license usages",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/billing/reports/billableusage"

		startDate := utils.GetFlag(cmd.Flags(), "time.Time", "startDate")
		if startDate != "" {
			queryParams["startDate"] = startDate
		}
		endDate := utils.GetFlag(cmd.Flags(), "time.Time", "endDate")
		if endDate != "" {
			queryParams["endDate"] = endDate
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
