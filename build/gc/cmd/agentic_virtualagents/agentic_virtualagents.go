package agentic_virtualagents

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
	Description = utils.FormatUsageDescription("agentic_virtualagents", "SWAGGER_OVERRIDE_/api/v2/agentic/virtualagents", "SWAGGER_OVERRIDE_/api/v2/agentic/virtualagents", "SWAGGER_OVERRIDE_/api/v2/agentic/virtualagents", "SWAGGER_OVERRIDE_/api/v2/agentic/virtualagents", )
	agentic_virtualagentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("agentic_virtualagents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(agentic_virtualagentsCmd)
}

func Cmdagentic_virtualagents() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/agentic/virtualagents", utils.FormatPermissions([]string{ "agentic:virtualAgent:add",  }), utils.GenerateDevCentreLink("POST", "AI Studio", "/api/v2/agentic/virtualagents")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateAgenticVirtualAgent"
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
        "$ref" : "#/components/schemas/AgenticVirtualAgent"
      }
    }
  }
}`)
	agentic_virtualagentsCmd.AddCommand(createCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/agentic/virtualagents/{virtualAgentId}", utils.FormatPermissions([]string{ "agentic:virtualAgent:view",  }), utils.GenerateDevCentreLink("GET", "AI Studio", "/api/v2/agentic/virtualagents/{virtualAgentId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AgenticVirtualAgent"
      }
    }
  }
}`)
	agentic_virtualagentsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter by matching name - case insensitive.")
	utils.AddFlag(listCmd.Flags(), "string", "nameContains", "", "Filter by name contains - case insensitive.")
	utils.AddFlag(listCmd.Flags(), "string", "status", "", "Filter by status. Valid values: Draft, Published")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "dateModified", "Sort by. Default value dateModified. Valid values: dateModified, name, status")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "desc", "Sort Order. Default value desc. Valid values: asc, desc")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number.")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size. The maximum page size is 100.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/agentic/virtualagents", utils.FormatPermissions([]string{ "agentic:virtualAgent:view",  }), utils.GenerateDevCentreLink("GET", "AI Studio", "/api/v2/agentic/virtualagents")))
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
	agentic_virtualagentsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/agentic/virtualagents/{virtualAgentId}", utils.FormatPermissions([]string{ "agentic:virtualAgent:edit",  }), utils.GenerateDevCentreLink("PATCH", "AI Studio", "/api/v2/agentic/virtualagents/{virtualAgentId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UpdateAgenticVirtualAgent"
      }
    }
  },
  "required" : true
}`)
	
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AgenticVirtualAgent"
      }
    }
  }
}`)
	agentic_virtualagentsCmd.AddCommand(updateCmd)
	return agentic_virtualagentsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a virtual agent.",
	Long:  "Create a virtual agent.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createagenticvirtualagent{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/agentic/virtualagents"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		headerParams := make(map[string]string)
		// to determine the Content-Type header
		localVarHttpContentTypes := []string{ "application/json",  }
		// set Content-Type header
		localVarHttpContentType := utils.SelectHeaderContentType(localVarHttpContentTypes)
		if localVarHttpContentType != "" {
			headerParams["Content-Type"] = localVarHttpContentType
		}
		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{
			"application/json",
		}
		// set Accept header
		localVarHttpHeaderAccept := utils.SelectHeaderAccept(localVarHttpHeaderAccepts)
		if localVarHttpHeaderAccept != "" {
			headerParams["Accept"] = localVarHttpHeaderAccept
		}

		const opId = "create"
		const httpMethod = "POST"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, headerParams, cmd, opId)
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
	Use:   "get [virtualAgentId]",
	Short: "Get virtual agent.",
	Long:  "Get virtual agent.",
	Args:  utils.DetermineArgs([]string{ "virtualAgentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/agentic/virtualagents/{virtualAgentId}"
		virtualAgentId, args := args[0], args[1:]
		path = strings.Replace(path, "{virtualAgentId}", fmt.Sprintf("%v", virtualAgentId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		headerParams := make(map[string]string)
		// to determine the Content-Type header
		localVarHttpContentTypes := []string{ "application/json",  }
		// set Content-Type header
		localVarHttpContentType := utils.SelectHeaderContentType(localVarHttpContentTypes)
		if localVarHttpContentType != "" {
			headerParams["Content-Type"] = localVarHttpContentType
		}
		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{
			"application/json",
		}
		// set Accept header
		localVarHttpHeaderAccept := utils.SelectHeaderAccept(localVarHttpHeaderAccepts)
		if localVarHttpHeaderAccept != "" {
			headerParams["Accept"] = localVarHttpHeaderAccept
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, headerParams, cmd, opId)
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
	Short: "Get all virtual agents.",
	Long:  "Get all virtual agents.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/agentic/virtualagents"

		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		nameContains := utils.GetFlag(cmd.Flags(), "string", "nameContains")
		if nameContains != "" {
			queryParams["nameContains"] = nameContains
		}
		status := utils.GetFlag(cmd.Flags(), "string", "status")
		if status != "" {
			queryParams["status"] = status
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		headerParams := make(map[string]string)
		// to determine the Content-Type header
		localVarHttpContentTypes := []string{ "application/json",  }
		// set Content-Type header
		localVarHttpContentType := utils.SelectHeaderContentType(localVarHttpContentTypes)
		if localVarHttpContentType != "" {
			headerParams["Content-Type"] = localVarHttpContentType
		}
		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{
			"application/json",
		}
		// set Accept header
		localVarHttpHeaderAccept := utils.SelectHeaderAccept(localVarHttpHeaderAccepts)
		if localVarHttpHeaderAccept != "" {
			headerParams["Accept"] = localVarHttpHeaderAccept
		}

		const opId = "list"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, headerParams, cmd, opId)
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
	Use:   "update [virtualAgentId]",
	Short: "Update a virtual agent.",
	Long:  "Update a virtual agent.",
	Args:  utils.DetermineArgs([]string{ "virtualAgentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updateagenticvirtualagent{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/agentic/virtualagents/{virtualAgentId}"
		virtualAgentId, args := args[0], args[1:]
		path = strings.Replace(path, "{virtualAgentId}", fmt.Sprintf("%v", virtualAgentId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		headerParams := make(map[string]string)
		// to determine the Content-Type header
		localVarHttpContentTypes := []string{ "application/json",  }
		// set Content-Type header
		localVarHttpContentType := utils.SelectHeaderContentType(localVarHttpContentTypes)
		if localVarHttpContentType != "" {
			headerParams["Content-Type"] = localVarHttpContentType
		}
		// to determine the Accept header
		localVarHttpHeaderAccepts := []string{
			"application/json",
		}
		// set Accept header
		localVarHttpHeaderAccept := utils.SelectHeaderAccept(localVarHttpHeaderAccepts)
		if localVarHttpHeaderAccept != "" {
			headerParams["Accept"] = localVarHttpHeaderAccept
		}

		const opId = "update"
		const httpMethod = "PATCH"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, headerParams, cmd, opId)
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
