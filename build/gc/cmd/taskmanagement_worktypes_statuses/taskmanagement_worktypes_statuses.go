package taskmanagement_worktypes_statuses

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
	Description = utils.FormatUsageDescription("taskmanagement_worktypes_statuses", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/statuses", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/statuses", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/statuses", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/statuses", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/statuses", )
	taskmanagement_worktypes_statusesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_worktypes_statuses"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_worktypes_statusesCmd)
}

func Cmdtaskmanagement_worktypes_statuses() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses", utils.FormatPermissions([]string{ "workitems:status:add",  }), utils.GenerateDevCentreLink("POST", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Status",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemStatusCreate"
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
        "$ref" : "#/components/schemas/WorkitemStatus"
      }
    }
  }
}`)
	taskmanagement_worktypes_statusesCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}", utils.FormatPermissions([]string{ "workitems:status:delete",  }), utils.GenerateDevCentreLink("DELETE", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Deleted successfully",
  "content" : { }
}`)
	taskmanagement_worktypes_statusesCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}", utils.FormatPermissions([]string{ "workitems:status:view",  }), utils.GenerateDevCentreLink("GET", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemStatus"
      }
    }
  }
}`)
	taskmanagement_worktypes_statusesCmd.AddCommand(getCmd)

	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses", utils.FormatPermissions([]string{ "workitems:status:view",  }), utils.GenerateDevCentreLink("GET", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemStatusListing"
      }
    }
  }
}`)
	taskmanagement_worktypes_statusesCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}", utils.FormatPermissions([]string{ "workitems:status:edit",  }), utils.GenerateDevCentreLink("PATCH", "Task Management", "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Status",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemStatusUpdate"
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
        "$ref" : "#/components/schemas/WorkitemStatus"
      }
    }
  }
}`)
	taskmanagement_worktypes_statusesCmd.AddCommand(updateCmd)
	return taskmanagement_worktypes_statusesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [worktypeId]",
	Short: "Add a status to a worktype",
	Long:  "Add a status to a worktype",
	Args:  utils.DetermineArgs([]string{ "worktypeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Workitemstatuscreate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses"
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
	Use:   "delete [worktypeId] [statusId]",
	Short: "Delete a status",
	Long:  "Delete a status",
	Args:  utils.DetermineArgs([]string{ "worktypeId", "statusId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)
		statusId, args := args[0], args[1:]
		path = strings.Replace(path, "{statusId}", fmt.Sprintf("%v", statusId), -1)

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
	Use:   "get [worktypeId] [statusId]",
	Short: "Get a status",
	Long:  "Get a status",
	Args:  utils.DetermineArgs([]string{ "worktypeId", "statusId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)
		statusId, args := args[0], args[1:]
		path = strings.Replace(path, "{statusId}", fmt.Sprintf("%v", statusId), -1)

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
	Short: "Get list of statuses for this worktype.",
	Long:  "Get list of statuses for this worktype.",
	Args:  utils.DetermineArgs([]string{ "worktypeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses"
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
	Use:   "update [worktypeId] [statusId]",
	Short: "Update the attributes of a status",
	Long:  "Update the attributes of a status",
	Args:  utils.DetermineArgs([]string{ "worktypeId", "statusId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Workitemstatusupdate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}"
		worktypeId, args := args[0], args[1:]
		path = strings.Replace(path, "{worktypeId}", fmt.Sprintf("%v", worktypeId), -1)
		statusId, args := args[0], args[1:]
		path = strings.Replace(path, "{statusId}", fmt.Sprintf("%v", statusId), -1)

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
