package conversations_videos_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_participants_communications"
)

func init() {
	conversations_videos_participantsCmd.AddCommand(conversations_videos_participants_communications.Cmdconversations_videos_participants_communications())
	conversations_videos_participantsCmd.Short = utils.GenerateCustomDescription(conversations_videos_participantsCmd.Short, conversations_videos_participants_communications.Description, )
	conversations_videos_participantsCmd.Long = conversations_videos_participantsCmd.Short
}
