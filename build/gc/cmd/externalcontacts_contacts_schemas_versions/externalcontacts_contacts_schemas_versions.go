package externalcontacts_contacts_schemas_versions

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
	Description = utils.FormatUsageDescription("externalcontacts_contacts_schemas_versions", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions", )
	externalcontacts_contacts_schemas_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_contacts_schemas_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_contacts_schemas_versionsCmd)
}

func Cmdexternalcontacts_contacts_schemas_versions() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}", utils.FormatPermissions([]string{ "externalContacts:customFields:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DataSchema"
  }
}`)
	externalcontacts_contacts_schemas_versionsCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions", utils.FormatPermissions([]string{ "externalContacts:customFields:view",  }), utils.GenerateDevCentreLink("GET", "External Contacts", "/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/DataSchema"
  }
}`)
	externalcontacts_contacts_schemas_versionsCmd.AddCommand(listCmd)
	
	return externalcontacts_contacts_schemas_versionsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [schemaId] [versionId]",
	Short: "Get a specific version of a schema",
	Long:  "Get a specific version of a schema",
	Args:  utils.DetermineArgs([]string{ "schemaId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}"
		schemaId, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)


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
var listCmd = &cobra.Command{
	Use:   "list [schemaId]",
	Short: "Get all versions of an external contact`s schema",
	Long:  "Get all versions of an external contact`s schema",
	Args:  utils.DetermineArgs([]string{ "schemaId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/schemas/{schemaId}/versions"
		schemaId, args := args[0], args[1:]
		path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)


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
