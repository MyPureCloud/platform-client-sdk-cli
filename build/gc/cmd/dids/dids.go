package dids
import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

var didsCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("dids"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud dids"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud dids`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(didsCmd)
}

func Cmddids() *cobra.Command { 
	getsdidCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getsdidCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/dids/{didId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(getsdidCmd.Flags(), "GET")
	didsCmd.AddCommand(getsdidCmd)
	
	utils.AddFlag(listsdidCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listsdidCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listsdidCmd.Flags(), "string", "sortBy", "number", "Sort by")
	utils.AddFlag(listsdidCmd.Flags(), "string", "sortOrder", "ASC", "Sort order")
	utils.AddFlag(listsdidCmd.Flags(), "string", "phoneNumber", "", "Filter by phoneNumber")
	utils.AddFlag(listsdidCmd.Flags(), "string", "ownerId", "", "Filter by the owner of a phone number")
	utils.AddFlag(listsdidCmd.Flags(), "string", "didPoolId", "", "Filter by the DID Pool assignment")
	utils.AddFlag(listsdidCmd.Flags(), "[]string", "id", "", "Filter by a specific list of ID`s")
	listsdidCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listsdidCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/dids", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(listsdidCmd.Flags(), "GET")
	didsCmd.AddCommand(listsdidCmd)
	
	updatesdidCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updatesdidCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/dids/{didId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(updatesdidCmd.Flags(), "PUT")
	didsCmd.AddCommand(updatesdidCmd)
	
	return didsCmd
}

var getsdidCmd = &cobra.Command{
	Use:   "getsdid [didId]",
	Short: "Get a DID by ID.",
	Long:  `Get a DID by ID.`,
	Args:  utils.DetermineArgs([]string{ "didId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/dids/{didId}"
		didId, args := args[0], args[1:]
		path = strings.Replace(path, "{didId}", fmt.Sprintf("%v", didId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "getsdid", urlString, "/api/v2/telephony/providers/edges/dids/{didId}")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listsdidCmd = &cobra.Command{
	Use:   "listsdid",
	Short: "Get a listing of DIDs",
	Long:  `Get a listing of DIDs`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/dids"

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
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		phoneNumber := utils.GetFlag(cmd.Flags(), "string", "phoneNumber")
		if phoneNumber != "" {
			queryParams["phoneNumber"] = phoneNumber
		}
		ownerId := utils.GetFlag(cmd.Flags(), "string", "ownerId")
		if ownerId != "" {
			queryParams["ownerId"] = ownerId
		}
		didPoolId := utils.GetFlag(cmd.Flags(), "string", "didPoolId")
		if didPoolId != "" {
			queryParams["didPoolId"] = didPoolId
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "listsdid", urlString, "/api/v2/telephony/providers/edges/dids")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updatesdidCmd = &cobra.Command{
	Use:   "updatesdid [didId]",
	Short: "Update a DID by ID.",
	Long:  `Update a DID by ID.`,
	Args:  utils.DetermineArgs([]string{ "didId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/dids/{didId}"
		didId, args := args[0], args[1:]
		path = strings.Replace(path, "{didId}", fmt.Sprintf("%v", didId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", "updatesdid", urlString, "/api/v2/telephony/providers/edges/dids/{didId}")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
