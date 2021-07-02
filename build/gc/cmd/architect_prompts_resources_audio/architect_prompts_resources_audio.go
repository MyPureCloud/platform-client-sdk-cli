package architect_prompts_resources_audio

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
	Description = utils.FormatUsageDescription("architect_prompts_resources_audio", "SWAGGER_OVERRIDE_/api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio", )
	architect_prompts_resources_audioCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("architect_prompts_resources_audio"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(architect_prompts_resources_audioCmd)
}

func Cmdarchitect_prompts_resources_audio() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio", utils.FormatPermissions([]string{ "architect:userPrompt:edit",  }), utils.GenerateDevCentreLink("DELETE", "Architect", "/api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Audio successfully deleted&quot;
}`)
	architect_prompts_resources_audioCmd.AddCommand(deleteCmd)
	
	return architect_prompts_resources_audioCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [promptId] [languageCode]",
	Short: "Delete specified user prompt resource audio",
	Long:  "Delete specified user prompt resource audio",
	Args:  utils.DetermineArgs([]string{ "promptId", "languageCode", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/architect/prompts/{promptId}/resources/{languageCode}/audio"
		promptId, args := args[0], args[1:]
		path = strings.Replace(path, "{promptId}", fmt.Sprintf("%v", promptId), -1)
		languageCode, args := args[0], args[1:]
		path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
