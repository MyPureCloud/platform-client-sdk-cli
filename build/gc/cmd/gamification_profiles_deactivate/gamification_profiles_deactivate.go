package gamification_profiles_deactivate

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
	Description = utils.FormatUsageDescription("gamification_profiles_deactivate", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles/{profileId}/deactivate", )
	gamification_profiles_deactivateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_profiles_deactivate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_profiles_deactivateCmd)
}

func Cmdgamification_profiles_deactivate() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/gamification/profiles/{profileId}/deactivate", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("POST", "Gamification", "/api/v2/gamification/profiles/{profileId}/deactivate")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PerformanceProfile"
      }
    }
  }
}`)
	gamification_profiles_deactivateCmd.AddCommand(createCmd)
	return gamification_profiles_deactivateCmd
}

var createCmd = &cobra.Command{
	Use:   "create [profileId]",
	Short: "Deactivate a performance profile",
	Long:  "Deactivate a performance profile",
	Args:  utils.DetermineArgs([]string{ "profileId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/{profileId}/deactivate"
		profileId, args := args[0], args[1:]
		path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "create"
		const httpMethod = "POST"
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
