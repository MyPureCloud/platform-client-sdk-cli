package users_development_activities

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("users_development_activities", "SWAGGER_OVERRIDE_/api/v2/users/development/activities", "SWAGGER_OVERRIDE_/api/v2/users/development/activities", )
	users_development_activitiesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_development_activities"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_development_activitiesCmd)
}

func Cmdusers_development_activities() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "varType", "", "Specifies the activity type. - REQUIRED Valid values: Informational, Coaching, AssessedContent, Questionnaire")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/users/development/activities/{activityId}", utils.FormatPermissions([]string{ "learning:assignment:view", "coaching:appointment:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("varType")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/DevelopmentActivity&quot;
  }
}`)
	users_development_activitiesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "[]string", "userId", "", "Specifies the list of user IDs to be queried, up to 100 user IDs. It searches for any relationship for the userId.")
	utils.AddFlag(listCmd.Flags(), "string", "moduleId", "", "Specifies the ID of the learning module.")
	utils.AddFlag(listCmd.Flags(), "string", "interval", "", "Specifies the dateDue range to be queried. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "string", "completionInterval", "", "Specifies the range of completion dates to be used for filtering. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "string", "overdue", "Any", "Specifies if non-overdue, overdue, or all activities are returned. If not specified, all activities are returned Valid values: True, False, Any")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "Desc", "Specifies result set sort order sorted by the date due; if not specified, default sort order is descending (Desc) Valid values: Asc, Desc")
	utils.AddFlag(listCmd.Flags(), "[]string", "types", "", "Specifies the activity types. Valid values: Informational, Coaching, AssessedContent, Questionnaire")
	utils.AddFlag(listCmd.Flags(), "[]string", "statuses", "", "Specifies the activity statuses to filter by Valid values: Planned, InProgress, Completed, InvalidSchedule")
	utils.AddFlag(listCmd.Flags(), "[]string", "relationship", "", "Specifies how the current user relation should be interpreted, and filters the activities returned to only the activities that have the specified relationship. If a value besides Attendee is specified, it will only return Coaching Appointments. If not specified, no filtering is applied. Valid values: Creator, Facilitator, Attendee")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/users/development/activities", utils.FormatPermissions([]string{ "learning:assignment:view", "coaching:appointment:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	users_development_activitiesCmd.AddCommand(listCmd)
	
	return users_development_activitiesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [activityId]",
	Short: "Get a Development Activity",
	Long:  "Get a Development Activity",
	Args:  utils.DetermineArgs([]string{ "activityId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/development/activities/{activityId}"
		activityId, args := args[0], args[1:]
		path = strings.Replace(path, "{activityId}", fmt.Sprintf("%v", activityId), -1)

		varType := utils.GetFlag(cmd.Flags(), "string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
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
	Short: "Get list of Development Activities",
	Long:  "Get list of Development Activities",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/development/activities"

		userId := utils.GetFlag(cmd.Flags(), "[]string", "userId")
		if userId != "" {
			queryParams["userId"] = userId
		}
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
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		types := utils.GetFlag(cmd.Flags(), "[]string", "types")
		if types != "" {
			queryParams["types"] = types
		}
		statuses := utils.GetFlag(cmd.Flags(), "[]string", "statuses")
		if statuses != "" {
			queryParams["statuses"] = statuses
		}
		relationship := utils.GetFlag(cmd.Flags(), "[]string", "relationship")
		if relationship != "" {
			queryParams["relationship"] = relationship
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
