package backgroundassistant

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/backgroundassistant_token"
)

func init() {
	backgroundassistantCmd.AddCommand(backgroundassistant_token.Cmdbackgroundassistant_token())
	backgroundassistantCmd.Short = utils.GenerateCustomDescription(backgroundassistantCmd.Short, backgroundassistant_token.Description, )
	backgroundassistantCmd.Long = backgroundassistantCmd.Short
}
