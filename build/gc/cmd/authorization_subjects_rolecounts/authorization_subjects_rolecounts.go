package authorization_subjects_rolecounts

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
	Description = utils.FormatUsageDescription("authorization_subjects_rolecounts", "SWAGGER_OVERRIDE_/api/v2/authorization/subjects/rolecounts", )
	authorization_subjects_rolecountsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_subjects_rolecounts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_subjects_rolecountsCmd)
}

func Cmdauthorization_subjects_rolecounts() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "id", "", "id")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/subjects/rolecounts", utils.FormatPermissions([]string{ "authorization:grant:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
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
	authorization_subjects_rolecountsCmd.AddCommand(getCmd)
	
	return authorization_subjects_rolecountsCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the count of roles granted to a list of subjects",
	Long:  "Get the count of roles granted to a list of subjects",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/subjects/rolecounts"

		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
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
