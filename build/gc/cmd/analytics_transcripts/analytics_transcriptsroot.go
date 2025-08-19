package analytics_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_transcripts_aggregates"
)

func init() {
	analytics_transcriptsCmd.AddCommand(analytics_transcripts_aggregates.Cmdanalytics_transcripts_aggregates())
	analytics_transcriptsCmd.Short = utils.GenerateCustomDescription(analytics_transcriptsCmd.Short, analytics_transcripts_aggregates.Description, )
	analytics_transcriptsCmd.Long = analytics_transcriptsCmd.Short
}
