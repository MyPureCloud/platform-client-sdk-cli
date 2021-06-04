package conversations_messages_inbound

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messages_inbound", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/inbound")
	conversations_messages_inboundCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_inbound"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_inboundCmd)
}

func Cmdconversations_messages_inbound() *cobra.Command {
	return conversations_messages_inboundCmd
}
