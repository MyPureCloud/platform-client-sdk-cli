package conversations_messages_inbound_open

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
	Description = utils.FormatUsageDescription("conversations_messages_inbound_open", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/inbound/open", )
	conversations_messages_inbound_openCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_inbound_open"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_inbound_openCmd)
}

func Cmdconversations_messages_inbound_open() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/messages/inbound/open", utils.FormatPermissions([]string{ "conversation:message:receive",  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/messages/inbound/open")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "NormalizedMessage",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/OpenNormalizedMessage"
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
        "$ref" : "#/components/schemas/OpenNormalizedMessage"
      }
    }
  }
}`)
	conversations_messages_inbound_openCmd.AddCommand(createCmd)
	return conversations_messages_inbound_openCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Send an inbound Open Message",
	Long:  "Send an inbound Open Message",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Opennormalizedmessage{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messages/inbound/open"

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
