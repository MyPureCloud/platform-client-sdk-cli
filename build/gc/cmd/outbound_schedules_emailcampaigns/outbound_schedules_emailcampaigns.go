package outbound_schedules_emailcampaigns

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
	Description = utils.FormatUsageDescription("outbound_schedules_emailcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/emailcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/emailcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/emailcampaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules/emailcampaigns", )
	outbound_schedules_emailcampaignsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_schedules_emailcampaigns"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_schedules_emailcampaignsCmd)
}

func Cmdoutbound_schedules_emailcampaigns() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}", utils.FormatPermissions([]string{ "outbound:emailCampaignSchedule:delete",  }), utils.GenerateDevCentreLink("DELETE", "Outbound", "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "Operation was successful."
}`)
	outbound_schedules_emailcampaignsCmd.AddCommand(deleteCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}", utils.FormatPermissions([]string{ "outbound:emailCampaignSchedule:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/EmailCampaignSchedule"
  }
}`)
	outbound_schedules_emailcampaignsCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/outbound/schedules/emailcampaigns", utils.FormatPermissions([]string{ "outbound:emailCampaignSchedule:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/schedules/emailcampaigns")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	outbound_schedules_emailcampaignsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}", utils.FormatPermissions([]string{ "outbound:emailCampaignSchedule:edit",  }), utils.GenerateDevCentreLink("PUT", "Outbound", "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "EmailCampaignSchedule",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/EmailCampaignSchedule"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/EmailCampaignSchedule"
  }
}`)
	outbound_schedules_emailcampaignsCmd.AddCommand(updateCmd)
	
	return outbound_schedules_emailcampaignsCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [emailCampaignId]",
	Short: "Delete an email campaign schedule.",
	Long:  "Delete an email campaign schedule.",
	Args:  utils.DetermineArgs([]string{ "emailCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}"
		emailCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{emailCampaignId}", fmt.Sprintf("%v", emailCampaignId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [emailCampaignId]",
	Short: "Get an email campaign schedule.",
	Long:  "Get an email campaign schedule.",
	Args:  utils.DetermineArgs([]string{ "emailCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}"
		emailCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{emailCampaignId}", fmt.Sprintf("%v", emailCampaignId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Query for a list of email campaign schedules.",
	Long:  "Query for a list of email campaign schedules.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/emailcampaigns"

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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [emailCampaignId]",
	Short: "Update an email campaign schedule.",
	Long:  "Update an email campaign schedule.",
	Args:  utils.DetermineArgs([]string{ "emailCampaignId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Emailcampaignschedule{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/schedules/emailcampaigns/{emailCampaignId}"
		emailCampaignId, args := args[0], args[1:]
		path = strings.Replace(path, "{emailCampaignId}", fmt.Sprintf("%v", emailCampaignId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
