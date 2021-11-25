package orgauthorization_trustors_users

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("orgauthorization_trustors_users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/users", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/users", )
	orgauthorization_trustors_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustors_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustors_usersCmd)
}

func Cmdorgauthorization_trustors_users() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:delete",  }), utils.GenerateDevCentreLink("DELETE", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Trust deleted"
}`)
	orgauthorization_trustors_usersCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:view",  }), utils.GenerateDevCentreLink("GET", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TrustUser"
  }
}`)
	orgauthorization_trustors_usersCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:view",  }), utils.GenerateDevCentreLink("GET", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	orgauthorization_trustors_usersCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:add",  }), utils.GenerateDevCentreLink("PUT", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", ``)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TrustUser"
  }
}`)
	orgauthorization_trustors_usersCmd.AddCommand(updateCmd)
	
	return orgauthorization_trustors_usersCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [trustorOrgId] [trusteeUserId]",
	Short: "Delete Trustee User",
	Long:  "Delete Trustee User",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}"
		trustorOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trustorOrgId}", fmt.Sprintf("%v", trustorOrgId), -1)
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

		const opId = "delete"
		const httpMethod = "DELETE"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
	Use:   "get [trustorOrgId] [trusteeUserId]",
	Short: "Get Trustee User",
	Long:  "Get Trustee User",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}"
		trustorOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trustorOrgId}", fmt.Sprintf("%v", trustorOrgId), -1)
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

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
	Use:   "list [trustorOrgId]",
	Short: "The list of users in the trustor organization (i.e. users granted access).",
	Long:  "The list of users in the trustor organization (i.e. users granted access).",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/users"
		trustorOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trustorOrgId}", fmt.Sprintf("%v", trustorOrgId), -1)

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

		const opId = "list"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
	Use:   "update [trustorOrgId] [trusteeUserId]",
	Short: "Add a Trustee user to the trust.",
	Long:  "Add a Trustee user to the trust.",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/users/{trusteeUserId}"
		trustorOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trustorOrgId}", fmt.Sprintf("%v", trustorOrgId), -1)
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

		const opId = "update"
		const httpMethod = "PUT"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
