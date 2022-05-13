package workforcemanagement_calendar_data_ics

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
	Description = utils.FormatUsageDescription("workforcemanagement_calendar_data_ics", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/calendar/data/ics", )
	workforcemanagement_calendar_data_icsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_calendar_data_ics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_calendar_data_icsCmd)
}

func Cmdworkforcemanagement_calendar_data_ics() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "calendarId", "", "The id of the ics-formatted calendar - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/calendar/data/ics", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/calendar/data/ics")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("calendarId")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "text/calendar" : {
      "schema" : {
        "type" : "string"
      }
    }
  }
}`)
	workforcemanagement_calendar_data_icsCmd.AddCommand(getCmd)
	return workforcemanagement_calendar_data_icsCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get ics formatted calendar based on shareable link",
	Long:  "Get ics formatted calendar based on shareable link",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/calendar/data/ics"

		calendarId := utils.GetFlag(cmd.Flags(), "string", "calendarId")
		if calendarId != "" {
			queryParams["calendarId"] = calendarId
		}
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
