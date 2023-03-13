package conversations_screenshares_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_screenshares_participants", "SWAGGER_OVERRIDE_/api/v2/conversations/screenshares/{conversationId}/participants")
	conversations_screenshares_participantsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_screenshares_participants"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_screenshares_participantsCmd)
}

func Cmdconversations_screenshares_participants() *cobra.Command {
	return conversations_screenshares_participantsCmd
}
