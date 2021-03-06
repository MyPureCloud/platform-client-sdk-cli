package outbound_messagingcampaigns_divisionviews

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("outbound_messagingcampaigns_divisionviews", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns/divisionviews", "SWAGGER_OVERRIDE_/api/v2/outbound/messagingcampaigns/divisionviews", )
	outbound_messagingcampaigns_divisionviewsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_messagingcampaigns_divisionviews"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_messagingcampaigns_divisionviewsCmd)
}

func Cmdoutbound_messagingcampaigns_divisionviews() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}", utils.FormatPermissions([]string{ "outbound:messagingCampaign:search",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/MessagingCampaignDivisionView&quot;
  }
}`)
	outbound_messagingcampaigns_divisionviewsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size. The max that will be returned is 100.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "a", "The direction to sort Valid values: ascending, descending")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "id")
	utils.AddFlag(listCmd.Flags(), "string", "senderSmsPhoneNumber", "", "Sender SMS Phone Number")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/messagingcampaigns/divisionviews", utils.FormatPermissions([]string{ "outbound:messagingCampaign:search",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/messagingcampaigns/divisionviews")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	outbound_messagingcampaigns_divisionviewsCmd.AddCommand(listCmd)
	
	return outbound_messagingcampaigns_divisionviewsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [messagingCampaignId]",
	Short: "Get a basic Messaging Campaign information object",
	Long:  "Get a basic Messaging Campaign information object",
	Args:  utils.DetermineArgs([]string{ "messagingCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/divisionviews/{messagingCampaignId}"
		messagingCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{messagingCampaignId}", fmt.Sprintf("%v", messagingCampaignId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Query a list of basic Messaging Campaign information objects",
	Long:  "Query a list of basic Messaging Campaign information objects",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/messagingcampaigns/divisionviews"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		senderSmsPhoneNumber := utils.GetFlag(cmd.Flags(), "string", "senderSmsPhoneNumber")
		if senderSmsPhoneNumber != "" {
			queryParams["senderSmsPhoneNumber"] = senderSmsPhoneNumber
		}
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
