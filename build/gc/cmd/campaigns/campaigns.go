package campaigns
import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

var campaignsCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("campaigns"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud campaigns"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud campaigns`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(campaignsCmd)
}

func Cmdcampaigns() *cobra.Command { 
	createCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/campaigns", utils.FormatPermissions([]string{ "outbound:campaign:add",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	campaignsCmd.AddCommand(createCmd)
	
	deleteCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/campaigns/{campaignId}", utils.FormatPermissions([]string{ "outbound:campaign:delete",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	campaignsCmd.AddCommand(deleteCmd)
	
	getCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/campaigns/{campaignId}", utils.FormatPermissions([]string{ "outbound:campaign:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	campaignsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size. The max that will be returned is 100.")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "filterType", "Prefix", "Filter type Valid values: Equals, RegEx, Contains, Prefix, LessThan, LessThanEqualTo, GreaterThan, GreaterThanEqualTo, BeginsWith, EndsWith")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "id")
	utils.AddFlag(listCmd.Flags(), "string", "contactListId", "", "Contact List ID")
	utils.AddFlag(listCmd.Flags(), "string", "dncListIds", "", "DNC list ID")
	utils.AddFlag(listCmd.Flags(), "string", "distributionQueueId", "", "Distribution queue ID")
	utils.AddFlag(listCmd.Flags(), "string", "edgeGroupId", "", "Edge group ID")
	utils.AddFlag(listCmd.Flags(), "string", "callAnalysisResponseSetId", "", "Call analysis response set ID")
	utils.AddFlag(listCmd.Flags(), "[]string", "divisionId", "", "Division ID(s)")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "a", "Sort order Valid values: ascending, descending")
	listCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/campaigns", utils.FormatPermissions([]string{ "outbound:campaign:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	campaignsCmd.AddCommand(listCmd)
	
	updateCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/outbound/campaigns/{campaignId}", utils.FormatPermissions([]string{ "outbound:campaign:edit",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	campaignsCmd.AddCommand(updateCmd)
	
	return campaignsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a campaign.",
	Long:  `Create a campaign.`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", "create", urlString, "/api/v2/outbound/campaigns")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [campaignId]",
	Short: "Delete a campaign.",
	Long:  `Delete a campaign.`,
	Args:  utils.DetermineArgs([]string{ "campaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns/{campaignId}"
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

		retryFunc := CommandService.DetermineAction("DELETE", "delete", urlString, "/api/v2/outbound/campaigns/{campaignId}")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [campaignId]",
	Short: "Get dialer campaign.",
	Long:  `Get dialer campaign.`,
	Args:  utils.DetermineArgs([]string{ "campaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns/{campaignId}"
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

		retryFunc := CommandService.DetermineAction("GET", "get", urlString, "/api/v2/outbound/campaigns/{campaignId}")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
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
	Short: "Query a list of dialer campaigns.",
	Long:  `Query a list of dialer campaigns.`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		filterType := utils.GetFlag(cmd.Flags(), "string", "filterType")
		if filterType != "" {
			queryParams["filterType"] = filterType
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		contactListId := utils.GetFlag(cmd.Flags(), "string", "contactListId")
		if contactListId != "" {
			queryParams["contactListId"] = contactListId
		}
		dncListIds := utils.GetFlag(cmd.Flags(), "string", "dncListIds")
		if dncListIds != "" {
			queryParams["dncListIds"] = dncListIds
		}
		distributionQueueId := utils.GetFlag(cmd.Flags(), "string", "distributionQueueId")
		if distributionQueueId != "" {
			queryParams["distributionQueueId"] = distributionQueueId
		}
		edgeGroupId := utils.GetFlag(cmd.Flags(), "string", "edgeGroupId")
		if edgeGroupId != "" {
			queryParams["edgeGroupId"] = edgeGroupId
		}
		callAnalysisResponseSetId := utils.GetFlag(cmd.Flags(), "string", "callAnalysisResponseSetId")
		if callAnalysisResponseSetId != "" {
			queryParams["callAnalysisResponseSetId"] = callAnalysisResponseSetId
		}
		divisionId := utils.GetFlag(cmd.Flags(), "[]string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "list", urlString, "/api/v2/outbound/campaigns")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [campaignId]",
	Short: "Update a campaign.",
	Long:  `Update a campaign.`,
	Args:  utils.DetermineArgs([]string{ "campaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/outbound/campaigns/{campaignId}"
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

		retryFunc := CommandService.DetermineAction("PUT", "update", urlString, "/api/v2/outbound/campaigns/{campaignId}")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
