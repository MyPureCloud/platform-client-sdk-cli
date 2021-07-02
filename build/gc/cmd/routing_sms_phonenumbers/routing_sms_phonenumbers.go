package routing_sms_phonenumbers

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
	Description = utils.FormatUsageDescription("routing_sms_phonenumbers", "SWAGGER_OVERRIDE_/api/v2/routing/sms/phonenumbers", "SWAGGER_OVERRIDE_/api/v2/routing/sms/phonenumbers", "SWAGGER_OVERRIDE_/api/v2/routing/sms/phonenumbers", "SWAGGER_OVERRIDE_/api/v2/routing/sms/phonenumbers", "SWAGGER_OVERRIDE_/api/v2/routing/sms/phonenumbers", )
	routing_sms_phonenumbersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_sms_phonenumbers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_sms_phonenumbersCmd)
}

func Cmdrouting_sms_phonenumbers() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/routing/sms/phonenumbers", utils.FormatPermissions([]string{ "sms:phoneNumber:add",  }), utils.GenerateDevCentreLink("POST", "Routing", "/api/v2/routing/sms/phonenumbers")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;SmsPhoneNumber&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SmsPhoneNumberProvision&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SmsPhoneNumber&quot;
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(createCmd)
	
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/routing/sms/phonenumbers/{addressId}", utils.FormatPermissions([]string{ "sms:phoneNumber:delete",  }), utils.GenerateDevCentreLink("DELETE", "Routing", "/api/v2/routing/sms/phonenumbers/{addressId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful&quot;
}`)
	routing_sms_phonenumbersCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/routing/sms/phonenumbers/{addressId}", utils.FormatPermissions([]string{ "sms:phoneNumber:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/sms/phonenumbers/{addressId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SmsPhoneNumber&quot;
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "phoneNumber", "", "Filter on phone number address. Allowable characters are the digits `0-9` and the wild card character `\\*`. If just digits are present, a contains search is done on the address pattern. For example, `317` could be matched anywhere in the address. An `\\*` will match multiple digits. For example, to match a specific area code within the US a pattern like `1317*` could be used.")
	utils.AddFlag(listCmd.Flags(), "string", "phoneNumberType", "", "Filter on phone number type Valid values: local, mobile, tollfree, shortcode")
	utils.AddFlag(listCmd.Flags(), "string", "phoneNumberStatus", "", "Filter on phone number status Valid values: active, invalid, porting")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/sms/phonenumbers", utils.FormatPermissions([]string{ "sms:phoneNumber:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/sms/phonenumbers")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/routing/sms/phonenumbers/{addressId}", utils.FormatPermissions([]string{ "sms:phoneNumber:edit",  }), utils.GenerateDevCentreLink("PUT", "Routing", "/api/v2/routing/sms/phonenumbers/{addressId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;SmsPhoneNumber&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SmsPhoneNumber&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SmsPhoneNumber&quot;
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(updateCmd)
	
	return routing_sms_phonenumbersCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Provision a phone number for SMS",
	Long:  "Provision a phone number for SMS",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers"

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
	Use:   "delete [addressId]",
	Short: "Delete a phone number provisioned for SMS.",
	Long:  "Delete a phone number provisioned for SMS.",
	Args:  utils.DetermineArgs([]string{ "addressId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers/{addressId}"
		addressId, args := args[0], args[1:]
		path = strings.Replace(path, "{addressId}", fmt.Sprintf("%v", addressId), -1)

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
	Use:   "get [addressId]",
	Short: "Get a phone number provisioned for SMS.",
	Long:  "Get a phone number provisioned for SMS.",
	Args:  utils.DetermineArgs([]string{ "addressId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers/{addressId}"
		addressId, args := args[0], args[1:]
		path = strings.Replace(path, "{addressId}", fmt.Sprintf("%v", addressId), -1)

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
	Short: "Get a list of provisioned phone numbers.",
	Long:  "Get a list of provisioned phone numbers.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers"

		phoneNumber := utils.GetFlag(cmd.Flags(), "string", "phoneNumber")
		if phoneNumber != "" {
			queryParams["phoneNumber"] = phoneNumber
		}
		phoneNumberType := utils.GetFlag(cmd.Flags(), "string", "phoneNumberType")
		if phoneNumberType != "" {
			queryParams["phoneNumberType"] = phoneNumberType
		}
		phoneNumberStatus := utils.GetFlag(cmd.Flags(), "string", "phoneNumberStatus")
		if phoneNumberStatus != "" {
			queryParams["phoneNumberStatus"] = phoneNumberStatus
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
	Use:   "update [addressId]",
	Short: "Update a phone number provisioned for SMS.",
	Long:  "Update a phone number provisioned for SMS.",
	Args:  utils.DetermineArgs([]string{ "addressId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers/{addressId}"
		addressId, args := args[0], args[1:]
		path = strings.Replace(path, "{addressId}", fmt.Sprintf("%v", addressId), -1)

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
