package voicemail

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_mailbox"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_userpolicies"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_queues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_policy"
)

func init() {
	voicemailCmd.AddCommand(voicemail_messages.Cmdvoicemail_messages())
	voicemailCmd.AddCommand(voicemail_mailbox.Cmdvoicemail_mailbox())
	voicemailCmd.AddCommand(voicemail_me.Cmdvoicemail_me())
	voicemailCmd.AddCommand(voicemail_groups.Cmdvoicemail_groups())
	voicemailCmd.AddCommand(voicemail_userpolicies.Cmdvoicemail_userpolicies())
	voicemailCmd.AddCommand(voicemail_search.Cmdvoicemail_search())
	voicemailCmd.AddCommand(voicemail_queues.Cmdvoicemail_queues())
	voicemailCmd.AddCommand(voicemail_policy.Cmdvoicemail_policy())
	voicemailCmd.Short = utils.GenerateCustomDescription(voicemailCmd.Short, voicemail_messages.Description, voicemail_mailbox.Description, voicemail_me.Description, voicemail_groups.Description, voicemail_userpolicies.Description, voicemail_search.Description, voicemail_queues.Description, voicemail_policy.Description, )
	voicemailCmd.Long = voicemailCmd.Short
}
