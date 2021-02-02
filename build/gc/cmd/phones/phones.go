package phones
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

var phonesCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("phones"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud phones"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud phones`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(phonesCmd)
}

func Cmdphones() *cobra.Command { 
	createCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/phones", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST")
	phonesCmd.AddCommand(createCmd)
	
	deleteCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE")
	phonesCmd.AddCommand(deleteCmd)
	
	getCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	phonesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "Value by which to sort")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ASC", "Sort order")
	utils.AddFlag(listCmd.Flags(), "string", "siteId", "", "Filter by site.id")
	utils.AddFlag(listCmd.Flags(), "string", "webRtcUserId", "", "Filter by webRtcUser.id")
	utils.AddFlag(listCmd.Flags(), "string", "phoneBaseSettingsId", "", "Filter by phoneBaseSettings.id")
	utils.AddFlag(listCmd.Flags(), "string", "linesLoggedInUserId", "", "Filter by lines.loggedInUser.id")
	utils.AddFlag(listCmd.Flags(), "string", "linesDefaultForUserId", "", "Filter by lines.defaultForUser.id")
	utils.AddFlag(listCmd.Flags(), "string", "phoneHardwareId", "", "Filter by phone_hardwareId")
	utils.AddFlag(listCmd.Flags(), "string", "linesId", "", "Filter by lines.id")
	utils.AddFlag(listCmd.Flags(), "string", "linesName", "", "Filter by lines.name")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Name of the Phone to filter by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Fields to expand in the response, comma-separated Valid values: properties, site, status, status.primaryEdgesStatus, status.secondaryEdgesStatus, phoneBaseSettings, lines")
	utils.AddFlag(listCmd.Flags(), "[]string", "fields", "", "Fields and properties to get, comma-separated Valid values: webRtcUser, properties.*, lines.loggedInUser, lines.defaultForUser")
	listCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/phones", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET")
	phonesCmd.AddCommand(listCmd)
	
	rebootCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", rebootCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/phones/{phoneId}/reboot", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(rebootCmd.Flags(), "POST")
	phonesCmd.AddCommand(rebootCmd)
	
	updateCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT")
	phonesCmd.AddCommand(updateCmd)
	
	return phonesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Phone",
	Long:  `Create a new Phone`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/phones"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", "create", urlString, "/api/v2/telephony/providers/edges/phones")
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
	Use:   "delete [phoneId]",
	Short: "Delete a Phone by ID",
	Long:  `Delete a Phone by ID`,
	Args:  utils.DetermineArgs([]string{ "phoneId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/phones/{phoneId}"
		phoneId, args := args[0], args[1:]
		path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", "delete", urlString, "/api/v2/telephony/providers/edges/phones/{phoneId}")
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
	Use:   "get [phoneId]",
	Short: "Get a Phone by ID",
	Long:  `Get a Phone by ID`,
	Args:  utils.DetermineArgs([]string{ "phoneId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/phones/{phoneId}"
		phoneId, args := args[0], args[1:]
		path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "get", urlString, "/api/v2/telephony/providers/edges/phones/{phoneId}")
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
	Short: "Get a list of Phone Instances",
	Long:  `Get a list of Phone Instances`,
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/phones"

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		siteId := utils.GetFlag(cmd.Flags(), "string", "siteId")
		if siteId != "" {
			queryParams["siteId"] = siteId
		}
		webRtcUserId := utils.GetFlag(cmd.Flags(), "string", "webRtcUserId")
		if webRtcUserId != "" {
			queryParams["webRtcUserId"] = webRtcUserId
		}
		phoneBaseSettingsId := utils.GetFlag(cmd.Flags(), "string", "phoneBaseSettingsId")
		if phoneBaseSettingsId != "" {
			queryParams["phoneBaseSettingsId"] = phoneBaseSettingsId
		}
		linesLoggedInUserId := utils.GetFlag(cmd.Flags(), "string", "linesLoggedInUserId")
		if linesLoggedInUserId != "" {
			queryParams["linesLoggedInUserId"] = linesLoggedInUserId
		}
		linesDefaultForUserId := utils.GetFlag(cmd.Flags(), "string", "linesDefaultForUserId")
		if linesDefaultForUserId != "" {
			queryParams["linesDefaultForUserId"] = linesDefaultForUserId
		}
		phoneHardwareId := utils.GetFlag(cmd.Flags(), "string", "phoneHardwareId")
		if phoneHardwareId != "" {
			queryParams["phoneHardwareId"] = phoneHardwareId
		}
		linesId := utils.GetFlag(cmd.Flags(), "string", "linesId")
		if linesId != "" {
			queryParams["linesId"] = linesId
		}
		linesName := utils.GetFlag(cmd.Flags(), "string", "linesName")
		if linesName != "" {
			queryParams["linesName"] = linesName
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		fields := utils.GetFlag(cmd.Flags(), "[]string", "fields")
		if fields != "" {
			queryParams["fields"] = fields
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", "list", urlString, "/api/v2/telephony/providers/edges/phones")
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
var rebootCmd = &cobra.Command{
	Use:   "reboot [phoneId]",
	Short: "Reboot a Phone",
	Long:  `Reboot a Phone`,
	Args:  utils.DetermineArgs([]string{ "phoneId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/phones/{phoneId}/reboot"
		phoneId, args := args[0], args[1:]
		path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", "reboot", urlString, "/api/v2/telephony/providers/edges/phones/{phoneId}/reboot")
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
	Use:   "update [phoneId]",
	Short: "Update a Phone by ID",
	Long:  `Update a Phone by ID`,
	Args:  utils.DetermineArgs([]string{ "phoneId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/phones/{phoneId}"
		phoneId, args := args[0], args[1:]
		path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", "update", urlString, "/api/v2/telephony/providers/edges/phones/{phoneId}")
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
