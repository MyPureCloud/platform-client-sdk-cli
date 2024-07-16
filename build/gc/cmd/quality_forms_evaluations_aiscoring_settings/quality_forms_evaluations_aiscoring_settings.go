package quality_forms_evaluations_aiscoring_settings

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
	Description = utils.FormatUsageDescription("quality_forms_evaluations_aiscoring_settings", "SWAGGER_OVERRIDE_/api/v2/quality/forms/evaluations/{formId}/aiscoring/settings", )
	quality_forms_evaluations_aiscoring_settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_forms_evaluations_aiscoring_settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_forms_evaluations_aiscoring_settingsCmd)
}

func Cmdquality_forms_evaluations_aiscoring_settings() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/quality/forms/evaluations/{formId}/aiscoring/settings", utils.FormatPermissions([]string{ "quality:evaluationForm:aiScoringEdit",  }), utils.GenerateDevCentreLink("PUT", "Quality", "/api/v2/quality/forms/evaluations/{formId}/aiscoring/settings")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "AI Scoring Settings",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AiScoringSettings"
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
        "$ref" : "#/components/schemas/AiScoringSettings"
      }
    }
  }
}`)
	quality_forms_evaluations_aiscoring_settingsCmd.AddCommand(updateCmd)
	return quality_forms_evaluations_aiscoring_settingsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var updateCmd = &cobra.Command{
	Use:   "update [formId]",
	Short: "Update the AI Scoring settings of an evaluation form.",
	Long:  "Update the AI Scoring settings of an evaluation form.",
	Args:  utils.DetermineArgs([]string{ "formId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Aiscoringsettings{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/forms/evaluations/{formId}/aiscoring/settings"
		formId, args := args[0], args[1:]
		path = strings.Replace(path, "{formId}", fmt.Sprintf("%v", formId), -1)

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
