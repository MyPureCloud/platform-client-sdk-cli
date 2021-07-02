package telephony_providers_edges_lines_template

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_lines_template", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/lines/template", )
	telephony_providers_edges_lines_templateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_lines_template"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_lines_templateCmd)
}

func Cmdtelephony_providers_edges_lines_template() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "lineBaseSettingsId", "", "The id of a Line Base Settings object upon which to base this Line - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/lines/template", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("GET", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/lines/template")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("lineBaseSettingsId")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Line&quot;
  }
}`)
	telephony_providers_edges_lines_templateCmd.AddCommand(getCmd)
	
	return telephony_providers_edges_lines_templateCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Line instance template based on a Line Base Settings object. This object can then be modified and saved as a new Line instance",
	Long:  "Get a Line instance template based on a Line Base Settings object. This object can then be modified and saved as a new Line instance",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/lines/template"

		lineBaseSettingsId := utils.GetFlag(cmd.Flags(), "string", "lineBaseSettingsId")
		if lineBaseSettingsId != "" {
			queryParams["lineBaseSettingsId"] = lineBaseSettingsId
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
