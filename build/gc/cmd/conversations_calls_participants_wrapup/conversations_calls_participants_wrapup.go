package conversations_calls_participants_wrapup

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("conversations_calls_participants_wrapup", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup", )
	conversations_calls_participants_wrapupCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_calls_participants_wrapup"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_calls_participants_wrapupCmd)
}

func Cmdconversations_calls_participants_wrapup() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "provisional", "false", "Indicates if the wrap-up code is provisional.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AssignedWrapupCode&quot;
  }
}`)
	conversations_calls_participants_wrapupCmd.AddCommand(getCmd)
	
	return conversations_calls_participants_wrapupCmd
}

var getCmd = &cobra.Command{
	Use:   "get [conversationId] [participantId]",
	Short: "Get the wrap-up for this conversation participant. ",
	Long:  "Get the wrap-up for this conversation participant. ",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)

		provisional := utils.GetFlag(cmd.Flags(), "bool", "provisional")
		if provisional != "" {
			queryParams["provisional"] = provisional
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
