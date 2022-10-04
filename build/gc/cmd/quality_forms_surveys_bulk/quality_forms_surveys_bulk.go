package quality_forms_surveys_bulk

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
	Description = utils.FormatUsageDescription("quality_forms_surveys_bulk", "SWAGGER_OVERRIDE_/api/v2/quality/forms/surveys/bulk", )
	quality_forms_surveys_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_forms_surveys_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_forms_surveys_bulkCmd)
}

func Cmdquality_forms_surveys_bulk() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "A comma-delimited list of valid survey form ids - REQUIRED")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/quality/forms/surveys/bulk", utils.FormatPermissions([]string{ "quality:surveyForm:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/forms/surveys/bulk")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("id")
	
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
	quality_forms_surveys_bulkCmd.AddCommand(listCmd)
	return quality_forms_surveys_bulkCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of survey forms by their ids",
	Long:  "Retrieve a list of survey forms by their ids",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/forms/surveys/bulk"

		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
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
