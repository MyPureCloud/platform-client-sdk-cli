package integrations_unifiedcommunications_thirdpartypresences

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
	Description = utils.FormatUsageDescription("integrations_unifiedcommunications_thirdpartypresences", "SWAGGER_OVERRIDE_/api/v2/integrations/unifiedcommunications/{ucIntegrationId}/thirdpartypresences", )
	integrations_unifiedcommunications_thirdpartypresencesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_unifiedcommunications_thirdpartypresences"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_unifiedcommunications_thirdpartypresencesCmd)
}

func Cmdintegrations_unifiedcommunications_thirdpartypresences() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/integrations/unifiedcommunications/{ucIntegrationId}/thirdpartypresences", utils.FormatPermissions([]string{ "integration:presence:edit",  }), utils.GenerateDevCentreLink("PUT", "Integrations", "/api/v2/integrations/unifiedcommunications/{ucIntegrationId}/thirdpartypresences")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "List of User presences",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "array",
        "items" : {
          "$ref" : "#/components/schemas/UCThirdPartyPresence"
        }
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "Accepted",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "string"
      }
    }
  }
}`)
	integrations_unifiedcommunications_thirdpartypresencesCmd.AddCommand(updateCmd)
	return integrations_unifiedcommunications_thirdpartypresencesCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var updateCmd = &cobra.Command{
	Use:   "update [ucIntegrationId]",
	Short: "Bulk integration presence ingestion",
	Long:  "Bulk integration presence ingestion",
	Args:  utils.DetermineArgs([]string{ "ucIntegrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Ucthirdpartypresence{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/unifiedcommunications/{ucIntegrationId}/thirdpartypresences"
		ucIntegrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{ucIntegrationId}", fmt.Sprintf("%v", ucIntegrationId), -1)

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
