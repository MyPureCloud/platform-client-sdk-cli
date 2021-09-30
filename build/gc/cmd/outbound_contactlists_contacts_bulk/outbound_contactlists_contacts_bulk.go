package outbound_contactlists_contacts_bulk

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
	Description = utils.FormatUsageDescription("outbound_contactlists_contacts_bulk", "SWAGGER_OVERRIDE_/api/v2/outbound/contactlists/{contactListId}/contacts/bulk", )
	outbound_contactlists_contacts_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_contactlists_contacts_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_contactlists_contacts_bulkCmd)
}

func Cmdoutbound_contactlists_contacts_bulk() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/outbound/contactlists/{contactListId}/contacts/bulk", utils.FormatPermissions([]string{ "outbound:contact:view",  }), utils.GenerateDevCentreLink("POST", "Outbound", "/api/v2/outbound/contactlists/{contactListId}/contacts/bulk")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "ContactIds to get.",
  "required" : true,
  "schema" : {
    "type" : "array",
    "items" : {
      "type" : "string"
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/DialerContact"
    }
  }
}`)
	outbound_contactlists_contacts_bulkCmd.AddCommand(createCmd)
	
	return outbound_contactlists_contacts_bulkCmd
}

var createCmd = &cobra.Command{
	Use:   "create [contactListId]",
	Short: "Get contacts from a contact list.",
	Long:  "Get contacts from a contact list.",
	Args:  utils.DetermineArgs([]string{ "contactListId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/outbound/contactlists/{contactListId}/contacts/bulk"
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
