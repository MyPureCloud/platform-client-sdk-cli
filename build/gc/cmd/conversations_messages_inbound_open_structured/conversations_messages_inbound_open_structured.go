package conversations_messages_inbound_open_structured

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messages_inbound_open_structured", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{integrationId}/inbound/open/structured")
	conversations_messages_inbound_open_structuredCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_inbound_open_structured"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_inbound_open_structuredCmd)
}

func Cmdconversations_messages_inbound_open_structured() *cobra.Command {
	return conversations_messages_inbound_open_structuredCmd
}
