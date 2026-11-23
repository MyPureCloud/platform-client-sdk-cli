package guides_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_transcripts_discovery"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/guides_transcripts_snapshots"
)

func init() {
	guides_transcriptsCmd.AddCommand(guides_transcripts_discovery.Cmdguides_transcripts_discovery())
	guides_transcriptsCmd.AddCommand(guides_transcripts_snapshots.Cmdguides_transcripts_snapshots())
	guides_transcriptsCmd.Short = utils.GenerateCustomDescription(guides_transcriptsCmd.Short, guides_transcripts_discovery.Description, guides_transcripts_snapshots.Description, )
	guides_transcriptsCmd.Long = guides_transcriptsCmd.Short
}
