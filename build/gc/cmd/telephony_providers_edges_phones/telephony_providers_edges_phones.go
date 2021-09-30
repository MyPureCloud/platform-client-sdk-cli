package telephony_providers_edges_phones

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_phones", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/phones", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/phones", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/phones", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/phones", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/phones", )
	telephony_providers_edges_phonesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_phones"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_phonesCmd)
}

func Cmdtelephony_providers_edges_phones() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/phones", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("POST", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/phones")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Phone",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/Phone"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Phone"
  }
}`)
	telephony_providers_edges_phonesCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("DELETE", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/phones/{phoneId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Operation was successful."
}`)
	telephony_providers_edges_phonesCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/phones/{phoneId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Phone"
  }
}`)
	telephony_providers_edges_phonesCmd.AddCommand(getCmd)
	
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
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/phones", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/phones")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	telephony_providers_edges_phonesCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/telephony/providers/edges/phones/{phoneId}", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("PUT", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/phones/{phoneId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "Phone",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/Phone"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Phone"
  }
}`)
	telephony_providers_edges_phonesCmd.AddCommand(updateCmd)
	
	return telephony_providers_edges_phonesCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Phone",
	Long:  "Create a new Phone",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Phone{}
			utils.Render(reqModel.String())
			
			return
		}

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
	Long:  "Delete a Phone by ID",
	Args:  utils.DetermineArgs([]string{ "phoneId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

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
	Long:  "Get a Phone by ID",
	Args:  utils.DetermineArgs([]string{ "phoneId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

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
	Long:  "Get a list of Phone Instances",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

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
var updateCmd = &cobra.Command{
	Use:   "update [phoneId]",
	Short: "Update a Phone by ID",
	Long:  "Update a Phone by ID",
	Args:  utils.DetermineArgs([]string{ "phoneId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Phone{}
			utils.Render(reqModel.String())
			
			return
		}

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
