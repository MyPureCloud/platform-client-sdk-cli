package conversations_participants_internalmessages_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_participants_internalmessages_users", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/internalmessages/users")
	conversations_participants_internalmessages_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_participants_internalmessages_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_participants_internalmessages_usersCmd)
}

func Cmdconversations_participants_internalmessages_users() *cobra.Command {
	return conversations_participants_internalmessages_usersCmd
}
