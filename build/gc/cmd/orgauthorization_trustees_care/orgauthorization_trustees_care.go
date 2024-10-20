package orgauthorization_trustees_care

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
	Description = utils.FormatUsageDescription("orgauthorization_trustees_care", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/care", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/care", )
	orgauthorization_trustees_careCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustees_care"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustees_careCmd)
}

func Cmdorgauthorization_trustees_care() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "assignDefaultRole", "", "Assign Admin role to default pairing with Customer Care")
	utils.AddFlag(createCmd.Flags(), "bool", "autoExpire", "", "Automatically expire pairing after 30 days")
	utils.AddFlag(createCmd.Flags(), "bool", "assignFullAccess", "", "Grant Customer Care full access to the organization")
	utils.AddFlag(createCmd.Flags(), "bool", "allowTrustedUserAccess", "", "Make Customer Care a Trusted User")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/orgauthorization/trustees/care", utils.FormatPermissions([]string{ "authorization:orgTrustee:add", "authorization:orgTrusteeUser:add",  }), utils.GenerateDevCentreLink("POST", "Organization Authorization", "/api/v2/orgauthorization/trustees/care")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TrustEntityListing"
      }
    }
  }
}`)
	orgauthorization_trustees_careCmd.AddCommand(createCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/orgauthorization/trustees/care", utils.FormatPermissions([]string{ "authorization:orgTrustee:view", "authorization:orgTrusteeUser:view",  }), utils.GenerateDevCentreLink("GET", "Organization Authorization", "/api/v2/orgauthorization/trustees/care")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TrusteeReferenceList"
      }
    }
  }
}`)
	orgauthorization_trustees_careCmd.AddCommand(getCmd)
	return orgauthorization_trustees_careCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new organization authorization trust with Customer Care. This is required to grant your regional Customer Care organization access to your organization.",
	Long:  "Create a new organization authorization trust with Customer Care. This is required to grant your regional Customer Care organization access to your organization.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/care"

		assignDefaultRole := utils.GetFlag(cmd.Flags(), "bool", "assignDefaultRole")
		if assignDefaultRole != "" {
			queryParams["assignDefaultRole"] = assignDefaultRole
		}
		autoExpire := utils.GetFlag(cmd.Flags(), "bool", "autoExpire")
		if autoExpire != "" {
			queryParams["autoExpire"] = autoExpire
		}
		assignFullAccess := utils.GetFlag(cmd.Flags(), "bool", "assignFullAccess")
		if assignFullAccess != "" {
			queryParams["assignFullAccess"] = assignFullAccess
		}
		allowTrustedUserAccess := utils.GetFlag(cmd.Flags(), "bool", "allowTrustedUserAccess")
		if allowTrustedUserAccess != "" {
			queryParams["allowTrustedUserAccess"] = allowTrustedUserAccess
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "create"
		const httpMethod = "POST"
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
	Use:   "get",
	Short: "Get Customer Care organization ids.",
	Long:  "Get Customer Care organization ids.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/care"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
