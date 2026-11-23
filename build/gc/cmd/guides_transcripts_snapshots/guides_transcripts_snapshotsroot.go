package guides_transcripts_snapshots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_transcripts_snapshots_generate"
)

func init() {
	guides_transcripts_snapshotsCmd.AddCommand(guides_transcripts_snapshots_generate.Cmdguides_transcripts_snapshots_generate())
	guides_transcripts_snapshotsCmd.Short = utils.GenerateCustomDescription(guides_transcripts_snapshotsCmd.Short, guides_transcripts_snapshots_generate.Description, )
	guides_transcripts_snapshotsCmd.Long = guides_transcripts_snapshotsCmd.Short
}
