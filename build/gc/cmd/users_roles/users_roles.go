package users_roles

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
	users_rolesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_roles"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud users_roles"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud users_roles`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_rolesCmd)
}

func Cmdusers_roles() *cobra.Command { 
	addCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", addCmd.UsageTemplate(), "PUT", "/api/v2/authorization/roles/{roleId}/users/add", utils.FormatPermissions([]string{ "authorization:grant:add",  })))
	utils.AddFileFlagIfUpsert(addCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;List of user IDs&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;type&quot; : &quot;string&quot;
    }
  }
}`)
	utils.AddPaginateFlagsIfListingResponse(addCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;type&quot; : &quot;string&quot;
    }
  }
}`)
	users_rolesCmd.AddCommand(addCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "PUT", "/api/v2/authorization/roles/{roleId}/users/remove", utils.FormatPermissions([]string{ "authorization:grant:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;List of user IDs&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;type&quot; : &quot;string&quot;
    }
  }
}`)
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;type&quot; : &quot;string&quot;
    }
  }
}`)
	users_rolesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(getCmd.Flags(), "int", "pageNumber", "1", "Page number")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/roles/{roleId}/users", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserEntityListing&quot;
  }
}`)
	users_rolesCmd.AddCommand(getCmd)
	
	return users_rolesCmd
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

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
