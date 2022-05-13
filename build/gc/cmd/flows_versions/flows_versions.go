package flows_versions

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
	Description = utils.FormatUsageDescription("flows_versions", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions", "SWAGGER_OVERRIDE_/api/v2/flows/{flowId}/versions", )
	flows_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_versionsCmd)
}

func Cmdflows_versions() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/{flowId}/versions", utils.FormatPermissions([]string{ "architect:flow:edit",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/{flowId}/versions")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "object",
        "additionalProperties" : {
          "type" : "object",
          "properties" : { }
        }
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
        "$ref" : "#/components/schemas/FlowVersion"
      }
    }
  }
}`)
	flows_versionsCmd.AddCommand(createCmd)

	utils.AddFlag(getCmd.Flags(), "string", "deleted", "", "Deleted flows")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}/versions/{versionId}", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/{flowId}/versions/{versionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/FlowVersion"
      }
    }
  }
}`)
	flows_versionsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "bool", "deleted", "", "Include Deleted flows")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/flows/{flowId}/versions", utils.FormatPermissions([]string{ "architect:flow:view",  }), utils.GenerateDevCentreLink("GET", "Architect", "/api/v2/flows/{flowId}/versions")))
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
	flows_versionsCmd.AddCommand(listCmd)
	return flows_versionsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [flowId]",
	Short: "Create flow version",
	Long:  "Create flow version",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Interface{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/versions"
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
var getCmd = &cobra.Command{
	Use:   "get [flowId] [versionId]",
	Short: "Get flow version",
	Long:  "Get flow version",
	Args:  utils.DetermineArgs([]string{ "flowId", "versionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/versions/{versionId}"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)
		versionId, args := args[0], args[1:]
		path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)

		deleted := utils.GetFlag(cmd.Flags(), "string", "deleted")
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
	Use:   "list [flowId]",
	Short: "Get flow version list",
	Long:  "Get flow version list",
	Args:  utils.DetermineArgs([]string{ "flowId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/{flowId}/versions"
		flowId, args := args[0], args[1:]
		path = strings.Replace(path, "{flowId}", fmt.Sprintf("%v", flowId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
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
