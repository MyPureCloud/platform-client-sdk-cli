package architect_prompts

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
	Description = utils.FormatUsageDescription("architect_prompts", "SWAGGER_OVERRIDE_/api/v2/architect/prompts", "SWAGGER_OVERRIDE_/api/v2/architect/prompts", "SWAGGER_OVERRIDE_/api/v2/architect/prompts", "SWAGGER_OVERRIDE_/api/v2/architect/prompts", "SWAGGER_OVERRIDE_/api/v2/architect/prompts", "SWAGGER_OVERRIDE_/api/v2/architect/prompts", )
	architect_promptsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_prompts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_promptsCmd)
}

func Cmdarchitect_prompts() *cobra.Command { 
	utils.AddFlag(batchdeleteCmd.Flags(), "[]string", "id", "", "List of Prompt IDs - REQUIRED")
	batchdeleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", batchdeleteCmd.UsageTemplate(), "DELETE", "/api/v2/architect/prompts", utils.FormatPermissions([]string{ "architect:userPrompt:delete",  }), utils.GenerateDevCentreLink("DELETE", "Architect", "/api/v2/architect/prompts")))
	utils.AddFileFlagIfUpsert(batchdeleteCmd.Flags(), "DELETE", ``)
	batchdeleteCmd.MarkFlagRequired("id")
	
	utils.AddPaginateFlagsIfListingResponse(batchdeleteCmd.Flags(), "DELETE", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Operation"
  }
}`)
	architect_promptsCmd.AddCommand(batchdeleteCmd)
	
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/architect/prompts", utils.FormatPermissions([]string{ "architect:userPrompt:add",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/architect/prompts")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/Prompt"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Prompt"
  }
}`)
	architect_promptsCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "bool", "allResources", "", "Whether or not to delete all the prompt resources")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/architect/prompts/{promptId}", utils.FormatPermissions([]string{ "architect:userPrompt:delete",  }), utils.GenerateDevCentreLink("DELETE", "Architect", "/api/v2/architect/prompts/{promptId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The request could not be understood by the server due to malformed syntax.",
  "schema" : {
    "$ref" : "#/definitions/ErrorBody"
  },
  "x-inin-error-codes" : {
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable"
  }
}`)
	architect_promptsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/architect/prompts/{promptId}", utils.FormatPermissions([]string{ "architect:userPrompt:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/architect/prompts/{promptId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Prompt"
  }
}`)
	architect_promptsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "[]string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "string", "description", "", "Description")
	utils.AddFlag(listCmd.Flags(), "string", "nameOrDescription", "", "Name or description")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "id", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "asc", "Sort order")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/architect/prompts", utils.FormatPermissions([]string{ "architect:userPrompt:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/architect/prompts")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	architect_promptsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/architect/prompts/{promptId}", utils.FormatPermissions([]string{ "architect:userPrompt:edit",  }), utils.GenerateDevCentreLink("PUT", "Architect", "/api/v2/architect/prompts/{promptId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/Prompt"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Prompt"
  }
}`)
	architect_promptsCmd.AddCommand(updateCmd)
	
	return architect_promptsCmd
}

var batchdeleteCmd = &cobra.Command{
	Use:   "batchdelete",
	Short: "Batch-delete a list of prompts",
	Long:  "Batch-delete a list of prompts",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/prompts"

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
	Short: "Create a new user prompt",
	Long:  "Create a new user prompt",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Prompt{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/prompts"

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
	Use:   "delete [promptId]",
	Short: "Delete specified user prompt",
	Long:  "Delete specified user prompt",
	Args:  utils.DetermineArgs([]string{ "promptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/prompts/{promptId}"
		promptId, args := args[0], args[1:]
		path = strings.Replace(path, "{promptId}", fmt.Sprintf("%v", promptId), -1)

		allResources := utils.GetFlag(cmd.Flags(), "bool", "allResources")
		if allResources != "" {
			queryParams["allResources"] = allResources
		}
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
	Use:   "get [promptId]",
	Short: "Get specified user prompt",
	Long:  "Get specified user prompt",
	Args:  utils.DetermineArgs([]string{ "promptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/prompts/{promptId}"
		promptId, args := args[0], args[1:]
		path = strings.Replace(path, "{promptId}", fmt.Sprintf("%v", promptId), -1)

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
	Short: "Get a pageable list of user prompts",
	Long:  "Get a pageable list of user prompts",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/prompts"

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		name := utils.GetFlag(cmd.Flags(), "[]string", "name")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [promptId]",
	Short: "Update specified user prompt",
	Long:  "Update specified user prompt",
	Args:  utils.DetermineArgs([]string{ "promptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Prompt{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/architect/prompts/{promptId}"
		promptId, args := args[0], args[1:]
		path = strings.Replace(path, "{promptId}", fmt.Sprintf("%v", promptId), -1)

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
