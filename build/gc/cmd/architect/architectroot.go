package architect

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_prompts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_systemprompts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_emergencygroups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_ivrs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_schedulegroups"
)

func init() {
	architectCmd.AddCommand(architect_prompts.Cmdarchitect_prompts())
	architectCmd.AddCommand(architect_systemprompts.Cmdarchitect_systemprompts())
	architectCmd.AddCommand(architect_emergencygroups.Cmdarchitect_emergencygroups())
	architectCmd.AddCommand(architect_ivrs.Cmdarchitect_ivrs())
	architectCmd.AddCommand(architect_dependencytracking.Cmdarchitect_dependencytracking())
	architectCmd.AddCommand(architect_schedules.Cmdarchitect_schedules())
	architectCmd.AddCommand(architect_schedulegroups.Cmdarchitect_schedulegroups())
	architectCmd.Short = utils.GenerateCustomDescription(architectCmd.Short, architect_prompts.Description, architect_systemprompts.Description, architect_emergencygroups.Description, architect_ivrs.Description, architect_dependencytracking.Description, architect_schedules.Description, architect_schedulegroups.Description, )
	architectCmd.Long = architectCmd.Short
}
