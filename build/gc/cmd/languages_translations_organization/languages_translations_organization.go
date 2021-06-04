package languages_translations_organization

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
	Description = utils.FormatUsageDescription("languages_translations_organization", "SWAGGER_OVERRIDE_/api/v2/languages/translations/organization", )
	languages_translations_organizationCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("languages_translations_organization"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(languages_translations_organizationCmd)
}

func Cmdlanguages_translations_organization() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "language", "", "The language of the translation to retrieve for the organization - REQUIRED")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/languages/translations/organization", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("language")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;object&quot;,
    &quot;additionalProperties&quot; : {
      &quot;type&quot; : &quot;object&quot;,
      &quot;properties&quot; : { }
    }
  }
}`)
	languages_translations_organizationCmd.AddCommand(getCmd)
	
	return languages_translations_organizationCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get effective translation for an organization by language",
	Long:  "Get effective translation for an organization by language",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/languages/translations/organization"

		language := utils.GetFlag(cmd.Flags(), "string", "language")
		if language != "" {
			queryParams["language"] = language
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
