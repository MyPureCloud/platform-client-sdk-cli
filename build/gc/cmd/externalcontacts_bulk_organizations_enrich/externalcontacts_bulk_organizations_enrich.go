package externalcontacts_bulk_organizations_enrich

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
	Description = utils.FormatUsageDescription("externalcontacts_bulk_organizations_enrich", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/bulk/organizations/enrich", )
	externalcontacts_bulk_organizations_enrichCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_bulk_organizations_enrich"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_bulk_organizations_enrichCmd)
}

func Cmdexternalcontacts_bulk_organizations_enrich() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "dryRun", "", "If true, the request will not make any modifications, but will show you what the end result *would* be.")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/externalcontacts/bulk/organizations/enrich", utils.FormatPermissions([]string{ "externalContacts:externalOrganization:enrich",  }), utils.GenerateDevCentreLink("POST", "External Contacts", "/api/v2/externalcontacts/bulk/organizations/enrich")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "External Organization Enrich Requests",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BulkOrganizationsEnrichRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BulkOrganizationsEnrichResponse"
      }
    }
  }
}`)
	externalcontacts_bulk_organizations_enrichCmd.AddCommand(createCmd)
	return externalcontacts_bulk_organizations_enrichCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Bulk enrich external organizations - Run up to 10 Enrich operations per request",
	Long:  "Bulk enrich external organizations - Run up to 10 Enrich operations per request",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Bulkorganizationsenrichrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/bulk/organizations/enrich"

		dryRun := utils.GetFlag(cmd.Flags(), "bool", "dryRun")
		if dryRun != "" {
			queryParams["dryRun"] = dryRun
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
