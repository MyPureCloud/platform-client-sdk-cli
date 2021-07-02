package conversations_recordingmetadata

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
	Description = utils.FormatUsageDescription("conversations_recordingmetadata", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/recordingmetadata", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/recordingmetadata", )
	conversations_recordingmetadataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_recordingmetadata"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_recordingmetadataCmd)
}

func Cmdconversations_recordingmetadata() *cobra.Command { 
	conversationmetadataCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", conversationmetadataCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/recordingmetadata", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/conversations/{conversationId}/recordingmetadata")))
	utils.AddFileFlagIfUpsert(conversationmetadataCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(conversationmetadataCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/RecordingMetadata&quot;
    }
  }
}`)
	conversations_recordingmetadataCmd.AddCommand(conversationmetadataCmd)
	
	recordingmetadataCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", recordingmetadataCmd.UsageTemplate(), "GET", "/api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}", utils.FormatPermissions([]string{ "recording:recording:view",  }), utils.GenerateDevCentreLink("GET", "Recording", "/api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}")))
	utils.AddFileFlagIfUpsert(recordingmetadataCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(recordingmetadataCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/RecordingMetadata&quot;
  }
}`)
	conversations_recordingmetadataCmd.AddCommand(recordingmetadataCmd)
	
	return conversations_recordingmetadataCmd
}

var conversationmetadataCmd = &cobra.Command{
	Use:   "conversationmetadata [conversationId]",
	Short: "Get recording metadata for a conversation. Does not return playable media.",
	Long:  "Get recording metadata for a conversation. Does not return playable media.",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/recordingmetadata"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

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
var recordingmetadataCmd = &cobra.Command{
	Use:   "recordingmetadata [conversationId] [recordingId]",
	Short: "Get metadata for a specific recording. Does not return playable media.",
	Long:  "Get metadata for a specific recording. Does not return playable media.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "recordingId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		recordingId, args := args[0], args[1:]
		path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)

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
