package integrations_actions_draft_testfile

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
	Description = utils.FormatUsageDescription("integrations_actions_draft_testfile", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/{actionId}/draft/test", )
	integrations_actions_draft_testfileCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_actions_draft_testfile"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_actions_draft_testfileCmd)
}

func Cmdintegrations_actions_draft_testfile() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/integrations/actions/{actionId}/draft/test", utils.FormatPermissions([]string{ "integrations:action:execute",  }), utils.GenerateDevCentreLink("POST", "Integrations", "/api/v2/integrations/actions/{actionId}/draft/test")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Map of parameters used for variable substitution.",
  "required" : true,
  "schema" : {
    "type" : "object",
    "additionalProperties" : {
      "type" : "object",
      "properties" : { }
    }
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/TestExecutionResult"
  }
}`)
	integrations_actions_draft_testfileCmd.AddCommand(createCmd)
	
	return integrations_actions_draft_testfileCmd
}

var createCmd = &cobra.Command{
	Use:   "create [actionId]",
	Short: "Test the execution of a draft. Responses will show execution steps broken out with intermediate results to help in debugging.",
	Long:  "Test the execution of a draft. Responses will show execution steps broken out with intermediate results to help in debugging.",
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

		path := "/api/v2/integrations/actions/{actionId}/draft/test"
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

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
