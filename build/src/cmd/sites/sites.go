package sites
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

var sitesCmd = &cobra.Command{
	Use:   "sites",
	Short: "Manages Genesys Cloud sites",
	Long:  `Manages Genesys Cloud sites`,
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(sitesCmd)
}

func Cmdsites() *cobra.Command { 
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	sitesCmd.AddCommand(createCmd)
	
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	sitesCmd.AddCommand(deleteCmd)
	
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	sitesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ASC", "Sort order")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "locationId", "", "Location Id")
	utils.AddFlag(listCmd.Flags(), "bool", "managed", "", "Filter by managed")
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	sitesCmd.AddCommand(listCmd)
	
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	sitesCmd.AddCommand(updateCmd)
	
	return sitesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Site.",
	Long:  `Create a Site.`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", "create", urlString)
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
var deleteCmd = &cobra.Command{
	Use:   "delete [siteId]",
	Short: "Delete a Site by ID",
	Long:  `Delete a Site by ID`,
	Args:  utils.DetermineArgs([]string{ "siteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", "delete", urlString)
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
var getCmd = &cobra.Command{
	Use:   "get [siteId]",
	Short: "Get a Site by ID.",
	Long:  `Get a Site by ID.`,
	Args:  utils.DetermineArgs([]string{ "siteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)

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
	Short: "Get the list of Sites.",
	Long:  `Get the list of Sites.`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites"

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
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		locationId := utils.GetFlag(cmd.Flags(), "string", "locationId")
		if locationId != "" {
			queryParams["locationId"] = locationId
		}
		managed := utils.GetFlag(cmd.Flags(), "bool", "managed")
		if managed != "" {
			queryParams["managed"] = managed
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
var updateCmd = &cobra.Command{
	Use:   "update [siteId]",
	Short: "Update a Site by ID.",
	Long:  `Update a Site by ID.`,
	Args:  utils.DetermineArgs([]string{ "siteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", "update", urlString)
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
