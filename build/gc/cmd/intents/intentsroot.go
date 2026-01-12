package intents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_categories"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_assignments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_customerintents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_sourceintents"
)

func init() {
	intentsCmd.AddCommand(intents_categories.Cmdintents_categories())
	intentsCmd.AddCommand(intents_assignments.Cmdintents_assignments())
	intentsCmd.AddCommand(intents_customerintents.Cmdintents_customerintents())
	intentsCmd.AddCommand(intents_sourceintents.Cmdintents_sourceintents())
	intentsCmd.Short = utils.GenerateCustomDescription(intentsCmd.Short, intents_categories.Description, intents_assignments.Description, intents_customerintents.Description, intents_sourceintents.Description, )
	intentsCmd.Long = intentsCmd.Short
}
