package languageunderstanding_miners_execute

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
	Description = utils.FormatUsageDescription("languageunderstanding_miners_execute", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/miners/{minerId}/execute", )
	languageunderstanding_miners_executeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_miners_execute"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_miners_executeCmd)
}

func Cmdlanguageunderstanding_miners_execute() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/languageunderstanding/miners/{minerId}/execute", utils.FormatPermissions([]string{ "languageUnderstanding:miner:execute",  }), utils.GenerateDevCentreLink("POST", "Language Understanding", "/api/v2/languageunderstanding/miners/{minerId}/execute")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : false,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/MinerExecuteRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Miner&quot;
  }
}`)
	languageunderstanding_miners_executeCmd.AddCommand(createCmd)
	
	return languageunderstanding_miners_executeCmd
}

var createCmd = &cobra.Command{
	Use:   "create [minerId]",
	Short: "Start the mining process. Specify date range pair with mediaType and queueIds for mining data from Genesys Cloud. Specify only uploadKey for mining through an external file.",
	Long:  "Start the mining process. Specify date range pair with mediaType and queueIds for mining data from Genesys Cloud. Specify only uploadKey for mining through an external file.",
	Args:  utils.DetermineArgs([]string{ "minerId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/miners/{minerId}/execute"
		minerId, args := args[0], args[1:]
		path = strings.Replace(path, "{minerId}", fmt.Sprintf("%v", minerId), -1)

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
