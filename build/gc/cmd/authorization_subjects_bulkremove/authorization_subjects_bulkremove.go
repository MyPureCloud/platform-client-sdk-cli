package authorization_subjects_bulkremove

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
	Description = utils.FormatUsageDescription("authorization_subjects_bulkremove", "SWAGGER_OVERRIDE_/api/v2/authorization/subjects/{subjectId}/bulkremove", )
	authorization_subjects_bulkremoveCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_subjects_bulkremove"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_subjects_bulkremoveCmd)
}

func Cmdauthorization_subjects_bulkremove() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/authorization/subjects/{subjectId}/bulkremove", utils.FormatPermissions([]string{ "authorization:grant:delete",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Pairs of role and division IDs&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/RoleDivisionGrants&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Bulk Grants Deleted&quot;
}`)
	authorization_subjects_bulkremoveCmd.AddCommand(createCmd)
	
	return authorization_subjects_bulkremoveCmd
}

var createCmd = &cobra.Command{
	Use:   "create [subjectId]",
	Short: "Bulk-remove grants from a subject.",
	Long:  "Bulk-remove grants from a subject.",
	Args:  utils.DetermineArgs([]string{ "subjectId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/subjects/{subjectId}/bulkremove"
		subjectId, args := args[0], args[1:]
		path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)

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
