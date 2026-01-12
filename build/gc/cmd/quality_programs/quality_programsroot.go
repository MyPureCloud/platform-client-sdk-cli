package quality_programs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_programs_agentscoringrules"
)

func init() {
	quality_programsCmd.AddCommand(quality_programs_agentscoringrules.Cmdquality_programs_agentscoringrules())
	quality_programsCmd.Short = utils.GenerateCustomDescription(quality_programsCmd.Short, quality_programs_agentscoringrules.Description, )
	quality_programsCmd.Long = quality_programsCmd.Short
}
