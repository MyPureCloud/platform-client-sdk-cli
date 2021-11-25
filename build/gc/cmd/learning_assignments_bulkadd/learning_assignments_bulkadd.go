package learning_assignments_bulkadd

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
	Description = utils.FormatUsageDescription("learning_assignments_bulkadd", "SWAGGER_OVERRIDE_/api/v2/learning/assignments/bulkadd", )
	learning_assignments_bulkaddCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_assignments_bulkadd"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_assignments_bulkaddCmd)
}

func Cmdlearning_assignments_bulkadd() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/learning/assignments/bulkadd", utils.FormatPermissions([]string{ "learning:assignment:add",  }), utils.GenerateDevCentreLink("POST", "Learning", "/api/v2/learning/assignments/bulkadd")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "The learning assignments to be created",
  "required" : false,
  "schema" : {
    "type" : "array",
    "items" : {
      "$ref" : "#/definitions/LearningAssignmentItem"
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/LearningAssignmentBulkAddResponse"
  }
}`)
	learning_assignments_bulkaddCmd.AddCommand(createCmd)
	
	return learning_assignments_bulkaddCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add multiple learning assignments",
	Long:  "Add multiple learning assignments",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Learningassignmentitem{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/learning/assignments/bulkadd"

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
