package workforcemanagement_alternativeshifts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_alternativeshifts", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/alternativeshifts")
	workforcemanagement_alternativeshiftsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_alternativeshifts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_alternativeshiftsCmd)
}

func Cmdworkforcemanagement_alternativeshifts() *cobra.Command {
	return workforcemanagement_alternativeshiftsCmd
}
