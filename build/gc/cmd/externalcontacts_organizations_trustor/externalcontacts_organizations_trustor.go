package externalcontacts_organizations_trustor

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
	Description = utils.FormatUsageDescription("externalcontacts_organizations_trustor", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor", )
	externalcontacts_organizations_trustorCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_organizations_trustor"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_organizations_trustorCmd)
}

func Cmdexternalcontacts_organizations_trustor() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor", utils.FormatPermissions([]string{ "externalContacts:externalOrganization:edit",  }), utils.GenerateDevCentreLink("DELETE", "External Contacts", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Trustor link has been deleted",
  "content" : { }
}`)
	externalcontacts_organizations_trustorCmd.AddCommand(deleteCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor/{trustorId}", utils.FormatPermissions([]string{ "externalContacts:externalOrganization:edit",  }), utils.GenerateDevCentreLink("PUT", "External Contacts", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor/{trustorId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", ``)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ExternalOrganizationTrustorLink"
      }
    }
  }
}`)
	externalcontacts_organizations_trustorCmd.AddCommand(updateCmd)
	return externalcontacts_organizations_trustorCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [externalOrganizationId]",
	Short: "Unlink the Trustor for this External Organization",
	Long:  "Unlink the Trustor for this External Organization",
	Args:  utils.DetermineArgs([]string{ "externalOrganizationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor"
		externalOrganizationId, args := args[0], args[1:]
		path = strings.Replace(path, "{externalOrganizationId}", fmt.Sprintf("%v", externalOrganizationId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [externalOrganizationId] [trustorId]",
	Short: "Links a Trustor with an External Organization",
	Long:  "Links a Trustor with an External Organization",
	Args:  utils.DetermineArgs([]string{ "externalOrganizationId", "trustorId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations/{externalOrganizationId}/trustor/{trustorId}"
		externalOrganizationId, args := args[0], args[1:]
		path = strings.Replace(path, "{externalOrganizationId}", fmt.Sprintf("%v", externalOrganizationId), -1)
		trustorId, args := args[0], args[1:]
		path = strings.Replace(path, "{trustorId}", fmt.Sprintf("%v", trustorId), -1)

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
