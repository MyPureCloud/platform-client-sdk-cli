package coaching_appointments_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments_aggregates_query"
)

func init() {
	coaching_appointments_aggregatesCmd.AddCommand(coaching_appointments_aggregates_query.Cmdcoaching_appointments_aggregates_query())
	coaching_appointments_aggregatesCmd.Short = utils.GenerateCustomDescription(coaching_appointments_aggregatesCmd.Short, coaching_appointments_aggregates_query.Description, )
	coaching_appointments_aggregatesCmd.Long = coaching_appointments_aggregatesCmd.Short
}
