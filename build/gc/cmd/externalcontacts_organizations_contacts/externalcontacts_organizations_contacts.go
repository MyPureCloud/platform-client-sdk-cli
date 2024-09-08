package externalcontacts_organizations_contacts

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
	Description = utils.FormatUsageDescription("externalcontacts_organizations_contacts", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts", )
	externalcontacts_organizations_contactsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_organizations_contacts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_organizations_contactsCmd)
}

func Cmdexternalcontacts_organizations_contacts() *cobra.Command { 
	utils.AddFlag(searchCmd.Flags(), "int", "pageSize", "20", "Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)")
	utils.AddFlag(searchCmd.Flags(), "int", "pageNumber", "1", "Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)")
	utils.AddFlag(searchCmd.Flags(), "string", "q", "", "User supplied search keywords (no special syntax is currently supported)")
	utils.AddFlag(searchCmd.Flags(), "string", "sortOrder", "", "The External Contact field to sort by. Any of: [firstName, lastName, middleName, title]. Direction: [asc, desc]. e.g. firstName:asc, title:desc")
	utils.AddFlag(searchCmd.Flags(), "[]string", "expand", "", "which fields, if any, to expand Valid values: externalOrganization, externalDataSources, identifiers, externalSources")
	searchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", searchCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts", utils.FormatPermissions([]string{ "relate:contact:view", "externalContacts:contact:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts")))
	utils.AddFileFlagIfUpsert(searchCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(searchCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	externalcontacts_organizations_contactsCmd.AddCommand(searchCmd)
	return externalcontacts_organizations_contactsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var searchCmd = &cobra.Command{
	Use:   "search [externalOrganizationId]",
	Short: "Search for external contacts in an external organization",
	Long:  "Search for external contacts in an external organization",
	Args:  utils.DetermineArgs([]string{ "externalOrganizationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations/{externalOrganizationId}/contacts"
		externalOrganizationId, args := args[0], args[1:]
		path = strings.Replace(path, "{externalOrganizationId}", fmt.Sprintf("%v", externalOrganizationId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		q := utils.GetFlag(cmd.Flags(), "string", "q")
		if q != "" {
			queryParams["q"] = q
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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

		const opId = "search"
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
