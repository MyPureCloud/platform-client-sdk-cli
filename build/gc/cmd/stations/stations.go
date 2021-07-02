package stations

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
	Description = utils.FormatUsageDescription("stations", "SWAGGER_OVERRIDE_/api/v2/stations", "SWAGGER_OVERRIDE_/api/v2/stations", )
	stationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("stations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(stationsCmd)
}

func Cmdstations() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/stations/{stationId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Stations", "/api/v2/stations/{stationId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Station&quot;
  }
}`)
	stationsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "userSelectable", "", "True for stations that the user can select otherwise false")
	utils.AddFlag(listCmd.Flags(), "string", "webRtcUserId", "", "Filter for the webRtc station of the webRtcUserId")
	utils.AddFlag(listCmd.Flags(), "string", "id", "", "Comma separated list of stationIds")
	utils.AddFlag(listCmd.Flags(), "string", "lineAppearanceId", "", "lineAppearanceId")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/stations", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Stations", "/api/v2/stations")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	stationsCmd.AddCommand(listCmd)
	
	return stationsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [stationId]",
	Short: "Get station.",
	Long:  "Get station.",
	Args:  utils.DetermineArgs([]string{ "stationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/stations/{stationId}"
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
	Short: "Get the list of available stations.",
	Long:  "Get the list of available stations.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/stations"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		userSelectable := utils.GetFlag(cmd.Flags(), "string", "userSelectable")
		if userSelectable != "" {
			queryParams["userSelectable"] = userSelectable
		}
		webRtcUserId := utils.GetFlag(cmd.Flags(), "string", "webRtcUserId")
		if webRtcUserId != "" {
			queryParams["webRtcUserId"] = webRtcUserId
		}
		id := utils.GetFlag(cmd.Flags(), "string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		lineAppearanceId := utils.GetFlag(cmd.Flags(), "string", "lineAppearanceId")
		if lineAppearanceId != "" {
			queryParams["lineAppearanceId"] = lineAppearanceId
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
