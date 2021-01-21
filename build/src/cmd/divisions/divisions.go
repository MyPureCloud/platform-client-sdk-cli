package divisions
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

var divisionsCmd = &cobra.Command{
	Use:   "divisions",
	Short: "Manages Genesys Cloud divisions",
	Long:  `Manages Genesys Cloud divisions`,
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(divisionsCmd)
}

func Cmddivisions() *cobra.Command { 
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	divisionsCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "bool", "force", "false", "Force delete this division as well as the grants and objects associated with it")
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	divisionsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "bool", "objectCount", "false", "Get count of objects in this division, grouped by type")
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	divisionsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "variable name requested by expand list")
	utils.AddFlag(listCmd.Flags(), "string", "nextPage", "", "next page token")
	utils.AddFlag(listCmd.Flags(), "string", "previousPage", "", "Previous page token")
	utils.AddFlag(listCmd.Flags(), "bool", "objectCount", "false", "Include the count of objects contained in the division")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "Optionally request specific divisions by their IDs")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Search term to filter by division name")
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	divisionsCmd.AddCommand(listCmd)
	
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	divisionsCmd.AddCommand(updateCmd)
	
	return divisionsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a division.",
	Long:  `Create a division.`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions"

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
	Use:   "delete [divisionId]",
	Short: "Delete a division.",
	Long:  `Delete a division.`,
	Args:  utils.DetermineArgs([]string{ "divisionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)

		force := utils.GetFlag(cmd.Flags(), "bool", "force")
		if force != "" {
			queryParams["force"] = force
		}
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
	Use:   "get [divisionId]",
	Short: "Returns an authorization division.",
	Long:  `Returns an authorization division.`,
	Args:  utils.DetermineArgs([]string{ "divisionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)

		objectCount := utils.GetFlag(cmd.Flags(), "bool", "objectCount")
		if objectCount != "" {
			queryParams["objectCount"] = objectCount
		}
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
	Short: "Retrieve a list of all divisions defined for the organization",
	Long:  `Retrieve a list of all divisions defined for the organization`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions"

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
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		nextPage := utils.GetFlag(cmd.Flags(), "string", "nextPage")
		if nextPage != "" {
			queryParams["nextPage"] = nextPage
		}
		previousPage := utils.GetFlag(cmd.Flags(), "string", "previousPage")
		if previousPage != "" {
			queryParams["previousPage"] = previousPage
		}
		objectCount := utils.GetFlag(cmd.Flags(), "bool", "objectCount")
		if objectCount != "" {
			queryParams["objectCount"] = objectCount
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
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
	Use:   "update [divisionId]",
	Short: "Update a division.",
	Long:  `Update a division.`,
	Args:  utils.DetermineArgs([]string{ "divisionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)

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
