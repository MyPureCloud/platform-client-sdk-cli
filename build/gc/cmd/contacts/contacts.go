package contacts

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
	contactsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contacts"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud contacts"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud contacts`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contactsCmd)
}

func Cmdcontacts() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/externalcontacts/contacts", utils.FormatPermissions([]string{ "externalContacts:contact:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExternalContact&quot;
  }
}`)
	contactsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/externalcontacts/contacts/{contactId}", utils.FormatPermissions([]string{ "externalContacts:contact:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Empty&quot;
  }
}`)
	contactsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "which fields, if any, to expand (externalOrganization,externalDataSources) Valid values: externalOrganization, externalDataSources")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/contacts/{contactId}", utils.FormatPermissions([]string{ "externalContacts:contact:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExternalContact&quot;
  }
}`)
	contactsCmd.AddCommand(getCmd)
	
	utils.AddFlag(scanCmd.Flags(), "int", "limit", "", "The number of contacts per page; must be between 10 and 200, default is 100)")
	utils.AddFlag(scanCmd.Flags(), "string", "cursor", "", "Indicates where to resume query results (not required for first page), each page returns a new cursor with a 24h TTL")
	scanCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", scanCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/scan/contacts", utils.FormatPermissions([]string{ "externalContacts:contact:view",  })))
	utils.AddFileFlagIfUpsert(scanCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(scanCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CursorContactListing&quot;
  }
}`)
	contactsCmd.AddCommand(scanCmd)
	
	utils.AddFlag(searchCmd.Flags(), "int", "pageSize", "20", "Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(searchCmd.Flags(), "int", "pageNumber", "1", "Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(searchCmd.Flags(), "string", "q", "", "User supplied search keywords (no special syntax is currently supported)")
	utils.AddFlag(searchCmd.Flags(), "string", "sortOrder", "", "Sort order")
	utils.AddFlag(searchCmd.Flags(), "[]string", "expand", "", "which fields, if any, to expand Valid values: externalOrganization, externalDataSources")
	searchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", searchCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/contacts", utils.FormatPermissions([]string{ "externalContacts:contact:view",  })))
	utils.AddFileFlagIfUpsert(searchCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(searchCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ContactListing&quot;
  }
}`)
	contactsCmd.AddCommand(searchCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/externalcontacts/contacts/{contactId}", utils.FormatPermissions([]string{ "externalContacts:contact:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExternalContact&quot;
  }
}`)
	contactsCmd.AddCommand(updateCmd)
	
	return contactsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an external contact",
	Long:  `Create an external contact`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
var deleteCmd = &cobra.Command{
	Use:   "delete [contactId]",
	Short: "Delete an external contact",
	Long:  `Delete an external contact`,
	Args:  utils.DetermineArgs([]string{ "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
var getCmd = &cobra.Command{
	Use:   "get [contactId]",
	Short: "Fetch an external contact",
	Long:  `Fetch an external contact`,
	Args:  utils.DetermineArgs([]string{ "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for external contacts using paging",
	Long:  `Scan for external contacts using paging`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/scan/contacts"

		limit := utils.GetFlag(cmd.Flags(), "int", "limit")
		if limit != "" {
			queryParams["limit"] = limit
		}
		cursor := utils.GetFlag(cmd.Flags(), "string", "cursor")
		if cursor != "" {
			queryParams["cursor"] = cursor
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
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for external contacts",
	Long:  `Search for external contacts`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts"

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
var updateCmd = &cobra.Command{
	Use:   "update [contactId]",
	Short: "Update an external contact",
	Long:  `Update an external contact`,
	Args:  utils.DetermineArgs([]string{ "contactId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/contacts/{contactId}"
		contactId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactId}", fmt.Sprintf("%v", contactId), -1)

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
