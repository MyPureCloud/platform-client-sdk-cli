package workforcemanagement_businessunits_scheduler

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_scheduler", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduler")
	workforcemanagement_businessunits_schedulerCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_scheduler"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_schedulerCmd)
}

func Cmdworkforcemanagement_businessunits_scheduler() *cobra.Command {
	return workforcemanagement_businessunits_schedulerCmd
}
