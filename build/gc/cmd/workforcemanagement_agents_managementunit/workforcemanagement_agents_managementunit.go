package workforcemanagement_agents_managementunit

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_managementunit", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/{agentId}/managementunit", )
	workforcemanagement_agents_managementunitCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_managementunit"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_managementunitCmd)
}

func Cmdworkforcemanagement_agents_managementunit() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/agents/{agentId}/managementunit", utils.FormatPermissions([]string{ "wfm:agent:view", "wfm:publishedSchedule:view", "wfm:schedule:view", "coaching:appointment:add", "coaching:appointment:edit",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/agents/{agentId}/managementunit")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AgentManagementUnitReference&quot;
  }
}`)
	workforcemanagement_agents_managementunitCmd.AddCommand(getCmd)
	
	return workforcemanagement_agents_managementunitCmd
}

var getCmd = &cobra.Command{
	Use:   "get [agentId]",
	Short: "Get the management unit to which the agent belongs",
	Long:  "Get the management unit to which the agent belongs",
	Args:  utils.DetermineArgs([]string{ "agentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/agents/{agentId}/managementunit"
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
