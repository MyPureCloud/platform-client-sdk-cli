package outbound_campaigns_callback_schedule

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
	Description = utils.FormatUsageDescription("outbound_campaigns_callback_schedule", "SWAGGER_OVERRIDE_/api/v2/outbound/campaigns/{campaignId}/callback/schedule", )
	outbound_campaigns_callback_scheduleCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_campaigns_callback_schedule"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_campaigns_callback_scheduleCmd)
}

func Cmdoutbound_campaigns_callback_schedule() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/campaigns/{campaignId}/callback/schedule", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/campaigns/{campaignId}/callback/schedule")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "ContactCallbackRequest",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ContactCallbackRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ContactCallbackRequest"
      }
    }
  }
}`)
	outbound_campaigns_callback_scheduleCmd.AddCommand(createCmd)
	return outbound_campaigns_callback_scheduleCmd
}

var createCmd = &cobra.Command{
	Use:   "create [campaignId]",
	Short: "Schedule a Callback for a Dialer Campaign (Deprecated)",
	Long:  "Schedule a Callback for a Dialer Campaign (Deprecated)",
	Args:  utils.DetermineArgs([]string{ "campaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Contactcallbackrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns/{campaignId}/callback/schedule"
		campaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{campaignId}", fmt.Sprintf("%v", campaignId), -1)

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
