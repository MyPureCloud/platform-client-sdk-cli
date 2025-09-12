package analytics_dataextraction_downloads_metadata

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
	Description = utils.FormatUsageDescription("analytics_dataextraction_downloads_metadata", "SWAGGER_OVERRIDE_/api/v2/analytics/dataextraction/downloads/metadata", )
	analytics_dataextraction_downloads_metadataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_dataextraction_downloads_metadata"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_dataextraction_downloads_metadataCmd)
}

func Cmdanalytics_dataextraction_downloads_metadata() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "before", "", "The cursor that points to the start of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned.")
	utils.AddFlag(listCmd.Flags(), "string", "pageSize", "", "Number of entities to return. Maximum of 200.")
	utils.AddFlag(listCmd.Flags(), "string", "dataSchema", "", "Data schema like conversations")
	utils.AddFlag(listCmd.Flags(), "time.Time", "dateStart", "", "Start DateTime filter. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z")
	utils.AddFlag(listCmd.Flags(), "time.Time", "dateEnd", "", "End DateTime filter. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/analytics/dataextraction/downloads/metadata", utils.FormatPermissions([]string{ "analytics:datawarehouse:view",  }), utils.GenerateDevCentreLink("GET", "Analytics", "/api/v2/analytics/dataextraction/downloads/metadata")))
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
	analytics_dataextraction_downloads_metadataCmd.AddCommand(listCmd)
	return analytics_dataextraction_downloads_metadataCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get metadata on files available for extraction",
	Long:  "Get metadata on files available for extraction",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/dataextraction/downloads/metadata"

		before := utils.GetFlag(cmd.Flags(), "string", "before")
		if before != "" {
			queryParams["before"] = before
		}
		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "string", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		dataSchema := utils.GetFlag(cmd.Flags(), "string", "dataSchema")
		if dataSchema != "" {
			queryParams["dataSchema"] = dataSchema
		}
		dateStart := utils.GetFlag(cmd.Flags(), "time.Time", "dateStart")
		if dateStart != "" {
			queryParams["dateStart"] = dateStart
		}
		dateEnd := utils.GetFlag(cmd.Flags(), "time.Time", "dateEnd")
		if dateEnd != "" {
			queryParams["dateEnd"] = dateEnd
		}
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
