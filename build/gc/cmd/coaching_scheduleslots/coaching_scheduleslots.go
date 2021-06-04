package coaching_scheduleslots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("coaching_scheduleslots", "SWAGGER_OVERRIDE_/api/v2/coaching/scheduleslots")
	coaching_scheduleslotsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("coaching_scheduleslots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(coaching_scheduleslotsCmd)
}

func Cmdcoaching_scheduleslots() *cobra.Command {
	return coaching_scheduleslotsCmd
}
