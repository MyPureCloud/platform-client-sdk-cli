package phones

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
	phonesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("phones"),
		Short: utils.FormatUsageDescription("Manages Genesys Cloud phones"),
		Long:  utils.FormatUsageDescription(`Manages Genesys Cloud phones`),
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(phonesCmd)
}

func Cmdphones() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/phones", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Phone&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Phone&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Phone&quot;
  }
}`)
	phonesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful.&quot;
}`)
	phonesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Phone&quot;
  }
}`)
	phonesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "name", "The field to sort by Valid values: name, status.operationalStatus, secondaryStatus.operationalStatus")
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
	utils.AddFlag(listCmd.Flags(), "string", "statusOperationalStatus", "", "The primary status to filter by")
	utils.AddFlag(listCmd.Flags(), "string", "secondaryStatusOperationalStatus", "", "The secondary status to filter by")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Fields to expand in the response, comma-separated Valid values: properties, site, status, status.primaryEdgesStatus, status.secondaryEdgesStatus, phoneBaseSettings, lines")
	utils.AddFlag(listCmd.Flags(), "[]string", "fields", "", "Fields and properties to get, comma-separated Valid values: webRtcUser, properties.*, lines.loggedInUser, lines.defaultForUser")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/phones", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/PhoneEntityListing&quot;
  }
}`)
	phonesCmd.AddCommand(listCmd)
	
	rebootCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", rebootCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/phones/{phoneId}/reboot", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(rebootCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(rebootCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ErrorBody&quot;
  },
  &quot;x-inin-error-codes&quot; : {
    &quot;bad.request&quot; : &quot;The request could not be understood by the server due to malformed syntax.&quot;,
    &quot;response.entity.too.large&quot; : &quot;The response is over the size limit. Reduce pageSize or expand list to reduce response size if applicable&quot;,
    &quot;invalid.date&quot; : &quot;Dates must be specified as ISO-8601 strings. For example: yyyy-MM-ddTHH:mm:ss.SSSZ&quot;,
    &quot;invalid.value&quot; : &quot;Value [%s] is not valid for field type [%s]. Allowable values are: %s&quot;
  }
}`)
	phonesCmd.AddCommand(rebootCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Phone&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Phone&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Phone&quot;
  }
}`)
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

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
		statusOperationalStatus := utils.GetFlag(cmd.Flags(), "string", "statusOperationalStatus")
		if statusOperationalStatus != "" {
			queryParams["statusOperationalStatus"] = statusOperationalStatus
		}
		secondaryStatusOperationalStatus := utils.GetFlag(cmd.Flags(), "string", "secondaryStatusOperationalStatus")
		if secondaryStatusOperationalStatus != "" {
			queryParams["secondaryStatusOperationalStatus"] = secondaryStatusOperationalStatus
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

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
