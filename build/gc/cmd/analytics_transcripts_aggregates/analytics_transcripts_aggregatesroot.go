package analytics_transcripts_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_transcripts_aggregates_query"
)

func init() {
	analytics_transcripts_aggregatesCmd.AddCommand(analytics_transcripts_aggregates_query.Cmdanalytics_transcripts_aggregates_query())
	analytics_transcripts_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_transcripts_aggregatesCmd.Short, analytics_transcripts_aggregates_query.Description, )
	analytics_transcripts_aggregatesCmd.Long = analytics_transcripts_aggregatesCmd.Short
}
