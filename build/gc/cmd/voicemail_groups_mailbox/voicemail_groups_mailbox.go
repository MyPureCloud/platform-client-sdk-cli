package voicemail_groups_mailbox

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
	Description = utils.FormatUsageDescription("voicemail_groups_mailbox", "SWAGGER_OVERRIDE_/api/v2/voicemail/groups/{groupId}/mailbox", )
	voicemail_groups_mailboxCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail_groups_mailbox"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemail_groups_mailboxCmd)
}

func Cmdvoicemail_groups_mailbox() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/voicemail/groups/{groupId}/mailbox", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Voicemail", "/api/v2/voicemail/groups/{groupId}/mailbox")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/VoicemailMailboxInfo"
  }
}`)
	voicemail_groups_mailboxCmd.AddCommand(getCmd)
	
	return voicemail_groups_mailboxCmd
}

var getCmd = &cobra.Command{
	Use:   "get [groupId]",
	Short: "Get the group`s mailbox information",
	Long:  "Get the group`s mailbox information",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/voicemail/groups/{groupId}/mailbox"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

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
