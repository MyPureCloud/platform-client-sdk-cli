package orgauthorization_trustees_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees_groups_roledivisions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees_groups_roles"
)

func init() {
	orgauthorization_trustees_groupsCmd.AddCommand(orgauthorization_trustees_groups_roledivisions.Cmdorgauthorization_trustees_groups_roledivisions())
	orgauthorization_trustees_groupsCmd.AddCommand(orgauthorization_trustees_groups_roles.Cmdorgauthorization_trustees_groups_roles())
	orgauthorization_trustees_groupsCmd.Short = utils.GenerateCustomDescription(orgauthorization_trustees_groupsCmd.Short, orgauthorization_trustees_groups_roledivisions.Description, orgauthorization_trustees_groups_roles.Description, )
	orgauthorization_trustees_groupsCmd.Long = orgauthorization_trustees_groupsCmd.Short
}
