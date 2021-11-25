package contentmanagement_workspaces_tagvalues

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
	Description = utils.FormatUsageDescription("contentmanagement_workspaces_tagvalues", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues", )
	contentmanagement_workspaces_tagvaluesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement_workspaces_tagvalues"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagement_workspaces_tagvaluesCmd)
}

func Cmdcontentmanagement_workspaces_tagvalues() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "tag",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/TagValue"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TagValue"
  }
}`)
	contentmanagement_workspaces_tagvaluesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}")))
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
	contentmanagement_workspaces_tagvaluesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: acl")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TagValue"
  }
}`)
	contentmanagement_workspaces_tagvaluesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "value", "", "filter the list of tags returned")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: acl")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	contentmanagement_workspaces_tagvaluesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PUT", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "Workspace",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/TagValue"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TagValue"
  }
}`)
	contentmanagement_workspaces_tagvaluesCmd.AddCommand(updateCmd)
	
	return contentmanagement_workspaces_tagvaluesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [workspaceId]",
	Short: "Create a workspace tag",
	Long:  "Create a workspace tag",
	Args:  utils.DetermineArgs([]string{ "workspaceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Tagvalue{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)

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
	Use:   "delete [workspaceId] [tagId]",
	Short: "Delete workspace tag",
	Long:  "Delete workspace tag",
	Args:  utils.DetermineArgs([]string{ "workspaceId", "tagId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)
		tagId, args := args[0], args[1:]
		path = strings.Replace(path, "{tagId}", fmt.Sprintf("%v", tagId), -1)

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
	Use:   "get [workspaceId] [tagId]",
	Short: "Get a workspace tag",
	Long:  "Get a workspace tag",
	Args:  utils.DetermineArgs([]string{ "workspaceId", "tagId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)
		tagId, args := args[0], args[1:]
		path = strings.Replace(path, "{tagId}", fmt.Sprintf("%v", tagId), -1)

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
	Use:   "list [workspaceId]",
	Short: "Get a list of workspace tags",
	Long:  "Get a list of workspace tags",
	Args:  utils.DetermineArgs([]string{ "workspaceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)

		value := utils.GetFlag(cmd.Flags(), "string", "value")
		if value != "" {
			queryParams["value"] = value
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
	Use:   "update [workspaceId] [tagId]",
	Short: "Update a workspace tag. Will update all documents with the new tag value.",
	Long:  "Update a workspace tag. Will update all documents with the new tag value.",
	Args:  utils.DetermineArgs([]string{ "workspaceId", "tagId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Tagvalue{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/{tagId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)
		tagId, args := args[0], args[1:]
		path = strings.Replace(path, "{tagId}", fmt.Sprintf("%v", tagId), -1)

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
