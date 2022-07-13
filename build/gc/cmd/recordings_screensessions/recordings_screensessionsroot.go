package recordings_screensessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_screensessions_metadata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_screensessions_acknowledge"
)

func init() {
	recordings_screensessionsCmd.AddCommand(recordings_screensessions_metadata.Cmdrecordings_screensessions_metadata())
	recordings_screensessionsCmd.AddCommand(recordings_screensessions_acknowledge.Cmdrecordings_screensessions_acknowledge())
	recordings_screensessionsCmd.Short = utils.GenerateCustomDescription(recordings_screensessionsCmd.Short, recordings_screensessions_metadata.Description, recordings_screensessions_acknowledge.Description, )
	recordings_screensessionsCmd.Long = recordings_screensessionsCmd.Short
}
