package processautomation_triggers_topics_testfile

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
	Description = utils.FormatUsageDescription("processautomation_triggers_topics_testfile", "SWAGGER_OVERRIDE_/api/v2/processautomation/triggers/topics/{topicName}/test", )
	processautomation_triggers_topics_testfileCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("processautomation_triggers_topics_testfile"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(processautomation_triggers_topics_testfileCmd)
}

func Cmdprocessautomation_triggers_topics_testfile() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/processautomation/triggers/topics/{topicName}/test", utils.FormatPermissions([]string{ "processautomation:trigger:test",  }), utils.GenerateDevCentreLink("POST", "Process Automation", "/api/v2/processautomation/triggers/topics/{topicName}/test")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "eventBody",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "string"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TestModeEventResults"
      }
    }
  }
}`)
	processautomation_triggers_topics_testfileCmd.AddCommand(createCmd)
	return processautomation_triggers_topics_testfileCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [topicName]",
	Short: "Test the matching of all organization Triggers on given topic using provided event body",
	Long:  "Test the matching of all organization Triggers on given topic using provided event body",
	Args:  utils.DetermineArgs([]string{ "topicName", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.String{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/processautomation/triggers/topics/{topicName}/test"
		topicName, args := args[0], args[1:]
		path = strings.Replace(path, "{topicName}", fmt.Sprintf("%v", topicName), -1)

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