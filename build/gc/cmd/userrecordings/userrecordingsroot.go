package userrecordings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/userrecordings_media"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/userrecordings_transcoding"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/userrecordings_summary"
)

func init() {
	userrecordingsCmd.AddCommand(userrecordings_media.Cmduserrecordings_media())
	userrecordingsCmd.AddCommand(userrecordings_transcoding.Cmduserrecordings_transcoding())
	userrecordingsCmd.AddCommand(userrecordings_summary.Cmduserrecordings_summary())
	userrecordingsCmd.Short = utils.GenerateCustomDescription(userrecordingsCmd.Short, userrecordings_media.Description, userrecordings_transcoding.Description, userrecordings_summary.Description, )
	userrecordingsCmd.Long = userrecordingsCmd.Short
}
