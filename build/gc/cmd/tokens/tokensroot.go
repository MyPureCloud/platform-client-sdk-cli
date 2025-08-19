package tokens

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/tokens_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/tokens_timeout"
)

func init() {
	tokensCmd.AddCommand(tokens_me.Cmdtokens_me())
	tokensCmd.AddCommand(tokens_timeout.Cmdtokens_timeout())
	tokensCmd.Short = utils.GenerateCustomDescription(tokensCmd.Short, tokens_me.Description, tokens_timeout.Description, )
	tokensCmd.Long = tokensCmd.Short
}
