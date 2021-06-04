package integrations_speech_tts_engines_voices

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
	Description = utils.FormatUsageDescription("integrations_speech_tts_engines_voices", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/tts/engines/{engineId}/voices", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/tts/engines/{engineId}/voices", )
	integrations_speech_tts_engines_voicesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_tts_engines_voices"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_tts_engines_voicesCmd)
}

func Cmdintegrations_speech_tts_engines_voices() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}", utils.FormatPermissions([]string{ "integrations:integration:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/TtsVoiceEntity&quot;
  }
}`)
	integrations_speech_tts_engines_voicesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/tts/engines/{engineId}/voices", utils.FormatPermissions([]string{ "integrations:integration:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	integrations_speech_tts_engines_voicesCmd.AddCommand(listCmd)
	
	return integrations_speech_tts_engines_voicesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [engineId] [voiceId]",
	Short: "Get details about a specific voice for a TTS engine",
	Long:  "Get details about a specific voice for a TTS engine",
	Args:  utils.DetermineArgs([]string{ "engineId", "voiceId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}"
		engineId, args := args[0], args[1:]
		path = strings.Replace(path, "{engineId}", fmt.Sprintf("%v", engineId), -1)
		voiceId, args := args[0], args[1:]
		path = strings.Replace(path, "{voiceId}", fmt.Sprintf("%v", voiceId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [engineId]",
	Short: "Get a list of voices for a TTS engine",
	Long:  "Get a list of voices for a TTS engine",
	Args:  utils.DetermineArgs([]string{ "engineId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/tts/engines/{engineId}/voices"
		engineId, args := args[0], args[1:]
		path = strings.Replace(path, "{engineId}", fmt.Sprintf("%v", engineId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
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
