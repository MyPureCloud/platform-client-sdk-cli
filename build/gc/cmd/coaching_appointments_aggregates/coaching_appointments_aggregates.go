package coaching_appointments_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("coaching_appointments_aggregates", "SWAGGER_OVERRIDE_/api/v2/coaching/appointments/aggregates")
	coaching_appointments_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_appointments_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_appointments_aggregatesCmd)
}

func Cmdcoaching_appointments_aggregates() *cobra.Command {
	return coaching_appointments_aggregatesCmd
}
