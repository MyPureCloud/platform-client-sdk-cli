package contentmanagement_workspaces_tagvalues_query

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
	Description = utils.FormatUsageDescription("contentmanagement_workspaces_tagvalues_query", "SWAGGER_OVERRIDE_/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query", )
	contentmanagement_workspaces_tagvalues_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("contentmanagement_workspaces_tagvalues_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(contentmanagement_workspaces_tagvalues_queryCmd)
}

func Cmdcontentmanagement_workspaces_tagvalues_query() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: acl")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Content Management", "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "query",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/TagQueryRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TagValueEntityListing"
  }
}`)
	contentmanagement_workspaces_tagvalues_queryCmd.AddCommand(createCmd)
	
	return contentmanagement_workspaces_tagvalues_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create [workspaceId]",
	Short: "Perform a prefix query on tags in the workspace",
	Long:  "Perform a prefix query on tags in the workspace",
	Args:  utils.DetermineArgs([]string{ "workspaceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Tagqueryrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query"
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
