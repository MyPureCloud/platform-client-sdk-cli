package intents_customerintents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_customerintents_sourceintents"
)

func init() {
	intents_customerintentsCmd.AddCommand(intents_customerintents_sourceintents.Cmdintents_customerintents_sourceintents())
	intents_customerintentsCmd.Short = utils.GenerateCustomDescription(intents_customerintentsCmd.Short, intents_customerintents_sourceintents.Description, )
	intents_customerintentsCmd.Long = intents_customerintentsCmd.Short
}
