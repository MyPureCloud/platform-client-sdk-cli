package guides_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("guides_transcripts", "SWAGGER_OVERRIDE_/api/v2/guides/transcripts")
	guides_transcriptsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("guides_transcripts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(guides_transcriptsCmd)
}

func Cmdguides_transcripts() *cobra.Command {
	return guides_transcriptsCmd
}
