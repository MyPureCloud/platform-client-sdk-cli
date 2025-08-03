package taskmanagement_workitems_users_wrapups

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
	Description = utils.FormatUsageDescription("taskmanagement_workitems_users_wrapups", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups", )
	taskmanagement_workitems_users_wrapupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_users_wrapups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_users_wrapupsCmd)
}

func Cmdtaskmanagement_workitems_users_wrapups() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "expands", "", "Which fields, if any, to expand. Valid values: wrapupCode")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Limit the number of entities to return. It is not guaranteed that the requested number of entities will be filled in a single request. If an `after` key is returned as part of the response it is possible that more entities that match the filter criteria exist. Maximum of 50.")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "descending", "Ascending or descending sort order Valid values: ascending, descending")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups", utils.FormatPermissions([]string{ "workitems:wrapup:view",  }), utils.GenerateDevCentreLink("GET", "Task Management", "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups")))
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
	taskmanagement_workitems_users_wrapupsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups", utils.FormatPermissions([]string{ "workitems:wrapup:edit",  }), utils.GenerateDevCentreLink("PATCH", "Task Management", "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "Request body to add/remove a wrapup code for a workitem",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkitemWrapupUpdate"
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
        "$ref" : "#/components/schemas/WorkitemWrapup"
      }
    }
  }
}`)
	taskmanagement_workitems_users_wrapupsCmd.AddCommand(updateCmd)
	return taskmanagement_workitems_users_wrapupsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [workitemId] [userId]",
	Short: "Get all wrapup codes added for the given user for a workitem.",
	Long:  "Get all wrapup codes added for the given user for a workitem.",
	Args:  utils.DetermineArgs([]string{ "workitemId", "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups"
		workitemId, args := args[0], args[1:]
		path = strings.Replace(path, "{workitemId}", fmt.Sprintf("%v", workitemId), -1)
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		expands := utils.GetFlag(cmd.Flags(), "string", "expands")
		if expands != "" {
			queryParams["expands"] = expands
		}
		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
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
	Use:   "update [workitemId] [userId]",
	Short: "Add/Remove a wrapup code for a given user in a workitem.",
	Long:  "Add/Remove a wrapup code for a given user in a workitem.",
	Args:  utils.DetermineArgs([]string{ "workitemId", "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Workitemwrapupupdate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups"
		workitemId, args := args[0], args[1:]
		path = strings.Replace(path, "{workitemId}", fmt.Sprintf("%v", workitemId), -1)
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

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
