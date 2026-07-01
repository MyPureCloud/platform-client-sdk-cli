package conversations_participants_screenmonitors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_screenmonitors_sessions"
)

func init() {
	conversations_participants_screenmonitorsCmd.AddCommand(conversations_participants_screenmonitors_sessions.Cmdconversations_participants_screenmonitors_sessions())
	conversations_participants_screenmonitorsCmd.Short = utils.GenerateCustomDescription(conversations_participants_screenmonitorsCmd.Short, conversations_participants_screenmonitors_sessions.Description, )
	conversations_participants_screenmonitorsCmd.Long = conversations_participants_screenmonitorsCmd.Short
}
