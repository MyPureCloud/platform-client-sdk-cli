package guides_sessions_turns

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
	Description = utils.FormatUsageDescription("guides_sessions_turns", "SWAGGER_OVERRIDE_/api/v2/guides/{guideId}/sessions/{guideSessionId}/turns", )
	guides_sessions_turnsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("guides_sessions_turns"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(guides_sessions_turnsCmd)
}

func Cmdguides_sessions_turns() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/guides/{guideId}/sessions/{guideSessionId}/turns", utils.FormatPermissions([]string{ "aiStudio:guideSessionTurn:add",  }), utils.GenerateDevCentreLink("POST", "AI Studio", "/api/v2/guides/{guideId}/sessions/{guideSessionId}/turns")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/GuideSessionTurnRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Successful Response",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/GuideSessionTurnResponse"
      }
    }
  }
}`)
	guides_sessions_turnsCmd.AddCommand(createCmd)
	return guides_sessions_turnsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [guideId] [guideSessionId]",
	Short: "Add a turn to a guide session.",
	Long:  "Add a turn to a guide session.",
	Args:  utils.DetermineArgs([]string{ "guideId", "guideSessionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Guidesessionturnrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/guides/{guideId}/sessions/{guideSessionId}/turns"
		guideId, args := args[0], args[1:]
		path = strings.Replace(path, "{guideId}", fmt.Sprintf("%v", guideId), -1)
		guideSessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{guideSessionId}", fmt.Sprintf("%v", guideSessionId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
