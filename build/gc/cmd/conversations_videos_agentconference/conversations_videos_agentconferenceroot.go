package conversations_videos_agentconference

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_videos_agentconference_communications"
)

func init() {
	conversations_videos_agentconferenceCmd.AddCommand(conversations_videos_agentconference_communications.Cmdconversations_videos_agentconference_communications())
	conversations_videos_agentconferenceCmd.Short = utils.GenerateCustomDescription(conversations_videos_agentconferenceCmd.Short, conversations_videos_agentconference_communications.Description, )
	conversations_videos_agentconferenceCmd.Long = conversations_videos_agentconferenceCmd.Short
}
