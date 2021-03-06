package stations_associateduser

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
	Description = utils.FormatUsageDescription("stations_associateduser", "SWAGGER_OVERRIDE_/api/v2/stations/{stationId}/associateduser", )
	stations_associateduserCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("stations_associateduser"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(stations_associateduserCmd)
}

func Cmdstations_associateduser() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/stations/{stationId}/associateduser", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Stations", "/api/v2/stations/{stationId}/associateduser")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful.&quot;
}`)
	stations_associateduserCmd.AddCommand(deleteCmd)
	
	return stations_associateduserCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [stationId]",
	Short: "Unassigns the user assigned to this station",
	Long:  "Unassigns the user assigned to this station",
	Args:  utils.DetermineArgs([]string{ "stationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/stations/{stationId}/associateduser"
		stationId, args := args[0], args[1:]
		path = strings.Replace(path, "{stationId}", fmt.Sprintf("%v", stationId), -1)

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
