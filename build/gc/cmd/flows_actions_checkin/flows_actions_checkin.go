package flows_actions_checkin

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
	Description = utils.FormatUsageDescription("flows_actions_checkin", "SWAGGER_OVERRIDE_/api/v2/flows/actions/checkin", )
	flows_actions_checkinCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_actions_checkin"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_actions_checkinCmd)
}

func Cmdflows_actions_checkin() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "string", "flow", "", "Flow ID - REQUIRED")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/flows/actions/checkin", utils.FormatPermissions([]string{ "architect:flow:edit", "architect:flow:unlock",  }), utils.GenerateDevCentreLink("POST", "Architect", "/api/v2/flows/actions/checkin")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	createCmd.MarkFlagRequired("flow")
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Operation"
  }
}`)
	flows_actions_checkinCmd.AddCommand(createCmd)
	
	return flows_actions_checkinCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Check-in flow",
	Long:  "Check-in flow",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/flows/actions/checkin"


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
