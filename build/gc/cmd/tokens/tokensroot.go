package tokens

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/tokens_me"
)

func init() {
	tokensCmd.AddCommand(tokens_me.Cmdtokens_me())
	tokensCmd.Short = utils.GenerateCustomDescription(tokensCmd.Short, tokens_me.Description, )
	tokensCmd.Long = tokensCmd.Short
}
