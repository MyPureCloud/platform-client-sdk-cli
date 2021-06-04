package voicemail_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("voicemail_me", "SWAGGER_OVERRIDE_/api/v2/voicemail/me")
	voicemail_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemail_meCmd)
}

func Cmdvoicemail_me() *cobra.Command {
	return voicemail_meCmd
}
