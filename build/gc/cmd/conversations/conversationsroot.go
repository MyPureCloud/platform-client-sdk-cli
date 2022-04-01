package conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_disconnect"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_assign"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_tags"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_faxes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_callbacks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_emails"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_recordingmetadata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_recordings"
)

func init() {
	conversationsCmd.AddCommand(conversations_disconnect.Cmdconversations_disconnect())
	conversationsCmd.AddCommand(conversations_assign.Cmdconversations_assign())
	conversationsCmd.AddCommand(conversations_tags.Cmdconversations_tags())
	conversationsCmd.AddCommand(conversations_faxes.Cmdconversations_faxes())
	conversationsCmd.AddCommand(conversations_callbacks.Cmdconversations_callbacks())
	conversationsCmd.AddCommand(conversations_calls.Cmdconversations_calls())
	conversationsCmd.AddCommand(conversations_chats.Cmdconversations_chats())
	conversationsCmd.AddCommand(conversations_cobrowsesessions.Cmdconversations_cobrowsesessions())
	conversationsCmd.AddCommand(conversations_emails.Cmdconversations_emails())
	conversationsCmd.AddCommand(conversations_messages.Cmdconversations_messages())
	conversationsCmd.AddCommand(conversations_messaging.Cmdconversations_messaging())
	conversationsCmd.AddCommand(conversations_participants.Cmdconversations_participants())
	conversationsCmd.AddCommand(conversations_recordingmetadata.Cmdconversations_recordingmetadata())
	conversationsCmd.AddCommand(conversations_recordings.Cmdconversations_recordings())
	conversationsCmd.Short = utils.GenerateCustomDescription(conversationsCmd.Short, conversations_disconnect.Description, conversations_assign.Description, conversations_tags.Description, conversations_faxes.Description, conversations_callbacks.Description, conversations_calls.Description, conversations_chats.Description, conversations_cobrowsesessions.Description, conversations_emails.Description, conversations_messages.Description, conversations_messaging.Description, conversations_participants.Description, conversations_recordingmetadata.Description, conversations_recordings.Description, )
	conversationsCmd.Long = conversationsCmd.Short
}
