package scripts_pages

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
	Description = utils.FormatUsageDescription("scripts_pages", "SWAGGER_OVERRIDE_/api/v2/scripts/{scriptId}/pages", "SWAGGER_OVERRIDE_/api/v2/scripts/{scriptId}/pages", )
	scripts_pagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scripts_pages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scripts_pagesCmd)
}

func Cmdscripts_pages() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "scriptDataVersion", "", "Advanced usage - controls the data version of the script")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/scripts/{scriptId}/pages/{pageId}", utils.FormatPermissions([]string{ "scripter:script:view",  }), utils.GenerateDevCentreLink("GET", "Scripts", "/api/v2/scripts/{scriptId}/pages/{pageId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Page"
  }
}`)
	scripts_pagesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "scriptDataVersion", "", "Advanced usage - controls the data version of the script")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/scripts/{scriptId}/pages", utils.FormatPermissions([]string{ "scripter:script:view",  }), utils.GenerateDevCentreLink("GET", "Scripts", "/api/v2/scripts/{scriptId}/pages")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/Page"
    }
  }
}`)
	scripts_pagesCmd.AddCommand(listCmd)
	
	return scripts_pagesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [scriptId] [pageId]",
	Short: "Get a page",
	Long:  "Get a page",
	Args:  utils.DetermineArgs([]string{ "scriptId", "pageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scripts/{scriptId}/pages/{pageId}"
		scriptId, args := args[0], args[1:]
		path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
		pageId, args := args[0], args[1:]
		path = strings.Replace(path, "{pageId}", fmt.Sprintf("%v", pageId), -1)

		scriptDataVersion := utils.GetFlag(cmd.Flags(), "string", "scriptDataVersion")
		if scriptDataVersion != "" {
			queryParams["scriptDataVersion"] = scriptDataVersion
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
	Use:   "list [scriptId]",
	Short: "Get the list of pages",
	Long:  "Get the list of pages",
	Args:  utils.DetermineArgs([]string{ "scriptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scripts/{scriptId}/pages"
		scriptId, args := args[0], args[1:]
		path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)

		scriptDataVersion := utils.GetFlag(cmd.Flags(), "string", "scriptDataVersion")
		if scriptDataVersion != "" {
			queryParams["scriptDataVersion"] = scriptDataVersion
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
