package integrations_speech_nuance_bots_launch_validate

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
	Description = utils.FormatUsageDescription("integrations_speech_nuance_bots_launch_validate", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/launch/validate", )
	integrations_speech_nuance_bots_launch_validateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_nuance_bots_launch_validate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_nuance_bots_launch_validateCmd)
}

func Cmdintegrations_speech_nuance_bots_launch_validate() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/launch/validate", utils.FormatPermissions([]string{ "integrations:integration:edit",  }), utils.GenerateDevCentreLink("POST", "Integrations", "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/launch/validate")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BotExecutionConfiguration"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Success",
  "content" : { }
}`)
	integrations_speech_nuance_bots_launch_validateCmd.AddCommand(createCmd)
	return integrations_speech_nuance_bots_launch_validateCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [nuanceIntegrationId]",
	Short: "Try out a single credential for a Nuance bot to know if the secret is correct",
	Long:  "Try out a single credential for a Nuance bot to know if the secret is correct",
	Args:  utils.DetermineArgs([]string{ "nuanceIntegrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Botexecutionconfiguration{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/launch/validate"
		nuanceIntegrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{nuanceIntegrationId}", fmt.Sprintf("%v", nuanceIntegrationId), -1)

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
