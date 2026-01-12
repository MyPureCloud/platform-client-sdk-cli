package intents_assignments_externalcontacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_assignments_externalcontacts_customerintents"
)

func init() {
	intents_assignments_externalcontactsCmd.AddCommand(intents_assignments_externalcontacts_customerintents.Cmdintents_assignments_externalcontacts_customerintents())
	intents_assignments_externalcontactsCmd.Short = utils.GenerateCustomDescription(intents_assignments_externalcontactsCmd.Short, intents_assignments_externalcontacts_customerintents.Description, )
	intents_assignments_externalcontactsCmd.Long = intents_assignments_externalcontactsCmd.Short
}
