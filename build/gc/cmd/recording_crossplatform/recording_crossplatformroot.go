package recording_crossplatform

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_crossplatform_mediaretentionpolicies"
)

func init() {
	recording_crossplatformCmd.AddCommand(recording_crossplatform_mediaretentionpolicies.Cmdrecording_crossplatform_mediaretentionpolicies())
	recording_crossplatformCmd.Short = utils.GenerateCustomDescription(recording_crossplatformCmd.Short, recording_crossplatform_mediaretentionpolicies.Description, )
	recording_crossplatformCmd.Long = recording_crossplatformCmd.Short
}
