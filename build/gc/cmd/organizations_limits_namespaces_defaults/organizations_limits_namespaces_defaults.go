package organizations_limits_namespaces_defaults

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
	Description = utils.FormatUsageDescription("organizations_limits_namespaces_defaults", "SWAGGER_OVERRIDE_/api/v2/organizations/limits/namespaces/{namespaceName}/defaults", )
	organizations_limits_namespaces_defaultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations_limits_namespaces_defaults"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizations_limits_namespaces_defaultsCmd)
}

func Cmdorganizations_limits_namespaces_defaults() *cobra.Command { 
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/organizations/limits/namespaces/{namespaceName}/defaults", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Organization", "/api/v2/organizations/limits/namespaces/{namespaceName}/defaults")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LimitsEntityListing&quot;
  }
}`)
	organizations_limits_namespaces_defaultsCmd.AddCommand(listCmd)
	
	return organizations_limits_namespaces_defaultsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [namespaceName]",
	Short: "Get the default limits in a namespace for an organization",
	Long:  "Get the default limits in a namespace for an organization",
	Args:  utils.DetermineArgs([]string{ "namespaceName", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/organizations/limits/namespaces/{namespaceName}/defaults"
		namespaceName, args := args[0], args[1:]
		path = strings.Replace(path, "{namespaceName}", fmt.Sprintf("%v", namespaceName), -1)

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
