package taskmanagement_worktypes_flows_oncreate_rules

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
	Description = utils.FormatUsageDescription("taskmanagement_worktypes_flows_oncreate_rules", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules", )
	taskmanagement_worktypes_flows_oncreate_rulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_worktypes_flows_oncreate_rules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_worktypes_flows_oncreate_rulesCmd)
}

func Cmdtaskmanagement_worktypes_flows_oncreate_rules() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules", utils.FormatPermissions([]string{ "workitems:flowRuleOnCreate:add",  }), utils.GenerateDevCentreLink("POST", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Rule",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemOnCreateRuleCreate"
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
        "$ref" : "#/components/schemas/WorkitemOnCreateRule"
      }
    }
  }
}`)
	taskmanagement_worktypes_flows_oncreate_rulesCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}", utils.FormatPermissions([]string{ "workitems:flowRuleOnCreate:delete",  }), utils.GenerateDevCentreLink("DELETE", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Deleted successfully",
  "content" : { }
}`)
	taskmanagement_worktypes_flows_oncreate_rulesCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}", utils.FormatPermissions([]string{ "workitems:flowRuleOnCreate:view",  }), utils.GenerateDevCentreLink("GET", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemOnCreateRule"
      }
    }
  }
}`)
	taskmanagement_worktypes_flows_oncreate_rulesCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an `after` key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 200.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules", utils.FormatPermissions([]string{ "workitems:flowRuleOnCreate:view",  }), utils.GenerateDevCentreLink("GET", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules")))
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
	taskmanagement_worktypes_flows_oncreate_rulesCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}", utils.FormatPermissions([]string{ "workitems:flowRuleOnCreate:edit",  }), utils.GenerateDevCentreLink("PATCH", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Rule",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemOnCreateRuleUpdate"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemOnCreateRule"
      }
    }
  }
}`)
	taskmanagement_worktypes_flows_oncreate_rulesCmd.AddCommand(updateCmd)
	return taskmanagement_worktypes_flows_oncreate_rulesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [worktypeId]",
	Short: "Add an on-create rule to a worktype",
	Long:  "Add an on-create rule to a worktype",
	Args:  utils.DetermineArgs([]string{ "worktypeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Workitemoncreaterulecreate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)

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
var deleteCmd = &cobra.Command{
	Use:   "delete [worktypeId] [ruleId]",
	Short: "Delete a rule",
	Long:  "Delete a rule",
	Args:  utils.DetermineArgs([]string{ "worktypeId", "ruleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)
		ruleId, args := args[0], args[1:]
		path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleId), -1)

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
var getCmd = &cobra.Command{
	Use:   "get [worktypeId] [ruleId]",
	Short: "Get an on-create rule",
	Long:  "Get an on-create rule",
	Args:  utils.DetermineArgs([]string{ "worktypeId", "ruleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)
		ruleId, args := args[0], args[1:]
		path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [worktypeId]",
	Short: "Get all on-create rules for a worktype",
	Long:  "Get all on-create rules for a worktype",
	Args:  utils.DetermineArgs([]string{ "worktypeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)

		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
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
var updateCmd = &cobra.Command{
	Use:   "update [worktypeId] [ruleId]",
	Short: "Update the attributes of a rule",
	Long:  "Update the attributes of a rule",
	Args:  utils.DetermineArgs([]string{ "worktypeId", "ruleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Workitemoncreateruleupdate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)
		ruleId, args := args[0], args[1:]
		path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleId), -1)

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
