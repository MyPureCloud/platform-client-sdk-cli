package integrations_speech_lex_bot_aliases

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
	Description = utils.FormatUsageDescription("integrations_speech_lex_bot_aliases", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/lex/bot/{botId}/aliases", )
	integrations_speech_lex_bot_aliasesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_lex_bot_aliases"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_lex_bot_aliasesCmd)
}

func Cmdintegrations_speech_lex_bot_aliases() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "string", "status", "", "Filter on alias status Valid values: READY, FAILED, BUILDING, NOT_BUILT")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter on alias name")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/integrations/speech/lex/bot/{botId}/aliases", utils.FormatPermissions([]string{ "integrations:integration:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/speech/lex/bot/{botId}/aliases")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	integrations_speech_lex_bot_aliasesCmd.AddCommand(listCmd)
	
	return integrations_speech_lex_bot_aliasesCmd
}

var listCmd = &cobra.Command{
	Use:   "list [botId]",
	Short: "Get a list of aliases for a bot in the customer`s AWS accounts",
	Long:  "Get a list of aliases for a bot in the customer`s AWS accounts",
	Args:  utils.DetermineArgs([]string{ "botId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/speech/lex/bot/{botId}/aliases"
		botId, args := args[0], args[1:]
		path = strings.Replace(path, "{botId}", fmt.Sprintf("%v", botId), -1)


		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		status := utils.GetFlag(cmd.Flags(), "string", "status")
		if status != "" {
			queryParams["status"] = status
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
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
