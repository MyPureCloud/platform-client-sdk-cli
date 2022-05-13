package groups_profile

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
	Description = utils.FormatUsageDescription("groups_profile", "SWAGGER_OVERRIDE_/api/v2/groups/{groupId}/profile", )
	groups_profileCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("groups_profile"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(groups_profileCmd)
}

func Cmdgroups_profile() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "fields", "", "Comma separated fields to return.  Allowable values can be found by querying /api/v2/fieldconfig?type=group and using the key for the elements returned by the fieldList")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/groups/{groupId}/profile", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Groups", "/api/v2/groups/{groupId}/profile")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/GroupProfile"
      }
    }
  }
}`)
	groups_profileCmd.AddCommand(getCmd)
	return groups_profileCmd
}

var getCmd = &cobra.Command{
	Use:   "get [groupId]",
	Short: "Get group profile",
	Long:  "Get group profile",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/groups/{groupId}/profile"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

		fields := utils.GetFlag(cmd.Flags(), "string", "fields")
		if fields != "" {
			queryParams["fields"] = fields
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
