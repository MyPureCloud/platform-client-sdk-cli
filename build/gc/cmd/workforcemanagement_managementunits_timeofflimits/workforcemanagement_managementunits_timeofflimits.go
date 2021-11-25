package workforcemanagement_managementunits_timeofflimits

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_timeofflimits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits", )
	workforcemanagement_managementunits_timeofflimitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_timeofflimits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_timeofflimitsCmd)
}

func Cmdworkforcemanagement_managementunits_timeofflimits() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits", utils.FormatPermissions([]string{ "wfm:timeOffLimit:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/CreateTimeOffLimitRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TimeOffLimit"
  }
}`)
	workforcemanagement_managementunits_timeofflimitsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}", utils.FormatPermissions([]string{ "wfm:timeOffLimit:delete",  }), utils.GenerateDevCentreLink("DELETE", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The time off limit object deleted successfully"
}`)
	workforcemanagement_managementunits_timeofflimitsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}", utils.FormatPermissions([]string{ "wfm:timeOffLimit:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TimeOffLimit"
  }
}`)
	workforcemanagement_managementunits_timeofflimitsCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits", utils.FormatPermissions([]string{ "wfm:timeOffLimit:view", "wfm:timeOffPlan:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TimeOffLimitListing"
  }
}`)
	workforcemanagement_managementunits_timeofflimitsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}", utils.FormatPermissions([]string{ "wfm:timeOffLimit:edit",  }), utils.GenerateDevCentreLink("PATCH", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/UpdateTimeOffLimitRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TimeOffLimit"
  }
}`)
	workforcemanagement_managementunits_timeofflimitsCmd.AddCommand(updateCmd)
	
	return workforcemanagement_managementunits_timeofflimitsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [managementUnitId]",
	Short: "Creates a new time off limit object under management unit.",
	Long:  "Creates a new time off limit object under management unit.",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createtimeofflimitrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [managementUnitId] [timeOffLimitId]",
	Short: "Deletes a time off limit object",
	Long:  "Deletes a time off limit object",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "timeOffLimitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		timeOffLimitId, args := args[0], args[1:]
		path = strings.Replace(path, "{timeOffLimitId}", fmt.Sprintf("%v", timeOffLimitId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "delete"
		const httpMethod = "DELETE"
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
var getCmd = &cobra.Command{
	Use:   "get [managementUnitId] [timeOffLimitId]",
	Short: "Gets a time off limit object",
	Long:  "Gets a time off limit object",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "timeOffLimitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		timeOffLimitId, args := args[0], args[1:]
		path = strings.Replace(path, "{timeOffLimitId}", fmt.Sprintf("%v", timeOffLimitId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [managementUnitId]",
	Short: "Gets a list of time off limit objects under management unit.",
	Long:  "Gets a list of time off limit objects under management unit.",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [managementUnitId] [timeOffLimitId]",
	Short: "Updates a time off limit object.",
	Long:  "Updates a time off limit object.",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "timeOffLimitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updatetimeofflimitrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		timeOffLimitId, args := args[0], args[1:]
		path = strings.Replace(path, "{timeOffLimitId}", fmt.Sprintf("%v", timeOffLimitId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PATCH"
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
