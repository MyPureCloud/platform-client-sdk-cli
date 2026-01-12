package intents_customerintents_sourceintents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_customerintents_sourceintents_bulk"
)

func init() {
	intents_customerintents_sourceintentsCmd.AddCommand(intents_customerintents_sourceintents_bulk.Cmdintents_customerintents_sourceintents_bulk())
	intents_customerintents_sourceintentsCmd.Short = utils.GenerateCustomDescription(intents_customerintents_sourceintentsCmd.Short, intents_customerintents_sourceintents_bulk.Description, )
	intents_customerintents_sourceintentsCmd.Long = intents_customerintents_sourceintentsCmd.Short
}
