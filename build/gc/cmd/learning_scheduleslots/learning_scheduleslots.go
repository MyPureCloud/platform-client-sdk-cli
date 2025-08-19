package learning_scheduleslots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("learning_scheduleslots", "SWAGGER_OVERRIDE_/api/v2/learning/scheduleslots")
	learning_scheduleslotsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_scheduleslots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_scheduleslotsCmd)
}

func Cmdlearning_scheduleslots() *cobra.Command {
	return learning_scheduleslotsCmd
}
