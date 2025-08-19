package workforcemanagement_businessunits_staffinggroups_planninggroups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_staffinggroups_planninggroups", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/planninggroups")
	workforcemanagement_businessunits_staffinggroups_planninggroupsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_staffinggroups_planninggroups"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_staffinggroups_planninggroupsCmd)
}

func Cmdworkforcemanagement_businessunits_staffinggroups_planninggroups() *cobra.Command {
	return workforcemanagement_businessunits_staffinggroups_planninggroupsCmd
}
