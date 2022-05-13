package integrations_types_configschemas

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
	Description = utils.FormatUsageDescription("integrations_types_configschemas", "SWAGGER_OVERRIDE_/api/v2/integrations/types/{typeId}/configschemas", )
	integrations_types_configschemasCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_types_configschemas"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_types_configschemasCmd)
}

func Cmdintegrations_types_configschemas() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/integrations/types/{typeId}/configschemas/{configType}", utils.FormatPermissions([]string{ "integrations:integration:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/types/{typeId}/configschemas/{configType}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/JsonSchemaDocument"
      }
    }
  }
}`)
	integrations_types_configschemasCmd.AddCommand(getCmd)
	return integrations_types_configschemasCmd
}

var getCmd = &cobra.Command{
	Use:   "get [typeId] [configType]",
	Short: "Get properties config schema for an integration type.",
	Long:  "Get properties config schema for an integration type.",
	Args:  utils.DetermineArgs([]string{ "typeId", "configType", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/types/{typeId}/configschemas/{configType}"
		typeId, args := args[0], args[1:]
		path = strings.Replace(path, "{typeId}", fmt.Sprintf("%v", typeId), -1)
		configType, args := args[0], args[1:]
		path = strings.Replace(path, "{configType}", fmt.Sprintf("%v", configType), -1)

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
