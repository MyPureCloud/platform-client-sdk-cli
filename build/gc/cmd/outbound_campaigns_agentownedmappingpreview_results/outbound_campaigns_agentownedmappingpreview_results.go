package outbound_campaigns_agentownedmappingpreview_results

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
	Description = utils.FormatUsageDescription("outbound_campaigns_agentownedmappingpreview_results", "SWAGGER_OVERRIDE_/api/v2/outbound/campaigns/{campaignId}/agentownedmappingpreview/results", )
	outbound_campaigns_agentownedmappingpreview_resultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_campaigns_agentownedmappingpreview_results"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_campaigns_agentownedmappingpreview_resultsCmd)
}

func Cmdoutbound_campaigns_agentownedmappingpreview_results() *cobra.Command { 
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/campaigns/{campaignId}/agentownedmappingpreview/results", utils.FormatPermissions([]string{ "outbound:campaign:view", "outbound:contact:view", "routing:queue:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/campaigns/{campaignId}/agentownedmappingpreview/results")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/AgentOwnedMappingPreviewListing"
  }
}`)
	outbound_campaigns_agentownedmappingpreview_resultsCmd.AddCommand(listCmd)
	
	return outbound_campaigns_agentownedmappingpreview_resultsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [campaignId]",
	Short: "Get a preview of how agents will be mapped to this campaign`s contact list.",
	Long:  "Get a preview of how agents will be mapped to this campaign`s contact list.",
	Args:  utils.DetermineArgs([]string{ "campaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns/{campaignId}/agentownedmappingpreview/results"
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

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
