package users_routingskills_bulk

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
	Description = utils.FormatUsageDescription("users_routingskills_bulk", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/routingskills/bulk", "SWAGGER_OVERRIDE_/api/v2/users/{userId}/routingskills/bulk", )
	users_routingskills_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_routingskills_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_routingskills_bulkCmd)
}

func Cmdusers_routingskills_bulk() *cobra.Command { 
	bulkaddCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", bulkaddCmd.UsageTemplate(), "PATCH", "/api/v2/users/{userId}/routingskills/bulk", utils.FormatPermissions([]string{ "routing:skill:assign",  })))
	utils.AddFileFlagIfUpsert(bulkaddCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Skill&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/UserRoutingSkillPost&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(bulkaddCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserSkillEntityListing&quot;
  }
}`)
	users_routingskills_bulkCmd.AddCommand(bulkaddCmd)
	
	bulkupdateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", bulkupdateCmd.UsageTemplate(), "PUT", "/api/v2/users/{userId}/routingskills/bulk", utils.FormatPermissions([]string{ "routing:skill:assign",  })))
	utils.AddFileFlagIfUpsert(bulkupdateCmd.Flags(), "PUT", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Skill&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/UserRoutingSkillPost&quot;
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(bulkupdateCmd.Flags(), "PUT", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserSkillEntityListing&quot;
  }
}`)
	users_routingskills_bulkCmd.AddCommand(bulkupdateCmd)
	
	return users_routingskills_bulkCmd
}

var bulkaddCmd = &cobra.Command{
	Use:   "bulkadd [userId]",
	Short: "Bulk add routing skills to user",
	Long:  "Bulk add routing skills to user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/routingskills/bulk"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
var bulkupdateCmd = &cobra.Command{
	Use:   "bulkupdate [userId]",
	Short: "Replace all routing skills assigned to a user",
	Long:  "Replace all routing skills assigned to a user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/{userId}/routingskills/bulk"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
