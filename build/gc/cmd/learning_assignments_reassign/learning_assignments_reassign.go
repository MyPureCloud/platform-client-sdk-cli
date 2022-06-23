package learning_assignments_reassign

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
	Description = utils.FormatUsageDescription("learning_assignments_reassign", "SWAGGER_OVERRIDE_/api/v2/learning/assignments/{assignmentId}/reassign", )
	learning_assignments_reassignCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_assignments_reassign"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_assignments_reassignCmd)
}

func Cmdlearning_assignments_reassign() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/learning/assignments/{assignmentId}/reassign", utils.FormatPermissions([]string{ "learning:assignment:add",  }), utils.GenerateDevCentreLink("POST", "Learning", "/api/v2/learning/assignments/{assignmentId}/reassign")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Reassigned version of assignment which can be taken again",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/LearningAssignment"
      }
    }
  }
}`)
	learning_assignments_reassignCmd.AddCommand(createCmd)
	return learning_assignments_reassignCmd
}

var createCmd = &cobra.Command{
	Use:   "create [assignmentId]",
	Short: "Reassign Learning Assignment",
	Long:  "Reassign Learning Assignment",
	Args:  utils.DetermineArgs([]string{ "assignmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments/{assignmentId}/reassign"
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
