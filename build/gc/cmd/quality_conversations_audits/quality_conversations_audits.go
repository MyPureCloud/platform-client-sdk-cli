package quality_conversations_audits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_conversations_audits", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/audits")
	quality_conversations_auditsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_conversations_audits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_conversations_auditsCmd)
}

func Cmdquality_conversations_audits() *cobra.Command {
	return quality_conversations_auditsCmd
}
