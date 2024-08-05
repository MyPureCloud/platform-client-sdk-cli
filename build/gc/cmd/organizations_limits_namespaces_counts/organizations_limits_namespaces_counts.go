package organizations_limits_namespaces_counts

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
	Description = utils.FormatUsageDescription("organizations_limits_namespaces_counts", "SWAGGER_OVERRIDE_/api/v2/organizations/limits/namespaces/{namespaceName}/counts", )
	organizations_limits_namespaces_countsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations_limits_namespaces_counts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizations_limits_namespaces_countsCmd)
}

func Cmdorganizations_limits_namespaces_counts() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "cursor", "", "Cursor provided when retrieving the last page")
	utils.AddFlag(listCmd.Flags(), "string", "entityId", "", "entity id of the count")
	utils.AddFlag(listCmd.Flags(), "string", "userId", "", "userid of the count")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/organizations/limits/namespaces/{namespaceName}/counts", utils.FormatPermissions([]string{ "limits:count:view",  }), utils.GenerateDevCentreLink("GET", "Organization", "/api/v2/organizations/limits/namespaces/{namespaceName}/counts")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	organizations_limits_namespaces_countsCmd.AddCommand(listCmd)
	return organizations_limits_namespaces_countsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [namespaceName]",
	Short: "Get estimated limit counts for a namespace. This is not a source of truth for limit values but a record of estimates to facilitate limit threshold tracking.",
	Long:  "Get estimated limit counts for a namespace. This is not a source of truth for limit values but a record of estimates to facilitate limit threshold tracking.",
	Args:  utils.DetermineArgs([]string{ "namespaceName", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/organizations/limits/namespaces/{namespaceName}/counts"
		namespaceName, args := args[0], args[1:]
		path = strings.Replace(path, "{namespaceName}", fmt.Sprintf("%v", namespaceName), -1)

		cursor := utils.GetFlag(cmd.Flags(), "string", "cursor")
		if cursor != "" {
			queryParams["cursor"] = cursor
		}
		entityId := utils.GetFlag(cmd.Flags(), "string", "entityId")
		if entityId != "" {
			queryParams["entityId"] = entityId
		}
		userId := utils.GetFlag(cmd.Flags(), "string", "userId")
		if userId != "" {
			queryParams["userId"] = userId
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
