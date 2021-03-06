package routing_email_domains_routes

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
	Description = utils.FormatUsageDescription("routing_email_domains_routes", "SWAGGER_OVERRIDE_/api/v2/routing/email/domains/{domainName}/routes", "SWAGGER_OVERRIDE_/api/v2/routing/email/domains/{domainName}/routes", "SWAGGER_OVERRIDE_/api/v2/routing/email/domains/{domainName}/routes", "SWAGGER_OVERRIDE_/api/v2/routing/email/domains/{domainName}/routes", "SWAGGER_OVERRIDE_/api/v2/routing/email/domains/{domainName}/routes", )
	routing_email_domains_routesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_email_domains_routes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_email_domains_routesCmd)
}

func Cmdrouting_email_domains_routes() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/routing/email/domains/{domainName}/routes", utils.FormatPermissions([]string{ "routing:email:manage",  }), utils.GenerateDevCentreLink("POST", "Routing", "/api/v2/routing/email/domains/{domainName}/routes")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Route&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/InboundRoute&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/InboundRoute&quot;
  }
}`)
	routing_email_domains_routesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/routing/email/domains/{domainName}/routes/{routeId}", utils.FormatPermissions([]string{ "routing:email:manage",  }), utils.GenerateDevCentreLink("DELETE", "Routing", "/api/v2/routing/email/domains/{domainName}/routes/{routeId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful.&quot;
}`)
	routing_email_domains_routesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/routing/email/domains/{domainName}/routes/{routeId}", utils.FormatPermissions([]string{ "routing:email:manage",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/email/domains/{domainName}/routes/{routeId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/InboundRoute&quot;
  }
}`)
	routing_email_domains_routesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "pattern", "", "Filter routes by the route`s pattern property")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/email/domains/{domainName}/routes", utils.FormatPermissions([]string{ "routing:email:manage",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/email/domains/{domainName}/routes")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	routing_email_domains_routesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/routing/email/domains/{domainName}/routes/{routeId}", utils.FormatPermissions([]string{ "routing:email:manage",  }), utils.GenerateDevCentreLink("PUT", "Routing", "/api/v2/routing/email/domains/{domainName}/routes/{routeId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Route&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/InboundRoute&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/InboundRoute&quot;
  }
}`)
	routing_email_domains_routesCmd.AddCommand(updateCmd)
	
	return routing_email_domains_routesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [domainName]",
	Short: "Create a route",
	Long:  "Create a route",
	Args:  utils.DetermineArgs([]string{ "domainName", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/email/domains/{domainName}/routes"
		domainName, args := args[0], args[1:]
		path = strings.Replace(path, "{domainName}", fmt.Sprintf("%v", domainName), -1)

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
	Use:   "delete [domainName] [routeId]",
	Short: "Delete a route",
	Long:  "Delete a route",
	Args:  utils.DetermineArgs([]string{ "domainName", "routeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/email/domains/{domainName}/routes/{routeId}"
		domainName, args := args[0], args[1:]
		path = strings.Replace(path, "{domainName}", fmt.Sprintf("%v", domainName), -1)
		routeId, args := args[0], args[1:]
		path = strings.Replace(path, "{routeId}", fmt.Sprintf("%v", routeId), -1)

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
	Use:   "get [domainName] [routeId]",
	Short: "Get a route",
	Long:  "Get a route",
	Args:  utils.DetermineArgs([]string{ "domainName", "routeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/email/domains/{domainName}/routes/{routeId}"
		domainName, args := args[0], args[1:]
		path = strings.Replace(path, "{domainName}", fmt.Sprintf("%v", domainName), -1)
		routeId, args := args[0], args[1:]
		path = strings.Replace(path, "{routeId}", fmt.Sprintf("%v", routeId), -1)

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
	Use:   "list [domainName]",
	Short: "Get routes",
	Long:  "Get routes",
	Args:  utils.DetermineArgs([]string{ "domainName", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/email/domains/{domainName}/routes"
		domainName, args := args[0], args[1:]
		path = strings.Replace(path, "{domainName}", fmt.Sprintf("%v", domainName), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pattern := utils.GetFlag(cmd.Flags(), "string", "pattern")
		if pattern != "" {
			queryParams["pattern"] = pattern
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
	Use:   "update [domainName] [routeId]",
	Short: "Update a route",
	Long:  "Update a route",
	Args:  utils.DetermineArgs([]string{ "domainName", "routeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/email/domains/{domainName}/routes/{routeId}"
		domainName, args := args[0], args[1:]
		path = strings.Replace(path, "{domainName}", fmt.Sprintf("%v", domainName), -1)
		routeId, args := args[0], args[1:]
		path = strings.Replace(path, "{routeId}", fmt.Sprintf("%v", routeId), -1)

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
