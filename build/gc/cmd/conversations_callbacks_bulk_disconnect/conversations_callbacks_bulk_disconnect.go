package conversations_callbacks_bulk_disconnect

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
	Description = utils.FormatUsageDescription("conversations_callbacks_bulk_disconnect", "SWAGGER_OVERRIDE_/api/v2/conversations/callbacks/bulk/disconnect", )
	conversations_callbacks_bulk_disconnectCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_callbacks_bulk_disconnect"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_callbacks_bulk_disconnectCmd)
}

func Cmdconversations_callbacks_bulk_disconnect() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/callbacks/bulk/disconnect", utils.FormatPermissions([]string{ "conversation:communication:disconnect",  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/callbacks/bulk/disconnect")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "BulkCallbackDisconnectRequest",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/BulkCallbackDisconnectRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Accepted"
}`)
	conversations_callbacks_bulk_disconnectCmd.AddCommand(createCmd)
	
	return conversations_callbacks_bulk_disconnectCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Disconnect multiple scheduled callbacks",
	Long:  "Disconnect multiple scheduled callbacks",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Bulkcallbackdisconnectrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/callbacks/bulk/disconnect"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
