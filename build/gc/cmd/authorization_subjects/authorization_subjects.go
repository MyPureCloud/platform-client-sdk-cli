package authorization_subjects

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
	Description = utils.FormatUsageDescription("authorization_subjects", "SWAGGER_OVERRIDE_/api/v2/authorization/subjects", )
	authorization_subjectsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_subjects"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_subjectsCmd)
}

func Cmdauthorization_subjects() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/authorization/subjects/{subjectId}", utils.FormatPermissions([]string{ "authorization:grant:view",  }), utils.GenerateDevCentreLink("GET", "Authorization", "/api/v2/authorization/subjects/{subjectId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/AuthzSubject"
  }
}`)
	authorization_subjectsCmd.AddCommand(getCmd)
	
	return authorization_subjectsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [subjectId]",
	Short: "Returns a listing of roles and permissions for a user.",
	Long:  "Returns a listing of roles and permissions for a user.",
	Args:  utils.DetermineArgs([]string{ "subjectId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/authorization/subjects/{subjectId}"
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
