package license_toggles

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
	Description = utils.FormatUsageDescription("license_toggles", "SWAGGER_OVERRIDE_/api/v2/license/toggles", "SWAGGER_OVERRIDE_/api/v2/license/toggles", )
	license_togglesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("license_toggles"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(license_togglesCmd)
}

func Cmdlicense_toggles() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/license/toggles/{featureName}", utils.FormatPermissions([]string{ "authorization:grant:add",  }), utils.GenerateDevCentreLink("POST", "License", "/api/v2/license/toggles/{featureName}")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LicenseOrgToggle&quot;
  }
}`)
	license_togglesCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/license/toggles/{featureName}", utils.FormatPermissions([]string{ "authorization:grant:add",  }), utils.GenerateDevCentreLink("GET", "License", "/api/v2/license/toggles/{featureName}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LicenseOrgToggle&quot;
  }
}`)
	license_togglesCmd.AddCommand(getCmd)
	
	return license_togglesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [featureName]",
	Short: "Switch PureCloud license feature toggle value.",
	Long:  "Switch PureCloud license feature toggle value.",
	Args:  utils.DetermineArgs([]string{ "featureName", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/license/toggles/{featureName}"
		featureName, args := args[0], args[1:]
		path = strings.Replace(path, "{featureName}", fmt.Sprintf("%v", featureName), -1)

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
var getCmd = &cobra.Command{
	Use:   "get [featureName]",
	Short: "Get PureCloud license feature toggle value.",
	Long:  "Get PureCloud license feature toggle value.",
	Args:  utils.DetermineArgs([]string{ "featureName", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/license/toggles/{featureName}"
		featureName, args := args[0], args[1:]
		path = strings.Replace(path, "{featureName}", fmt.Sprintf("%v", featureName), -1)

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
