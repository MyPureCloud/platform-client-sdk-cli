package users_me

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
	Description = utils.FormatUsageDescription("users_me", "SWAGGER_OVERRIDE_/api/v2/users/me", )
	users_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_meCmd)
}

func Cmdusers_me() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand. Valid values: routingStatus, presence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, profileSkills, certifications, locations, groups, skills, languages, languagePreference, employerInfo, biography, date, geolocationsettings, organization, presencedefinitions, locationdefinitions, orgauthorization, orgproducts, favorites, superiors, directreports, adjacents, routingskills, routinglanguages, fieldconfigs, token, trustors, logCapture")
	utils.AddFlag(getCmd.Flags(), "string", "integrationPresenceSource", "", "Get your presence for a given integration. This parameter will only be used when presence is provided as an expand. Valid values: MicrosoftTeams, ZoomPhone, RingCentral")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/users/me", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Users", "/api/v2/users/me")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UserMe&quot;
  }
}`)
	users_meCmd.AddCommand(getCmd)
	
	return users_meCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current user details.",
	Long:  "Get current user details.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/users/me"

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
