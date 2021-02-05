package didpools
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

var didpoolsCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("didpools"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud didpools"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud didpools`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(didpoolsCmd)
}

func Cmddidpools() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/didpools", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	didpoolsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/telephony/providers/edges/didpools/{didPoolId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	didpoolsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/didpools/{didPoolId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	didpoolsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "number", "Sort by")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "Filter by a specific list of ID`s")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/didpools", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	didpoolsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/didpools/{didPoolId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	didpoolsCmd.AddCommand(updateCmd)
	
	return didpoolsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new DID pool",
	Long:  `Create a new DID pool`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/didpools"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", "create", urlString, "/api/v2/telephony/providers/edges/didpools")
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
var deleteCmd = &cobra.Command{
	Use:   "delete [didPoolId]",
	Short: "Delete a DID Pool by ID.",
	Long:  `Delete a DID Pool by ID.`,
	Args:  utils.DetermineArgs([]string{ "didPoolId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/didpools/{didPoolId}"
		didPoolId, args := args[0], args[1:]
		path = strings.Replace(path, "{didPoolId}", fmt.Sprintf("%v", didPoolId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", "delete", urlString, "/api/v2/telephony/providers/edges/didpools/{didPoolId}")
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
var getCmd = &cobra.Command{
	Use:   "get [didPoolId]",
	Short: "Get a DID Pool by ID.",
	Long:  `Get a DID Pool by ID.`,
	Args:  utils.DetermineArgs([]string{ "didPoolId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/didpools/{didPoolId}"
		didPoolId, args := args[0], args[1:]
		path = strings.Replace(path, "{didPoolId}", fmt.Sprintf("%v", didPoolId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "get", urlString, "/api/v2/telephony/providers/edges/didpools/{didPoolId}")
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a listing of DID Pools",
	Long:  `Get a listing of DID Pools`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/didpools"

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

		retryFunc := CommandService.DetermineAction("GET", "list", urlString, "/api/v2/telephony/providers/edges/didpools")
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
var updateCmd = &cobra.Command{
	Use:   "update [didPoolId]",
	Short: "Update a DID Pool by ID.",
	Long:  `Update a DID Pool by ID.`,
	Args:  utils.DetermineArgs([]string{ "didPoolId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/didpools/{didPoolId}"
		didPoolId, args := args[0], args[1:]
		path = strings.Replace(path, "{didPoolId}", fmt.Sprintf("%v", didPoolId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", "update", urlString, "/api/v2/telephony/providers/edges/didpools/{didPoolId}")
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
