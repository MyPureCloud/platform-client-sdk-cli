package learning_assignments_reset

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
	Description = utils.FormatUsageDescription("learning_assignments_reset", "SWAGGER_OVERRIDE_/api/v2/learning/assignments/{assignmentId}/reset", )
	learning_assignments_resetCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_assignments_reset"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_assignments_resetCmd)
}

func Cmdlearning_assignments_reset() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/learning/assignments/{assignmentId}/reset", utils.FormatPermissions([]string{ "learning:assignment:reset",  }), utils.GenerateDevCentreLink("POST", "Learning", "/api/v2/learning/assignments/{assignmentId}/reset")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Reset version of assignment which can be taken again",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/LearningAssignment"
      }
    }
  }
}`)
	learning_assignments_resetCmd.AddCommand(createCmd)
	return learning_assignments_resetCmd
}

var createCmd = &cobra.Command{
	Use:   "create [assignmentId]",
	Short: "Reset Learning Assignment",
	Long:  "Reset Learning Assignment",
	Args:  utils.DetermineArgs([]string{ "assignmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments/{assignmentId}/reset"
		assignmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{assignmentId}", fmt.Sprintf("%v", assignmentId), -1)

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
