package externalcontacts_organizations_relationships

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
	Description = utils.FormatUsageDescription("externalcontacts_organizations_relationships", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/organizations/{externalOrganizationId}/relationships", )
	externalcontacts_organizations_relationshipsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_organizations_relationships"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_organizations_relationshipsCmd)
}

func Cmdexternalcontacts_organizations_relationships() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "20", "Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(listCmd.Flags(), "string", "expand", "", "which fields, if any, to expand Valid values: externalDataSources")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "", "Sort order")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/relationships", utils.FormatPermissions([]string{ "relate:externalOrganization:view", "externalContacts:externalOrganization:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/organizations/{externalOrganizationId}/relationships")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	externalcontacts_organizations_relationshipsCmd.AddCommand(listCmd)
	
	return externalcontacts_organizations_relationshipsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [externalOrganizationId]",
	Short: "Fetch a relationship for an external organization",
	Long:  "Fetch a relationship for an external organization",
	Args:  utils.DetermineArgs([]string{ "externalOrganizationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations/{externalOrganizationId}/relationships"
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
		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
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
