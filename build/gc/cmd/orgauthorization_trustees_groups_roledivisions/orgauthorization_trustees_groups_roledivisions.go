package orgauthorization_trustees_groups_roledivisions

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
	Description = utils.FormatUsageDescription("orgauthorization_trustees_groups_roledivisions", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions", )
	orgauthorization_trustees_groups_roledivisionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustees_groups_roledivisions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustees_groups_roledivisionsCmd)
}

func Cmdorgauthorization_trustees_groups_roledivisions() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions", utils.FormatPermissions([]string{ "authorization:orgTrusteeGroup:edit",  }), utils.GenerateDevCentreLink("PUT", "Organization Authorization", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "Set of roles with corresponding divisions to apply",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/RoleDivisionGrants"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UserAuthorization"
      }
    }
  }
}`)
	orgauthorization_trustees_groups_roledivisionsCmd.AddCommand(updateCmd)
	return orgauthorization_trustees_groups_roledivisionsCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [trusteeOrgId] [trusteeGroupId]",
	Short: "Update Trustee Group Roles",
	Long:  "Update Trustee Group Roles",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", "trusteeGroupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Roledivisiongrants{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roledivisions"
		trusteeOrgId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeOrgId}", fmt.Sprintf("%v", trusteeOrgId), -1)
		trusteeGroupId, args := args[0], args[1:]
		path = strings.Replace(path, "{trusteeGroupId}", fmt.Sprintf("%v", trusteeGroupId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
