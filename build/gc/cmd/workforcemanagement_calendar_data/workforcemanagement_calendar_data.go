package workforcemanagement_calendar_data

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_calendar_data", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/calendar/data")
	workforcemanagement_calendar_dataCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_calendar_data"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_calendar_dataCmd)
}

func Cmdworkforcemanagement_calendar_data() *cobra.Command {
	return workforcemanagement_calendar_dataCmd
}
