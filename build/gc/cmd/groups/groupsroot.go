package groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_greetings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_members"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_individuals"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_profile"
)

func init() {
	groupsCmd.AddCommand(groups_greetings.Cmdgroups_greetings())
	groupsCmd.AddCommand(groups_search.Cmdgroups_search())
	groupsCmd.AddCommand(groups_members.Cmdgroups_members())
	groupsCmd.AddCommand(groups_individuals.Cmdgroups_individuals())
	groupsCmd.AddCommand(groups_profile.Cmdgroups_profile())
	groupsCmd.Short = utils.GenerateCustomDescription(groupsCmd.Short, groups_greetings.Description, groups_search.Description, groups_members.Description, groups_individuals.Description, groups_profile.Description, )
	groupsCmd.Long = groupsCmd.Short
}
