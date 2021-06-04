package speechandtextanalytics_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("speechandtextanalytics_transcripts", "SWAGGER_OVERRIDE_/api/v2/speechandtextanalytics/transcripts")
	speechandtextanalytics_transcriptsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("speechandtextanalytics_transcripts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(speechandtextanalytics_transcriptsCmd)
}

func Cmdspeechandtextanalytics_transcripts() *cobra.Command {
	return speechandtextanalytics_transcriptsCmd
}
