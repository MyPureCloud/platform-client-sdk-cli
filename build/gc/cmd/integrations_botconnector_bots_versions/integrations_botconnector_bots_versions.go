package integrations_botconnector_bots_versions

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
	Description = utils.FormatUsageDescription("integrations_botconnector_bots_versions", "SWAGGER_OVERRIDE_/api/v2/integrations/botconnector/{integrationId}/bots/{botId}/versions", )
	integrations_botconnector_bots_versionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_botconnector_bots_versions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_botconnector_bots_versionsCmd)
}

func Cmdintegrations_botconnector_bots_versions() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/integrations/botconnector/{integrationId}/bots/{botId}/versions", utils.FormatPermissions([]string{ "integration:botconnector:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/botconnector/{integrationId}/bots/{botId}/versions")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	integrations_botconnector_bots_versionsCmd.AddCommand(listCmd)
	
	return integrations_botconnector_bots_versionsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [integrationId] [botId]",
	Short: "Get a list of bot versions for a bot",
	Long:  "Get a list of bot versions for a bot",
	Args:  utils.DetermineArgs([]string{ "integrationId", "botId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/botconnector/{integrationId}/bots/{botId}/versions"
		integrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)
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
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "list"
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
