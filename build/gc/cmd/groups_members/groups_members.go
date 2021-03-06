package groups_members

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
	Description = utils.FormatUsageDescription("groups_members", "SWAGGER_OVERRIDE_/api/v2/groups/{groupId}/members", "SWAGGER_OVERRIDE_/api/v2/groups/{groupId}/members", "SWAGGER_OVERRIDE_/api/v2/groups/{groupId}/members", )
	groups_membersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("groups_members"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(groups_membersCmd)
}

func Cmdgroups_members() *cobra.Command { 
	addCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", addCmd.UsageTemplate(), "POST", "/api/v2/groups/{groupId}/members", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Groups", "/api/v2/groups/{groupId}/members")))
	utils.AddFileFlagIfUpsert(addCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Add members&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/GroupMembersUpdate&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(addCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Success, group membership was updated&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Empty&quot;
  }
}`)
	groups_membersCmd.AddCommand(addCmd)
	
	utils.AddFlag(deleteCmd.Flags(), "string", "ids", "", "Comma separated list of userIds to remove - REQUIRED")
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/groups/{groupId}/members", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Groups", "/api/v2/groups/{groupId}/members")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	deleteCmd.MarkFlagRequired("ids")
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Success, group membership was updated&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Empty&quot;
  }
}`)
	groups_membersCmd.AddCommand(deleteCmd)
	
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ASC", "Ascending or descending sort order Valid values: ascending, descending")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: routingStatus, presence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, profileSkills, certifications, locations, groups, skills, languages, languagePreference, employerInfo, biography")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/groups/{groupId}/members", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Groups", "/api/v2/groups/{groupId}/members")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	groups_membersCmd.AddCommand(listCmd)
	
	return groups_membersCmd
}

var addCmd = &cobra.Command{
	Use:   "add [groupId]",
	Short: "Add members",
	Long:  "Add members",
	Args:  utils.DetermineArgs([]string{ "groupId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/groups/{groupId}/members"
		groupId, args := args[0], args[1:]
		path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)

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
var deleteCmd = &cobra.Command{
	Use:   "delete [groupId]",
	Short: "Remove members",
	Long:  "Remove members",
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

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
	Use:   "list [groupId]",
	Short: "Get group members, includes individuals, owners, and dynamically included people",
	Long:  "Get group members, includes individuals, owners, and dynamically included people",
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
