package userrecordings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/userrecordings_summary"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/userrecordings_media"
)

func init() {
	userrecordingsCmd.AddCommand(userrecordings_summary.Cmduserrecordings_summary())
	userrecordingsCmd.AddCommand(userrecordings_media.Cmduserrecordings_media())
	userrecordingsCmd.Short = utils.GenerateCustomDescription(userrecordingsCmd.Short, userrecordings_summary.Description, userrecordings_media.Description, )
	userrecordingsCmd.Long = userrecordingsCmd.Short
}
