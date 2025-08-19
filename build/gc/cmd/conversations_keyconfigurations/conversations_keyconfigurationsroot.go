package conversations_keyconfigurations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_keyconfigurations_validate"
)

func init() {
	conversations_keyconfigurationsCmd.AddCommand(conversations_keyconfigurations_validate.Cmdconversations_keyconfigurations_validate())
	conversations_keyconfigurationsCmd.Short = utils.GenerateCustomDescription(conversations_keyconfigurationsCmd.Short, conversations_keyconfigurations_validate.Description, )
	conversations_keyconfigurationsCmd.Long = conversations_keyconfigurationsCmd.Short
}
