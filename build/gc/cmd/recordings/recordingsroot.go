package recordings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_deletionprotection"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_screensessions"
)

func init() {
	recordingsCmd.AddCommand(recordings_deletionprotection.Cmdrecordings_deletionprotection())
	recordingsCmd.AddCommand(recordings_screensessions.Cmdrecordings_screensessions())
	recordingsCmd.Short = utils.GenerateCustomDescription(recordingsCmd.Short, recordings_deletionprotection.Description, recordings_screensessions.Description, )
	recordingsCmd.Long = recordingsCmd.Short
}
