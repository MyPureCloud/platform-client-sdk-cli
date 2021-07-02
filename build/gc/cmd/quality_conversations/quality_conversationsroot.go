package quality_conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_conversations_audits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_conversations_evaluations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_conversations_surveys"
)

func init() {
	quality_conversationsCmd.AddCommand(quality_conversations_audits.Cmdquality_conversations_audits())
	quality_conversationsCmd.AddCommand(quality_conversations_evaluations.Cmdquality_conversations_evaluations())
	quality_conversationsCmd.AddCommand(quality_conversations_surveys.Cmdquality_conversations_surveys())
	quality_conversationsCmd.Short = utils.GenerateCustomDescription(quality_conversationsCmd.Short, quality_conversations_audits.Description, quality_conversations_evaluations.Description, quality_conversations_surveys.Description, )
	quality_conversationsCmd.Long = quality_conversationsCmd.Short
}
