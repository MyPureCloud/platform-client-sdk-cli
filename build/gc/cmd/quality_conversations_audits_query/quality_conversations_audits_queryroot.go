package quality_conversations_audits_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_conversations_audits_query_results"
)

func init() {
	quality_conversations_audits_queryCmd.AddCommand(quality_conversations_audits_query_results.Cmdquality_conversations_audits_query_results())
	quality_conversations_audits_queryCmd.Short = utils.GenerateCustomDescription(quality_conversations_audits_queryCmd.Short, quality_conversations_audits_query_results.Description, )
	quality_conversations_audits_queryCmd.Long = quality_conversations_audits_queryCmd.Short
}
