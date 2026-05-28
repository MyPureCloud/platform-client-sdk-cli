package workforcemanagement_agents_opportunities_enrollments_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_opportunities_enrollments_bulk", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/opportunities/enrollments/bulk")
	workforcemanagement_agents_opportunities_enrollments_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_opportunities_enrollments_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_opportunities_enrollments_bulkCmd)
}

func Cmdworkforcemanagement_agents_opportunities_enrollments_bulk() *cobra.Command {
	return workforcemanagement_agents_opportunities_enrollments_bulkCmd
}
