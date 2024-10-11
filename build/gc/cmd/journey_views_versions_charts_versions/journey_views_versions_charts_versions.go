package journey_views_versions_charts_versions

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
	Description = utils.FormatUsageDescription("journey_views_versions_charts_versions", "SWAGGER_OVERRIDE_/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/charts/{chartId}/versions", )
	journey_views_versions_charts_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_views_versions_charts_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_views_versions_charts_versionsCmd)
}

func Cmdjourney_views_versions_charts_versions() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/charts/{chartId}/versions/{chartVersion}", utils.FormatPermissions([]string{ "journey:views:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/charts/{chartId}/versions/{chartVersion}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/JourneyViewChart"
      }
    }
  }
}`)
	journey_views_versions_charts_versionsCmd.AddCommand(getCmd)
	return journey_views_versions_charts_versionsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [viewId] [journeyViewVersion] [chartId] [chartVersion]",
	Short: "Get a Chart by ID and version",
	Long:  "Get a Chart by ID and version",
	Args:  utils.DetermineArgs([]string{ "viewId", "journeyViewVersion", "chartId", "chartVersion", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/journey/views/{viewId}/versions/{journeyViewVersion}/charts/{chartId}/versions/{chartVersion}"
		viewId, args := args[0], args[1:]
		path = strings.Replace(path, "{viewId}", fmt.Sprintf("%v", viewId), -1)
		journeyViewVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{journeyViewVersion}", fmt.Sprintf("%v", journeyViewVersion), -1)
		chartId, args := args[0], args[1:]
		path = strings.Replace(path, "{chartId}", fmt.Sprintf("%v", chartId), -1)
		chartVersion, args := args[0], args[1:]
		path = strings.Replace(path, "{chartVersion}", fmt.Sprintf("%v", chartVersion), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
