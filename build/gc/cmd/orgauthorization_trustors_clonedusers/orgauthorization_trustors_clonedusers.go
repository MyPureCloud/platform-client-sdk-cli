package orgauthorization_trustors_clonedusers

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
	Description = utils.FormatUsageDescription("orgauthorization_trustors_clonedusers", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers", )
	orgauthorization_trustors_clonedusersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustors_clonedusers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustors_clonedusersCmd)
}

func Cmdorgauthorization_trustors_clonedusers() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:delete",  }), utils.GenerateDevCentreLink("DELETE", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Trust deleted",
  "content" : { }
}`)
	orgauthorization_trustors_clonedusersCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:view",  }), utils.GenerateDevCentreLink("GET", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ClonedUser"
      }
    }
  }
}`)
	orgauthorization_trustors_clonedusersCmd.AddCommand(getCmd)

	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:view",  }), utils.GenerateDevCentreLink("GET", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ClonedUserEntityListing"
      }
    }
  }
}`)
	orgauthorization_trustors_clonedusersCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:add",  }), utils.GenerateDevCentreLink("PUT", "Organization Authorization", "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", ``)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ClonedUser"
      }
    }
  }
}`)
	orgauthorization_trustors_clonedusersCmd.AddCommand(updateCmd)
	return orgauthorization_trustors_clonedusersCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [trustorOrgId] [trusteeUserId]",
	Short: "Delete Cloned User",
	Long:  "Delete Cloned User",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}"
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [trustorOrgId] [trusteeUserId]",
	Short: "Get Cloned User",
	Long:  "Get Cloned User",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}"
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list [trustorOrgId]",
	Short: "The list of cloned users in the trustor organization (i.e. users with a native user record).",
	Long:  "The list of cloned users in the trustor organization (i.e. users with a native user record).",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers"
		trustorOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trustorOrgId}", fmt.Sprintf("%v", trustorOrgId), -1)

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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [trustorOrgId] [trusteeUserId]",
	Short: "Creates a clone of the trustee user in the trustor org.",
	Long:  "Creates a clone of the trustee user in the trustor org.",
	Args:  utils.DetermineArgs([]string{ "trustorOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers/{trusteeUserId}"
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
