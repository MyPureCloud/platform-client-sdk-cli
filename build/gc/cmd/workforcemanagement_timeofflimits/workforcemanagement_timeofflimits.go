package workforcemanagement_timeofflimits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_timeofflimits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/timeofflimits")
	workforcemanagement_timeofflimitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_timeofflimits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_timeofflimitsCmd)
}

func Cmdworkforcemanagement_timeofflimits() *cobra.Command {
	return workforcemanagement_timeofflimitsCmd
}
