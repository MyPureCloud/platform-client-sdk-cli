
package users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/bulkupdate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/queues_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/skills_users"
)

func init() {
	usersCmd.AddCommand(bulkupdate.Cmdbulkupdate())
	usersCmd.AddCommand(queues_users.Cmdqueues_users())
	usersCmd.AddCommand(skills_users.Cmdskills_users())
}
