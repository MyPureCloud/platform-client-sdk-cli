package workforcemanagement_businessunits_intraday_planninggroups

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_intraday_planninggroups", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups", )
	workforcemanagement_businessunits_intraday_planninggroupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_intraday_planninggroups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_intraday_planninggroupsCmd)
}

func Cmdworkforcemanagement_businessunits_intraday_planninggroups() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "time.Time", "date", "", "yyyy-MM-dd date string interpreted in the configured business unit time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups", utils.FormatPermissions([]string{ "wfm:intraday:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("date")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/WfmIntradayPlanningGroupListing"
  }
}`)
	workforcemanagement_businessunits_intraday_planninggroupsCmd.AddCommand(listCmd)
	
	return workforcemanagement_businessunits_intraday_planninggroupsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [businessUnitId]",
	Short: "Get intraday planning groups for the given date",
	Long:  "Get intraday planning groups for the given date",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)

		date := utils.GetFlag(cmd.Flags(), "time.Time", "date")
		if date != "" {
			queryParams["date"] = date
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
