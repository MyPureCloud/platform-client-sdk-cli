package taskmanagement_workitems_acd_cancel

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
	Description = utils.FormatUsageDescription("taskmanagement_workitems_acd_cancel", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/workitems/{workitemId}/acd/cancel", )
	taskmanagement_workitems_acd_cancelCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_workitems_acd_cancel"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_workitems_acd_cancelCmd)
}

func Cmdtaskmanagement_workitems_acd_cancel() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/taskmanagement/workitems/{workitemId}/acd/cancel", utils.FormatPermissions([]string{ "workitems:workitem:cancelRouting",  }), utils.GenerateDevCentreLink("POST", "Task Management", "/api/v2/taskmanagement/workitems/{workitemId}/acd/cancel")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Workitem"
      }
    }
  }
}`)
	taskmanagement_workitems_acd_cancelCmd.AddCommand(createCmd)
	return taskmanagement_workitems_acd_cancelCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [workitemId]",
	Short: "Cancel the assignment process for a workitem that is currently queued for assignment through ACD.",
	Long:  "Cancel the assignment process for a workitem that is currently queued for assignment through ACD.",
	Args:  utils.DetermineArgs([]string{ "workitemId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/taskmanagement/workitems/{workitemId}/acd/cancel"
		workitemId, args := args[0], args[1:]
		path = strings.Replace(path, "{workitemId}", fmt.Sprintf("%v", workitemId), -1)

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
