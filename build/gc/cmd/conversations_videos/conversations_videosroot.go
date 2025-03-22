package conversations_videos

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_recordingstate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_agentconference"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_meetings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_details"
)

func init() {
	conversations_videosCmd.AddCommand(conversations_videos_recordingstate.Cmdconversations_videos_recordingstate())
	conversations_videosCmd.AddCommand(conversations_videos_participants.Cmdconversations_videos_participants())
	conversations_videosCmd.AddCommand(conversations_videos_agentconference.Cmdconversations_videos_agentconference())
	conversations_videosCmd.AddCommand(conversations_videos_meetings.Cmdconversations_videos_meetings())
	conversations_videosCmd.AddCommand(conversations_videos_details.Cmdconversations_videos_details())
	conversations_videosCmd.Short = utils.GenerateCustomDescription(conversations_videosCmd.Short, conversations_videos_recordingstate.Description, conversations_videos_participants.Description, conversations_videos_agentconference.Description, conversations_videos_meetings.Description, conversations_videos_details.Description, )
	conversations_videosCmd.Long = conversations_videosCmd.Short
}
