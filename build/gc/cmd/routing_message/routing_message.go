package routing_message

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_message", "SWAGGER_OVERRIDE_/api/v2/routing/message")
	routing_messageCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_message"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_messageCmd)
}

func Cmdrouting_message() *cobra.Command {
	return routing_messageCmd
}
