package workforcemanagement_calendar

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_calendar", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/calendar")
	workforcemanagement_calendarCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_calendar"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_calendarCmd)
}

func Cmdworkforcemanagement_calendar() *cobra.Command {
	return workforcemanagement_calendarCmd
}
