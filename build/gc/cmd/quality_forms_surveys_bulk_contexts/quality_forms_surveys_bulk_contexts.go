package quality_forms_surveys_bulk_contexts

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
	Description = utils.FormatUsageDescription("quality_forms_surveys_bulk_contexts", "SWAGGER_OVERRIDE_/api/v2/quality/forms/surveys/bulk/contexts", )
	quality_forms_surveys_bulk_contextsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_forms_surveys_bulk_contexts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_forms_surveys_bulk_contextsCmd)
}

func Cmdquality_forms_surveys_bulk_contexts() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "[]string", "contextId", "", "A comma-delimited list of valid survey form context ids - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "bool", "published", "true", "If true, the latest published version will be included. If false, only the unpublished version will be included.")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/quality/forms/surveys/bulk/contexts", utils.FormatPermissions([]string{ "quality:surveyForm:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/forms/surveys/bulk/contexts")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("contextId")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/SurveyForm"
        }
      }
    }
  }
}`)
	quality_forms_surveys_bulk_contextsCmd.AddCommand(listCmd)
	return quality_forms_surveys_bulk_contextsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve a list of the latest form versions by context ids",
	Long:  "Retrieve a list of the latest form versions by context ids",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/forms/surveys/bulk/contexts"

		contextId := utils.GetFlag(cmd.Flags(), "[]string", "contextId")
		if contextId != "" {
			queryParams["contextId"] = contextId
		}
		published := utils.GetFlag(cmd.Flags(), "bool", "published")
		if published != "" {
			queryParams["published"] = published
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

		utils.Render(results)
	},
}
