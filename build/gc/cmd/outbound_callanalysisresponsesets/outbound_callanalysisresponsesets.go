package outbound_callanalysisresponsesets

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
	Description = utils.FormatUsageDescription("outbound_callanalysisresponsesets", "SWAGGER_OVERRIDE_/api/v2/outbound/callanalysisresponsesets", "SWAGGER_OVERRIDE_/api/v2/outbound/callanalysisresponsesets", "SWAGGER_OVERRIDE_/api/v2/outbound/callanalysisresponsesets", "SWAGGER_OVERRIDE_/api/v2/outbound/callanalysisresponsesets", "SWAGGER_OVERRIDE_/api/v2/outbound/callanalysisresponsesets", )
	outbound_callanalysisresponsesetsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_callanalysisresponsesets"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_callanalysisresponsesetsCmd)
}

func Cmdoutbound_callanalysisresponsesets() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/callanalysisresponsesets", utils.FormatPermissions([]string{ "outbound:responseSet:add",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/callanalysisresponsesets")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "ResponseSet",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ResponseSet"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ResponseSet"
      }
    }
  }
}`)
	outbound_callanalysisresponsesetsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}", utils.FormatPermissions([]string{ "outbound:responseSet:delete",  }), utils.GenerateDevCentreLink("DELETE", "Outbound", "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Operation was successful.",
  "content" : { }
}`)
	outbound_callanalysisresponsesetsCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}", utils.FormatPermissions([]string{ "outbound:responseSet:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ResponseSet"
      }
    }
  }
}`)
	outbound_callanalysisresponsesetsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size. The max that will be returned is 100.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "bool", "allowEmptyResult", "false", "Whether to return an empty page when there are no results for that page")
	utils.AddFlag(listCmd.Flags(), "string", "filterType", "Prefix", "Filter type Valid values: Equals, RegEx, Contains, Prefix, LessThan, LessThanEqualTo, GreaterThan, GreaterThanEqualTo, BeginsWith, EndsWith")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "a", "Sort order Valid values: ascending, descending")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/callanalysisresponsesets", utils.FormatPermissions([]string{ "outbound:responseSet:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/callanalysisresponsesets")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	outbound_callanalysisresponsesetsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}", utils.FormatPermissions([]string{ "outbound:responseSet:edit",  }), utils.GenerateDevCentreLink("PUT", "Outbound", "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "ResponseSet",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ResponseSet"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ResponseSet"
      }
    }
  }
}`)
	outbound_callanalysisresponsesetsCmd.AddCommand(updateCmd)
	return outbound_callanalysisresponsesetsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a dialer call analysis response set.",
	Long:  "Create a dialer call analysis response set.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Responseset{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/callanalysisresponsesets"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "create"
		const httpMethod = "POST"
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
var deleteCmd = &cobra.Command{
	Use:   "delete [callAnalysisSetId]",
	Short: "Delete a dialer call analysis response set.",
	Long:  "Delete a dialer call analysis response set.",
	Args:  utils.DetermineArgs([]string{ "callAnalysisSetId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}"
		callAnalysisSetId, args := args[0], args[1:]
		path = strings.Replace(path, "{callAnalysisSetId}", fmt.Sprintf("%v", callAnalysisSetId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "delete"
		const httpMethod = "DELETE"
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
var getCmd = &cobra.Command{
	Use:   "get [callAnalysisSetId]",
	Short: "Get a dialer call analysis response set.",
	Long:  "Get a dialer call analysis response set.",
	Args:  utils.DetermineArgs([]string{ "callAnalysisSetId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}"
		callAnalysisSetId, args := args[0], args[1:]
		path = strings.Replace(path, "{callAnalysisSetId}", fmt.Sprintf("%v", callAnalysisSetId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Query a list of dialer call analysis response sets.",
	Long:  "Query a list of dialer call analysis response sets.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/callanalysisresponsesets"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		allowEmptyResult := utils.GetFlag(cmd.Flags(), "bool", "allowEmptyResult")
		if allowEmptyResult != "" {
			queryParams["allowEmptyResult"] = allowEmptyResult
		}
		filterType := utils.GetFlag(cmd.Flags(), "string", "filterType")
		if filterType != "" {
			queryParams["filterType"] = filterType
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
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

		const opId = "list"
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
var updateCmd = &cobra.Command{
	Use:   "update [callAnalysisSetId]",
	Short: "Update a dialer call analysis response set.",
	Long:  "Update a dialer call analysis response set.",
	Args:  utils.DetermineArgs([]string{ "callAnalysisSetId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Responseset{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/callanalysisresponsesets/{callAnalysisSetId}"
		callAnalysisSetId, args := args[0], args[1:]
		path = strings.Replace(path, "{callAnalysisSetId}", fmt.Sprintf("%v", callAnalysisSetId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PUT"
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
