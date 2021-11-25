package integrations_actions_schemas

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
	Description = utils.FormatUsageDescription("integrations_actions_schemas", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/schemas", )
	integrations_actions_schemasCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_actions_schemas"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_actions_schemasCmd)
}

func Cmdintegrations_actions_schemas() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/actions/{actionId}/schemas/{fileName}", utils.FormatPermissions([]string{ "integrations:action:view", "bridge:actions:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/actions/{actionId}/schemas/{fileName}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/JsonSchemaDocument"
  }
}`)
	integrations_actions_schemasCmd.AddCommand(getCmd)
	
	return integrations_actions_schemasCmd
}

var getCmd = &cobra.Command{
	Use:   "get [actionId] [fileName]",
	Short: "Retrieve schema for an action based on filename.",
	Long:  "Retrieve schema for an action based on filename.",
	Args:  utils.DetermineArgs([]string{ "actionId", "fileName", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/schemas/{fileName}"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
		fileName, args := args[0], args[1:]
		path = strings.Replace(path, "{fileName}", fmt.Sprintf("%v", fileName), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
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
