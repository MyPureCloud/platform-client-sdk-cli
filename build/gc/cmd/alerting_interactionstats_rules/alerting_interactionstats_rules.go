package alerting_interactionstats_rules

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
	Description = utils.FormatUsageDescription("alerting_interactionstats_rules", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/rules", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/rules", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/rules", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/rules", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/rules", )
	alerting_interactionstats_rulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("alerting_interactionstats_rules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(alerting_interactionstats_rulesCmd)
}

func Cmdalerting_interactionstats_rules() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: notificationUsers")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/alerting/interactionstats/rules", utils.FormatPermissions([]string{ "alerting:rule:add",  }), utils.GenerateDevCentreLink("POST", "Alerting", "/api/v2/alerting/interactionstats/rules")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "AlertingRule",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/InteractionStatsRule"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/InteractionStatsRule"
  }
}`)
	alerting_interactionstats_rulesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/alerting/interactionstats/rules/{ruleId}", utils.FormatPermissions([]string{ "alerting:rule:delete",  }), utils.GenerateDevCentreLink("DELETE", "Alerting", "/api/v2/alerting/interactionstats/rules/{ruleId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Interaction stats rule deleted"
}`)
	alerting_interactionstats_rulesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: notificationUsers")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/alerting/interactionstats/rules/{ruleId}", utils.FormatPermissions([]string{ "alerting:rule:view",  }), utils.GenerateDevCentreLink("GET", "Alerting", "/api/v2/alerting/interactionstats/rules/{ruleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/InteractionStatsRule"
  }
}`)
	alerting_interactionstats_rulesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: notificationUsers")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/alerting/interactionstats/rules", utils.FormatPermissions([]string{ "alerting:rule:view",  }), utils.GenerateDevCentreLink("GET", "Alerting", "/api/v2/alerting/interactionstats/rules")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	alerting_interactionstats_rulesCmd.AddCommand(listCmd)
	
	utils.AddFlag(updateCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: notificationUsers")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/alerting/interactionstats/rules/{ruleId}", utils.FormatPermissions([]string{ "alerting:rule:edit",  }), utils.GenerateDevCentreLink("PUT", "Alerting", "/api/v2/alerting/interactionstats/rules/{ruleId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "AlertingRule",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/InteractionStatsRule"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/InteractionStatsRule"
  }
}`)
	alerting_interactionstats_rulesCmd.AddCommand(updateCmd)
	
	return alerting_interactionstats_rulesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an interaction stats rule.",
	Long:  "Create an interaction stats rule.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Interactionstatsrule{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/rules"


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

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
var deleteCmd = &cobra.Command{
	Use:   "delete [ruleId]",
	Short: "Delete an interaction stats rule.",
	Long:  "Delete an interaction stats rule.",
	Args:  utils.DetermineArgs([]string{ "ruleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/rules/{ruleId}"
		ruleId, args := args[0], args[1:]
		path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleId), -1)


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
	Use:   "get [ruleId]",
	Short: "Get an interaction stats rule.",
	Long:  "Get an interaction stats rule.",
	Args:  utils.DetermineArgs([]string{ "ruleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/rules/{ruleId}"
		ruleId, args := args[0], args[1:]
		path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleId), -1)


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
	Short: "Get an interaction stats rule list.",
	Long:  "Get an interaction stats rule list.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/rules"


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
	Use:   "update [ruleId]",
	Short: "Update an interaction stats rule",
	Long:  "Update an interaction stats rule",
	Args:  utils.DetermineArgs([]string{ "ruleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Interactionstatsrule{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/rules/{ruleId}"
		ruleId, args := args[0], args[1:]
		path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleId), -1)


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
