package coaching_appointments_me

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
	Description = utils.FormatUsageDescription("coaching_appointments_me", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/me", )
	coaching_appointments_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_appointments_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_appointments_meCmd)
}

func Cmdcoaching_appointments_me() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "interval", "", "Interval to filter data by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "[]string", "statuses", "", "Appointment Statuses to filter by Valid values: Scheduled, InProgress, Completed, InvalidSchedule")
	utils.AddFlag(listCmd.Flags(), "[]string", "facilitatorIds", "", "The facilitator IDs for which to retrieve appointments")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "", "Sort (by due date) either Asc or Desc Valid values: Desc, Asc")
	utils.AddFlag(listCmd.Flags(), "[]string", "relationships", "", "Relationships to filter by Valid values: Creator, Facilitator, Attendee")
	utils.AddFlag(listCmd.Flags(), "string", "completionInterval", "", "Appointment completion start and end to filter by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss")
	utils.AddFlag(listCmd.Flags(), "string", "overdue", "", "Overdue status to filter by Valid values: True, False, Any")
	utils.AddFlag(listCmd.Flags(), "string", "intervalCondition", "", "Filter condition for interval Valid values: StartsIn, Overlaps")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/coaching/appointments/me", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Coaching", "/api/v2/coaching/appointments/me")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "Get my coaching appointments successful",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	coaching_appointments_meCmd.AddCommand(listCmd)
	return coaching_appointments_meCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get my appointments for a given date range",
	Long:  "Get my appointments for a given date range",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/me"

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
		intervalCondition := utils.GetFlag(cmd.Flags(), "string", "intervalCondition")
		if intervalCondition != "" {
			queryParams["intervalCondition"] = intervalCondition
		}
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
