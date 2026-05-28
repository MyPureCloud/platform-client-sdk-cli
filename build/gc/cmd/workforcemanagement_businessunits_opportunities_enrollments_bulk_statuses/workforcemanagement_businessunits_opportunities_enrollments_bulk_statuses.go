package workforcemanagement_businessunits_opportunities_enrollments_bulk_statuses

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_opportunities_enrollments_bulk_statuses", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/opportunities/enrollments/bulk/statuses")
	workforcemanagement_businessunits_opportunities_enrollments_bulk_statusesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_opportunities_enrollments_bulk_statuses"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_opportunities_enrollments_bulk_statusesCmd)
}

func Cmdworkforcemanagement_businessunits_opportunities_enrollments_bulk_statuses() *cobra.Command {
	return workforcemanagement_businessunits_opportunities_enrollments_bulk_statusesCmd
}
