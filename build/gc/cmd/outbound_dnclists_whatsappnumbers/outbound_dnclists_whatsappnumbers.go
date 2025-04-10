package outbound_dnclists_whatsappnumbers

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
	Description = utils.FormatUsageDescription("outbound_dnclists_whatsappnumbers", "SWAGGER_OVERRIDE_/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers", "SWAGGER_OVERRIDE_/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers", )
	outbound_dnclists_whatsappnumbersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_dnclists_whatsappnumbers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_dnclists_whatsappnumbersCmd)
}

func Cmdoutbound_dnclists_whatsappnumbers() *cobra.Command { 
	utils.AddFlag(deleteCmd.Flags(), "bool", "expiredOnly", "false", "Set to true to only remove DNC whatsApp numbers that are expired")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers", utils.FormatPermissions([]string{ "outbound:dnc:delete",  }), utils.GenerateDevCentreLink("DELETE", "Outbound", "/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "DNC whatsApp numbers deleted.",
  "content" : { }
}`)
	outbound_dnclists_whatsappnumbersCmd.AddCommand(deleteCmd)

	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers", utils.FormatPermissions([]string{ "outbound:dnc:edit",  }), utils.GenerateDevCentreLink("PATCH", "Outbound", "/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "DNC whatsApp numbers",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/DncPatchWhatsAppNumbersRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "No Content",
  "content" : { }
}`)
	outbound_dnclists_whatsappnumbersCmd.AddCommand(updateCmd)
	return outbound_dnclists_whatsappnumbersCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [dncListId]",
	Short: "Deletes all or expired whatsApp numbers from a DNC list.",
	Long:  "Deletes all or expired whatsApp numbers from a DNC list.",
	Args:  utils.DetermineArgs([]string{ "dncListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers"
		dncListId, args := args[0], args[1:]
		path = strings.Replace(path, "{dncListId}", fmt.Sprintf("%v", dncListId), -1)

		expiredOnly := utils.GetFlag(cmd.Flags(), "bool", "expiredOnly")
		if expiredOnly != "" {
			queryParams["expiredOnly"] = expiredOnly
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
var updateCmd = &cobra.Command{
	Use:   "update [dncListId]",
	Short: "Add entries to or delete entries from a DNC list.",
	Long:  "Add entries to or delete entries from a DNC list.",
	Args:  utils.DetermineArgs([]string{ "dncListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Dncpatchwhatsappnumbersrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/dnclists/{dncListId}/whatsappnumbers"
		dncListId, args := args[0], args[1:]
		path = strings.Replace(path, "{dncListId}", fmt.Sprintf("%v", dncListId), -1)

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
		const httpMethod = "PATCH"
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
