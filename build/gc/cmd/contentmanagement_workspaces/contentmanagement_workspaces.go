package contentmanagement_workspaces

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
	Description = utils.FormatUsageDescription("contentmanagement_workspaces", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces", )
	contentmanagement_workspacesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement_workspaces"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagement_workspacesCmd)
}

func Cmdcontentmanagement_workspaces() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/contentmanagement/workspaces", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Content Management", "/api/v2/contentmanagement/workspaces")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Workspace&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WorkspaceCreate&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Workspace&quot;
  }
}`)
	contentmanagement_workspacesCmd.AddCommand(createCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "string", "moveChildrenToWorkspaceId", "", "New location for objects in deleted workspace.")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/contentmanagement/workspaces/{workspaceId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Accepted - Processing Delete&quot;
}`)
	contentmanagement_workspacesCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: summary, acl")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/workspaces/{workspaceId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Workspace&quot;
  }
}`)
	contentmanagement_workspacesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "[]string", "access", "", "Requested access level. Valid values: content, admin, document:create, document:viewContent, document:viewMetadata, document:download, document:delete, document:update, document:share, document:shareView, document:email, document:print, document:auditView, document:replace, document:tag, tag:create, tag:view, tag:update, tag:apply, tag:remove, tag:delete")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: summary, acl")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/workspaces", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/workspaces")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	contentmanagement_workspacesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/contentmanagement/workspaces/{workspaceId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PUT", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Workspace&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Workspace&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Workspace&quot;
  }
}`)
	contentmanagement_workspacesCmd.AddCommand(updateCmd)
	
	return contentmanagement_workspacesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a group workspace",
	Long:  "Create a group workspace",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
	Use:   "delete [workspaceId]",
	Short: "Delete a workspace",
	Long:  "Delete a workspace",
	Args:  utils.DetermineArgs([]string{ "workspaceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)

		moveChildrenToWorkspaceId := utils.GetFlag(cmd.Flags(), "string", "moveChildrenToWorkspaceId")
		if moveChildrenToWorkspaceId != "" {
			queryParams["moveChildrenToWorkspaceId"] = moveChildrenToWorkspaceId
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
	Use:   "get [workspaceId]",
	Short: "Get a workspace.",
	Long:  "Get a workspace.",
	Args:  utils.DetermineArgs([]string{ "workspaceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of workspaces.",
	Long:  "Get a list of workspaces.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		access := utils.GetFlag(cmd.Flags(), "[]string", "access")
		if access != "" {
			queryParams["access"] = access
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
var updateCmd = &cobra.Command{
	Use:   "update [workspaceId]",
	Short: "Update a workspace",
	Long:  "Update a workspace",
	Args:  utils.DetermineArgs([]string{ "workspaceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}"
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
