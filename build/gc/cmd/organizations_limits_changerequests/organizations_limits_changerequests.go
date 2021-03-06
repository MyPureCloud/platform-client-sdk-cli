package organizations_limits_changerequests

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
	Description = utils.FormatUsageDescription("organizations_limits_changerequests", "SWAGGER_OVERRIDE_/api/v2/organizations/limits/changerequests", "SWAGGER_OVERRIDE_/api/v2/organizations/limits/changerequests", )
	organizations_limits_changerequestsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("organizations_limits_changerequests"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(organizations_limits_changerequestsCmd)
}

func Cmdorganizations_limits_changerequests() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/organizations/limits/changerequests/{requestId}", utils.FormatPermissions([]string{ "limits:organization:view",  }), utils.GenerateDevCentreLink("GET", "Organization", "/api/v2/organizations/limits/changerequests/{requestId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/LimitChangeRequestDetails&quot;
  }
}`)
	organizations_limits_changerequestsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "after", "", "Timestamp indicating the date to begin after when searching for requests.")
	utils.AddFlag(listCmd.Flags(), "int", "before", "", "Timestamp indicating the date to end before when searching for requests.")
	utils.AddFlag(listCmd.Flags(), "string", "status", "", "Status of the request to be filtered by Valid values: Open, Approved, ImplementingChange, ChangeImplemented, Rejected, Rollback, ImplementingRollback, RollbackImplemented")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page Size")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: statusHistory")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/organizations/limits/changerequests", utils.FormatPermissions([]string{ "limits:organization:view",  }), utils.GenerateDevCentreLink("GET", "Organization", "/api/v2/organizations/limits/changerequests")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	organizations_limits_changerequestsCmd.AddCommand(listCmd)
	
	return organizations_limits_changerequestsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [requestId]",
	Short: "Get a limit change request",
	Long:  "Get a limit change request",
	Args:  utils.DetermineArgs([]string{ "requestId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/organizations/limits/changerequests/{requestId}"
		requestId, args := args[0], args[1:]
		path = strings.Replace(path, "{requestId}", fmt.Sprintf("%v", requestId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the available limit change requests",
	Long:  "Get the available limit change requests",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/organizations/limits/changerequests"

		after := utils.GetFlag(cmd.Flags(), "int", "after")
		if after != "" {
			queryParams["after"] = after
		}
		before := utils.GetFlag(cmd.Flags(), "int", "before")
		if before != "" {
			queryParams["before"] = before
		}
		status := utils.GetFlag(cmd.Flags(), "string", "status")
		if status != "" {
			queryParams["status"] = status
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
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
