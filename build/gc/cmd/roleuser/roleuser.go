package roleuser
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

var roleuserCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("roleuser"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud roleuser"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud roleuser`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(roleuserCmd)
}

func Cmdroleuser() *cobra.Command { 
	addCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", addCmd.UsageTemplate(), "PUT", "/api/v2/authorization/roles/{roleId}/users/add", utils.FormatPermissions([]string{ "authorization:grant:add",  })))
	utils.AddFileFlagIfUpsert(addCmd.Flags(), "PUT")
	roleuserCmd.AddCommand(addCmd)
	
	deleteCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "PUT", "/api/v2/authorization/roles/{roleId}/users/remove", utils.FormatPermissions([]string{ "authorization:grant:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "PUT")
	roleuserCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(getCmd.Flags(), "int", "pageNumber", "1", "Page number")
	getCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/roles/{roleId}/users", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	roleuserCmd.AddCommand(getCmd)
	
	return roleuserCmd
}

var addCmd = &cobra.Command{
	Use:   "add [roleId]",
	Short: "Sets the users for the role",
	Long:  `Sets the users for the role`,
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}/users/add"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", "add", urlString, "/api/v2/authorization/roles/{roleId}/users/add")
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
	Use:   "delete [roleId]",
	Short: "Removes the users from the role",
	Long:  `Removes the users from the role`,
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}/users/remove"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", "delete", urlString, "/api/v2/authorization/roles/{roleId}/users/remove")
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
	Use:   "get [roleId]",
	Short: "Get a list of the users in a specified role.",
	Long:  `Get a list of the users in a specified role.`,
	Args:  utils.DetermineArgs([]string{ "roleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/roles/{roleId}/users"
		roleId, args := args[0], args[1:]
		path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "get", urlString, "/api/v2/authorization/roles/{roleId}/users")
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
