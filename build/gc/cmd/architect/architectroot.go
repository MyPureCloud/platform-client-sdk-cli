package architect

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_dependencytracking"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_emergencygroups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_grammars"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_ivrs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_prompts"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_schedulegroups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_systemprompts"
)

func init() {
	architectCmd.AddCommand(architect_dependencytracking.Cmdarchitect_dependencytracking())
	architectCmd.AddCommand(architect_emergencygroups.Cmdarchitect_emergencygroups())
	architectCmd.AddCommand(architect_grammars.Cmdarchitect_grammars())
	architectCmd.AddCommand(architect_ivrs.Cmdarchitect_ivrs())
	architectCmd.AddCommand(architect_prompts.Cmdarchitect_prompts())
	architectCmd.AddCommand(architect_schedulegroups.Cmdarchitect_schedulegroups())
	architectCmd.AddCommand(architect_schedules.Cmdarchitect_schedules())
	architectCmd.AddCommand(architect_systemprompts.Cmdarchitect_systemprompts())
	architectCmd.Short = utils.GenerateCustomDescription(architectCmd.Short, architect_dependencytracking.Description, architect_emergencygroups.Description, architect_grammars.Description, architect_ivrs.Description, architect_prompts.Description, architect_schedulegroups.Description, architect_schedules.Description, architect_systemprompts.Description, )
	architectCmd.Long = architectCmd.Short
}
