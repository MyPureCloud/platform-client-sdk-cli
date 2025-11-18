package recordings_screensessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_screensessions_acknowledge"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_screensessions_details"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_screensessions_metadata"
)

func init() {
	recordings_screensessionsCmd.AddCommand(recordings_screensessions_acknowledge.Cmdrecordings_screensessions_acknowledge())
	recordings_screensessionsCmd.AddCommand(recordings_screensessions_details.Cmdrecordings_screensessions_details())
	recordings_screensessionsCmd.AddCommand(recordings_screensessions_metadata.Cmdrecordings_screensessions_metadata())
	recordings_screensessionsCmd.Short = utils.GenerateCustomDescription(recordings_screensessionsCmd.Short, recordings_screensessions_acknowledge.Description, recordings_screensessions_details.Description, recordings_screensessions_metadata.Description, )
	recordings_screensessionsCmd.Long = recordings_screensessionsCmd.Short
}
