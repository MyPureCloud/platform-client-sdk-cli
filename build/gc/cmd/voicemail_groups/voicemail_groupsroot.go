package voicemail_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_groups_mailbox"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_groups_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/voicemail_groups_policy"
)

func init() {
	voicemail_groupsCmd.AddCommand(voicemail_groups_mailbox.Cmdvoicemail_groups_mailbox())
	voicemail_groupsCmd.AddCommand(voicemail_groups_messages.Cmdvoicemail_groups_messages())
	voicemail_groupsCmd.AddCommand(voicemail_groups_policy.Cmdvoicemail_groups_policy())
	voicemail_groupsCmd.Short = utils.GenerateCustomDescription(voicemail_groupsCmd.Short, voicemail_groups_mailbox.Description, voicemail_groups_messages.Description, voicemail_groups_policy.Description, )
	voicemail_groupsCmd.Long = voicemail_groupsCmd.Short
}
