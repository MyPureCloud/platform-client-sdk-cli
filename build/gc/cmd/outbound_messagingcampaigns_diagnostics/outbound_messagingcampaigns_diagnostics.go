package outbound_messagingcampaigns_diagnostics

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
	Description = utils.FormatUsageDescription("outbound_messagingcampaigns_diagnostics", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/diagnostics", )
	outbound_messagingcampaigns_diagnosticsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_messagingcampaigns_diagnostics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_messagingcampaigns_diagnosticsCmd)
}

func Cmdoutbound_messagingcampaigns_diagnostics() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/diagnostics", utils.FormatPermissions([]string{ "outbound:messagingCampaign:view", "outbound:emailCampaign:view", "outbound:whatsAppCampaign:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/diagnostics")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/MessagingCampaignDiagnostics"
      }
    }
  }
}`)
	outbound_messagingcampaigns_diagnosticsCmd.AddCommand(getCmd)
	return outbound_messagingcampaigns_diagnosticsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [messagingCampaignId]",
	Short: "Get messaging campaign diagnostics",
	Long:  "Get messaging campaign diagnostics",
	Args:  utils.DetermineArgs([]string{ "messagingCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/{messagingCampaignId}/diagnostics"
		messagingCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{messagingCampaignId}", fmt.Sprintf("%v", messagingCampaignId), -1)

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
