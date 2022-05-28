package coaching_scheduleslots_query

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
	Description = utils.FormatUsageDescription("coaching_scheduleslots_query", "SWAGGER_OVERRIDE_/api/v2/coaching/scheduleslots/query", )
	coaching_scheduleslots_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_scheduleslots_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_scheduleslots_queryCmd)
}

func Cmdcoaching_scheduleslots_query() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/coaching/scheduleslots/query", utils.FormatPermissions([]string{ "coaching:scheduleSlot:view",  }), utils.GenerateDevCentreLink("POST", "Coaching", "/api/v2/coaching/scheduleslots/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The slot search request",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingSlotsRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Coaching slots retrieved",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CoachingSlotsResponse"
      }
    }
  }
}`)
	coaching_scheduleslots_queryCmd.AddCommand(createCmd)
	return coaching_scheduleslots_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Get list of possible slots where a coaching appointment can be scheduled.",
	Long:  "Get list of possible slots where a coaching appointment can be scheduled.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Coachingslotsrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/coaching/scheduleslots/query"

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
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
