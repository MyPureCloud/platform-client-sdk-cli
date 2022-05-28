package routing_sms_availablephonenumbers

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
	Description = utils.FormatUsageDescription("routing_sms_availablephonenumbers", "SWAGGER_OVERRIDE_/api/v2/routing/sms/availablephonenumbers", )
	routing_sms_availablephonenumbersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_sms_availablephonenumbers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_sms_availablephonenumbersCmd)
}

func Cmdrouting_sms_availablephonenumbers() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "countryCode", "", "The ISO 3166-1 alpha-2 country code of the county for which available phone numbers should be returned - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "string", "region", "", "Region/province/state that can be used to restrict the numbers returned")
	utils.AddFlag(listCmd.Flags(), "string", "city", "", "City that can be used to restrict the numbers returned")
	utils.AddFlag(listCmd.Flags(), "string", "areaCode", "", "Area code that can be used to restrict the numbers returned")
	utils.AddFlag(listCmd.Flags(), "string", "phoneNumberType", "", "Type of available phone numbers searched - REQUIRED Valid values: local, mobile, tollfree")
	utils.AddFlag(listCmd.Flags(), "string", "pattern", "", "A pattern to match phone numbers. Valid characters are `*` and [0-9a-zA-Z]. The `*` character will match any single digit.")
	utils.AddFlag(listCmd.Flags(), "string", "addressRequirement", "", "This indicates whether the phone number requires to have an Address registered. Valid values: none, any, local, foreign")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/routing/sms/availablephonenumbers", utils.FormatPermissions([]string{ "sms:phoneNumber:add",  }), utils.GenerateDevCentreLink("GET", "Routing", "/api/v2/routing/sms/availablephonenumbers")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("countryCode")
	listCmd.MarkFlagRequired("phoneNumberType")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SMSAvailablePhoneNumberEntityListing"
      }
    }
  }
}`)
	routing_sms_availablephonenumbersCmd.AddCommand(listCmd)
	return routing_sms_availablephonenumbersCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of available phone numbers for SMS provisioning.",
	Long:  "Get a list of available phone numbers for SMS provisioning.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/routing/sms/availablephonenumbers"

		countryCode := utils.GetFlag(cmd.Flags(), "string", "countryCode")
		if countryCode != "" {
			queryParams["countryCode"] = countryCode
		}
		region := utils.GetFlag(cmd.Flags(), "string", "region")
		if region != "" {
			queryParams["region"] = region
		}
		city := utils.GetFlag(cmd.Flags(), "string", "city")
		if city != "" {
			queryParams["city"] = city
		}
		areaCode := utils.GetFlag(cmd.Flags(), "string", "areaCode")
		if areaCode != "" {
			queryParams["areaCode"] = areaCode
		}
		phoneNumberType := utils.GetFlag(cmd.Flags(), "string", "phoneNumberType")
		if phoneNumberType != "" {
			queryParams["phoneNumberType"] = phoneNumberType
		}
		pattern := utils.GetFlag(cmd.Flags(), "string", "pattern")
		if pattern != "" {
			queryParams["pattern"] = pattern
		}
		addressRequirement := utils.GetFlag(cmd.Flags(), "string", "addressRequirement")
		if addressRequirement != "" {
			queryParams["addressRequirement"] = addressRequirement
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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

		utils.Render(results)
	},
}
