package voicemail_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_users_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_users_mailbox"
)

func init() {
	voicemail_usersCmd.AddCommand(voicemail_users_messages.Cmdvoicemail_users_messages())
	voicemail_usersCmd.AddCommand(voicemail_users_mailbox.Cmdvoicemail_users_mailbox())
	voicemail_usersCmd.Short = utils.GenerateCustomDescription(voicemail_usersCmd.Short, voicemail_users_messages.Description, voicemail_users_mailbox.Description, )
	voicemail_usersCmd.Long = voicemail_usersCmd.Short
}
