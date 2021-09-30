package telephony_providers_edges_lines

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_lines", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/lines", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/{edgeId}/lines", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/lines", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/{edgeId}/lines", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/{edgeId}/lines", )
	telephony_providers_edges_linesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_lines"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_linesCmd)
}

func Cmdtelephony_providers_edges_lines() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/lines/{lineId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/lines/{lineId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Line"
  }
}`)
	telephony_providers_edges_linesCmd.AddCommand(getCmd)
	
	getedgelineCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getedgelineCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}")))
	utils.AddFileFlagIfUpsert(getedgelineCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getedgelineCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/EdgeLine"
  }
}`)
	telephony_providers_edges_linesCmd.AddCommand(getedgelineCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "Value by which to sort")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Fields to expand in the response, comma-separated Valid values: properties, site, edgeGroup, primaryEdge, secondaryEdge, edges, assignedUser")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/lines", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/lines")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	telephony_providers_edges_linesCmd.AddCommand(listCmd)
	
	utils.AddFlag(listedgelinesCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listedgelinesCmd.Flags(), "int", "pageNumber", "1", "Page number")
	listedgelinesCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listedgelinesCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/{edgeId}/lines", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/{edgeId}/lines")))
	utils.AddFileFlagIfUpsert(listedgelinesCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listedgelinesCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	telephony_providers_edges_linesCmd.AddCommand(listedgelinesCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("PUT", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "Line",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/EdgeLine"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/EdgeLine"
  }
}`)
	telephony_providers_edges_linesCmd.AddCommand(updateCmd)
	
	return telephony_providers_edges_linesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [lineId]",
	Short: "Get a Line by ID",
	Long:  "Get a Line by ID",
	Args:  utils.DetermineArgs([]string{ "lineId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/lines/{lineId}"
		lineId, args := args[0], args[1:]
		path = strings.Replace(path, "{lineId}", fmt.Sprintf("%v", lineId), -1)


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
var getedgelineCmd = &cobra.Command{
	Use:   "getedgeline [edgeId] [lineId]",
	Short: "Get line",
	Long:  "Get line",
	Args:  utils.DetermineArgs([]string{ "edgeId", "lineId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}"
		edgeId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
		lineId, args := args[0], args[1:]
		path = strings.Replace(path, "{lineId}", fmt.Sprintf("%v", lineId), -1)


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
	Short: "Get a list of Lines",
	Long:  "Get a list of Lines",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/lines"


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
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
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
var listedgelinesCmd = &cobra.Command{
	Use:   "listedgelines [edgeId]",
	Short: "Get the list of lines.",
	Long:  "Get the list of lines.",
	Args:  utils.DetermineArgs([]string{ "edgeId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/{edgeId}/lines"
		edgeId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)


		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
	Use:   "update [edgeId] [lineId]",
	Short: "Update a line.",
	Long:  "Update a line.",
	Args:  utils.DetermineArgs([]string{ "edgeId", "lineId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Edgeline{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}"
		edgeId, args := args[0], args[1:]
		path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
		lineId, args := args[0], args[1:]
		path = strings.Replace(path, "{lineId}", fmt.Sprintf("%v", lineId), -1)


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
