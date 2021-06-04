package orgauthorization_trustees_users_roles

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
	Description = utils.FormatUsageDescription("orgauthorization_trustees_users_roles", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles", )
	orgauthorization_trustees_users_rolesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustees_users_roles"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustees_users_rolesCmd)
}

func Cmdorgauthorization_trustees_users_roles() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Roles deleted&quot;
}`)
	orgauthorization_trustees_users_rolesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserAuthorization&quot;
  }
}`)
	orgauthorization_trustees_users_rolesCmd.AddCommand(getCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;List of roles&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;type&quot; : &quot;string&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserAuthorization&quot;
  }
}`)
	orgauthorization_trustees_users_rolesCmd.AddCommand(updateCmd)
	
	return orgauthorization_trustees_users_rolesCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [trusteeOrgId] [trusteeUserId]",
	Short: "Delete Trustee User Roles",
	Long:  "Delete Trustee User Roles",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles"
		trusteeOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeOrgId}", fmt.Sprintf("%v", trusteeOrgId), -1)
		trusteeUserId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeUserId}", fmt.Sprintf("%v", trusteeUserId), -1)

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
var getCmd = &cobra.Command{
	Use:   "get [trusteeOrgId] [trusteeUserId]",
	Short: "Get Trustee User Roles",
	Long:  "Get Trustee User Roles",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles"
		trusteeOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeOrgId}", fmt.Sprintf("%v", trusteeOrgId), -1)
		trusteeUserId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeUserId}", fmt.Sprintf("%v", trusteeUserId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [trusteeOrgId] [trusteeUserId]",
	Short: "Update Trustee User Roles",
	Long:  "Update Trustee User Roles",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles"
		trusteeOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeOrgId}", fmt.Sprintf("%v", trusteeOrgId), -1)
		trusteeUserId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeUserId}", fmt.Sprintf("%v", trusteeUserId), -1)

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
