package workforcemanagement_managementunits_agents

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_agents", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents", )
	workforcemanagement_managementunits_agentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_agents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_agentsCmd)
}

func Cmdworkforcemanagement_managementunits_agents() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "excludeCapabilities", "", "Excludes all capabilities of the agent such as queues, languages, and skills")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}", utils.FormatPermissions([]string{ "wfm:agent:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WfmAgent"
      }
    }
  }
}`)
	workforcemanagement_managementunits_agentsCmd.AddCommand(getCmd)
	return workforcemanagement_managementunits_agentsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [managementUnitId] [agentId]",
	Short: "Get data for agent in the management unit",
	Long:  "Get data for agent in the management unit",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "agentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		agentId, args := args[0], args[1:]
		path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)

		excludeCapabilities := utils.GetFlag(cmd.Flags(), "bool", "excludeCapabilities")
		if excludeCapabilities != "" {
			queryParams["excludeCapabilities"] = excludeCapabilities
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "get"
		const httpMethod = "GET"
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
