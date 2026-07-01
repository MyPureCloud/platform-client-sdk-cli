package conversations_participants_screenmonitors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_participants_screenmonitors", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/screenmonitors")
	conversations_participants_screenmonitorsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_participants_screenmonitors"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_participants_screenmonitorsCmd)
}

func Cmdconversations_participants_screenmonitors() *cobra.Command {
	return conversations_participants_screenmonitorsCmd
}
