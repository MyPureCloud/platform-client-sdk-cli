package organizations

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
	organizationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud organizations"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud organizations`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizationsCmd)
}

func Cmdorganizations() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/externalcontacts/organizations", utils.FormatPermissions([]string{ "relate:externalOrganization:add", "externalContacts:externalOrganization:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExternalOrganization&quot;
  }
}`)
	organizationsCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/externalcontacts/organizations/{externalOrganizationId}", utils.FormatPermissions([]string{ "relate:externalOrganization:delete", "externalContacts:externalOrganization:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Empty&quot;
  }
}`)
	organizationsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "which fields, if any, to expand (externalDataSources) Valid values: externalDataSources")
	utils.AddFlag(getCmd.Flags(), "bool", "includeTrustors", "", "(true or false) whether or not to include trustor information embedded in the externalOrganization")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/organizations/{externalOrganizationId}", utils.FormatPermissions([]string{ "relate:externalOrganization:view", "externalContacts:externalOrganization:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExternalOrganization&quot;
  }
}`)
	organizationsCmd.AddCommand(getCmd)
	
	utils.AddFlag(scanCmd.Flags(), "int", "limit", "", "The number of organizations per page; must be between 10 and 200, default is 100)")
	utils.AddFlag(scanCmd.Flags(), "string", "cursor", "", "Indicates where to resume query results (not required for first page), each page returns a new cursor with a 24h TTL")
	scanCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", scanCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/scan/organizations", utils.FormatPermissions([]string{ "externalContacts:externalOrganization:view",  })))
	utils.AddFileFlagIfUpsert(scanCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(scanCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CursorOrganizationListing&quot;
  }
}`)
	organizationsCmd.AddCommand(scanCmd)
	
	utils.AddFlag(searchCmd.Flags(), "int", "pageSize", "20", "Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(searchCmd.Flags(), "int", "pageNumber", "1", "Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be &lt;= 1,000)")
	utils.AddFlag(searchCmd.Flags(), "string", "q", "", "Search query")
	utils.AddFlag(searchCmd.Flags(), "[]string", "trustorId", "", "Search for external organizations by trustorIds (limit 25). If supplied, the `q` parameters is ignored. Items are returned in the order requested")
	utils.AddFlag(searchCmd.Flags(), "string", "sortOrder", "", "Sort order")
	utils.AddFlag(searchCmd.Flags(), "[]string", "expand", "", "which fields, if any, to expand Valid values: externalDataSources")
	utils.AddFlag(searchCmd.Flags(), "bool", "includeTrustors", "", "(true or false) whether or not to include trustor information embedded in the externalOrganization")
	searchCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", searchCmd.UsageTemplate(), "GET", "/api/v2/externalcontacts/organizations", utils.FormatPermissions([]string{ "relate:externalOrganization:view", "externalContacts:externalOrganization:view",  })))
	utils.AddFileFlagIfUpsert(searchCmd.Flags(), "GET")
	utils.AddPaginateFlagsIfListingResponse(searchCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExternalOrganizationListing&quot;
  }
}`)
	organizationsCmd.AddCommand(searchCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/externalcontacts/organizations/{externalOrganizationId}", utils.FormatPermissions([]string{ "relate:externalOrganization:edit", "externalContacts:externalOrganization:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExternalOrganization&quot;
  }
}`)
	organizationsCmd.AddCommand(updateCmd)
	
	return organizationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an external organization",
	Long:  `Create an external organization`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations"

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
	Use:   "delete [externalOrganizationId]",
	Short: "Delete an external organization",
	Long:  `Delete an external organization`,
	Args:  utils.DetermineArgs([]string{ "externalOrganizationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations/{externalOrganizationId}"
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
	Use:   "get [externalOrganizationId]",
	Short: "Fetch an external organization",
	Long:  `Fetch an external organization`,
	Args:  utils.DetermineArgs([]string{ "externalOrganizationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations/{externalOrganizationId}"
		externalOrganizationId, args := args[0], args[1:]
		path = strings.Replace(path, "{externalOrganizationId}", fmt.Sprintf("%v", externalOrganizationId), -1)

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		includeTrustors := utils.GetFlag(cmd.Flags(), "bool", "includeTrustors")
		if includeTrustors != "" {
			queryParams["includeTrustors"] = includeTrustors
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
	Short: "Scan for external organizations using paging",
	Long:  `Scan for external organizations using paging`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/scan/organizations"

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
	Short: "Search for external organizations",
	Long:  `Search for external organizations`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations"

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
		trustorId := utils.GetFlag(cmd.Flags(), "[]string", "trustorId")
		if trustorId != "" {
			queryParams["trustorId"] = trustorId
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		includeTrustors := utils.GetFlag(cmd.Flags(), "bool", "includeTrustors")
		if includeTrustors != "" {
			queryParams["includeTrustors"] = includeTrustors
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
	Use:   "update [externalOrganizationId]",
	Short: "Update an external organization",
	Long:  `Update an external organization`,
	Args:  utils.DetermineArgs([]string{ "externalOrganizationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/externalcontacts/organizations/{externalOrganizationId}"
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
