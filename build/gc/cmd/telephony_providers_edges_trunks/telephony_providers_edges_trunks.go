package telephony_providers_edges_trunks

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_trunks", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/trunks", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/trunks", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/{edgeId}/trunks", )
	telephony_providers_edges_trunksCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_trunks"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_trunksCmd)
}

func Cmdtelephony_providers_edges_trunks() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/trunks/{trunkId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/trunks/{trunkId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Trunk&quot;
  }
}`)
	telephony_providers_edges_trunksCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "Value by which to sort")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ASC", "Sort order")
	utils.AddFlag(listCmd.Flags(), "string", "edgeId", "", "Filter by Edge Ids")
	utils.AddFlag(listCmd.Flags(), "string", "trunkBaseId", "", "Filter by Trunk Base Ids")
	utils.AddFlag(listCmd.Flags(), "string", "trunkType", "", "Filter by a Trunk type Valid values: EXTERNAL, PHONE, EDGE")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/trunks", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/trunks")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	telephony_providers_edges_trunksCmd.AddCommand(listCmd)
	
	utils.AddFlag(listedgetrunksCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listedgetrunksCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listedgetrunksCmd.Flags(), "string", "sortBy", "name", "Value by which to sort")
	utils.AddFlag(listedgetrunksCmd.Flags(), "string", "sortOrder", "ASC", "Sort order")
	utils.AddFlag(listedgetrunksCmd.Flags(), "string", "trunkBaseId", "", "Filter by Trunk Base Ids")
	utils.AddFlag(listedgetrunksCmd.Flags(), "string", "trunkType", "", "Filter by a Trunk type Valid values: EXTERNAL, PHONE, EDGE")
	listedgetrunksCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listedgetrunksCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/{edgeId}/trunks", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/{edgeId}/trunks")))
	utils.AddFileFlagIfUpsert(listedgetrunksCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listedgetrunksCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	telephony_providers_edges_trunksCmd.AddCommand(listedgetrunksCmd)
	
	return telephony_providers_edges_trunksCmd
}

var getCmd = &cobra.Command{
	Use:   "get [trunkId]",
	Short: "Get a Trunk by ID",
	Long:  "Get a Trunk by ID",
	Args:  utils.DetermineArgs([]string{ "trunkId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/trunks/{trunkId}"
		trunkId, args := args[0], args[1:]
		path = strings.Replace(path, "{trunkId}", fmt.Sprintf("%v", trunkId), -1)

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
	Short: "Get the list of available trunks.",
	Long:  "Get the list of available trunks.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/trunks"

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		edgeId := utils.GetFlag(cmd.Flags(), "string", "edgeId")
		if edgeId != "" {
			queryParams["edgeId"] = edgeId
		}
		trunkBaseId := utils.GetFlag(cmd.Flags(), "string", "trunkBaseId")
		if trunkBaseId != "" {
			queryParams["trunkBaseId"] = trunkBaseId
		}
		trunkType := utils.GetFlag(cmd.Flags(), "string", "trunkType")
		if trunkType != "" {
			queryParams["trunkType"] = trunkType
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
var listedgetrunksCmd = &cobra.Command{
	Use:   "listedgetrunks [edgeId]",
	Short: "Get the list of available trunks for the given Edge.",
	Long:  "Get the list of available trunks for the given Edge.",
	Args:  utils.DetermineArgs([]string{ "edgeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/{edgeId}/trunks"
		edgeId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		trunkBaseId := utils.GetFlag(cmd.Flags(), "string", "trunkBaseId")
		if trunkBaseId != "" {
			queryParams["trunkBaseId"] = trunkBaseId
		}
		trunkType := utils.GetFlag(cmd.Flags(), "string", "trunkType")
		if trunkType != "" {
			queryParams["trunkType"] = trunkType
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
