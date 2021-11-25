package profile_users

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
	Description = utils.FormatUsageDescription("profile_users", "SWAGGER_OVERRIDE_/api/v2/profiles/users", )
	profile_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("profile_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(profile_usersCmd)
}

func Cmdprofile_users() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "[]string", "id", "", "id")
	utils.AddFlag(listCmd.Flags(), "[]string", "jid", "", "jid")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ASC", "Ascending or descending sort order Valid values: ascending, descending")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: routingStatus, presence, conversationSummary, outOfOffice, geolocation, station, authorization")
	utils.AddFlag(listCmd.Flags(), "string", "integrationPresenceSource", "", "Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an expand. Valid values: MicrosoftTeams, ZoomPhone, RingCentral")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/profiles/users", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Users", "/api/v2/profiles/users")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/SWAGGER_OVERRIDE_list"
  }
}`)
	profile_usersCmd.AddCommand(listCmd)
	
	return profile_usersCmd
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a user profile listing",
	Long:  "Get a user profile listing",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/profiles/users"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		id := utils.GetFlag(cmd.Flags(), "[]string", "id")
		if id != "" {
			queryParams["id"] = id
		}
		jid := utils.GetFlag(cmd.Flags(), "[]string", "jid")
		if jid != "" {
			queryParams["jid"] = jid
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		integrationPresenceSource := utils.GetFlag(cmd.Flags(), "string", "integrationPresenceSource")
		if integrationPresenceSource != "" {
			queryParams["integrationPresenceSource"] = integrationPresenceSource
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "list"
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
