package screenrecording

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenrecording_token"
)

func init() {
	screenrecordingCmd.AddCommand(screenrecording_token.Cmdscreenrecording_token())
	screenrecordingCmd.Short = utils.GenerateCustomDescription(screenrecordingCmd.Short, screenrecording_token.Description, )
	screenrecordingCmd.Long = screenrecordingCmd.Short
}
