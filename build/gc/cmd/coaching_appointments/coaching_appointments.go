package coaching_appointments

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
	Description = utils.FormatUsageDescription("coaching_appointments", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments", )
	coaching_appointmentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_appointments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_appointmentsCmd)
}

func Cmdcoaching_appointments() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/coaching/appointments", utils.FormatPermissions([]string{ "coaching:appointment:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The appointment to add&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CreateCoachingAppointmentRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Appointment created&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CoachingAppointmentResponse&quot;
  }
}`)
	coaching_appointmentsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/coaching/appointments/{appointmentId}", utils.FormatPermissions([]string{ "coaching:appointment:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Appointment delete request accepted.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CoachingAppointmentReference&quot;
  }
}`)
	coaching_appointmentsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/coaching/appointments/{appointmentId}", utils.FormatPermissions([]string{ "coaching:appointment:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;Retrieved appointment&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CoachingAppointmentResponse&quot;
  }
}`)
	coaching_appointmentsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "[]string", "userIds", "", "The user IDs for which to retrieve appointments - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "string", "interval", "", "Interval to filter data by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "[]string", "statuses", "", "Appointment Statuses to filter by Valid values: Scheduled, InProgress, Completed, InvalidSchedule")
	utils.AddFlag(listCmd.Flags(), "[]string", "facilitatorIds", "", "The facilitator IDs for which to retrieve appointments")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "", "Sort (by due date) either Asc or Desc Valid values: Desc, Asc")
	utils.AddFlag(listCmd.Flags(), "[]string", "relationships", "", "Relationships to filter by Valid values: Creator, Facilitator, Attendee")
	utils.AddFlag(listCmd.Flags(), "string", "completionInterval", "", "Appointment completion start and end to filter by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "string", "overdue", "", "Overdue status to filter by Valid values: True, False, Any")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/coaching/appointments", utils.FormatPermissions([]string{ "coaching:appointment:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("userIds")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;Get users coaching appointments successful&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	coaching_appointmentsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/coaching/appointments/{appointmentId}", utils.FormatPermissions([]string{ "coaching:appointment:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;The new version of the appointment&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UpdateCoachingAppointmentRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;Appointment updated&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CoachingAppointmentResponse&quot;
  }
}`)
	coaching_appointmentsCmd.AddCommand(updateCmd)
	
	return coaching_appointmentsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new appointment",
	Long:  "Create a new appointment",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments"

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
	Use:   "delete [appointmentId]",
	Short: "Delete an existing appointment",
	Long:  "Delete an existing appointment",
	Args:  utils.DetermineArgs([]string{ "appointmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)

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
	Use:   "get [appointmentId]",
	Short: "Retrieve an appointment",
	Long:  "Retrieve an appointment",
	Args:  utils.DetermineArgs([]string{ "appointmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)

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
	Short: "Get appointments for users and optional date range",
	Long:  "Get appointments for users and optional date range",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments"

		userIds := utils.GetFlag(cmd.Flags(), "[]string", "userIds")
		if userIds != "" {
			queryParams["userIds"] = userIds
		}
		interval := utils.GetFlag(cmd.Flags(), "string", "interval")
		if interval != "" {
			queryParams["interval"] = interval
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		statuses := utils.GetFlag(cmd.Flags(), "[]string", "statuses")
		if statuses != "" {
			queryParams["statuses"] = statuses
		}
		facilitatorIds := utils.GetFlag(cmd.Flags(), "[]string", "facilitatorIds")
		if facilitatorIds != "" {
			queryParams["facilitatorIds"] = facilitatorIds
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		relationships := utils.GetFlag(cmd.Flags(), "[]string", "relationships")
		if relationships != "" {
			queryParams["relationships"] = relationships
		}
		completionInterval := utils.GetFlag(cmd.Flags(), "string", "completionInterval")
		if completionInterval != "" {
			queryParams["completionInterval"] = completionInterval
		}
		overdue := utils.GetFlag(cmd.Flags(), "string", "overdue")
		if overdue != "" {
			queryParams["overdue"] = overdue
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
	Use:   "update [appointmentId]",
	Short: "Update an existing appointment",
	Long:  "Update an existing appointment",
	Args:  utils.DetermineArgs([]string{ "appointmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)

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
