package workforcemanagement_businessunits_divisionviews

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_divisionviews", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/divisionviews", )
	workforcemanagement_businessunits_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_divisionviewsCmd)
}

func Cmdworkforcemanagement_businessunits_divisionviews() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionId", "", "The divisionIds to filter by. If omitted, will return business units in all divisions")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/divisionviews", utils.FormatPermissions([]string{ "wfm:businessUnit:search",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/divisionviews")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/BusinessUnitListing"
  }
}`)
	workforcemanagement_businessunits_divisionviewsCmd.AddCommand(listCmd)
	
	return workforcemanagement_businessunits_divisionviewsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get business units across divisions",
	Long:  "Get business units across divisions",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/divisionviews"


		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
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
