package learning_assignments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_reschedule"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_reset"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_reassign"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_bulkadd"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_bulkremove"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_steps"
)

func init() {
	learning_assignmentsCmd.AddCommand(learning_assignments_aggregates.Cmdlearning_assignments_aggregates())
	learning_assignmentsCmd.AddCommand(learning_assignments_reschedule.Cmdlearning_assignments_reschedule())
	learning_assignmentsCmd.AddCommand(learning_assignments_me.Cmdlearning_assignments_me())
	learning_assignmentsCmd.AddCommand(learning_assignments_reset.Cmdlearning_assignments_reset())
	learning_assignmentsCmd.AddCommand(learning_assignments_reassign.Cmdlearning_assignments_reassign())
	learning_assignmentsCmd.AddCommand(learning_assignments_bulkadd.Cmdlearning_assignments_bulkadd())
	learning_assignmentsCmd.AddCommand(learning_assignments_bulkremove.Cmdlearning_assignments_bulkremove())
	learning_assignmentsCmd.AddCommand(learning_assignments_steps.Cmdlearning_assignments_steps())
	learning_assignmentsCmd.Short = utils.GenerateCustomDescription(learning_assignmentsCmd.Short, learning_assignments_aggregates.Description, learning_assignments_reschedule.Description, learning_assignments_me.Description, learning_assignments_reset.Description, learning_assignments_reassign.Description, learning_assignments_bulkadd.Description, learning_assignments_bulkremove.Description, learning_assignments_steps.Description, )
	learning_assignmentsCmd.Long = learning_assignmentsCmd.Short
}
