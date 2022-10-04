package outbound_contactlists_timezonemappingpreview

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
	Description = utils.FormatUsageDescription("outbound_contactlists_timezonemappingpreview", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview", )
	outbound_contactlists_timezonemappingpreviewCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlists_timezonemappingpreview"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlists_timezonemappingpreviewCmd)
}

func Cmdoutbound_contactlists_timezonemappingpreview() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview", utils.FormatPermissions([]string{ "outbound:contactList:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TimeZoneMappingPreview"
      }
    }
  }
}`)
	outbound_contactlists_timezonemappingpreviewCmd.AddCommand(getCmd)
	return outbound_contactlists_timezonemappingpreviewCmd
}

var getCmd = &cobra.Command{
	Use:   "get [contactListId]",
	Short: "Preview the result of applying Automatic Time Zone Mapping to a contact list",
	Long:  "Preview the result of applying Automatic Time Zone Mapping to a contact list",
	Args:  utils.DetermineArgs([]string{ "contactListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview"
		contactListId, args := args[0], args[1:]
		path = strings.Replace(path, "{contactListId}", fmt.Sprintf("%v", contactListId), -1)

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
