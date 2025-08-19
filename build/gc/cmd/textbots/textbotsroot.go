package textbots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots_botflows"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/textbots_bots"
)

func init() {
	textbotsCmd.AddCommand(textbots_botflows.Cmdtextbots_botflows())
	textbotsCmd.AddCommand(textbots_bots.Cmdtextbots_bots())
	textbotsCmd.Short = utils.GenerateCustomDescription(textbotsCmd.Short, textbots_botflows.Description, textbots_bots.Description, )
	textbotsCmd.Long = textbotsCmd.Short
}
