package voicemail_messages_media

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
	Description = utils.FormatUsageDescription("voicemail_messages_media", "SWAGGER_OVERRIDE_/api/v2/voicemail/messages/{messageId}/media", )
	voicemail_messages_mediaCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail_messages_media"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemail_messages_mediaCmd)
}

func Cmdvoicemail_messages_media() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "formatId", "WEBM", "The desired media format. Valid values: WAV, WEBM, WAV_ULAW, OGG_VORBIS, OGG_OPUS, MP3, NONE")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/voicemail/messages/{messageId}/media", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/VoicemailMediaInfo&quot;
  }
}`)
	voicemail_messages_mediaCmd.AddCommand(getCmd)
	
	return voicemail_messages_mediaCmd
}

var getCmd = &cobra.Command{
	Use:   "get [messageId]",
	Short: "Get media playback URI for this voicemail message",
	Long:  "Get media playback URI for this voicemail message",
	Args:  utils.DetermineArgs([]string{ "messageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/voicemail/messages/{messageId}/media"
		messageId, args := args[0], args[1:]
		path = strings.Replace(path, "{messageId}", fmt.Sprintf("%v", messageId), -1)

		formatId := utils.GetFlag(cmd.Flags(), "string", "formatId")
		if formatId != "" {
			queryParams["formatId"] = formatId
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
