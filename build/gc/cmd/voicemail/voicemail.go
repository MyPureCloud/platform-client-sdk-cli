package voicemail

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("voicemail", "SWAGGER_OVERRIDE_/api/v2/voicemail")
	voicemailCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("voicemail"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(voicemailCmd)
}

func Cmdvoicemail() *cobra.Command {
	return voicemailCmd
}
