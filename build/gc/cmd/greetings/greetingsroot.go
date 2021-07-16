package greetings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_media"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_defaults"
)

func init() {
	greetingsCmd.AddCommand(greetings_media.Cmdgreetings_media())
	greetingsCmd.AddCommand(greetings_defaults.Cmdgreetings_defaults())
	greetingsCmd.Short = utils.GenerateCustomDescription(greetingsCmd.Short, greetings_media.Description, greetings_defaults.Description, )
	greetingsCmd.Long = greetingsCmd.Short
}
