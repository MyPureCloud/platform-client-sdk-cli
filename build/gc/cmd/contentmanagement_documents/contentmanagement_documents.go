package contentmanagement_documents

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
	Description = utils.FormatUsageDescription("contentmanagement_documents", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/documents", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/documents", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/documents", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/documents", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/documents", )
	contentmanagement_documentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement_documents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagement_documentsCmd)
}

func Cmdcontentmanagement_documents() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "copySource", "", "Copy a document within a workspace or to a new workspace. Provide a document ID as the copy source.")
	utils.AddFlag(createCmd.Flags(), "string", "moveSource", "", "Move a document to a new workspace. Provide a document ID as the move source.")
	utils.AddFlag(createCmd.Flags(), "bool", "override", "", "Override any lock on the source document")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/contentmanagement/documents", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Content Management", "/api/v2/contentmanagement/documents")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Document",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/DocumentUpload"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Document"
  }
}`)
	contentmanagement_documentsCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "bool", "override", "", "Override any lock on the document")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/contentmanagement/documents/{documentId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Content Management", "/api/v2/contentmanagement/documents/{documentId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Accepted - Processing Delete"
}`)
	contentmanagement_documentsCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: lockInfo, acl, workspace")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/documents/{documentId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/documents/{documentId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Document"
  }
}`)
	contentmanagement_documentsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "workspaceId", "", "Workspace ID - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: acl, workspace")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "name or dateCreated")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ascending", "ascending or descending")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/documents", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/documents")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("workspaceId")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	contentmanagement_documentsCmd.AddCommand(listCmd)
	
	utils.AddFlag(updateCmd.Flags(), "string", "expand", "", "Expand some document fields Valid values: acl")
	utils.AddFlag(updateCmd.Flags(), "bool", "override", "", "Override any lock on the document")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "POST", "/api/v2/contentmanagement/documents/{documentId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Content Management", "/api/v2/contentmanagement/documents/{documentId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Document",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/DocumentUpdate"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Document"
  }
}`)
	contentmanagement_documentsCmd.AddCommand(updateCmd)
	
	return contentmanagement_documentsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a document.",
	Long:  "Add a document.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Documentupload{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/documents"

		copySource := utils.GetFlag(cmd.Flags(), "string", "copySource")
		if copySource != "" {
			queryParams["copySource"] = copySource
		}
		moveSource := utils.GetFlag(cmd.Flags(), "string", "moveSource")
		if moveSource != "" {
			queryParams["moveSource"] = moveSource
		}
		override := utils.GetFlag(cmd.Flags(), "bool", "override")
		if override != "" {
			queryParams["override"] = override
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
	Use:   "delete [documentId]",
	Short: "Delete a document.",
	Long:  "Delete a document.",
	Args:  utils.DetermineArgs([]string{ "documentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/documents/{documentId}"
		documentId, args := args[0], args[1:]
		path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)

		override := utils.GetFlag(cmd.Flags(), "bool", "override")
		if override != "" {
			queryParams["override"] = override
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
	Use:   "get [documentId]",
	Short: "Get a document.",
	Long:  "Get a document.",
	Args:  utils.DetermineArgs([]string{ "documentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/documents/{documentId}"
		documentId, args := args[0], args[1:]
		path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)

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
	Use:   "list",
	Short: "Get a list of documents.",
	Long:  "Get a list of documents.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/documents"

		workspaceId := utils.GetFlag(cmd.Flags(), "string", "workspaceId")
		if workspaceId != "" {
			queryParams["workspaceId"] = workspaceId
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
	Use:   "update [documentId]",
	Short: "Update a document.",
	Long:  "Update a document.",
	Args:  utils.DetermineArgs([]string{ "documentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Documentupdate{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/documents/{documentId}"
		documentId, args := args[0], args[1:]
		path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		override := utils.GetFlag(cmd.Flags(), "bool", "override")
		if override != "" {
			queryParams["override"] = override
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
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
