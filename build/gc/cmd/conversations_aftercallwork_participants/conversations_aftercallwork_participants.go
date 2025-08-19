package conversations_aftercallwork_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_aftercallwork_participants", "SWAGGER_OVERRIDE_/api/v2/conversations/aftercallwork/{conversationId}/participants")
	conversations_aftercallwork_participantsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_aftercallwork_participants"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_aftercallwork_participantsCmd)
}

func Cmdconversations_aftercallwork_participants() *cobra.Command {
	return conversations_aftercallwork_participantsCmd
}
