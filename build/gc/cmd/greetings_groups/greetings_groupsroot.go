package greetings_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_groups_downloads"
)

func init() {
	greetings_groupsCmd.AddCommand(greetings_groups_downloads.Cmdgreetings_groups_downloads())
	greetings_groupsCmd.Short = utils.GenerateCustomDescription(greetings_groupsCmd.Short, greetings_groups_downloads.Description, )
	greetings_groupsCmd.Long = greetings_groupsCmd.Short
}
