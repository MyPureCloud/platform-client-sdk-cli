package scripts_export

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
	Description = utils.FormatUsageDescription("scripts_export", "SWAGGER_OVERRIDE_/api/v2/scripts/{scriptId}/export", )
	scripts_exportCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("scripts_export"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(scripts_exportCmd)
}

func Cmdscripts_export() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/scripts/{scriptId}/export", utils.FormatPermissions([]string{ "scripter:script:view",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : false,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExportScriptRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ExportScriptResponse&quot;
  }
}`)
	scripts_exportCmd.AddCommand(createCmd)
	
	return scripts_exportCmd
}

var createCmd = &cobra.Command{
	Use:   "create [scriptId]",
	Short: "Export a script via download service.",
	Long:  "Export a script via download service.",
	Args:  utils.DetermineArgs([]string{ "scriptId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/scripts/{scriptId}/export"
		scriptId, args := args[0], args[1:]
		path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)

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
