package learning_assignments_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_assignments_aggregates_query"
)

func init() {
	learning_assignments_aggregatesCmd.AddCommand(learning_assignments_aggregates_query.Cmdlearning_assignments_aggregates_query())
	learning_assignments_aggregatesCmd.Short = utils.GenerateCustomDescription(learning_assignments_aggregatesCmd.Short, learning_assignments_aggregates_query.Description, )
	learning_assignments_aggregatesCmd.Long = learning_assignments_aggregatesCmd.Short
}
