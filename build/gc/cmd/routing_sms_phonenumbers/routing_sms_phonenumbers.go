package routing_sms_phonenumbers

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
  "description" : "SmsPhoneNumber",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SmsPhoneNumberProvision"
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
        "$ref" : "#/components/schemas/SmsPhoneNumber"
      }
    }
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/routing/sms/phonenumbers/{addressId}", utils.FormatPermissions([]string{ "sms:phoneNumber:delete",  }), utils.GenerateDevCentreLink("DELETE", "Routing", "/api/v2/routing/sms/phonenumbers/{addressId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Accepted -The phone number delete is in progress.",
  "content" : { }
}`)
	routing_sms_phonenumbersCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "Expand response with additional information Valid values: compliance, supportedContent")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/routing/sms/phonenumbers/{addressId}", utils.FormatPermissions([]string{ "sms:phoneNumber:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/sms/phonenumbers/{addressId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SmsPhoneNumber"
      }
    }
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "phoneNumber", "", "Filter on phone number address. Allowable characters are the digits `0-9` and the wild card character `\\*`. If just digits are present, a contains search is done on the address pattern. For example, `317` could be matched anywhere in the address. An `\\*` will match multiple digits. For example, to match a specific area code within the US a pattern like `1317*` could be used.")
	utils.AddFlag(listCmd.Flags(), "[]string", "phoneNumberType", "", "Filter on phone number type Valid values: local, mobile, tollfree, shortcode, alphanumeric")
	utils.AddFlag(listCmd.Flags(), "[]string", "phoneNumberStatus", "", "Filter on phone number status Valid values: active, invalid, initiated, porting, pending, pending-cancellation")
	utils.AddFlag(listCmd.Flags(), "[]string", "countryCode", "", "Filter on country code")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "Optional field to sort results Valid values: phoneNumber, countryCode, country, dateCreated, dateModified, phoneNumberStatus, phoneNumberType, purchaseDate, supportsMms, supportsSms, supportsVoice")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "", "Sort order Valid values: ascending, descending")
	utils.AddFlag(listCmd.Flags(), "string", "language", "en-US", "A language tag (which is sometimes referred to as a locale identifier) to use to localize country field and sort operations")
	utils.AddFlag(listCmd.Flags(), "string", "integrationId", "", "Filter on the Genesys Cloud integration id to which the phone number belongs to")
	utils.AddFlag(listCmd.Flags(), "string", "supportedContentId", "", "Filter based on the supported content ID")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/sms/phonenumbers", utils.FormatPermissions([]string{ "sms:phoneNumber:view",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/sms/phonenumbers")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(listCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/routing/sms/phonenumbers/{addressId}", utils.FormatPermissions([]string{ "sms:phoneNumber:edit",  }), utils.GenerateDevCentreLink("PUT", "Routing", "/api/v2/routing/sms/phonenumbers/{addressId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "SmsPhoneNumber",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SmsPhoneNumber"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SmsPhoneNumber"
      }
    }
  }
}`)
	routing_sms_phonenumbersCmd.AddCommand(updateCmd)
	return routing_sms_phonenumbersCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Provision a phone number for SMS",
	Long:  "Provision a phone number for SMS",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Smsphonenumberprovision{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers"

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
var deleteCmd = &cobra.Command{
	Use:   "delete [addressId]",
	Short: "Delete a phone number provisioned for SMS.",
	Long:  "Delete a phone number provisioned for SMS.",
	Args:  utils.DetermineArgs([]string{ "addressId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers/{addressId}"
		addressId, args := args[0], args[1:]
		path = strings.Replace(path, "{addressId}", fmt.Sprintf("%v", addressId), -1)

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

		const opId = "delete"
		const httpMethod = "DELETE"
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
var getCmd = &cobra.Command{
	Use:   "get [addressId]",
	Short: "Get a phone number provisioned for SMS.",
	Long:  "Get a phone number provisioned for SMS.",
	Args:  utils.DetermineArgs([]string{ "addressId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers/{addressId}"
		addressId, args := args[0], args[1:]
		path = strings.Replace(path, "{addressId}", fmt.Sprintf("%v", addressId), -1)

		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of provisioned phone numbers.",
	Long:  "Get a list of provisioned phone numbers.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers"

		phoneNumber := utils.GetFlag(cmd.Flags(), "string", "phoneNumber")
		if phoneNumber != "" {
			queryParams["phoneNumber"] = phoneNumber
		}
		phoneNumberType := utils.GetFlag(cmd.Flags(), "[]string", "phoneNumberType")
		if phoneNumberType != "" {
			queryParams["phoneNumberType"] = phoneNumberType
		}
		phoneNumberStatus := utils.GetFlag(cmd.Flags(), "[]string", "phoneNumberStatus")
		if phoneNumberStatus != "" {
			queryParams["phoneNumberStatus"] = phoneNumberStatus
		}
		countryCode := utils.GetFlag(cmd.Flags(), "[]string", "countryCode")
		if countryCode != "" {
			queryParams["countryCode"] = countryCode
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		language := utils.GetFlag(cmd.Flags(), "string", "language")
		if language != "" {
			queryParams["language"] = language
		}
		integrationId := utils.GetFlag(cmd.Flags(), "string", "integrationId")
		if integrationId != "" {
			queryParams["integration.id"] = integrationId
		}
		supportedContentId := utils.GetFlag(cmd.Flags(), "string", "supportedContentId")
		if supportedContentId != "" {
			queryParams["supportedContent.id"] = supportedContentId
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

		const opId = "list"
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
var updateCmd = &cobra.Command{
	Use:   "update [addressId]",
	Short: "Update a phone number provisioned for SMS.",
	Long:  "Update a phone number provisioned for SMS.",
	Args:  utils.DetermineArgs([]string{ "addressId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Smsphonenumber{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/phonenumbers/{addressId}"
		addressId, args := args[0], args[1:]
		path = strings.Replace(path, "{addressId}", fmt.Sprintf("%v", addressId), -1)

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

		const opId = "update"
		const httpMethod = "PUT"
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
