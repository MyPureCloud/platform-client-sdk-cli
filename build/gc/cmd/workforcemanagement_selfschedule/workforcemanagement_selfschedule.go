package workforcemanagement_selfschedule

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_selfschedule", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/selfschedule")
	workforcemanagement_selfscheduleCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_selfschedule"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_selfscheduleCmd)
}

func Cmdworkforcemanagement_selfschedule() *cobra.Command {
	return workforcemanagement_selfscheduleCmd
}
