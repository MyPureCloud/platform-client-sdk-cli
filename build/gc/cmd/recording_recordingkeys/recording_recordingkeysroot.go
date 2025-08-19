package recording_recordingkeys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_recordingkeys_rotationschedule"
)

func init() {
	recording_recordingkeysCmd.AddCommand(recording_recordingkeys_rotationschedule.Cmdrecording_recordingkeys_rotationschedule())
	recording_recordingkeysCmd.Short = utils.GenerateCustomDescription(recording_recordingkeysCmd.Short, recording_recordingkeys_rotationschedule.Description, )
	recording_recordingkeysCmd.Long = recording_recordingkeysCmd.Short
}
