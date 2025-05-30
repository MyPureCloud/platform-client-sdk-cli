package analytics_reporting_settings_users_dashboards

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
	Description = utils.FormatUsageDescription("analytics_reporting_settings_users_dashboards", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/settings/users/{userId}/dashboards", )
	analytics_reporting_settings_users_dashboardsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_settings_users_dashboards"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_settings_users_dashboardsCmd)
}

func Cmdanalytics_reporting_settings_users_dashboards() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "asc", "")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "50", "")
	utils.AddFlag(listCmd.Flags(), "bool", "publicOnly", "", "If true, retrieve only public dashboards")
	utils.AddFlag(listCmd.Flags(), "bool", "favoriteOnly", "", "If true, retrieve only favorite dashboards")
	utils.AddFlag(listCmd.Flags(), "bool", "deletedOnly", "", "If true, retrieve only deleted dashboards that are still recoverable")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "retrieve dashboards that match with given name")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/analytics/reporting/settings/users/{userId}/dashboards", utils.FormatPermissions([]string{ "analytics:dashboardConfigurations:viewPrivate",  }), utils.GenerateDevCentreLink("GET", "Analytics", "/api/v2/analytics/reporting/settings/users/{userId}/dashboards")))
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
	analytics_reporting_settings_users_dashboardsCmd.AddCommand(listCmd)
	return analytics_reporting_settings_users_dashboardsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [userId]",
	Short: "Get list of dashboards for an user",
	Long:  "Get list of dashboards for an user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/reporting/settings/users/{userId}/dashboards"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		publicOnly := utils.GetFlag(cmd.Flags(), "bool", "publicOnly")
		if publicOnly != "" {
			queryParams["publicOnly"] = publicOnly
		}
		favoriteOnly := utils.GetFlag(cmd.Flags(), "bool", "favoriteOnly")
		if favoriteOnly != "" {
			queryParams["favoriteOnly"] = favoriteOnly
		}
		deletedOnly := utils.GetFlag(cmd.Flags(), "bool", "deletedOnly")
		if deletedOnly != "" {
			queryParams["deletedOnly"] = deletedOnly
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
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
