package groups_greetings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_greetings_defaults"
)

func init() {
	groups_greetingsCmd.AddCommand(groups_greetings_defaults.Cmdgroups_greetings_defaults())
	groups_greetingsCmd.Short = utils.GenerateCustomDescription(groups_greetingsCmd.Short, groups_greetings_defaults.Description, )
	groups_greetingsCmd.Long = groups_greetingsCmd.Short
}
