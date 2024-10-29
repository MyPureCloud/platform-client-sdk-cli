package analytics_reporting_dashboards_users_bulk_remove

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
	Description = utils.FormatUsageDescription("analytics_reporting_dashboards_users_bulk_remove", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/dashboards/users/bulk/remove", )
	analytics_reporting_dashboards_users_bulk_removeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_dashboards_users_bulk_remove"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_dashboards_users_bulk_removeCmd)
}

func Cmdanalytics_reporting_dashboards_users_bulk_remove() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/analytics/reporting/dashboards/users/bulk/remove", utils.FormatPermissions([]string{ "analytics:dashboardConfigurations:deleteActive", "analytics:dashboardConfigurations:deleteInactive",  }), utils.GenerateDevCentreLink("POST", "Analytics", "/api/v2/analytics/reporting/dashboards/users/bulk/remove")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "List of userIds",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "type" : "string"
        }
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Dashboards were deleted successfully.",
  "content" : { }
}`)
	analytics_reporting_dashboards_users_bulk_removeCmd.AddCommand(createCmd)
	return analytics_reporting_dashboards_users_bulk_removeCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Bulk soft delete dashboards owned by other user(s)",
	Long:  "Bulk soft delete dashboards owned by other user(s)",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.String{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/reporting/dashboards/users/bulk/remove"

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
