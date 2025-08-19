package routing_skillgroups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_skillgroups_members"
)

func init() {
	routing_skillgroupsCmd.AddCommand(routing_skillgroups_members.Cmdrouting_skillgroups_members())
	routing_skillgroupsCmd.Short = utils.GenerateCustomDescription(routing_skillgroupsCmd.Short, routing_skillgroups_members.Description, )
	routing_skillgroupsCmd.Long = routing_skillgroupsCmd.Short
}
