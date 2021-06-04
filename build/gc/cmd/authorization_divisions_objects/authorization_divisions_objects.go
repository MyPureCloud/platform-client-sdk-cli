package authorization_divisions_objects

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
	Description = utils.FormatUsageDescription("authorization_divisions_objects", "SWAGGER_OVERRIDE_/api/v2/authorization/divisions/{divisionId}/objects", )
	authorization_divisions_objectsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("authorization_divisions_objects"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(authorization_divisions_objectsCmd)
}

func Cmdauthorization_divisions_objects() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/authorization/divisions/{divisionId}/objects/{objectType}", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Object Id List&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;type&quot; : &quot;string&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;The divisions were updated successfully&quot;
}`)
	authorization_divisions_objectsCmd.AddCommand(createCmd)
	
	return authorization_divisions_objectsCmd
}

var createCmd = &cobra.Command{
	Use:   "create [divisionId] [objectType]",
	Short: "Assign a list of objects to a division",
	Long:  "Assign a list of objects to a division",
	Args:  utils.DetermineArgs([]string{ "divisionId", "objectType", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/authorization/divisions/{divisionId}/objects/{objectType}"
		divisionId, args := args[0], args[1:]
		path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
		objectType, args := args[0], args[1:]
		path = strings.Replace(path, "{objectType}", fmt.Sprintf("%v", objectType), -1)

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
