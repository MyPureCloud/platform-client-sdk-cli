package outbound_contactlists_divisionviews

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("outbound_contactlists_divisionviews", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlists/divisionviews", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlists/divisionviews", )
	outbound_contactlists_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlists_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlists_divisionviewsCmd)
}

func Cmdoutbound_contactlists_divisionviews() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "includeImportStatus", "false", "Include import status")
	utils.AddFlag(getCmd.Flags(), "bool", "includeSize", "false", "Include size")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/contactlists/divisionviews/{contactListId}", utils.FormatPermissions([]string{ "outbound:contactList:search",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/contactlists/divisionviews/{contactListId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ContactListDivisionView&quot;
  }
}`)
	outbound_contactlists_divisionviewsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "bool", "includeImportStatus", "false", "Include import status")
	utils.AddFlag(listCmd.Flags(), "bool", "includeSize", "false", "Include size")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size. The max that will be returned is 100.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "filterType", "Prefix", "Filter type Valid values: Equals, RegEx, Contains, Prefix, LessThan, LessThanEqualTo, GreaterThan, GreaterThanEqualTo, BeginsWith, EndsWith")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "id")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "a", "Sort order Valid values: ascending, descending")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/contactlists/divisionviews", utils.FormatPermissions([]string{ "outbound:contactList:search",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/contactlists/divisionviews")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	outbound_contactlists_divisionviewsCmd.AddCommand(listCmd)
	
	return outbound_contactlists_divisionviewsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [contactListId]",
	Short: "Get a basic ContactList information object",
	Long:  "Get a basic ContactList information object",
	Args:  utils.DetermineArgs([]string{ "contactListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/divisionviews/{contactListId}"
		contactListId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactListId}", fmt.Sprintf("%v", contactListId), -1)

		includeImportStatus := utils.GetFlag(cmd.Flags(), "bool", "includeImportStatus")
		if includeImportStatus != "" {
			queryParams["includeImportStatus"] = includeImportStatus
		}
		includeSize := utils.GetFlag(cmd.Flags(), "bool", "includeSize")
		if includeSize != "" {
			queryParams["includeSize"] = includeSize
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Query a list of simplified contact list objects.",
	Long:  "Query a list of simplified contact list objects.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/divisionviews"

		includeImportStatus := utils.GetFlag(cmd.Flags(), "bool", "includeImportStatus")
		if includeImportStatus != "" {
			queryParams["includeImportStatus"] = includeImportStatus
		}
		includeSize := utils.GetFlag(cmd.Flags(), "bool", "includeSize")
		if includeSize != "" {
			queryParams["includeSize"] = includeSize
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		filterType := utils.GetFlag(cmd.Flags(), "string", "filterType")
		if filterType != "" {
			queryParams["filterType"] = filterType
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
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
