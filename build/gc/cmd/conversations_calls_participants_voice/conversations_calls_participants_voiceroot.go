package conversations_calls_participants_voice

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_voice_consult"
)

func init() {
	conversations_calls_participants_voiceCmd.AddCommand(conversations_calls_participants_voice_consult.Cmdconversations_calls_participants_voice_consult())
	conversations_calls_participants_voiceCmd.Short = utils.GenerateCustomDescription(conversations_calls_participants_voiceCmd.Short, conversations_calls_participants_voice_consult.Description, )
	conversations_calls_participants_voiceCmd.Long = conversations_calls_participants_voiceCmd.Short
}
