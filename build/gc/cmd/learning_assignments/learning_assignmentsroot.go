package learning_assignments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_bulkadd"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_bulkremove"
)

func init() {
	learning_assignmentsCmd.AddCommand(learning_assignments_aggregates.Cmdlearning_assignments_aggregates())
	learning_assignmentsCmd.AddCommand(learning_assignments_me.Cmdlearning_assignments_me())
	learning_assignmentsCmd.AddCommand(learning_assignments_bulkadd.Cmdlearning_assignments_bulkadd())
	learning_assignmentsCmd.AddCommand(learning_assignments_bulkremove.Cmdlearning_assignments_bulkremove())
	learning_assignmentsCmd.Short = utils.GenerateCustomDescription(learning_assignmentsCmd.Short, learning_assignments_aggregates.Description, learning_assignments_me.Description, learning_assignments_bulkadd.Description, learning_assignments_bulkremove.Description, )
	learning_assignmentsCmd.Long = learning_assignmentsCmd.Short
}
