package quality_conversations_audits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_conversations_audits_query"
)

func init() {
	quality_conversations_auditsCmd.AddCommand(quality_conversations_audits_query.Cmdquality_conversations_audits_query())
	quality_conversations_auditsCmd.Short = utils.GenerateCustomDescription(quality_conversations_auditsCmd.Short, quality_conversations_audits_query.Description, )
	quality_conversations_auditsCmd.Long = quality_conversations_auditsCmd.Short
}
