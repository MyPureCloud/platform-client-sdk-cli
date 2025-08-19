package routing_skillgroups_members

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_skillgroups_members_divisions"
)

func init() {
	routing_skillgroups_membersCmd.AddCommand(routing_skillgroups_members_divisions.Cmdrouting_skillgroups_members_divisions())
	routing_skillgroups_membersCmd.Short = utils.GenerateCustomDescription(routing_skillgroups_membersCmd.Short, routing_skillgroups_members_divisions.Description, )
	routing_skillgroups_membersCmd.Long = routing_skillgroups_membersCmd.Short
}
