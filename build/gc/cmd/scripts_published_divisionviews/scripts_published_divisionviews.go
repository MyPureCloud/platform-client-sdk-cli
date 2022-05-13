package scripts_published_divisionviews

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
	Description = utils.FormatUsageDescription("scripts_published_divisionviews", "SWAGGER_OVERRIDE_/api/v2/scripts/published/divisionviews", )
	scripts_published_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scripts_published_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scripts_published_divisionviewsCmd)
}

func Cmdscripts_published_divisionviews() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "expand", "", "Expand")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name filter")
	utils.AddFlag(listCmd.Flags(), "string", "feature", "", "Feature filter")
	utils.AddFlag(listCmd.Flags(), "string", "flowId", "", "Secure flow id filter")
	utils.AddFlag(listCmd.Flags(), "string", "scriptDataVersion", "", "Advanced usage - controls the data version of the script")
	utils.AddFlag(listCmd.Flags(), "string", "divisionIds", "", "Filters scripts to requested divisionIds")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/scripts/published/divisionviews", utils.FormatPermissions([]string{ "scripter:publishedScript:search",  }), utils.GenerateDevCentreLink("GET", "Scripts", "/api/v2/scripts/published/divisionviews")))
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
	scripts_published_divisionviewsCmd.AddCommand(listCmd)
	return scripts_published_divisionviewsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the published scripts metadata.",
	Long:  "Get the published scripts metadata.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/scripts/published/divisionviews"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		feature := utils.GetFlag(cmd.Flags(), "string", "feature")
		if feature != "" {
			queryParams["feature"] = feature
		}
		flowId := utils.GetFlag(cmd.Flags(), "string", "flowId")
		if flowId != "" {
			queryParams["flowId"] = flowId
		}
		scriptDataVersion := utils.GetFlag(cmd.Flags(), "string", "scriptDataVersion")
		if scriptDataVersion != "" {
			queryParams["scriptDataVersion"] = scriptDataVersion
		}
		divisionIds := utils.GetFlag(cmd.Flags(), "string", "divisionIds")
		if divisionIds != "" {
			queryParams["divisionIds"] = divisionIds
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
