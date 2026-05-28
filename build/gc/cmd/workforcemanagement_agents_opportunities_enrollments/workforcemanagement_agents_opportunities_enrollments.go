package workforcemanagement_agents_opportunities_enrollments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_opportunities_enrollments", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/opportunities/enrollments")
	workforcemanagement_agents_opportunities_enrollmentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_opportunities_enrollments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_opportunities_enrollmentsCmd)
}

func Cmdworkforcemanagement_agents_opportunities_enrollments() *cobra.Command {
	return workforcemanagement_agents_opportunities_enrollmentsCmd
}
