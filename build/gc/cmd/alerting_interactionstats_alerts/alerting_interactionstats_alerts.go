package alerting_interactionstats_alerts

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
	Description = utils.FormatUsageDescription("alerting_interactionstats_alerts", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/alerts", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/alerts", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/alerts", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/alerts", )
	alerting_interactionstats_alertsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("alerting_interactionstats_alerts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(alerting_interactionstats_alertsCmd)
}

func Cmdalerting_interactionstats_alerts() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/alerting/interactionstats/alerts/{alertId}", utils.FormatPermissions([]string{ "alerting:alert:delete",  }), utils.GenerateDevCentreLink("DELETE", "Alerting", "/api/v2/alerting/interactionstats/alerts/{alertId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Interaction stats alert deleted"
}`)
	alerting_interactionstats_alertsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: notificationUsers")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/alerting/interactionstats/alerts/{alertId}", utils.FormatPermissions([]string{ "alerting:alert:view",  }), utils.GenerateDevCentreLink("GET", "Alerting", "/api/v2/alerting/interactionstats/alerts/{alertId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/InteractionStatsAlert"
  }
}`)
	alerting_interactionstats_alertsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: notificationUsers")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/alerting/interactionstats/alerts", utils.FormatPermissions([]string{ "alerting:alert:view",  }), utils.GenerateDevCentreLink("GET", "Alerting", "/api/v2/alerting/interactionstats/alerts")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	alerting_interactionstats_alertsCmd.AddCommand(listCmd)
	
	utils.AddFlag(updateCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: notificationUsers")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/alerting/interactionstats/alerts/{alertId}", utils.FormatPermissions([]string{ "alerting:alert:edit",  }), utils.GenerateDevCentreLink("PUT", "Alerting", "/api/v2/alerting/interactionstats/alerts/{alertId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "InteractionStatsAlert",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/UnreadStatus"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/UnreadStatus"
  }
}`)
	alerting_interactionstats_alertsCmd.AddCommand(updateCmd)
	
	return alerting_interactionstats_alertsCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [alertId]",
	Short: "Delete an interaction stats alert",
	Long:  "Delete an interaction stats alert",
	Args:  utils.DetermineArgs([]string{ "alertId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/alerts/{alertId}"
		alertId, args := args[0], args[1:]
		path = strings.Replace(path, "{alertId}", fmt.Sprintf("%v", alertId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
	Use:   "get [alertId]",
	Short: "Get an interaction stats alert",
	Long:  "Get an interaction stats alert",
	Args:  utils.DetermineArgs([]string{ "alertId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/alerts/{alertId}"
		alertId, args := args[0], args[1:]
		path = strings.Replace(path, "{alertId}", fmt.Sprintf("%v", alertId), -1)


		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
	Use:   "list",
	Short: "Get interaction stats alert list.",
	Long:  "Get interaction stats alert list.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/alerts"


		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
var updateCmd = &cobra.Command{
	Use:   "update [alertId]",
	Short: "Update an interaction stats alert read status",
	Long:  "Update an interaction stats alert read status",
	Args:  utils.DetermineArgs([]string{ "alertId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Unreadstatus{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/alerts/{alertId}"
		alertId, args := args[0], args[1:]
		path = strings.Replace(path, "{alertId}", fmt.Sprintf("%v", alertId), -1)


		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
