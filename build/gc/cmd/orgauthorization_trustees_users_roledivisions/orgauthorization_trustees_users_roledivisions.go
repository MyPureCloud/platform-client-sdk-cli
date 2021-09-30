package orgauthorization_trustees_users_roledivisions

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
	Description = utils.FormatUsageDescription("orgauthorization_trustees_users_roledivisions", "SWAGGER_OVERRIDE_/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roledivisions", )
	orgauthorization_trustees_users_roledivisionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("orgauthorization_trustees_users_roledivisions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(orgauthorization_trustees_users_roledivisionsCmd)
}

func Cmdorgauthorization_trustees_users_roledivisions() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roledivisions", utils.FormatPermissions([]string{ "authorization:orgTrusteeUser:edit",  }), utils.GenerateDevCentreLink("PUT", "Organization Authorization", "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roledivisions")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "Set of roles with corresponding divisions to apply",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/RoleDivisionGrants"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/UserAuthorization"
  }
}`)
	orgauthorization_trustees_users_roledivisionsCmd.AddCommand(updateCmd)
	
	return orgauthorization_trustees_users_roledivisionsCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [trusteeOrgId] [trusteeUserId]",
	Short: "Update Trustee User Roles",
	Long:  "Update Trustee User Roles",
	Args:  utils.DetermineArgs([]string{ "trusteeOrgId", "trusteeUserId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Roledivisiongrants{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roledivisions"
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
