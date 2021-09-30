package learning_assignments

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
	Description = utils.FormatUsageDescription("learning_assignments", "SWAGGER_OVERRIDE_/api/v2/learning/assignments", "SWAGGER_OVERRIDE_/api/v2/learning/assignments", "SWAGGER_OVERRIDE_/api/v2/learning/assignments", "SWAGGER_OVERRIDE_/api/v2/learning/assignments", "SWAGGER_OVERRIDE_/api/v2/learning/assignments", )
	learning_assignmentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_assignments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_assignmentsCmd)
}

func Cmdlearning_assignments() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/learning/assignments", utils.FormatPermissions([]string{ "learning:assignment:add",  }), utils.GenerateDevCentreLink("POST", "Learning", "/api/v2/learning/assignments")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "The Learning Assignment to be created",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/LearningAssignmentCreate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/LearningAssignment"
  }
}`)
	learning_assignmentsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/learning/assignments/{assignmentId}", utils.FormatPermissions([]string{ "learning:assignment:delete",  }), utils.GenerateDevCentreLink("DELETE", "Learning", "/api/v2/learning/assignments/{assignmentId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The learning assignment was deleted successfully"
}`)
	learning_assignmentsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Fields to expand in response Valid values: module, assessment, assessmentForm")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/learning/assignments/{assignmentId}", utils.FormatPermissions([]string{ "learning:assignment:view",  }), utils.GenerateDevCentreLink("GET", "Learning", "/api/v2/learning/assignments/{assignmentId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/LearningAssignment"
  }
}`)
	learning_assignmentsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "moduleId", "", "Specifies the ID of the learning module. Fetch assignments for learning module ID")
	utils.AddFlag(listCmd.Flags(), "string", "interval", "", "Specifies the range of dueDates to be queried. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "string", "completionInterval", "", "Specifies the range of completion dates to be used for filtering. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "string", "overdue", "Any", "Specifies if only the non-overdue (overdue is False) or overdue (overdue is True) assignments are returned. If overdue is Any or if the overdue parameter is not supplied, all assignments are returned Valid values: True, False, Any")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "pass", "Any", "Specifies if only the failed (pass is False) or passed (pass is True) assignments (completed with assessment)are returned. If pass is Any or if the pass parameter is not supplied, all assignments are returned Valid values: True, False, Any")
	utils.AddFlag(listCmd.Flags(), "float32", "minPercentageScore", "", "The minimum assessment score for an assignment (completed with assessment) to be included in the results (inclusive)")
	utils.AddFlag(listCmd.Flags(), "float32", "maxPercentageScore", "", "The maximum assessment score for an assignment (completed with assessment) to be included in the results (inclusive)")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "Desc", "Specifies result set sort order; if not specified, default sort order is descending (Desc) Valid values: Asc, Desc")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Specifies which field to sort the results by, default sort is by recommendedCompletionDate Valid values: RecommendedCompletionDate, DateModified")
	utils.AddFlag(listCmd.Flags(), "[]string", "userId", "", "Specifies the list of user IDs to be queried, up to 100 user IDs.")
	utils.AddFlag(listCmd.Flags(), "[]string", "types", "", "Specifies the assignment types, currently not supported and will be ignored. For now, all learning assignments regardless of types will be returned Valid values: Informational, AssessedContent, Assessment")
	utils.AddFlag(listCmd.Flags(), "[]string", "states", "", "Specifies the assignment states to filter by Valid values: Assigned, InProgress, Completed")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Specifies the expand option for returning additional information Valid values: ModuleSummary")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/learning/assignments", utils.FormatPermissions([]string{ "learning:assignment:view",  }), utils.GenerateDevCentreLink("GET", "Learning", "/api/v2/learning/assignments")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	learning_assignmentsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/learning/assignments/{assignmentId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "Learning", "/api/v2/learning/assignments/{assignmentId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "The Learning Assignment to be updated",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/LearningAssignmentUpdate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/LearningAssignment"
  }
}`)
	learning_assignmentsCmd.AddCommand(updateCmd)
	
	return learning_assignmentsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Learning Assignment",
	Long:  "Create Learning Assignment",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Learningassignmentcreate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments"


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
	Use:   "delete [assignmentId]",
	Short: "Delete a learning assignment",
	Long:  "Delete a learning assignment",
	Args:  utils.DetermineArgs([]string{ "assignmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments/{assignmentId}"
		assignmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{assignmentId}", fmt.Sprintf("%v", assignmentId), -1)


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
	Use:   "get [assignmentId]",
	Short: "Get Learning Assignment",
	Long:  "Get Learning Assignment",
	Args:  utils.DetermineArgs([]string{ "assignmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments/{assignmentId}"
		assignmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{assignmentId}", fmt.Sprintf("%v", assignmentId), -1)


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
	Short: "List of Learning module Assignments",
	Long:  "List of Learning module Assignments",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments"


		moduleId := utils.GetFlag(cmd.Flags(), "string", "moduleId")
		if moduleId != "" {
			queryParams["moduleId"] = moduleId
		}
		interval := utils.GetFlag(cmd.Flags(), "string", "interval")
		if interval != "" {
			queryParams["interval"] = interval
		}
		completionInterval := utils.GetFlag(cmd.Flags(), "string", "completionInterval")
		if completionInterval != "" {
			queryParams["completionInterval"] = completionInterval
		}
		overdue := utils.GetFlag(cmd.Flags(), "string", "overdue")
		if overdue != "" {
			queryParams["overdue"] = overdue
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pass := utils.GetFlag(cmd.Flags(), "string", "pass")
		if pass != "" {
			queryParams["pass"] = pass
		}
		minPercentageScore := utils.GetFlag(cmd.Flags(), "float32", "minPercentageScore")
		if minPercentageScore != "" {
			queryParams["minPercentageScore"] = minPercentageScore
		}
		maxPercentageScore := utils.GetFlag(cmd.Flags(), "float32", "maxPercentageScore")
		if maxPercentageScore != "" {
			queryParams["maxPercentageScore"] = maxPercentageScore
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		userId := utils.GetFlag(cmd.Flags(), "[]string", "userId")
		if userId != "" {
			queryParams["userId"] = userId
		}
		types := utils.GetFlag(cmd.Flags(), "[]string", "types")
		if types != "" {
			queryParams["types"] = types
		}
		states := utils.GetFlag(cmd.Flags(), "[]string", "states")
		if states != "" {
			queryParams["states"] = states
		}
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
	Use:   "update [assignmentId]",
	Short: "Update Learning Assignment",
	Long:  "Update Learning Assignment",
	Args:  utils.DetermineArgs([]string{ "assignmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Learningassignmentupdate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments/{assignmentId}"
		assignmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{assignmentId}", fmt.Sprintf("%v", assignmentId), -1)


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
