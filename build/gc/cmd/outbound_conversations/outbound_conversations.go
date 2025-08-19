package outbound_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_conversations", "SWAGGER_OVERRIDE_/api/v2/outbound/conversations")
	outbound_conversationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_conversations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_conversationsCmd)
}

func Cmdoutbound_conversations() *cobra.Command {
	return outbound_conversationsCmd
}
