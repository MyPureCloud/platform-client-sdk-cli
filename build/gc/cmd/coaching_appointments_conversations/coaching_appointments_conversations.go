package coaching_appointments_conversations

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
	Description = utils.FormatUsageDescription("coaching_appointments_conversations", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/{appointmentId}/conversations", )
	coaching_appointments_conversationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_appointments_conversations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_appointments_conversationsCmd)
}

func Cmdcoaching_appointments_conversations() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/coaching/appointments/{appointmentId}/conversations", utils.FormatPermissions([]string{ "coaching:appointment:edit", "coaching:appointmentConversation:add",  }), utils.GenerateDevCentreLink("POST", "Coaching", "/api/v2/coaching/appointments/{appointmentId}/conversations")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/AddConversationRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Conversation Added",
  "schema" : {
    "$ref" : "#/definitions/AddConversationResponse"
  }
}`)
	coaching_appointments_conversationsCmd.AddCommand(createCmd)
	
	return coaching_appointments_conversationsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [appointmentId]",
	Short: "Add a conversation to an appointment",
	Long:  "Add a conversation to an appointment",
	Args:  utils.DetermineArgs([]string{ "appointmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Addconversationrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/appointments/{appointmentId}/conversations"
		appointmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{appointmentId}", fmt.Sprintf("%v", appointmentId), -1)


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
