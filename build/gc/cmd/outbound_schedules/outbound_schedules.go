package outbound_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_schedules", "SWAGGER_OVERRIDE_/api/v2/outbound/schedules")
	outbound_schedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_schedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_schedulesCmd)
}

func Cmdoutbound_schedules() *cobra.Command {
	return outbound_schedulesCmd
}
