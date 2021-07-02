package telephony_providers_edges_sites_outboundroutes

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_sites_outboundroutes", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes", )
	telephony_providers_edges_sites_outboundroutesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_sites_outboundroutes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_sites_outboundroutesCmd)
}

func Cmdtelephony_providers_edges_sites_outboundroutes() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("POST", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;OutboundRoute&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/OutboundRouteBase&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/OutboundRouteBase&quot;
  }
}`)
	telephony_providers_edges_sites_outboundroutesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("DELETE", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful.&quot;
}`)
	telephony_providers_edges_sites_outboundroutesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/OutboundRouteBase&quot;
  }
}`)
	telephony_providers_edges_sites_outboundroutesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "externalTrunkBasesIds", "", "externalTrunkBases.ids")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "Sort by")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	telephony_providers_edges_sites_outboundroutesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("PUT", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;OutboundRoute&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/OutboundRouteBase&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/OutboundRouteBase&quot;
  }
}`)
	telephony_providers_edges_sites_outboundroutesCmd.AddCommand(updateCmd)
	
	return telephony_providers_edges_sites_outboundroutesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [siteId]",
	Short: "Create outbound route",
	Long:  "Create outbound route",
	Args:  utils.DetermineArgs([]string{ "siteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)

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
	Use:   "delete [siteId] [outboundRouteId]",
	Short: "Delete Outbound Route",
	Long:  "Delete Outbound Route",
	Args:  utils.DetermineArgs([]string{ "siteId", "outboundRouteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
		outboundRouteId, args := args[0], args[1:]
		path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)

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
	Use:   "get [siteId] [outboundRouteId]",
	Short: "Get an outbound route",
	Long:  "Get an outbound route",
	Args:  utils.DetermineArgs([]string{ "siteId", "outboundRouteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
		outboundRouteId, args := args[0], args[1:]
		path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)

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
	Use:   "list [siteId]",
	Short: "Get outbound routes",
	Long:  "Get outbound routes",
	Args:  utils.DetermineArgs([]string{ "siteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		externalTrunkBasesIds := utils.GetFlag(cmd.Flags(), "string", "externalTrunkBasesIds")
		if externalTrunkBasesIds != "" {
			queryParams["externalTrunkBasesIds"] = externalTrunkBasesIds
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
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
	Use:   "update [siteId] [outboundRouteId]",
	Short: "Update outbound route",
	Long:  "Update outbound route",
	Args:  utils.DetermineArgs([]string{ "siteId", "outboundRouteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
		outboundRouteId, args := args[0], args[1:]
		path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)

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
