package agentic_virtualagents_versions

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
	Description = utils.FormatUsageDescription("agentic_virtualagents_versions", "SWAGGER_OVERRIDE_/api/v2/agentic/virtualagents/{virtualAgentId}/versions", "SWAGGER_OVERRIDE_/api/v2/agentic/virtualagents/{virtualAgentId}/versions", "SWAGGER_OVERRIDE_/api/v2/agentic/virtualagents/{virtualAgentId}/versions", )
	agentic_virtualagents_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("agentic_virtualagents_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(agentic_virtualagents_versionsCmd)
}

func Cmdagentic_virtualagents_versions() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "validateOnly", "false", "Validate the request without creating the version.")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/agentic/virtualagents/{virtualAgentId}/versions", utils.FormatPermissions([]string{ "agentic:virtualAgentVersion:add",  }), utils.GenerateDevCentreLink("POST", "AI Studio", "/api/v2/agentic/virtualagents/{virtualAgentId}/versions")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateAgenticVirtualAgentVersion"
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
        "$ref" : "#/components/schemas/AgenticVirtualAgentVersion"
      }
    }
  }
}`)
	agentic_virtualagents_versionsCmd.AddCommand(createCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/agentic/virtualagents/{virtualAgentId}/versions/{versionId}", utils.FormatPermissions([]string{ "agentic:virtualAgentVersion:view",  }), utils.GenerateDevCentreLink("GET", "AI Studio", "/api/v2/agentic/virtualagents/{virtualAgentId}/versions/{versionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AgenticVirtualAgentVersion"
      }
    }
  }
}`)
	agentic_virtualagents_versionsCmd.AddCommand(getCmd)

	utils.AddFlag(updateCmd.Flags(), "bool", "validateOnly", "false", "Validate the update without saving the version.")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/agentic/virtualagents/{virtualAgentId}/versions/{versionId}", utils.FormatPermissions([]string{ "agentic:virtualAgentVersion:edit",  }), utils.GenerateDevCentreLink("PATCH", "AI Studio", "/api/v2/agentic/virtualagents/{virtualAgentId}/versions/{versionId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UpdateAgenticVirtualAgentVersion"
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
        "$ref" : "#/components/schemas/AgenticVirtualAgentVersion"
      }
    }
  }
}`)
	agentic_virtualagents_versionsCmd.AddCommand(updateCmd)
	return agentic_virtualagents_versionsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [virtualAgentId]",
	Short: "Create a virtual agent version.",
	Long:  "Create a virtual agent version.",
	Args:  utils.DetermineArgs([]string{ "virtualAgentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createagenticvirtualagentversion{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/agentic/virtualagents/{virtualAgentId}/versions"
		virtualAgentId, args := args[0], args[1:]
		path = strings.Replace(path, "{virtualAgentId}", fmt.Sprintf("%v", virtualAgentId), -1)

		validateOnly := utils.GetFlag(cmd.Flags(), "bool", "validateOnly")
		if validateOnly != "" {
			queryParams["validateOnly"] = validateOnly
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
	Use:   "get [virtualAgentId] [versionId]",
	Short: "Get a virtual agent version.",
	Long:  "Get a virtual agent version.",
	Args:  utils.DetermineArgs([]string{ "virtualAgentId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/agentic/virtualagents/{virtualAgentId}/versions/{versionId}"
		virtualAgentId, args := args[0], args[1:]
		path = strings.Replace(path, "{virtualAgentId}", fmt.Sprintf("%v", virtualAgentId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [virtualAgentId] [versionId]",
	Short: "Update a virtual agent version.",
	Long:  "Update a virtual agent version.",
	Args:  utils.DetermineArgs([]string{ "virtualAgentId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updateagenticvirtualagentversion{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/agentic/virtualagents/{virtualAgentId}/versions/{versionId}"
		virtualAgentId, args := args[0], args[1:]
		path = strings.Replace(path, "{virtualAgentId}", fmt.Sprintf("%v", virtualAgentId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)

		validateOnly := utils.GetFlag(cmd.Flags(), "bool", "validateOnly")
		if validateOnly != "" {
			queryParams["validateOnly"] = validateOnly
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
