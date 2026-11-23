package guides_transcripts_snapshots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("guides_transcripts_snapshots", "SWAGGER_OVERRIDE_/api/v2/guides/transcripts/snapshots")
	guides_transcripts_snapshotsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("guides_transcripts_snapshots"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(guides_transcripts_snapshotsCmd)
}

func Cmdguides_transcripts_snapshots() *cobra.Command {
	return guides_transcripts_snapshotsCmd
}
