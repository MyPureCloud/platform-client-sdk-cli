package routing_email

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_email", "SWAGGER_OVERRIDE_/api/v2/routing/email")
	routing_emailCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_email"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_emailCmd)
}

func Cmdrouting_email() *cobra.Command {
	return routing_emailCmd
}
