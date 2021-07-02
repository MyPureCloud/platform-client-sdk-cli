package speechandtextanalytics_conversations_communications_transcripturl

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
	Description = utils.FormatUsageDescription("speechandtextanalytics_conversations_communications_transcripturl", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl", )
	speechandtextanalytics_conversations_communications_transcripturlCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_conversations_communications_transcripturl"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_conversations_communications_transcripturlCmd)
}

func Cmdspeechandtextanalytics_conversations_communications_transcripturl() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl", utils.FormatPermissions([]string{ "recording:recording:view",  }), utils.GenerateDevCentreLink("GET", "Speech &amp; Text Analytics", "/api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/TranscriptUrl&quot;
  }
}`)
	speechandtextanalytics_conversations_communications_transcripturlCmd.AddCommand(getCmd)
	
	return speechandtextanalytics_conversations_communications_transcripturlCmd
}

var getCmd = &cobra.Command{
	Use:   "get [conversationId] [communicationId]",
	Short: "Get the pre-signed S3 URL for the transcript of a specific communication of a conversation",
	Long:  "Get the pre-signed S3 URL for the transcript of a specific communication of a conversation",
	Args:  utils.DetermineArgs([]string{ "conversationId", "communicationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		communicationId, args := args[0], args[1:]
		path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)

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
