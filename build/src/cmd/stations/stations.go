package stations
import (
	"fmt"
	"gc/retry"
	"gc/utils"
	"gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"strings"
)

var stationsCmd = &cobra.Command{
	Use:   "stations",
	Short: "Manages Genesys Cloud stations",
	Long:  `Manages Genesys Cloud stations`,
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(stationsCmd)
}

func Cmdstations() *cobra.Command { 
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	stationsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "userSelectable", "", "True for stations that the user can select otherwise false")
	utils.AddFlag(listCmd.Flags(), "string", "webRtcUserId", "", "Filter for the webRtc station of the webRtcUserId")
	utils.AddFlag(listCmd.Flags(), "string", "id", "", "Comma separated list of stationIds")
	utils.AddFlag(listCmd.Flags(), "string", "lineAppearanceId", "", "lineAppearanceId")
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	stationsCmd.AddCommand(listCmd)
	
	return stationsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [stationId]",
	Short: "Get station.",
	Long:  `Get station.`,
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

		retryFunc := CommandService.DetermineAction("GET", "get", urlString)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the list of available stations.",
	Long:  `Get the list of available stations.`,
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

		retryFunc := CommandService.DetermineAction("GET", "list", urlString)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		utils.Render(results)
	},
}
