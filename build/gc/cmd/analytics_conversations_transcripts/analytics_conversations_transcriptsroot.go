package analytics_conversations_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_transcripts_query"
)

func init() {
	analytics_conversations_transcriptsCmd.AddCommand(analytics_conversations_transcripts_query.Cmdanalytics_conversations_transcripts_query())
	analytics_conversations_transcriptsCmd.Short = utils.GenerateCustomDescription(analytics_conversations_transcriptsCmd.Short, analytics_conversations_transcripts_query.Description, )
	analytics_conversations_transcriptsCmd.Long = analytics_conversations_transcriptsCmd.Short
}
