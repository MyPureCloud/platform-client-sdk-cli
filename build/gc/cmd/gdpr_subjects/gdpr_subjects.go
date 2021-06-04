package gdpr_subjects

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
	Description = utils.FormatUsageDescription("gdpr_subjects", "SWAGGER_OVERRIDE_/api/v2/gdpr/subjects", )
	gdpr_subjectsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gdpr_subjects"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gdpr_subjectsCmd)
}

func Cmdgdpr_subjects() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "searchType", "", "Search Type - REQUIRED Valid values: NAME, ADDRESS, PHONE, EMAIL, TWITTER")
	utils.AddFlag(listCmd.Flags(), "string", "searchValue", "", "Search Value - REQUIRED")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/gdpr/subjects", utils.FormatPermissions([]string{ "gdpr:subject:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("searchType")
	listCmd.MarkFlagRequired("searchValue")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	gdpr_subjectsCmd.AddCommand(listCmd)
	
	return gdpr_subjectsCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get GDPR subjects",
	Long:  "Get GDPR subjects",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gdpr/subjects"

		searchType := utils.GetFlag(cmd.Flags(), "string", "searchType")
		if searchType != "" {
			queryParams["searchType"] = searchType
		}
		searchValue := utils.GetFlag(cmd.Flags(), "string", "searchValue")
		if searchValue != "" {
			queryParams["searchValue"] = searchValue
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
