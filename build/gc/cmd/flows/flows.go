package flows

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
	Description = utils.FormatUsageDescription("flows", "SWAGGER_OVERRIDE_/api/v2/flows", "SWAGGER_OVERRIDE_/api/v2/flows", "SWAGGER_OVERRIDE_/api/v2/flows", "SWAGGER_OVERRIDE_/api/v2/flows", "SWAGGER_OVERRIDE_/api/v2/flows", "SWAGGER_OVERRIDE_/api/v2/flows", )
	flowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flowsCmd)
}

func Cmdflows() *cobra.Command { 
	utils.AddFlag(batchdeleteCmd.Flags(), "[]string", "id", "", "List of Flow IDs - REQUIRED")
	batchdeleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", batchdeleteCmd.UsageTemplate(), "DELETE", "/api/v2/flows", utils.FormatPermissions([]string{ "architect:flow:delete",  }), utils.GenerateDevCentreLink("DELETE", "Architect", "/api/v2/flows")))
	utils.AddFileFlagIfUpsert(batchdeleteCmd.Flags(), "DELETE", ``)
	batchdeleteCmd.MarkFlagRequired("id")
	
	utils.AddPaginateFlagsIfListingResponse(batchdeleteCmd.Flags(), "DELETE", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Operation"
      }
    }
  }
}`)
	flowsCmd.AddCommand(batchdeleteCmd)

	utils.AddFlag(createCmd.Flags(), "string", "language", "", "Language")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows", utils.FormatPermissions([]string{ "architect:flow:add",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Flow"
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
        "$ref" : "#/components/schemas/Flow"
      }
    }
  }
}`)
	flowsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/flows/{flowId}", utils.FormatPermissions([]string{ "architect:flow:delete",  }), utils.GenerateDevCentreLink("DELETE", "Architect", "/api/v2/flows/{flowId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ErrorBody"
      }
    }
  },
  "x-inin-error-codes" : {
    "architect.request.header.missing" : "A required request header is missing or empty.",
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable"
  }
}`)
	flowsCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "bool", "deleted", "false", "Deleted flows")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/{flowId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Flow"
      }
    }
  }
}`)
	flowsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "[]string", "varType", "", "Type Valid values: bot, commonmodule, digitalbot, inboundcall, inboundchat, inboundemail, inboundshortmessage, outboundcall, inqueuecall, inqueueemail, inqueueshortmessage, speech, securecall, surveyinvite, voice, voicemail, workflow, workitem")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "id", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "asc", "Sort order")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "ID")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "description", "", "Description")
	utils.AddFlag(listCmd.Flags(), "string", "nameOrDescription", "", "Name or description")
	utils.AddFlag(listCmd.Flags(), "string", "publishVersionId", "", "Publish version ID")
	utils.AddFlag(listCmd.Flags(), "string", "editableBy", "", "Editable by")
	utils.AddFlag(listCmd.Flags(), "string", "lockedBy", "", "Locked by")
	utils.AddFlag(listCmd.Flags(), "string", "lockedByClientId", "", "Locked by client ID")
	utils.AddFlag(listCmd.Flags(), "string", "secure", "", "Secure Valid values: any, checkedin, published")
	utils.AddFlag(listCmd.Flags(), "bool", "deleted", "false", "Include deleted")
	utils.AddFlag(listCmd.Flags(), "bool", "includeSchemas", "false", "Include variable schemas")
	utils.AddFlag(listCmd.Flags(), "string", "publishedAfter", "", "Published after")
	utils.AddFlag(listCmd.Flags(), "string", "publishedBefore", "", "Published before")
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionId", "", "division ID(s)")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows")))
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
	flowsCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/flows/{flowId}", utils.FormatPermissions([]string{ "architect:flow:edit",  }), utils.GenerateDevCentreLink("PUT", "Architect", "/api/v2/flows/{flowId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Flow"
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
        "$ref" : "#/components/schemas/Flow"
      }
    }
  }
}`)
	flowsCmd.AddCommand(updateCmd)
	return flowsCmd
}

var batchdeleteCmd = &cobra.Command{
	Use:   "batchdelete",
	Short: "Batch-delete a list of flows",
	Long:  "Batch-delete a list of flows",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows"

		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "batchdelete"
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create flow",
	Long:  "Create flow",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Flow{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows"

		language := utils.GetFlag(cmd.Flags(), "string", "language")
		if language != "" {
			queryParams["language"] = language
		}
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [flowId]",
	Short: "Delete flow",
	Long:  "Delete flow",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [flowId]",
	Short: "Get flow",
	Long:  "Get flow",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)

		deleted := utils.GetFlag(cmd.Flags(), "bool", "deleted")
		if deleted != "" {
			queryParams["deleted"] = deleted
		}
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a pageable list of flows, filtered by query parameters",
	Long:  "Get a pageable list of flows, filtered by query parameters",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows"

		varType := utils.GetFlag(cmd.Flags(), "[]string", "varType")
		if varType != "" {
			queryParams["varType"] = varType
		}
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
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		description := utils.GetFlag(cmd.Flags(), "string", "description")
		if description != "" {
			queryParams["description"] = description
		}
		nameOrDescription := utils.GetFlag(cmd.Flags(), "string", "nameOrDescription")
		if nameOrDescription != "" {
			queryParams["nameOrDescription"] = nameOrDescription
		}
		publishVersionId := utils.GetFlag(cmd.Flags(), "string", "publishVersionId")
		if publishVersionId != "" {
			queryParams["publishVersionId"] = publishVersionId
		}
		editableBy := utils.GetFlag(cmd.Flags(), "string", "editableBy")
		if editableBy != "" {
			queryParams["editableBy"] = editableBy
		}
		lockedBy := utils.GetFlag(cmd.Flags(), "string", "lockedBy")
		if lockedBy != "" {
			queryParams["lockedBy"] = lockedBy
		}
		lockedByClientId := utils.GetFlag(cmd.Flags(), "string", "lockedByClientId")
		if lockedByClientId != "" {
			queryParams["lockedByClientId"] = lockedByClientId
		}
		secure := utils.GetFlag(cmd.Flags(), "string", "secure")
		if secure != "" {
			queryParams["secure"] = secure
		}
		deleted := utils.GetFlag(cmd.Flags(), "bool", "deleted")
		if deleted != "" {
			queryParams["deleted"] = deleted
		}
		includeSchemas := utils.GetFlag(cmd.Flags(), "bool", "includeSchemas")
		if includeSchemas != "" {
			queryParams["includeSchemas"] = includeSchemas
		}
		publishedAfter := utils.GetFlag(cmd.Flags(), "string", "publishedAfter")
		if publishedAfter != "" {
			queryParams["publishedAfter"] = publishedAfter
		}
		publishedBefore := utils.GetFlag(cmd.Flags(), "string", "publishedBefore")
		if publishedBefore != "" {
			queryParams["publishedBefore"] = publishedBefore
		}
		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [flowId]",
	Short: "Update flow",
	Long:  "Update flow",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Flow{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
