package workforcemanagement_calendar_url

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_calendar_url", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/calendar/url")
	workforcemanagement_calendar_urlCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_calendar_url"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_calendar_urlCmd)
}

func Cmdworkforcemanagement_calendar_url() *cobra.Command {
	return workforcemanagement_calendar_urlCmd
}
