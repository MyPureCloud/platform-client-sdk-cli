package speechandtextanalytics_conversations_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_conversations_communications", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/conversations/{conversationId}/communications")
	speechandtextanalytics_conversations_communicationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_conversations_communications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_conversations_communicationsCmd)
}

func Cmdspeechandtextanalytics_conversations_communications() *cobra.Command {
	return speechandtextanalytics_conversations_communicationsCmd
}
