package flows_actions_checkout

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
	Description = utils.FormatUsageDescription("flows_actions_checkout", "SWAGGER_OVERRIDE_/api/v2/flows/actions/checkout", )
	flows_actions_checkoutCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_actions_checkout"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_actions_checkoutCmd)
}

func Cmdflows_actions_checkout() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "flow", "", "Flow ID - REQUIRED")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/actions/checkout", utils.FormatPermissions([]string{ "architect:flow:edit",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/actions/checkout")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	createCmd.MarkFlagRequired("flow")
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Flow"
      }
    }
  }
}`)
	flows_actions_checkoutCmd.AddCommand(createCmd)
	return flows_actions_checkoutCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Check-out flow",
	Long:  "Check-out flow",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/actions/checkout"

		flow := utils.GetFlag(cmd.Flags(), "string", "flow")
		if flow != "" {
			queryParams["flow"] = flow
		}
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
