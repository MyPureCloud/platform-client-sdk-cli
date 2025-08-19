package voicemail_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_queues_messages"
)

func init() {
	voicemail_queuesCmd.AddCommand(voicemail_queues_messages.Cmdvoicemail_queues_messages())
	voicemail_queuesCmd.Short = utils.GenerateCustomDescription(voicemail_queuesCmd.Short, voicemail_queues_messages.Description, )
	voicemail_queuesCmd.Long = voicemail_queuesCmd.Short
}
