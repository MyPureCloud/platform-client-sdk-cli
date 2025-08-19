package quality_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_conversations", "SWAGGER_OVERRIDE_/api/v2/quality/conversations")
	quality_conversationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_conversations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_conversationsCmd)
}

func Cmdquality_conversations() *cobra.Command {
	return quality_conversationsCmd
}
