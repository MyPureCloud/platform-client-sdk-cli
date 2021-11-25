package workforcemanagement_adherence

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
	Description = utils.FormatUsageDescription("workforcemanagement_adherence", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/adherence", )
	workforcemanagement_adherenceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_adherence"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_adherenceCmd)
}

func Cmdworkforcemanagement_adherence() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "userId", "", "User Id(s) for which to fetch current schedule adherence information.  Min 1, Max of 100 userIds per request - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/adherence", utils.FormatPermissions([]string{ "wfm:realtimeAdherence:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/adherence")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("userId")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/UserScheduleAdherence"
    }
  }
}`)
	workforcemanagement_adherenceCmd.AddCommand(getCmd)
	
	return workforcemanagement_adherenceCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list of UserScheduleAdherence records for the requested users",
	Long:  "Get a list of UserScheduleAdherence records for the requested users",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/adherence"

		userId := utils.GetFlag(cmd.Flags(), "[]string", "userId")
		if userId != "" {
			queryParams["userId"] = userId
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
