package recording_localkeys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/recording_localkeys_settings"
)

func init() {
	recording_localkeysCmd.AddCommand(recording_localkeys_settings.Cmdrecording_localkeys_settings())
	recording_localkeysCmd.Short = utils.GenerateCustomDescription(recording_localkeysCmd.Short, recording_localkeys_settings.Description, )
	recording_localkeysCmd.Long = recording_localkeysCmd.Short
}
