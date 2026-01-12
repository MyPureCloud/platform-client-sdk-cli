package intents_assignments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/intents_assignments_externalcontacts"
)

func init() {
	intents_assignmentsCmd.AddCommand(intents_assignments_externalcontacts.Cmdintents_assignments_externalcontacts())
	intents_assignmentsCmd.Short = utils.GenerateCustomDescription(intents_assignmentsCmd.Short, intents_assignments_externalcontacts.Description, )
	intents_assignmentsCmd.Long = intents_assignmentsCmd.Short
}
