package orgauthorization_trustees_users

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
	Description = utils.FormatUsageDescription("orgauthorization_trustees_users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users", )
	orgauthorization_trustees_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustees_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustees_usersCmd)
}

func Cmdorgauthorization_trustees_users() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Trust&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/TrustMemberCreate&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/TrustUser&quot;
  }
}`)
	orgauthorization_trustees_usersCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Trust deleted&quot;
}`)
	orgauthorization_trustees_usersCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/TrustUser&quot;
  }
}`)
	orgauthorization_trustees_usersCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	orgauthorization_trustees_usersCmd.AddCommand(listCmd)
	
	return orgauthorization_trustees_usersCmd
}

var createCmd = &cobra.Command{
	Use:   "create [trusteeOrgId]",
	Short: "Add a user to the trust.",
	Long:  "Add a user to the trust.",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users"
		trusteeOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeOrgId}", fmt.Sprintf("%v", trusteeOrgId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
	Use:   "delete [trusteeOrgId] [trusteeUserId]",
	Short: "Delete Trustee User",
	Long:  "Delete Trustee User",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}"
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
	Short: "Get Trustee User",
	Long:  "Get Trustee User",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}"
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
var listCmd = &cobra.Command{
	Use:   "list [trusteeOrgId]",
	Short: "The list of trustee users for this organization (i.e. users granted access to this organization).",
	Long:  "The list of trustee users for this organization (i.e. users granted access to this organization).",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users"
		trusteeOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeOrgId}", fmt.Sprintf("%v", trusteeOrgId), -1)

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
