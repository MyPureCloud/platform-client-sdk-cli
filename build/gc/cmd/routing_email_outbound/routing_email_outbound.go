package routing_email_outbound

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_email_outbound", "SWAGGER_OVERRIDE_/api/v2/routing/email/outbound")
	routing_email_outboundCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_email_outbound"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_email_outboundCmd)
}

func Cmdrouting_email_outbound() *cobra.Command {
	return routing_email_outboundCmd
}
