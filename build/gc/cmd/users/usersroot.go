package users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/queue"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/skill"
)

func init() {
	usersCmd.AddCommand(queue.Cmdqueue())
	usersCmd.AddCommand(skill.Cmdskill())
}