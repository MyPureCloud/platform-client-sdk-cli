package voicemail_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_me_mailbox"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_me_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_me_policy"
)

func init() {
	voicemail_meCmd.AddCommand(voicemail_me_mailbox.Cmdvoicemail_me_mailbox())
	voicemail_meCmd.AddCommand(voicemail_me_messages.Cmdvoicemail_me_messages())
	voicemail_meCmd.AddCommand(voicemail_me_policy.Cmdvoicemail_me_policy())
	voicemail_meCmd.Short = utils.GenerateCustomDescription(voicemail_meCmd.Short, voicemail_me_mailbox.Description, voicemail_me_messages.Description, voicemail_me_policy.Description, )
	voicemail_meCmd.Long = voicemail_meCmd.Short
}
