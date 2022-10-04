package integrations_actions_testfile

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
	Description = utils.FormatUsageDescription("integrations_actions_testfile", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/test", )
	integrations_actions_testfileCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_actions_testfile"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_actions_testfileCmd)
}

func Cmdintegrations_actions_testfile() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/integrations/actions/{actionId}/test", utils.FormatPermissions([]string{ "integrations:action:execute", "bridge:actions:execute",  }), utils.GenerateDevCentreLink("POST", "Integrations", "/api/v2/integrations/actions/{actionId}/test")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "Map of parameters used for variable substitution.",
  "content" : {
    "application/json" : {
      "schema" : {
        "type" : "object",
        "additionalProperties" : {
          "type" : "object",
          "properties" : { }
        }
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
        "$ref" : "#/components/schemas/TestExecutionResult"
      }
    }
  }
}`)
	integrations_actions_testfileCmd.AddCommand(createCmd)
	return integrations_actions_testfileCmd
}

var createCmd = &cobra.Command{
	Use:   "create [actionId]",
	Short: "Test the execution of an action. Responses will show execution steps broken out with intermediate results to help in debugging.",
	Long:  "Test the execution of an action. Responses will show execution steps broken out with intermediate results to help in debugging.",
	Args:  utils.DetermineArgs([]string{ "actionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Interface{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/integrations/actions/{actionId}/test"
		actionId, args := args[0], args[1:]
		path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)

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
