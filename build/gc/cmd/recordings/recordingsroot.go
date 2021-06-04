package recordings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_screensessions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recordings_deletionprotection"
)

func init() {
	recordingsCmd.AddCommand(recordings_screensessions.Cmdrecordings_screensessions())
	recordingsCmd.AddCommand(recordings_deletionprotection.Cmdrecordings_deletionprotection())
	recordingsCmd.Short = utils.GenerateCustomDescription(recordingsCmd.Short, recordings_screensessions.Description, recordings_deletionprotection.Description, )
	recordingsCmd.Long = recordingsCmd.Short
}
