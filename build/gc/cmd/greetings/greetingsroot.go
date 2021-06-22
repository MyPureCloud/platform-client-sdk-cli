package greetings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_defaults"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_media"
)

func init() {
	greetingsCmd.AddCommand(greetings_defaults.Cmdgreetings_defaults())
	greetingsCmd.AddCommand(greetings_media.Cmdgreetings_media())
	greetingsCmd.Short = utils.GenerateCustomDescription(greetingsCmd.Short, greetings_defaults.Description, greetings_media.Description, )
	greetingsCmd.Long = greetingsCmd.Short
}
