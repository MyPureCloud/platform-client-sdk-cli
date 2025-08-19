package workforcemanagement_businessunits_scheduling

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_scheduling", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling")
	workforcemanagement_businessunits_schedulingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_scheduling"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_schedulingCmd)
}

func Cmdworkforcemanagement_businessunits_scheduling() *cobra.Command {
	return workforcemanagement_businessunits_schedulingCmd
}
