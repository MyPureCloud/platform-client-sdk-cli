package members
import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

var membersCmd = &cobra.Command{
	Use:   utils.FormatUsageDescription("members"),
	Short: utils.FormatUsageDescription("Manages Genesys Cloud members"),
	Long:  utils.FormatUsageDescription(`Manages Genesys Cloud members`),
}
var CommandService services.CommandService

func init() {
	CommandService = services.NewCommandService(membersCmd)
}

func Cmdmembers() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(getCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(getCmd.Flags(), "string", "sortOrder", "ASC", "Ascending or descending sort order Valid values: ascending, descending")
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: routingStatus, presence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, profileSkills, certifications, locations, groups, skills, languages, languagePreference, employerInfo, biography")
	getCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/groups/{groupId}/members", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET")
	membersCmd.AddCommand(getCmd)
	
	utils.AddFlag(removeCmd.Flags(), "string", "ids", "", "Comma separated list of userIds to remove - REQUIRED")
	removeCmd.SetHelpTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", removeCmd.UsageTemplate(), "DELETE", "/api/v2/groups/{groupId}/members", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(removeCmd.Flags(), "DELETE")
	membersCmd.AddCommand(removeCmd)
	
	return membersCmd
}

var getCmd = &cobra.Command{
	Use:   "get [groupId]",
	Short: "Get group members, includes individuals, owners, and dynamically included people",
	Long:  `Get group members, includes individuals, owners, and dynamically included people`,
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/groups/{groupId}/members"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
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

		retryFunc := CommandService.DetermineAction("GET", "get", urlString, "/api/v2/groups/{groupId}/members")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
var removeCmd = &cobra.Command{
	Use:   "remove [groupId]",
	Short: "Remove members",
	Long:  `Remove members`,
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/groups/{groupId}/members"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

		ids := utils.GetFlag(cmd.Flags(), "string", "ids")
		if ids != "" {
			queryParams["ids"] = ids
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", "remove", urlString, "/api/v2/groups/{groupId}/members")
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			MaxRetriesBeforeQuitting: 3,
			MaxRetryTimeSec: 10,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
