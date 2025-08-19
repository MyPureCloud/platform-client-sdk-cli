package groups_dynamicsettings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/groups_dynamicsettings_preview"
)

func init() {
	groups_dynamicsettingsCmd.AddCommand(groups_dynamicsettings_preview.Cmdgroups_dynamicsettings_preview())
	groups_dynamicsettingsCmd.Short = utils.GenerateCustomDescription(groups_dynamicsettingsCmd.Short, groups_dynamicsettings_preview.Description, )
	groups_dynamicsettingsCmd.Long = groups_dynamicsettingsCmd.Short
}
