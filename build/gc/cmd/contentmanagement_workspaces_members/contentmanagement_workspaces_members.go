package contentmanagement_workspaces_members

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
	Description = utils.FormatUsageDescription("contentmanagement_workspaces_members", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/members", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/members", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/members", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/members", )
	contentmanagement_workspaces_membersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement_workspaces_members"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagement_workspaces_membersCmd)
}

func Cmdcontentmanagement_workspaces_members() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}")))
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
    "bad.request" : "The request could not be understood by the server due to malformed syntax.",
    "response.entity.too.large" : "The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable"
  }
}`)
	contentmanagement_workspaces_membersCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: member")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkspaceMember"
      }
    }
  }
}`)
	contentmanagement_workspaces_membersCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: member")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/contentmanagement/workspaces/{workspaceId}/members", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/members")))
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
	contentmanagement_workspaces_membersCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PUT", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "Workspace Member",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkspaceMember"
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
        "$ref" : "#/components/schemas/WorkspaceMember"
      }
    }
  }
}`)
	contentmanagement_workspaces_membersCmd.AddCommand(updateCmd)
	return contentmanagement_workspaces_membersCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [workspaceId] [memberId]",
	Short: "Delete a member from a workspace",
	Long:  "Delete a member from a workspace",
	Args:  utils.DetermineArgs([]string{ "workspaceId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

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
	Use:   "get [workspaceId] [memberId]",
	Short: "Get a workspace member",
	Long:  "Get a workspace member",
	Args:  utils.DetermineArgs([]string{ "workspaceId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

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
	Short: "Get a list workspace members",
	Long:  "Get a list workspace members",
	Args:  utils.DetermineArgs([]string{ "workspaceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/members"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)

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
	Use:   "update [workspaceId] [memberId]",
	Short: "Add a member to a workspace",
	Long:  "Add a member to a workspace",
	Args:  utils.DetermineArgs([]string{ "workspaceId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Workspacemember{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}"
		workspaceId, args := args[0], args[1:]
		path = strings.Replace(path, "{workspaceId}", fmt.Sprintf("%v", workspaceId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

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
