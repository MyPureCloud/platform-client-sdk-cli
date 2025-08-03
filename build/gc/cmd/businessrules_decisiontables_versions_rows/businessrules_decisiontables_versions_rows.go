package businessrules_decisiontables_versions_rows

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
	Description = utils.FormatUsageDescription("businessrules_decisiontables_versions_rows", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows", "SWAGGER_OVERRIDE_/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows", )
	businessrules_decisiontables_versions_rowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("businessrules_decisiontables_versions_rows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(businessrules_decisiontables_versions_rowsCmd)
}

func Cmdbusinessrules_decisiontables_versions_rows() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows", utils.FormatPermissions([]string{ "businessrules:decisionTableRow:add", "routing:queue:view",  }), utils.GenerateDevCentreLink("POST", "Business Rules", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Create decision table row request",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateDecisionTableRowRequest"
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
        "$ref" : "#/components/schemas/DecisionTableRow"
      }
    }
  }
}`)
	businessrules_decisiontables_versions_rowsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}", utils.FormatPermissions([]string{ "businessrules:decisionTableRow:delete", "routing:queue:view",  }), utils.GenerateDevCentreLink("DELETE", "Business Rules", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Deleted successfully",
  "content" : { }
}`)
	businessrules_decisiontables_versions_rowsCmd.AddCommand(deleteCmd)

	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}", utils.FormatPermissions([]string{ "businessrules:decisionTableRow:view",  }), utils.GenerateDevCentreLink("GET", "Business Rules", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DecisionTableRow"
      }
    }
  }
}`)
	businessrules_decisiontables_versions_rowsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "pageNumber", "", "Page number of the entities to return. Defaults to 1.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "", "Number of entities to return. Maximum of 100. Defaults to 25.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows", utils.FormatPermissions([]string{ "businessrules:decisionTableRow:view",  }), utils.GenerateDevCentreLink("GET", "Business Rules", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows")))
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
	businessrules_decisiontables_versions_rowsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}", utils.FormatPermissions([]string{ "businessrules:decisionTableRow:edit", "routing:queue:view",  }), utils.GenerateDevCentreLink("PUT", "Business Rules", "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "Full update decision table row request",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PutDecisionTableRowRequest"
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
        "$ref" : "#/components/schemas/DecisionTableRow"
      }
    }
  }
}`)
	businessrules_decisiontables_versions_rowsCmd.AddCommand(updateCmd)
	return businessrules_decisiontables_versions_rowsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [tableId] [tableVersion]",
	Short: "Create a decision table row",
	Long:  "Create a decision table row",
	Args:  utils.DetermineArgs([]string{ "tableId", "tableVersion", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createdecisiontablerowrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows"
		tableId, args := args[0], args[1:]
		path = strings.Replace(path, "{tableId}", fmt.Sprintf("%v", tableId), -1)
		tableVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{tableVersion}", fmt.Sprintf("%v", tableVersion), -1)

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
	Use:   "delete [tableId] [tableVersion] [rowId]",
	Short: "Delete a decision table row",
	Long:  "Delete a decision table row",
	Args:  utils.DetermineArgs([]string{ "tableId", "tableVersion", "rowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}"
		tableId, args := args[0], args[1:]
		path = strings.Replace(path, "{tableId}", fmt.Sprintf("%v", tableId), -1)
		tableVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{tableVersion}", fmt.Sprintf("%v", tableVersion), -1)
		rowId, args := args[0], args[1:]
		path = strings.Replace(path, "{rowId}", fmt.Sprintf("%v", rowId), -1)

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
	Use:   "get [tableId] [tableVersion] [rowId]",
	Short: "Get a decision table row",
	Long:  "Get a decision table row",
	Args:  utils.DetermineArgs([]string{ "tableId", "tableVersion", "rowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}"
		tableId, args := args[0], args[1:]
		path = strings.Replace(path, "{tableId}", fmt.Sprintf("%v", tableId), -1)
		tableVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{tableVersion}", fmt.Sprintf("%v", tableVersion), -1)
		rowId, args := args[0], args[1:]
		path = strings.Replace(path, "{rowId}", fmt.Sprintf("%v", rowId), -1)

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
	Use:   "list [tableId] [tableVersion]",
	Short: "Get a list of decision table rows.",
	Long:  "Get a list of decision table rows.",
	Args:  utils.DetermineArgs([]string{ "tableId", "tableVersion", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows"
		tableId, args := args[0], args[1:]
		path = strings.Replace(path, "{tableId}", fmt.Sprintf("%v", tableId), -1)
		tableVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{tableVersion}", fmt.Sprintf("%v", tableVersion), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "string", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
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
	Use:   "update [tableId] [tableVersion] [rowId]",
	Short: "Full update a decision table row",
	Long:  "Full update a decision table row",
	Args:  utils.DetermineArgs([]string{ "tableId", "tableVersion", "rowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Putdecisiontablerowrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/businessrules/decisiontables/{tableId}/versions/{tableVersion}/rows/{rowId}"
		tableId, args := args[0], args[1:]
		path = strings.Replace(path, "{tableId}", fmt.Sprintf("%v", tableId), -1)
		tableVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{tableVersion}", fmt.Sprintf("%v", tableVersion), -1)
		rowId, args := args[0], args[1:]
		path = strings.Replace(path, "{rowId}", fmt.Sprintf("%v", rowId), -1)

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
