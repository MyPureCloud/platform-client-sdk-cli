package outbound_campaigns_linedistribution

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
	Description = utils.FormatUsageDescription("outbound_campaigns_linedistribution", "SWAGGER_OVERRIDE_/api/v2/outbound/campaigns/{campaignId}/linedistribution", )
	outbound_campaigns_linedistributionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_campaigns_linedistribution"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_campaigns_linedistributionCmd)
}

func Cmdoutbound_campaigns_linedistribution() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "includeOnlyActiveCampaigns", "true", "If true will return only active Campaigns")
	utils.AddFlag(getCmd.Flags(), "string", "edgeGroupId", "", "Edge group to be used in line distribution calculations instead of current Campaign`s Edge Group. Campaign`s Site and Edge Group are mutually exclusive.")
	utils.AddFlag(getCmd.Flags(), "string", "siteId", "", "Site to be used in line distribution calculations instead of current Campaign`s Site.  Campaign`s Site and Edge Group are mutually exclusive.")
	utils.AddFlag(getCmd.Flags(), "bool", "useWeight", "", "Enable usage of weight, this value overrides current Campaign`s setting in line distribution calculations")
	utils.AddFlag(getCmd.Flags(), "int", "relativeWeight", "", "Relative weight to be used in line distribution calculations instead of current Campaign`s relative weight")
	utils.AddFlag(getCmd.Flags(), "int", "outboundLineCount", "", "The number of outbound lines to be used in line distribution calculations, instead of current Campaign`s Outbound Lines Count")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/campaigns/{campaignId}/linedistribution", utils.FormatPermissions([]string{ "outbound:lineDistribution:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/campaigns/{campaignId}/linedistribution")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CampaignOutboundLinesDistribution"
      }
    }
  }
}`)
	outbound_campaigns_linedistributionCmd.AddCommand(getCmd)
	return outbound_campaigns_linedistributionCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [campaignId]",
	Short: "Get line distribution information for campaigns using same Edge Group or Site as given campaign",
	Long:  "Get line distribution information for campaigns using same Edge Group or Site as given campaign",
	Args:  utils.DetermineArgs([]string{ "campaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns/{campaignId}/linedistribution"
		campaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{campaignId}", fmt.Sprintf("%v", campaignId), -1)

		includeOnlyActiveCampaigns := utils.GetFlag(cmd.Flags(), "bool", "includeOnlyActiveCampaigns")
		if includeOnlyActiveCampaigns != "" {
			queryParams["includeOnlyActiveCampaigns"] = includeOnlyActiveCampaigns
		}
		edgeGroupId := utils.GetFlag(cmd.Flags(), "string", "edgeGroupId")
		if edgeGroupId != "" {
			queryParams["edgeGroupId"] = edgeGroupId
		}
		siteId := utils.GetFlag(cmd.Flags(), "string", "siteId")
		if siteId != "" {
			queryParams["siteId"] = siteId
		}
		useWeight := utils.GetFlag(cmd.Flags(), "bool", "useWeight")
		if useWeight != "" {
			queryParams["useWeight"] = useWeight
		}
		relativeWeight := utils.GetFlag(cmd.Flags(), "int", "relativeWeight")
		if relativeWeight != "" {
			queryParams["relativeWeight"] = relativeWeight
		}
		outboundLineCount := utils.GetFlag(cmd.Flags(), "int", "outboundLineCount")
		if outboundLineCount != "" {
			queryParams["outboundLineCount"] = outboundLineCount
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
