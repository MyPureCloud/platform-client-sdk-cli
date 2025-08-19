package groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_individuals"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_members"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_profile"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_greetings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_dynamicsettings"
)

func init() {
	groupsCmd.AddCommand(groups_individuals.Cmdgroups_individuals())
	groupsCmd.AddCommand(groups_members.Cmdgroups_members())
	groupsCmd.AddCommand(groups_profile.Cmdgroups_profile())
	groupsCmd.AddCommand(groups_search.Cmdgroups_search())
	groupsCmd.AddCommand(groups_greetings.Cmdgroups_greetings())
	groupsCmd.AddCommand(groups_dynamicsettings.Cmdgroups_dynamicsettings())
	groupsCmd.Short = utils.GenerateCustomDescription(groupsCmd.Short, groups_individuals.Description, groups_members.Description, groups_profile.Description, groups_search.Description, groups_greetings.Description, groups_dynamicsettings.Description, )
	groupsCmd.Long = groupsCmd.Short
}
