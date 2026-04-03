package workforcemanagement_businessunits_shifttrading_weeks_summary

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_weeks_summary", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/shifttrading/weeks/summary")
	workforcemanagement_businessunits_shifttrading_weeks_summaryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_weeks_summary"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_shifttrading_weeks_summaryCmd)
}

func Cmdworkforcemanagement_businessunits_shifttrading_weeks_summary() *cobra.Command {
	return workforcemanagement_businessunits_shifttrading_weeks_summaryCmd
}
