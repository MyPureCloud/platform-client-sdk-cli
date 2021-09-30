package workforcemanagement_businessunits_planninggroups

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_planninggroups", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups", )
	workforcemanagement_businessunits_planninggroupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_planninggroups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_planninggroupsCmd)
}

func Cmdworkforcemanagement_businessunits_planninggroups() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups", utils.FormatPermissions([]string{ "wfm:planningGroup:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/CreatePlanningGroupRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/PlanningGroup"
  }
}`)
	workforcemanagement_businessunits_planninggroupsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}", utils.FormatPermissions([]string{ "wfm:planningGroup:delete",  }), utils.GenerateDevCentreLink("DELETE", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The planning group was deleted successfully"
}`)
	workforcemanagement_businessunits_planninggroupsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}", utils.FormatPermissions([]string{ "wfm:planningGroup:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/PlanningGroup"
  }
}`)
	workforcemanagement_businessunits_planninggroupsCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups", utils.FormatPermissions([]string{ "wfm:planningGroup:view", "wfm:shortTermForecast:view", "wfm:intraday:view", "wfm:agent:view", "wfm:schedule:view", "wfm:publishedSchedule:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/PlanningGroupList"
  }
}`)
	workforcemanagement_businessunits_planninggroupsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}", utils.FormatPermissions([]string{ "wfm:planningGroup:edit",  }), utils.GenerateDevCentreLink("PATCH", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/UpdatePlanningGroupRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/PlanningGroup"
  }
}`)
	workforcemanagement_businessunits_planninggroupsCmd.AddCommand(updateCmd)
	
	return workforcemanagement_businessunits_planninggroupsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [businessUnitId]",
	Short: "Adds a new planning group",
	Long:  "Adds a new planning group",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createplanninggrouprequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)


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
	Use:   "delete [businessUnitId] [planningGroupId]",
	Short: "Deletes the planning group",
	Long:  "Deletes the planning group",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "planningGroupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		planningGroupId, args := args[0], args[1:]
		path = strings.Replace(path, "{planningGroupId}", fmt.Sprintf("%v", planningGroupId), -1)


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
	Use:   "get [businessUnitId] [planningGroupId]",
	Short: "Get a planning group",
	Long:  "Get a planning group",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "planningGroupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		planningGroupId, args := args[0], args[1:]
		path = strings.Replace(path, "{planningGroupId}", fmt.Sprintf("%v", planningGroupId), -1)


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
	Use:   "list [businessUnitId]",
	Short: "Gets list of planning groups",
	Long:  "Gets list of planning groups",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)


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
	Use:   "update [businessUnitId] [planningGroupId]",
	Short: "Updates the planning group",
	Long:  "Updates the planning group",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "planningGroupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updateplanninggrouprequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		planningGroupId, args := args[0], args[1:]
		path = strings.Replace(path, "{planningGroupId}", fmt.Sprintf("%v", planningGroupId), -1)


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
