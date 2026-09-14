package workforcemanagement_agents_schedulingpreferences

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_schedulingpreferences", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/{agentId}/schedulingpreferences")
	workforcemanagement_agents_schedulingpreferencesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_schedulingpreferences"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_schedulingpreferencesCmd)
}

func Cmdworkforcemanagement_agents_schedulingpreferences() *cobra.Command {
	return workforcemanagement_agents_schedulingpreferencesCmd
}
