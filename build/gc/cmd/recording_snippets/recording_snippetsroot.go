package recording_snippets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_snippets_settings"
)

func init() {
	recording_snippetsCmd.AddCommand(recording_snippets_settings.Cmdrecording_snippets_settings())
	recording_snippetsCmd.Short = utils.GenerateCustomDescription(recording_snippetsCmd.Short, recording_snippets_settings.Description, )
	recording_snippetsCmd.Long = recording_snippetsCmd.Short
}
