package greetings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_media"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_downloads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_defaults"
)

func init() {
	greetingsCmd.AddCommand(greetings_media.Cmdgreetings_media())
	greetingsCmd.AddCommand(greetings_groups.Cmdgreetings_groups())
	greetingsCmd.AddCommand(greetings_downloads.Cmdgreetings_downloads())
	greetingsCmd.AddCommand(greetings_users.Cmdgreetings_users())
	greetingsCmd.AddCommand(greetings_defaults.Cmdgreetings_defaults())
	greetingsCmd.Short = utils.GenerateCustomDescription(greetingsCmd.Short, greetings_media.Description, greetings_groups.Description, greetings_downloads.Description, greetings_users.Description, greetings_defaults.Description, )
	greetingsCmd.Long = greetingsCmd.Short
}
