package conversations_socials_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_socials_participants", "SWAGGER_OVERRIDE_/api/v2/conversations/socials/{conversationId}/participants")
	conversations_socials_participantsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_socials_participants"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_socials_participantsCmd)
}

func Cmdconversations_socials_participants() *cobra.Command {
	return conversations_socials_participantsCmd
}
