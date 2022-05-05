package recording_keyconfigurations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_keyconfigurations_validate"
)

func init() {
	recording_keyconfigurationsCmd.AddCommand(recording_keyconfigurations_validate.Cmdrecording_keyconfigurations_validate())
	recording_keyconfigurationsCmd.Short = utils.GenerateCustomDescription(recording_keyconfigurationsCmd.Short, recording_keyconfigurations_validate.Description, )
	recording_keyconfigurationsCmd.Long = recording_keyconfigurationsCmd.Short
}
