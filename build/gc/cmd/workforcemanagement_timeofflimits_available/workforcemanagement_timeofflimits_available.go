package workforcemanagement_timeofflimits_available

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_timeofflimits_available", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/timeofflimits/available")
	workforcemanagement_timeofflimits_availableCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_timeofflimits_available"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_timeofflimits_availableCmd)
}

func Cmdworkforcemanagement_timeofflimits_available() *cobra.Command {
	return workforcemanagement_timeofflimits_availableCmd
}
