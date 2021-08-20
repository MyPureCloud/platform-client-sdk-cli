package coaching_appointments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments_annotations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments_status"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments_statuses"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/coaching_appointments_me"
)

func init() {
	coaching_appointmentsCmd.AddCommand(coaching_appointments_annotations.Cmdcoaching_appointments_annotations())
	coaching_appointmentsCmd.AddCommand(coaching_appointments_conversations.Cmdcoaching_appointments_conversations())
	coaching_appointmentsCmd.AddCommand(coaching_appointments_status.Cmdcoaching_appointments_status())
	coaching_appointmentsCmd.AddCommand(coaching_appointments_statuses.Cmdcoaching_appointments_statuses())
	coaching_appointmentsCmd.AddCommand(coaching_appointments_aggregates.Cmdcoaching_appointments_aggregates())
	coaching_appointmentsCmd.AddCommand(coaching_appointments_me.Cmdcoaching_appointments_me())
	coaching_appointmentsCmd.Short = utils.GenerateCustomDescription(coaching_appointmentsCmd.Short, coaching_appointments_annotations.Description, coaching_appointments_conversations.Description, coaching_appointments_status.Description, coaching_appointments_statuses.Description, coaching_appointments_aggregates.Description, coaching_appointments_me.Description, )
	coaching_appointmentsCmd.Long = coaching_appointmentsCmd.Short
}
