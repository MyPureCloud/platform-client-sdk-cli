package outbound_contactlists_importstatus

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
	Description = utils.FormatUsageDescription("outbound_contactlists_importstatus", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlists/{contactListId}/importstatus", )
	outbound_contactlists_importstatusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlists_importstatus"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlists_importstatusCmd)
}

func Cmdoutbound_contactlists_importstatus() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/outbound/contactlists/{contactListId}/importstatus", utils.FormatPermissions([]string{ "outbound:contactList:view",  }), utils.GenerateDevCentreLink("GET", "Outbound", "/api/v2/outbound/contactlists/{contactListId}/importstatus")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/ImportStatus"
  }
}`)
	outbound_contactlists_importstatusCmd.AddCommand(getCmd)
	
	return outbound_contactlists_importstatusCmd
}

var getCmd = &cobra.Command{
	Use:   "get [contactListId]",
	Short: "Get dialer contactList import status.",
	Long:  "Get dialer contactList import status.",
	Args:  utils.DetermineArgs([]string{ "contactListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/importstatus"
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
