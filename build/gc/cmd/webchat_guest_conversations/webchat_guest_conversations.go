package webchat_guest_conversations

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
	Description = utils.FormatUsageDescription("webchat_guest_conversations", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations", )
	webchat_guest_conversationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webchat_guest_conversations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webchat_guest_conversationsCmd)
}

func Cmdwebchat_guest_conversations() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/webchat/guest/conversations", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "WebChat", "/api/v2/webchat/guest/conversations")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "CreateConversationRequest",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateWebChatConversationRequest"
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
        "$ref" : "#/components/schemas/CreateWebChatConversationResponse"
      }
    }
  }
}`)
	webchat_guest_conversationsCmd.AddCommand(createCmd)
	return webchat_guest_conversationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an ACD chat conversation from an external customer.",
	Long:  "Create an ACD chat conversation from an external customer.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createwebchatconversationrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations"

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
