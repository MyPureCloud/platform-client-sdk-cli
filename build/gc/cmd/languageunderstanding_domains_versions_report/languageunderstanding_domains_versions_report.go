package languageunderstanding_domains_versions_report

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
	Description = utils.FormatUsageDescription("languageunderstanding_domains_versions_report", "SWAGGER_OVERRIDE_/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report", )
	languageunderstanding_domains_versions_reportCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languageunderstanding_domains_versions_report"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languageunderstanding_domains_versions_reportCmd)
}

func Cmdlanguageunderstanding_domains_versions_report() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report", utils.FormatPermissions([]string{ "languageUnderstanding:nluDomainVersion:view", "dialog:botVersion:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;Find quality report for NLU Domain Version.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/NluDomainVersionQualityReport&quot;
  }
}`)
	languageunderstanding_domains_versions_reportCmd.AddCommand(getCmd)
	
	return languageunderstanding_domains_versions_reportCmd
}

var getCmd = &cobra.Command{
	Use:   "get [domainId] [domainVersionId]",
	Short: "Retrieved quality report for the specified NLU Domain Version",
	Long:  "Retrieved quality report for the specified NLU Domain Version",
	Args:  utils.DetermineArgs([]string{ "domainId", "domainVersionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report"
		domainId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
		domainVersionId, args := args[0], args[1:]
		path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)

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
