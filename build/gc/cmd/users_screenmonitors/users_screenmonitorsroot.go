package users_screenmonitors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_screenmonitors_sessions"
)

func init() {
	users_screenmonitorsCmd.AddCommand(users_screenmonitors_sessions.Cmdusers_screenmonitors_sessions())
	users_screenmonitorsCmd.Short = utils.GenerateCustomDescription(users_screenmonitorsCmd.Short, users_screenmonitors_sessions.Description, )
	users_screenmonitorsCmd.Long = users_screenmonitorsCmd.Short
}
