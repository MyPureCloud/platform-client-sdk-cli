package voicemail_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("voicemail_queues", "SWAGGER_OVERRIDE_/api/v2/voicemail/queues")
	voicemail_queuesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail_queues"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemail_queuesCmd)
}

func Cmdvoicemail_queues() *cobra.Command {
	return voicemail_queuesCmd
}
