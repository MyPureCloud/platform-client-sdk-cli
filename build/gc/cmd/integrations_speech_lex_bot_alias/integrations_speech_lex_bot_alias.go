package integrations_speech_lex_bot_alias

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
	Description = utils.FormatUsageDescription("integrations_speech_lex_bot_alias", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/lex/bot/alias", )
	integrations_speech_lex_bot_aliasCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_lex_bot_alias"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_lex_bot_aliasCmd)
}

func Cmdintegrations_speech_lex_bot_alias() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/lex/bot/alias/{aliasId}", utils.FormatPermissions([]string{ "integrations:integration:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LexBotAlias&quot;
  }
}`)
	integrations_speech_lex_bot_aliasCmd.AddCommand(getCmd)
	
	return integrations_speech_lex_bot_aliasCmd
}

var getCmd = &cobra.Command{
	Use:   "get [aliasId]",
	Short: "Get details about a Lex bot alias",
	Long:  "Get details about a Lex bot alias",
	Args:  utils.DetermineArgs([]string{ "aliasId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/lex/bot/alias/{aliasId}"
		aliasId, args := args[0], args[1:]
		path = strings.Replace(path, "{aliasId}", fmt.Sprintf("%v", aliasId), -1)

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
