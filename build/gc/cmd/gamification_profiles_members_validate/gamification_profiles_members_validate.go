package gamification_profiles_members_validate

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
	Description = utils.FormatUsageDescription("gamification_profiles_members_validate", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles/{performanceProfileId}/members/validate", )
	gamification_profiles_members_validateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_profiles_members_validate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_profiles_members_validateCmd)
}

func Cmdgamification_profiles_members_validate() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/gamification/profiles/{performanceProfileId}/members/validate", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("POST", "Gamification", "/api/v2/gamification/profiles/{performanceProfileId}/members/validate")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "memberAssignments",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ValidateAssignUsers"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AssignmentValidation"
      }
    }
  }
}`)
	gamification_profiles_members_validateCmd.AddCommand(createCmd)
	return gamification_profiles_members_validateCmd
}

var createCmd = &cobra.Command{
	Use:   "create [performanceProfileId]",
	Short: "Validate member assignment",
	Long:  "Validate member assignment",
	Args:  utils.DetermineArgs([]string{ "performanceProfileId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Validateassignusers{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/{performanceProfileId}/members/validate"
		performanceProfileId, args := args[0], args[1:]
		path = strings.Replace(path, "{performanceProfileId}", fmt.Sprintf("%v", performanceProfileId), -1)

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
