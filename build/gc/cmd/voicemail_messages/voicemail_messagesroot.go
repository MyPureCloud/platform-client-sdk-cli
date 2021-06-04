package voicemail_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_messages_media"
)

func init() {
	voicemail_messagesCmd.AddCommand(voicemail_messages_media.Cmdvoicemail_messages_media())
	voicemail_messagesCmd.Short = utils.GenerateCustomDescription(voicemail_messagesCmd.Short, voicemail_messages_media.Description, )
	voicemail_messagesCmd.Long = voicemail_messagesCmd.Short
}
