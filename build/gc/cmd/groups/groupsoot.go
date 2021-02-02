package groups

import "github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/members"

func init() {
	groupsCmd.AddCommand(members.Cmdmembers())
}