package conversations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowse"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_keyconfigurations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_tags"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_utilizationlabel"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_secureattributes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_barge"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_disconnect"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_assign"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_faxes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_aftercallwork"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_callbacks"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_emails"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_screenshares"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_socials"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_suggestions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_summaries"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_recordingmetadata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_recordings"
)

func init() {
	conversationsCmd.AddCommand(conversations_cobrowse.Cmdconversations_cobrowse())
	conversationsCmd.AddCommand(conversations_keyconfigurations.Cmdconversations_keyconfigurations())
	conversationsCmd.AddCommand(conversations_participants.Cmdconversations_participants())
	conversationsCmd.AddCommand(conversations_tags.Cmdconversations_tags())
	conversationsCmd.AddCommand(conversations_utilizationlabel.Cmdconversations_utilizationlabel())
	conversationsCmd.AddCommand(conversations_secureattributes.Cmdconversations_secureattributes())
	conversationsCmd.AddCommand(conversations_barge.Cmdconversations_barge())
	conversationsCmd.AddCommand(conversations_disconnect.Cmdconversations_disconnect())
	conversationsCmd.AddCommand(conversations_assign.Cmdconversations_assign())
	conversationsCmd.AddCommand(conversations_faxes.Cmdconversations_faxes())
	conversationsCmd.AddCommand(conversations_settings.Cmdconversations_settings())
	conversationsCmd.AddCommand(conversations_aftercallwork.Cmdconversations_aftercallwork())
	conversationsCmd.AddCommand(conversations_callbacks.Cmdconversations_callbacks())
	conversationsCmd.AddCommand(conversations_calls.Cmdconversations_calls())
	conversationsCmd.AddCommand(conversations_chats.Cmdconversations_chats())
	conversationsCmd.AddCommand(conversations_cobrowsesessions.Cmdconversations_cobrowsesessions())
	conversationsCmd.AddCommand(conversations_emails.Cmdconversations_emails())
	conversationsCmd.AddCommand(conversations_messages.Cmdconversations_messages())
	conversationsCmd.AddCommand(conversations_screenshares.Cmdconversations_screenshares())
	conversationsCmd.AddCommand(conversations_socials.Cmdconversations_socials())
	conversationsCmd.AddCommand(conversations_videos.Cmdconversations_videos())
	conversationsCmd.AddCommand(conversations_messaging.Cmdconversations_messaging())
	conversationsCmd.AddCommand(conversations_suggestions.Cmdconversations_suggestions())
	conversationsCmd.AddCommand(conversations_summaries.Cmdconversations_summaries())
	conversationsCmd.AddCommand(conversations_recordingmetadata.Cmdconversations_recordingmetadata())
	conversationsCmd.AddCommand(conversations_recordings.Cmdconversations_recordings())
	conversationsCmd.Short = utils.GenerateCustomDescription(conversationsCmd.Short, conversations_cobrowse.Description, conversations_keyconfigurations.Description, conversations_participants.Description, conversations_tags.Description, conversations_utilizationlabel.Description, conversations_secureattributes.Description, conversations_barge.Description, conversations_disconnect.Description, conversations_assign.Description, conversations_faxes.Description, conversations_settings.Description, conversations_aftercallwork.Description, conversations_callbacks.Description, conversations_calls.Description, conversations_chats.Description, conversations_cobrowsesessions.Description, conversations_emails.Description, conversations_messages.Description, conversations_screenshares.Description, conversations_socials.Description, conversations_videos.Description, conversations_messaging.Description, conversations_suggestions.Description, conversations_summaries.Description, conversations_recordingmetadata.Description, conversations_recordings.Description, )
	conversationsCmd.Long = conversationsCmd.Short
}
